package options

import "github.com/spf13/pflag"

type HTTP struct {
	Addr    string `json:"addr" mapstructure:"addr"`
	OpenTLS bool   `json:"open_tls" mapstructure:"open_tls"`
	TLS     *TLS   `json:"tls" mapstructure:"tls"`
}

type TLS struct {
	CertFile string `json:"cert_file" mapstructure:"cert_file"`
	KeyFile  string `json:"key_file" mapstructure:"key_file"`
}

func (h *HTTP) ApplyFlags(fs *pflag.FlagSet) {
	fs.StringVar(&h.Addr, "http.addr", h.Addr, "http server address")
	fs.BoolVar(&h.OpenTLS, "http.open-tls", h.OpenTLS, "open tls")
	fs.StringVar(&h.TLS.CertFile, "http.tls.cert", h.TLS.CertFile, "tls cert file")
	fs.StringVar(&h.TLS.KeyFile, "http.tls.key", h.TLS.KeyFile, "tls key file")
}

func NewHTTPOptions() *HTTP {
	return &HTTP{
		Addr:    "127.0.0.1:8080",
		OpenTLS: false,
		TLS: &TLS{
			CertFile: "",
			KeyFile:  "",
		},
	}
}
