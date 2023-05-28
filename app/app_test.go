package app

import "testing"

func TestNewApp(t *testing.T) {

	// 	t.Run("create new app with valid configPath", func(t *testing.T) {
	// 		configPath := "../config.json"
	// 		_, err := NewApp(configPath)
	// 		if err != nil {
	// 			t.Errorf("Should've created a new app, but got the following error:  %+v", err)
	// 		}

	// 	})

	// 	t.Run("create new app with invalid configPath", func(t *testing.T) {
	// 		configPath := "/config.json"
	// 		_, err := NewApp(configPath)
	// 		if err == nil {
	// 			t.Errorf("Should've gotten an ivalid path error but got :  %+v", err)
	// 		}

	// 	})

	// }

	// func TestListenAndServe(t *testing.T) {

	// 	configPath := "../config.json"
	// 	a, err := NewApp(configPath)
	// 	if err != nil {
	// 		t.Errorf("Should've created a new app, but got the following error:  %+v", err)
	// 	}

	// err = a.ListenAndServe()
	//
	//	if err != nil {
	//		t.Errorf("Should've started a server, but got the following error:  %+v", err)
	//	}
}
