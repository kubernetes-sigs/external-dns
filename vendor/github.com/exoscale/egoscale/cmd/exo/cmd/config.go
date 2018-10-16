package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigFileName = "exoscale"
	defaultEndpoint       = "https://api.exoscale.ch/compute"
	defaultTemplate       = "Linux Ubuntu 18.04 LTS 64-bit"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate config file for this cli",
}

func configCmdRun(cmd *cobra.Command, args []string) error {
	if viper.ConfigFileUsed() != "" {
		fmt.Println("Good day! exo is already configured with accounts:")
		listAccounts()
		return addNewAccount(false)
	}
	csPath, ok := isCloudstackINIFileExist()
	if ok {
		resp, ok, err := askCloudstackINIMigration(csPath)
		if err != nil {
			return err
		}
		if !ok {
			return addNewAccount(true)
		}

		cfgPath, err := createConfigFile(defaultConfigFileName)
		if err != nil {
			return err
		}
		if err := importCloudstackINI(resp, csPath, cfgPath); err != nil {
			return err
		}
		return addNewAccount(false)
	}
	fmt.Print(`
Hi happy Exoscalian, some configuration is required to use exo.

We now need some very important information, find them there.
	<https://portal.exoscale.com/account/profile/api>

`)
	return addNewAccount(true)
}

func addNewAccount(firstRun bool) error {

	config := &config{}

	if firstRun {
		filePath, err := createConfigFile(defaultConfigFileName)
		if err != nil {
			return err
		}

		viper.SetConfigFile(filePath)

		newAccount, err := getAccount()
		if err != nil {
			return err
		}
		config.DefaultAccount = newAccount.Name
		config.Accounts = []account{*newAccount}
		viper.Set("defaultAccount", newAccount.Name)
	}

	for askQuestion("Do you wish to add another account?") {
		newAccount, err := getAccount()
		if err != nil {
			return err
		}
		config.Accounts = append(config.Accounts, *newAccount)
		if askQuestion("Make [" + newAccount.Name + "] your default profile?") {
			config.DefaultAccount = newAccount.Name
			viper.Set("defaultAccount", newAccount.Name)
		}
	}

	if len(config.Accounts) == 0 {
		return nil
	}

	return addAccount(viper.ConfigFileUsed(), config)
}

func getAccount() (*account, error) {
	reader := bufio.NewReader(os.Stdin)

	var client *egoscale.Client

	account := &account{
		Endpoint: defaultEndpoint,
		Key:      "",
		Secret:   "",
	}

	for i := 0; ; i++ {
		if i > 0 {
			endpoint, err := readInput(reader, "API Endpoint", account.Endpoint)
			if err != nil {
				return nil, err
			}
			if endpoint != account.Endpoint {
				account.Endpoint = endpoint
			}
		}

		apiKey, err := readInput(reader, "API Key", account.Key)
		if err != nil {
			return nil, err
		}
		if apiKey != account.Key {
			account.Key = apiKey
		}

		secret := ""
		if account.Secret != "" && len(account.Secret) > 10 {
			secret = account.Secret[0:7] + "..."
		}
		secretKey, err := readInput(reader, "Secret Key", secret)
		if err != nil {
			return nil, err
		}
		if secretKey != account.Secret && secretKey != secret {
			account.Secret = secretKey
		}

		client = egoscale.NewClient(account.Endpoint, account.Key, account.Secret)

		fmt.Printf("Checking the credentials of %q...", account.Key)
		acc := &egoscale.Account{}
		err = client.GetWithContext(gContext, acc)
		if err != nil {
			fmt.Print(` failure.

Let's start over.

`)
		} else {
			fmt.Print(" success!\n\n")
			account.Name = acc.Name
			account.Account = acc.Name
			break
		}
	}

	name, err := readInput(reader, "Account name", account.Name)
	if err != nil {
		return nil, err
	}
	if name != "" {
		account.Name = name
	}

	for isAccountExist(account.Name) {
		fmt.Printf("Account name [%s] already exist\n", name)
		name, err = readInput(reader, "Account name", account.Name)
		if err != nil {
			return nil, err
		}

		account.Name = name
	}

	defaultZone, err := chooseZone(account.Name, client)
	if err != nil {
		return nil, err
	}

	account.DefaultZone = defaultZone
	account.DNSEndpoint = strings.Replace(account.Endpoint, "/compute", "/dns", 1)

	return account, nil
}

