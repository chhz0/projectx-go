package options

import "github.com/spf13/pflag"

type Log struct {
	Level      string   `json:"level" mapstructure:"level"`
	Encoding   string   `json:"encoding" mapstructure:"encoding"`
	Caller     bool     `json:"caller" mapstructure:"caller"`
	CallerSkip int      `json:"caller_skip" mapstructure:"caller_skip"`
	Output     []string `json:"output" mapstructure:"output"`
}

func (l *Log) ApplyFlags(fs *pflag.FlagSet) {
	fs.StringVar(&l.Level, "log.level", l.Level, "log level")
	fs.StringVar(&l.Encoding, "log.encoding", l.Encoding, "log encoding")
	fs.BoolVar(&l.Caller, "log.caller", l.Caller, "log caller")
	fs.IntVar(&l.CallerSkip, "log.caller_skip", l.CallerSkip, "log caller skip")
	fs.StringSliceVar(&l.Output, "log.output", l.Output, "log output")
}

func NewLogOptions() *Log {
	return &Log{
		Level:      "info",
		Encoding:   "console",
		Caller:     true,
		CallerSkip: 1,
		Output:     []string{"stdout"},
	}
}
