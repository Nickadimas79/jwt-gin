package models

import (
	"html"
	"strings"

	"github.com/Nickadimas79/jwt-gin/controllers"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username, password string) (string, error) {
	var err error
	user := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, username)
	if err != nil {
		return "", err
	}

	// TODO:: Look into better naming/package placement here
	t, err := controllers.Generate(user.ID)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (u *User) Save() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) BeforeSave() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPass)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