func addAccount(filePath string, newAccounts *config) error {

	accountsSize := 0
	currentAccounts := []account{}
	if gAllAccount != nil {
		accountsSize = len(gAllAccount.Accounts)
		currentAccounts = gAllAccount.Accounts
	}

	newAccountsSize := 0

	if newAccounts != nil {
		newAccountsSize = len(newAccounts.Accounts)
	}

	accounts := make([]map[string]string, accountsSize+newAccountsSize)

	conf := &config{}

	for i, acc := range currentAccounts {

		accounts[i] = map[string]string{}

		accounts[i]["name"] = acc.Name
		accounts[i]["computeEndpoint"] = acc.Endpoint
		accounts[i]["dnsEndpoint"] = acc.DNSEndpoint
		accounts[i]["key"] = acc.Key
		accounts[i]["secret"] = acc.Secret
		accounts[i]["defaultZone"] = acc.DefaultZone
		accounts[i]["defaultTemplate"] = acc.DefaultTemplate
		accounts[i]["account"] = acc.Account

		conf.Accounts = append(conf.Accounts, acc)
	}

	if newAccounts != nil {

		for i, acc := range newAccounts.Accounts {

			accounts[accountsSize+i] = map[string]string{}

			accounts[accountsSize+i]["name"] = acc.Name
			accounts[accountsSize+i]["computeEndpoint"] = acc.Endpoint
			accounts[accountsSize+i]["dnsEndpoint"] = acc.DNSEndpoint
			accounts[accountsSize+i]["key"] = acc.Key
			accounts[accountsSize+i]["secret"] = acc.Secret
			accounts[accountsSize+i]["defaultZone"] = acc.DefaultZone
			accounts[accountsSize+i]["defaultTemplate"] = acc.DefaultTemplate
			accounts[accountsSize+i]["account"] = acc.Account
			conf.Accounts = append(conf.Accounts, acc)
		}
	}

	viper.SetConfigType("toml")
	viper.SetConfigFile(filePath)

	viper.Set("accounts", accounts)

	if err := viper.WriteConfig(); err != nil {
		return err
	}

	conf.DefaultAccount = viper.Get("defaultAccount").(string)
	gAllAccount = conf

	return nil

}

func isCloudstackINIFileExist() (string, bool) {

	envConfigPath := os.Getenv("CLOUDSTACK_CONFIG")

	usr, _ := user.Current()

	localConfig, _ := filepath.Abs("cloudstack.ini")
	inis := []string{
		localConfig,
		filepath.Join(usr.HomeDir, ".cloudstack.ini"),
		filepath.Join(gConfigFolder, "cloudstack.ini"),
		envConfigPath,
	}

	cfgPath := ""

	for _, i := range inis {
		if _, err := os.Stat(i); err != nil {
			continue
		}
		cfgPath = i
		break
	}

	if cfgPath == "" {
		return "", false
	}
	return cfgPath, true
}

func askCloudstackINIMigration(csFilePath string) (string, bool, error) {

	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, csFilePath)
	if err != nil {
		return "", false, err
	}

	if len(cfg.Sections()) <= 0 {
		return "", false, nil
	}

	fmt.Printf("We've found a %q configuration file with the following configurations:\n", "cloudstack.ini")
	for i, acc := range cfg.Sections() {
		if i == 0 {
			continue
		}
		fmt.Printf("- [%s] %s\n", acc.Name(), acc.Key("key").String())
	}
	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)

	resp, err := readInput(reader, "Which one should we import?", "All, some, none")
	if err != nil {
		return "", false, err
	}

	resp = strings.ToLower(resp)
	if resp == "" {
		resp = "all"
	}

	return resp, (resp == "all" || resp == "some"), nil
}

