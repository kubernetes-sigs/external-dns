package credentials

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/spf13/viper"
)

type FileOpt func(*FileProvider)

// FileOptWithFilename returns a FileOpt overriding the default filename.
func FileOptWithFilename(filename string) FileOpt {
	return func(f *FileProvider) {
		f.filename = filename
	}
}

// FileOptWithAccount returns a FileOpt overriding the default account.
func FileOptWithAccount(account string) FileOpt {
	return func(f *FileProvider) {
		f.account = account
	}
}

type FileProvider struct {
	filename  string
	account   string
	retrieved bool

	// TODO: export some fields from the config like: default zone...etc.
}

func NewFileCredentials(opts ...FileOpt) *Credentials {
	fp := &FileProvider{}
	for _, opt := range opts {
		opt(fp)
	}
	return NewCredentials(fp)
}

func (f *FileProvider) Retrieve() (Value, error) {
	f.retrieved = false

	viperConf, err := f.retrieveViperConfig()
	if err != nil {
		return Value{}, err
	}

	if err := viperConf.ReadInConfig(); err != nil {
		return Value{}, err
	}

	config := Config{}
	if err := viperConf.Unmarshal(&config); err != nil {
		return Value{}, fmt.Errorf("file provider: couldn't read config: %w", err)
	}

	if len(config.Accounts) == 0 {
		return Value{}, fmt.Errorf("file provider: no accounts were found into %q", viper.ConfigFileUsed())
	}

	if f.account == "" && config.DefaultAccount == "" {
		return Value{}, fmt.Errorf("file provider: no account defined")
	}

	accountName := config.DefaultAccount
	if f.account != "" {
		accountName = f.account
	}

	account := Account{}
	for i, a := range config.Accounts {
		if a.Name == accountName {
			account = config.Accounts[i]
			break
		}
	}

	v := Value{
		APIKey:    account.Key,
		APISecret: account.Secret,
	}

	if !v.IsSet() {
		return Value{}, fmt.Errorf("file provider: account %q: %w", accountName, ErrMissingIncomplete)
	}

	f.retrieved = true

	return v, nil
}

// IsExpired returns if the shared credentials have expired.
func (f *FileProvider) IsExpired() bool {
	return !f.retrieved
}

type Account struct {
	Name                string
	Account             string
	SosEndpoint         string
	Environment         string
	Key                 string
	Secret              string
	SecretCommand       []string
	DefaultZone         string
	DefaultSSHKey       string
	DefaultTemplate     string
	DefaultOutputFormat string
	ClientTimeout       int
	CustomHeaders       map[string]string
}

type Config struct {
	DefaultAccount      string
	DefaultOutputFormat string
	Accounts            []Account
}

func (f *FileProvider) retrieveViperConfig() (*viper.Viper, error) {
	config := viper.New()

	if f.filename != "" {
		config.SetConfigFile(f.filename)
		return config, nil
	}

	cfgdir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("could not find configuration directory: %s", err)
	}

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	config.SetConfigName("exoscale")
	config.SetConfigType("toml")
	config.AddConfigPath(path.Join(cfgdir, "exoscale"))
	config.AddConfigPath(path.Join(usr.HomeDir, ".exoscale"))
	config.AddConfigPath(usr.HomeDir)
	config.AddConfigPath(".")

	return config, nil
}
