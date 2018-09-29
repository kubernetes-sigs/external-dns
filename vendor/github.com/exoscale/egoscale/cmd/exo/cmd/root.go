package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
	"path"
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gContext context.Context

var gConfigFolder string
var gConfigFilePath string

//current Account information
var gAccountName string
var gCurrentAccount *account

var gAllAccount *config

//egoscale client
var cs *egoscale.Client
var csDNS *egoscale.Client

//Aliases
var gListAlias = []string{"ls"}
var gRemoveAlias = []string{"rm"}
var gDeleteAlias = []string{"del"}
var gShowAlias = []string{"get"}
var gCreateAlias = []string{"add"}
var gUploadAlias = []string{"up"}
var gDissociateAlias = []string{"disassociate", "dissoc"}
var gAssociateAlias = []string{"assoc"}

type account struct {
	Name            string
	Account         string
	Endpoint        string
	DNSEndpoint     string
	Key             string
	Secret          string
	DefaultZone     string
	DefaultTemplate string
}

type config struct {
	DefaultAccount string
	Accounts       []account
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "exo",
	Short:         "A simple CLI to use CloudStack using egoscale lib",
	SilenceUsage:  true,
	SilenceErrors: true,
	//Long:  `A simple CLI to use CloudStack using egoscale lib`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of exo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("exo version %s\n", egoscale.Version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	gContext = ctx

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&gConfigFilePath, "config", "C", "", "Specify an alternate config file [env EXOSCALE_CONFIG]")
	RootCmd.PersistentFlags().StringVarP(&gAccountName, "account", "A", "", "Account to use in config file [env EXOSCALE_ACCOUNT]")
	RootCmd.AddCommand(versionCmd)

	cobra.OnInitialize(initConfig, buildClient)
}

var ignoreClientBuild = false

func buildClient() {
	if ignoreClientBuild {
		return
	}

	if cs != nil {
		return
	}

	csDNS = egoscale.NewClient(gCurrentAccount.DNSEndpoint, gCurrentAccount.Key, gCurrentAccount.Secret)

	cs = egoscale.NewClient(gCurrentAccount.Endpoint, gCurrentAccount.Key, gCurrentAccount.Secret)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	envs := map[string]string{
		"EXOSCALE_CONFIG":  "config",
		"EXOSCALE_ACCOUNT": "account",
	}

	for env, flag := range envs {
		flag := RootCmd.Flags().Lookup(flag)
		if value, ok := os.LookupEnv(env); ok {
			if err := flag.Value.Set(value); err != nil {
				log.Fatal(err)
			}
		}
	}

	// an attempt to mimic existing behaviours

	envEndpoint := readFromEnv(
		"EXOSCALE_ENDPOINT",
		"EXOSCALE_COMPUTE_ENDPOINT",
		"CLOUDSTACK_ENDPOINT")

	envKey := readFromEnv(
		"EXOSCALE_KEY",
		"EXOSCALE_API_KEY",
		"CLOUDSTACK_KEY",
		"CLOUSTACK_API_KEY",
	)

	envSecret := readFromEnv(
		"EXOSCALE_SECRET",
		"EXOSCALE_API_SECRET",
		"EXOSCALE_SECRET_KEY",
		"CLOUDSTACK_SECRET",
		"CLOUDSTACK_SECRET_KEY",
	)

	if envEndpoint != "" && envKey != "" && envSecret != "" {
		cs = egoscale.NewClient(envEndpoint, envKey, envSecret)
		return
	}

	config := &config{}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	gConfigFolder = path.Join(usr.HomeDir, ".exoscale")

	if gConfigFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(gConfigFilePath)
	} else {
		// Search config in home directory with name ".cobra_test" (without extension).
		viper.SetConfigName("exoscale")
		viper.AddConfigPath(path.Join(usr.HomeDir, ".exoscale"))
		viper.AddConfigPath(usr.HomeDir)
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil && getCmdPosition("config") == 1 {
		ignoreClientBuild = true
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatal(fmt.Errorf("couldn't read config: %s", err))
	}

	if config.DefaultAccount == "" && gAccountName == "" {
		log.Fatalf("default account not defined")
	}

	if gAccountName == "" {
		gAccountName = config.DefaultAccount
	}

	gAllAccount = config
	gAllAccount.DefaultAccount = gAccountName

	for i, acc := range config.Accounts {
		if acc.Name == gAccountName {
			gCurrentAccount = &config.Accounts[i]
			break
		}
	}

	if gCurrentAccount == nil {
		log.Fatalf("could't find any account with name: %q", gAccountName)
	}

	if gCurrentAccount.Endpoint == "" {
		gCurrentAccount.Endpoint = defaultEndpoint
	}

	if gCurrentAccount.DNSEndpoint == "" {
		gCurrentAccount.DNSEndpoint = strings.Replace(gCurrentAccount.Endpoint, "/compute", "/dns", 1)
	}
	if gCurrentAccount.DefaultTemplate == "" {
		gCurrentAccount.DefaultTemplate = defaultTemplate
	}

	gCurrentAccount.Endpoint = strings.TrimRight(gCurrentAccount.Endpoint, "/")
	gCurrentAccount.DNSEndpoint = strings.TrimRight(gCurrentAccount.DNSEndpoint, "/")
}

// getCmdPosition returns a command position by fetching os.args and ignoring flags
//
// example: "$ exo -r preprod vm create" vm position is 1 and create is 2
//
func getCmdPosition(cmd string) int {

	count := 1

	isFlagParam := false

	for _, arg := range os.Args[1:] {

		if strings.HasPrefix(arg, "-") {

			flag := RootCmd.Flags().Lookup(strings.Trim(arg, "-"))
			if flag == nil {
				flag = RootCmd.Flags().ShorthandLookup(strings.Trim(arg, "-"))
			}

			if flag != nil && (flag.Value.Type() != "bool") {
				isFlagParam = true
			}
			continue
		}

		if isFlagParam {
			isFlagParam = false
			continue
		}

		if arg == cmd {
			break
		}
		count++
	}

	return count
}

// readFromEnv is a os.Getenv on steroids
func readFromEnv(keys ...string) string {
	for _, key := range keys {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
	}
	return ""
}
