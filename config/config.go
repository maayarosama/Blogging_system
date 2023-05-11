package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Configuration struct {
	Server     Server     `json:"server"`
	Database   DB         `json:"database"`
	Token      JwtToken   `json:"token"`
	MailSender MailSender `json:"mailSender"`
}

type MailSender struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Timeout  int    `json:"timeout"`
}

// JWT Authentation token data
type JwtToken struct {
	Secret  string `json:"secret"`
	Timeout int    `json:"timeout"`
}

// Server struct to hold server's information
type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// DB struct to hold database's file path
type DB struct {
	Path string `json:"file"`
}

// Read config file to bytes
func ReadConfigFile(path string) ([]byte, error) {
	confFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer confFile.Close()
	conf, err := io.ReadAll(confFile)
	if err != nil {
		return conf, err
	}
	return conf, nil
}

// parse config file
func ParseConfigFile(conf []byte) (Configuration, error) {
	var myConfig Configuration
	err := json.Unmarshal(conf, &myConfig)
	if err != nil {
		return myConfig, err
	}

	if myConfig.Server.Host == "" || myConfig.Server.Port == "" {
		return myConfig, errors.New("server host and port are required")
	}

	if myConfig.Database.Path == "" {
		return myConfig, errors.New("database file is required")
	}
	if myConfig.Token.Secret == "" || myConfig.Token.Timeout == 0 {
		return myConfig, errors.New("jwt token configuration is required")
	}

	if myConfig.MailSender.Email == "" || myConfig.MailSender.Password == "" || myConfig.MailSender.Timeout == 0 {
		return myConfig, errors.New("mail sender configuration is required")
	}

	return myConfig, nil
}
