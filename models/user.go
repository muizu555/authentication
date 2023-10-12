package models

import (
	"auth-jwt/utils/token"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func (u User) Save() (User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}

// 自動的に呼び出される
func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Username = strings.ToLower(u.Username)

	return nil
}

func (u User) PrepareOutput() User {
	fmt.Println(u)
	u.Password = ""
	return u
}

func GenerateToken(username string, password string) (string, error) {
	var user User

	err := DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func extractTokenString(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValid(c *gin.Context) error {
	tokenString := extractTokenString(c)

	_, err := parseToken(tokenString)

	if err != nil {
		return err
	}

	return nil
}

func ExtractTokenId(c *gin.Context) (uint, error) {
	tokenString := extractTokenString(c)

	token, err := parseToken(tokenString)

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userId, ok := claims["user_id"].(float64)

		if !ok {
			return 0, nil
		}

		return uint(userId), nil
	}

	return 0, nil
}
