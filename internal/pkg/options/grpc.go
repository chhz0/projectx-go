package options

import "github.com/spf13/pflag"

type GRPC struct {
	Addr           string `json:"addr" mapstructure:"addr"`
	MaxMessageSize int    `json:"max_message_size" mapstructure:"max_message_size"`
}

func (g *GRPC) ApplyFlags(fs *pflag.FlagSet) {
	fs.StringVar(&g.Addr, "grpc.addr", g.Addr, "grpc server address")
	fs.IntVar(&g.MaxMessageSize, "grpc.max-message-size", g.MaxMessageSize, "grpc max message size")
}

func NewGRPCOptions() *GRPC {
	return &GRPC{
		Addr:           "0.0.0.0:443",
		MaxMessageSize: 1024 * 1024 * 4, // 4M
	}
}
