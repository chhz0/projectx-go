package app

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	// ffile mean full-file must include path, name and type. e.g. ./config/app.yaml
	ffile string

	Viper *viper.Viper
}

func NewConfig(fname, ftype string, paths ...string) *Config {
	v := viper.New()

	v.SetConfigName(fname)
	v.SetConfigType(ftype)
	for _, p := range paths {
		v.AddConfigPath(p)
	}

	return &Config{
		ffile: "",
		Viper: v,
	}
}

func (c *Config) SetEnv(prefix string, allowEmpty bool) {
	c.Viper.AutomaticEnv()
	c.Viper.SetEnvPrefix(prefix)
	c.Viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	c.Viper.AllowEmptyEnv(allowEmpty)
}

func (c *Config) Read(errHandler func(error)) {
	if c.ffile != "" {
		c.Viper.SetConfigFile(c.ffile)
	}

	if err := c.Viper.ReadInConfig(); err != nil {
		errHandler(err)
	}
}

func (c *Config) Unmarshal(to any, errHandler func(error)) {
	if err := c.Viper.Unmarshal(to); err != nil {
		errHandler(err)
	}
}

func (c *Config) Watch(onChange func(v *viper.Viper, in fsnotify.Event)) {
	c.Viper.WatchConfig()
	c.Viper.OnConfigChange(func(in fsnotify.Event) {
		onChange(c.Viper, in)
	})
}

func (c *Config) SetFlag(fs *pflag.FlagSet) {
	fs.StringVarP(&c.ffile, "config", "c", "", "A complete configuration file, including the path, name, and type.")
	_ = c.Viper.BindPFlags(fs)
}
