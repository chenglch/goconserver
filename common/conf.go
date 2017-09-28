package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

var (
	serverConfig *ServerConfig
	clientConfig *ClientConfig
)

type ServerConfig struct {
	Global struct {
		Host          string `yaml:"host"`
		SSLKeyFile    string `yaml:"ssl_key_file"`
		SSLCertFile   string `yaml:"ssl_cert_file"`
		SSLCACertFile string `yaml:"ssl_ca_cert_file"`
		LogFile       string `yaml:"logfile"`
		Worker        int    `yaml:"worker"`
	}
	API struct {
		Port        string `yaml:"port"`
		HttpTimeout int    `yaml:"http_timeout"`
	}
	Console struct {
		Port          string `yaml:"port"`
		DataDir       string `yaml:"datadir"`
		LogDir        string `yaml:"logdir"`
		ClientTimeout int    `yaml:"client_timeout"`
		TargetTimeout int    `yaml:"target_timeout"`
	}
}

func InitServerConfig(confFile string) (*ServerConfig, error) {
	serverConfig.Global.Host = "0.0.0.0"
	serverConfig.Global.LogFile = ""
	serverConfig.Global.Worker = 1
	serverConfig.API.Port = "8089"
	serverConfig.API.HttpTimeout = 10
	serverConfig.Console.Port = "12430"
	serverConfig.Console.DataDir = "/var/lib/consoleserver/"
	serverConfig.Console.LogDir = "/var/log/consoleserver/nodes/"
	serverConfig.Console.ClientTimeout = 30
	serverConfig.Console.TargetTimeout = 30
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		return serverConfig, nil
	}
	err = yaml.Unmarshal(data, &serverConfig)
	if err != nil {
		return serverConfig, err
	}
	return serverConfig, nil
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}

type ClientConfig struct {
	SSLKeyFile     string
	SSLCertFile    string
	SSLCACertFile  string
	HTTPUrl        string
	ConsolePort    string
	ConsoleTimeout int
	ServerHost     string
}

func NewClientConfig() (*ClientConfig, error) {
	var err error
	clientConfig = new(ClientConfig)
	clientConfig.HTTPUrl = "http://127.0.0.1:8089"
	clientConfig.ServerHost = "127.0.0.1"
	clientConfig.ConsolePort = "12430"
	clientConfig.ConsoleTimeout = 30
	if os.Getenv("CONGO_URL") != "" {
		clientConfig.HTTPUrl = os.Getenv("CONGO_URL")
	}
	if os.Getenv("CONGO_SERVER_HOST") != "" {
		clientConfig.ServerHost = os.Getenv("CONGO_SERVER_HOST")
	}
	if os.Getenv("CONGO_PORT") != "" {
		clientConfig.ConsolePort = os.Getenv("CONGO_PORT")
	}
	if os.Getenv("CONGO_CONSOLE_TIMEOUT") != "" {
		clientConfig.ConsoleTimeout, err = strconv.Atoi(os.Getenv("CONGO_CONSOLE_TIMEOUT"))
		if err != nil {
			return nil, err
		}
	}
	if os.Getenv("CONGO_SSL_KEY") != "" {
		clientConfig.SSLKeyFile = os.Getenv("CONGO_SSL_KEY")
	}
	if os.Getenv("CONGO_SSL_CERT") != "" {
		clientConfig.SSLCertFile = os.Getenv("CONGO_SSL_CERT")
	}
	if os.Getenv("CONGO_SSL_CA_CERT") != "" {
		clientConfig.SSLCACertFile = os.Getenv("CONGO_SSL_CA_CERT")
	}
	return clientConfig, nil
}

func GetClientConfig() *ClientConfig {
	return clientConfig
}
