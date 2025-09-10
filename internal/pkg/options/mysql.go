package options

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DSN string `json:"dsn" mapstructure:"dsn"`

	Host     string `json:"host" mapstructure:"host"`
	Port     string `json:"port" mapstructure:"port"`
	User     string `json:"user" mapstructure:"user"`
	Password string `json:"password" mapstructure:"password"`
	DB       string `json:"db" mapstructure:"db"`

	MaxIdleConns int           `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int           `json:"max_open_conns" mapstructure:"max_open_conns"`
	MaxConnsLife time.Duration `json:"max_conns_life" mapstructure:"max_conns_life"`

	LogLevel string `json:"log_level" mapstructure:"log_level"`
	LogFile  string `json:"log_file" mapstructure:"log_file"`
}

func (m *MySQL) GetDSN() string {
	if m.DSN == "" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.User, m.Password, m.Host, m.Port, m.DB,
		)
	}
	return m.DSN
}

// OpenDB open mysql database
func (m *MySQL) OpenDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(m.GetDSN()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(m.MaxConnsLife)

	return db, nil
}

func (m *MySQL) ApplyFlags(fs *pflag.FlagSet) {
	fs.StringVar(&m.DSN, "mysql.dsn", m.DSN, "mysql dsn [$USER:$PWD@tcp($HOST:$PORT)/$DB?...]")
	fs.StringVar(&m.LogLevel, "mysql.log-level", m.LogLevel, "mysql log level")
}

func (m *MySQL) Validate() []error {
	return nil
}

func NewMySQLOptions() *MySQL {
	return &MySQL{
		DSN:          "",
		Host:         "127.0.0.1",
		Port:         "3306",
		User:         "root",
		Password:     "",
		DB:           "",
		MaxIdleConns: 100,
		MaxOpenConns: 100,
		MaxConnsLife: time.Second * 10,
		LogLevel:     "info",
		LogFile:      "",
	}
}
