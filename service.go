package psqlsrv

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	rscsrv "github.com/lab259/go-rscsrv"
	_ "github.com/lib/pq"
)

type SSLMode string

const (
	SSLModeDisable    SSLMode = ""
	SSLModeRequire    SSLMode = "require"
	SSLModeVerifyCA   SSLMode = "verify-ca"
	SSLModeVerifyFull SSLMode = "verify-full"
)

type Configuration struct {
	Database    string        `yaml:"database"`
	Username    string        `yaml:"username"`
	Password    string        `yaml:"password"`
	Host        string        `yaml:"host"`
	Port        int           `yaml:"port"`
	SSLMode     SSLMode       `yaml:"ssl_mode"`
	Timeout     time.Duration `yaml:"timeout"`
	SSLCert     string        `yaml:"ssl_cert"`
	SSLKey      string        `yaml:"ssl_key"`
	SSLRootCert string        `yaml:"ssl_root_cert"`
}

func (config Configuration) formatParameter(name, value string) string {
	safeValue := strings.Replace(value, "'", `\'`, -1)
	return fmt.Sprintf("%s='%s'", name, safeValue)
}

func (config Configuration) ConnectionString() string {
	var parameters []string
	if config.Database != "" {
		parameters = append(parameters, config.formatParameter("dbname", config.Database))
	}

	if config.Username != "" {
		parameters = append(parameters, config.formatParameter("user", config.Username))
	}

	if config.Password != "" {
		parameters = append(parameters, config.formatParameter("password", config.Password))
	}

	if config.Host != "" {
		parameters = append(parameters, config.formatParameter("host", config.Host))
	}

	if config.Port != 0 {
		parameters = append(parameters, config.formatParameter("port", strconv.Itoa(config.Port)))
	}

	parameters = append(parameters, config.formatParameter("sslmode", string(config.SSLMode)))

	if config.Timeout > 0 {
		parameters = append(parameters, config.formatParameter("connect_timeout", strconv.Itoa(int(config.Timeout/time.Second))))
	}

	if config.SSLCert != "" {
		parameters = append(parameters, config.formatParameter("sslcert", config.SSLCert))
	}

	if config.SSLKey != "" {
		parameters = append(parameters, config.formatParameter("sslkey", config.SSLKey))
	}

	if config.SSLRootCert != "" {
		parameters = append(parameters, config.formatParameter("sslrootcert", config.SSLRootCert))
	}

	return strings.Join(parameters, " ")
}

type PsqlService struct {
	running       bool
	db            *sql.DB
	Configuration Configuration
}

func (srv *PsqlService) Name() string {
	return "Psql Service"
}

func (srv *PsqlService) LoadConfiguration() (interface{}, error) {
	return nil, errors.New("not implemented")
}

func (srv *PsqlService) ApplyConfiguration(configuration interface{}) error {
	switch c := configuration.(type) {
	case Configuration:
		srv.Configuration = c
	case *Configuration:
		srv.Configuration = *c
	default:
		return rscsrv.ErrWrongConfigurationInformed
	}

	return nil
}

func (srv *PsqlService) Restart() error {
	if srv.running {
		if err := srv.Stop(); err != nil {
			return err
		}
	}
	return srv.Start()
}

func (srv *PsqlService) Start() error {
	if srv.running {
		return nil
	}

	db, err := sql.Open("postgres", srv.Configuration.ConnectionString())
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	srv.db = db
	srv.running = true
	return nil
}

func (srv *PsqlService) Stop() error {
	if srv.running {
		if err := srv.db.Close(); err != nil {
			return err
		}
		srv.running = false
	}
	return nil
}

func (srv *PsqlService) Ping() error {
	if !srv.running {
		return rscsrv.ErrServiceNotRunning
	}

	return srv.db.Ping()
}

func (srv *PsqlService) DB() (*sql.DB, error) {
	if !srv.running {
		return nil, rscsrv.ErrServiceNotRunning
	}

	return srv.db, nil
}
