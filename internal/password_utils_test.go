package internal

import "testing"

func TestHashPassword(t *testing.T) {

	t.Run("create password hash", func(t *testing.T) {
		password := "password"
		_, err := HashPassword(password)
		if err != nil {
			t.Errorf("Should've gotten a hashed password, but got %+v", err)
		}

	})

	t.Run("verify correct hashed password", func(t *testing.T) {
		password := "password"
		_, err := HashPassword(password)
		if err != nil {
			t.Errorf("Should've gotten a hashed password, but got %+v", err)
		}
		match := VerifyPassword(password, "password")

		if match == nil {
			t.Errorf("password isn't correct")

		}

	})

	t.Run("verify incorrect hashed password", func(t *testing.T) {
		password := "password"
		_, err := HashPassword(password)
		if err != nil {
			t.Errorf("Should've gotten a hashed password, but got %+v", err)
		}
		match := VerifyPassword(password, "password33")

		if match == nil {
			t.Errorf("Should'nt have passed, password is incorrect")

		}

	})

}
