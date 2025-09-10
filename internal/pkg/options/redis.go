package options

import "github.com/spf13/pflag"

type Redis struct {
	Url      string `json:"url" mapstructure:"url"`
	User     string `json:"user" mapstructure:"user"`
	Passwrod string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}

func (r *Redis) LocalFlags(fs *pflag.FlagSet) {
	fs.StringVar(&r.Url, "redis.url", r.Url, "redis url")
	fs.StringVar(&r.User, "redis.user", r.User, "redis user")
	fs.StringVar(&r.Passwrod, "redis.password", r.Passwrod, "redis password")
	fs.IntVar(&r.DB, "redis.db", r.DB, "redis database")

}

func (r *Redis) Validate() error {
	return nil
}

func NewRedisOptions() *Redis {
	return &Redis{
		Url:      "127.0.0.1:6379",
		User:     "",
		Passwrod: "",
		DB:       0,
	}
}