func importCloudstackINI(option, csPath, cfgPath string) error {
	cfg, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, csPath)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)

	config := &config{}

	setdefaultAccount := 1
	for i, acc := range cfg.Sections() {
		if i == 0 {
			continue
		}

		if option == "some" {
			if !askQuestion(fmt.Sprintf("Do you want to import [%s] %s?", acc.Name(), acc.Key("key").String())) {
				if viper.Get("defaultAccount") == nil {
					setdefaultAccount = i + 1
				}
				continue
			}
		}

		csAccount := account{
			Name:     acc.Name(),
			Endpoint: acc.Key("endpoint").String(),
			Key:      acc.Key("key").String(),
			Secret:   acc.Key("secret").String(),
		}

		csClient := egoscale.NewClient(csAccount.Endpoint, csAccount.Key, csAccount.Secret)

		fmt.Printf("Checking the credentials of %q...", csAccount.Key)
		a := &egoscale.Account{}
		err := csClient.GetWithContext(gContext, a)
		if err != nil {
			fmt.Println(" failure.")
			if !askQuestion(fmt.Sprintf("Do you want to keep %s?", acc.Name())) {
				continue
			}
		} else {
			fmt.Println(" success!")
			csAccount.Name = a.Name
			csAccount.Account = a.Name
		}
		fmt.Println("")

		name, err := readInput(reader, "Account name", csAccount.Name)
		if err != nil {
			return err
		}
		if name != "" {
			csAccount.Name = name
		}

		for isAccountExist(csAccount.Name) {
			fmt.Printf("Account name [%s] already exist\n", csAccount.Name)
			name, err = readInput(reader, "Account name", csAccount.Name)
			if err != nil {
				return err
			}

			csAccount.Name = name
		}

		defaultZone, err := chooseZone(csAccount.Name, csClient)
		if err != nil {
			return err
		}

		csAccount.DefaultZone = defaultZone

		isDefault := false
		if askQuestion(fmt.Sprintf("Is %q your default profile?", csAccount.Name)) {
			isDefault = true
		}

		config.Accounts = append(config.Accounts, csAccount)

		if i == setdefaultAccount || isDefault {
			config.DefaultAccount = csAccount.Name
			viper.Set("defaultAccount", csAccount.Name)
		}
		gAllAccount = config
	}

	gAllAccount = nil
	return addAccount(cfgPath, config)
}

func isAccountExist(name string) bool {

	if gAllAccount == nil {
		return false
	}

	for _, acc := range gAllAccount.Accounts {
		if acc.Name == name {
			return true
		}
	}

	return false
}

func createConfigFile(fileName string) (string, error) {
	if _, err := os.Stat(gConfigFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(gConfigFolder, os.ModePerm); err != nil {
			return "", err
		}
	}

	filepath := path.Join(gConfigFolder, fileName+".toml")

	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		return "", fmt.Errorf("%q exists already", filepath)
	}
	return filepath, nil
}

func readInput(reader *bufio.Reader, text, def string) (string, error) {
	if def == "" {
		fmt.Printf("[+] %s [%s]: ", text, "none")
	} else {
		fmt.Printf("[+] %s [%s]: ", text, def)
	}
	c := make(chan bool)
	defer close(c)

	input := ""
	var err error
	go func() {
		input, err = reader.ReadString('\n')
		c <- true
	}()

	select {
	case <-c:
	case <-gContext.Done():
		err = fmt.Errorf("")
	}

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		input = def
	}
	return input, nil
}

func askQuestion(text string) bool {

	reader := bufio.NewReader(os.Stdin)

	resp, err := readInput(reader, text, "yN")
	if err != nil {
		log.Fatal(err)
	}

	return (strings.ToLower(resp) == "y" || strings.ToLower(resp) == "yes")
}

func listAccounts() {
	if gAllAccount == nil {
		return
	}
	for _, acc := range gAllAccount.Accounts {
		print("- ", acc.Name)
		if acc.Name == gAllAccount.DefaultAccount {
			print(" [Default]")
		}
		fmt.Println("")
	}
}

func getAccountByName(name string) *account {
	if gAllAccount == nil {
		return nil
	}
	for i, acc := range gAllAccount.Accounts {
		if acc.Name == name {
			return &gAllAccount.Accounts[i]
		}
	}
	return nil
}

func getSelectedZone(number string, zones map[string]string) (string, bool) {
	zName, ok := zones[number]
	if !ok {
		return "", false
	}
	return zName, true
}

func chooseZone(accountName string, cs *egoscale.Client) (string, error) {

	reader := bufio.NewReader(os.Stdin)

	zonesResp, err := cs.ListWithContext(gContext, &egoscale.Zone{})
	if err != nil {
		return "", err
	}

	zones := map[string]string{}

	if len(zonesResp) == 0 {
		return "", fmt.Errorf("no zones were found")
	}

	fmt.Printf("Choose %q default zone:\n", accountName)

	for i, z := range zonesResp {
		zone := z.(*egoscale.Zone)

		zName := strings.ToLower(zone.Name)

		n := fmt.Sprintf("%d", i+1)

		zones[n] = zName

		fmt.Printf("%d: %s\n", i+1, zName)
	}

	zoneNumber, err := readInput(reader, "Select", "1")
	if err != nil {
		return "", err
	}

	defaultZone, ok := getSelectedZone(zoneNumber, zones)
	for !ok {
		fmt.Println("Error: Invalid zone number")
		defaultZone, err = chooseZone(accountName, cs)
		if err == nil {
			break
		}
	}
	return defaultZone, nil
}

func init() {

	configCmd.RunE = configCmdRun
	RootCmd.AddCommand(configCmd)
}
