package internal

import (
	"os"
	"testing"
)

var config = `{
	"server": {
		"host": "localhost",
		"port": ":3000"
	},
	"database": {
		"file": "./test.db"
	},
	"token": {
		"secret": "secret",
		"timeout": 200
	},
	"mailSender": {
		"email" : "g.gmail@com",
		"password": "password",
		"timeout": 10
	}

}
`

func TestReadConfig(t *testing.T) {

	t.Run("config file doesn't exist", func(t *testing.T) {

		configPath := "/config.json"

		_, err := ReadConfigFile(configPath)
		if err == nil {
			t.Errorf("File doesn't exist: %v", err)
		}
	})
	t.Run("read config file", func(t *testing.T) {

		configPath := t.TempDir() + "/config.json"
		err := os.WriteFile(configPath, []byte(config), 0644)

		if err != nil {
			t.Errorf("Couldn't write config file: %v", err)
		}

		_, err = ReadConfigFile(configPath)
		if err != nil {
			t.Errorf("Couldn't read config file: %v", err)
		}
	})

}

func TestParseConfig(t *testing.T) {

	t.Run("read correct data from config file", func(t *testing.T) {

		configPath := t.TempDir() + "/config.json"
		err := os.WriteFile(configPath, []byte(config), 0644)

		if err != nil {
			t.Errorf("Couldn't write config file: %v", err)
		}

		data, err := ReadConfigFile(configPath)
		if err != nil {
			t.Errorf("Couldn't read config file: %v", err)
		}

		got, _ := ParseConfigFile(data)
		want := Configuration{
			Server: Server{
				Host: "localhost",
				Port: ":3000",
			},
			Database: DB{
				Path: "./test.db",
			},
			Token: JwtToken{
				Secret:  "secret",
				Timeout: 200,
			},
			MailSender: MailSender{
				Email:    "g.gmail@com",
				Password: "password",
				Timeout:  10,
			},
		}

		if got != want {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}

	})

	t.Run("read incorrect data from config file", func(t *testing.T) {

		configPath := t.TempDir() + "/config.json"
		err := os.WriteFile(configPath, []byte(config), 0644)

		if err != nil {
			t.Errorf("Couldn't write config file: %v", err)
		}

		data, err := ReadConfigFile(configPath)
		if err != nil {
			t.Errorf("Couldn't read config file: %v", err)
		}

		got, _ := ParseConfigFile(data)
		want := Configuration{
			Server: Server{
				Host: "localhost",
				Port: ":3000",
			},
			Database: DB{
				Path: "./test.db",
			},
			Token: JwtToken{
				Secret:  "secret",
				Timeout: 200,
			},
			MailSender: MailSender{
				Email:    "g.@com",
				Password: "password",
				Timeout:  10,
			},
		}

		if got == want {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}

	})

	// t.Run("read data from config file missing server", func(t *testing.T) {

	// 	var config = `{
	// 		"database": {
	// 			"file": "./test.db"
	// 		},
	// 		"token": {
	// 			"secret": "secret",
	// 			"timeout": 200
	// 		},
	// 		"mailSender": {
	// 			"email" : "g.gmail@com",
	// 			"password": "password",
	// 			"timeout": 10
	// 		}

	// 	}
	// 	`

	// 	configPath := t.TempDir() + "/config.json"
	// 	err := os.WriteFile(configPath, []byte(config), 0644)

	// 	if err != nil {
	// 		t.Errorf("Couldn't write config file: %v", err)
	// 	}

	// 	data, err := ReadConfigFile(configPath)
	// 	if err != nil {
	// 		t.Errorf("Couldn't read config file: %v", err)
	// 	}

	// 	_, got := ParseConfigFile(data)
	// 	want := errors.New("server host and port are required")
	// 	// if err == nil {
	// 	// 	t.Errorf("Should not have passed : %v", err)
	// 	// }
	// 	if got != want {
	// 		t.Errorf("Got: %v. Expected: %v", got, want)
	// 	}

	// })

	// t.Run("read data from config file missing server", func(t *testing.T) {

	// 	var config = `{
	// 		"database": {
	// 			"file": "./test.db"
	// 		},
	// 		"token": {
	// 			"secret": "secret",
	// 			"timeout": 200
	// 		},
	// 		"mailSender": {
	// 			"email" : "g.gmail@com",
	// 			"password": "password",
	// 			"timeout": 10
	// 		}

	// 	}
	// 	`

	// 	configPath := t.TempDir() + "/config.json"
	// 	err := os.WriteFile(configPath, []byte(config), 0644)

	// 	if err != nil {
	// 		t.Errorf("Couldn't write config file: %v", err)
	// 	}

	// 	data, err := ReadConfigFile(configPath)
	// 	if err != nil {
	// 		t.Errorf("Couldn't read config file: %v", err)
	// 	}

	// 	got, err := ParseConfigFile(data)
	// 	want := Configuration{
	// 		Server: Server{
	// 			Host: "localhost",
	// 			Port: ":3000",
	// 		},
	// 		Database: DB{
	// 			Path: "./test.db",
	// 		},
	// 		Token: JwtToken{
	// 			Secret:  "secret",
	// 			Timeout: 200,
	// 		},
	// 		MailSender: MailSender{
	// 			Email:    "g.@com",
	// 			Password: "password",
	// 			Timeout:  10,
	// 		},
	// 	}

	// 	if err == nil {
	// 		t.Errorf("Should not have passed : %v", err)
	// 	}
	// 	if got == want {
	// 		t.Errorf("Got: %v. Expected: %v", got, want)
	// 	}

	// })

}
