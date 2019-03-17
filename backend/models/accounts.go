package models

import (
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type (
	// Token struct
	Token struct {
		UserID unit
		jwt.StandardClaims
	}
	// Account struct
	Account struct {
		gorm.Model
		Email    string `json:"email"`
		Password string `json:"password"`
		Token    string `json:"token";sql:"-"`
	}
)

// Validate invoming user details
func (account *Account) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}
	if len(account.Password) < 6 {
		return u.Message(false, "Password length must be longer than 6 characters"), false
	}
	// Email must be unique
	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retyr"), false
	}
	if err != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}
	return u.Message(false, "Requirement passed"), true
}

// Craete returns a jwt token when successful
func (account *Account) Create() map[string]interface{} {
	if response, ok := account.Validate(); !ok {
		return response
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	GetDB().Create(account)
	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

// Login returns a map when login is successful
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	account.Password = ""
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	response := u.Message(true, "Logged In")
	response["account"] = account
	return response
}

// GetUser returns a pointer to an Account
func GetUser(u unit) *Account {
	account := &Account{}
	GetDB().Table("account").Where("id = ?", u).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}
