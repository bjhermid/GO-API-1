package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("1234")
	if err != nil {
		t.Errorf("failed to hashing password, err :%v", err)
	}
	if hash == "" {
		t.Errorf("hash should be not empty")
	}
	if hash == "1234" {
		t.Errorf("password is not hashed")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("1234")
	if err != nil {
		t.Errorf("failed to hashing password, err :%v", err)
	}

	if !ComparePassword(hash, []byte("1234")) {
		t.Errorf("password should be matching to hash")
	}
	if ComparePassword(hash, []byte("wrongpassword")) {
		t.Errorf("Password should be not matching to hash")
	}
}
