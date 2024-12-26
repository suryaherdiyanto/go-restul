package token

import (
	"testing"
	"time"

	"github.com/go-restful/app/model"
)

func TestCreateToken(t *testing.T) {
	secret := "thetokensecret"
	user := &model.User{
		Id:        1,
		Email:     "jogndoe@gmail.com",
		FirstName: "John",
	}

	_, err := GenerateToken(user, secret, time.Minute*15)
	if err != nil {
		t.Errorf("Generate Token Faield: %v", err)
	}

}

func TestValidateToken(t *testing.T) {
	secret := "thetokensecret"
	user := &model.User{
		Id:        1,
		Email:     "jogndoe@gmail.com",
		FirstName: "John",
	}

	token, err := GenerateToken(user, secret, time.Minute*15)
	if err != nil {
		t.Errorf("Generate Token Faield: %v", err)
	}

	claims, err := ValidateToken(token, secret)

	if err != nil {
		t.Errorf("Validate Token Faield: %v", err)
	}

	if claims.Email != user.Email {
		t.Errorf("Email is not same, expected: %s, got: %s", user.Email, claims.Email)
	}

}

func TestInvalidToken(t *testing.T) {
	secret := "thetokensecret"

	user := &model.User{
		Id:        1,
		Email:     "jogndoe@gmail.com",
		FirstName: "John",
	}

	token, err := GenerateToken(user, "wrongsecret", time.Minute*15)
	if err != nil {
		t.Errorf("Generate Token Faield: %v", err)
	}

	_, err = ValidateToken(token, secret)

	if err == nil {
		t.Errorf("Invalid Token should fail")
	}

}
