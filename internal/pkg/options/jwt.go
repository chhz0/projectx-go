package options

import (
	"time"

	"github.com/spf13/pflag"
)

type JWT struct {
	Realm   string        `json:"realm" mapstructure:"realm"`
	Key     string        `json:"key" mapstructure:"key"`
	Expire  time.Duration `json:"expire" mapstructure:"expire"`
	Refresh time.Duration `json:"refresh" mapstructure:"refresh"`
}

func (j *JWT) ApplyFlags(fs *pflag.FlagSet) {
	fs.StringVar(&j.Realm, "jwt.realm", j.Realm, "jwt realm")
	fs.StringVar(&j.Key, "jwt.key", j.Key, "jwt key")
	fs.DurationVar(&j.Expire, "jwt.expire", j.Expire, "jwt expire")
	fs.DurationVar(&j.Refresh, "jwt.refresh", j.Refresh, "jwt refresh")
}

func NewJWTOptions() *JWT {
	return &JWT{
		Realm:   "jwt realm",
		Key:     "jwt key",
		Expire:  time.Hour * 24,
		Refresh: time.Hour * 24,
	}
}
