package internal

import "testing"

func TestGenerateRandomCode(t *testing.T) {

	got := GenerateRandomCode()
	if got < 1000 || got > 9999 {
		t.Errorf("Generated code should be between 1000 and 9999, but got : %v", got)
	}

}
