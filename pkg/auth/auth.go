package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	adminUsername = "admin"
	adminPassword = "admin"
	jwtSecret     = "iu4fcn0qnua"
)

type LoginRequest struct {
	Username string
	Password string
}

func VerifyUser(c *gin.Context) {

	token, err := ParseTokenFromHeader(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	_, err = VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
}

func VerifyToken(token *jwt.Token) (string, error) {
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "", errors.New("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("Invalid token")
		}
		return username, nil
	}
	return "", nil
}

func ParseTokenFromHeader(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	tokenArr := strings.Split(tokenString, " ")
	if len(tokenArr) != 2 {
		err := errors.New("invalid token")
		return nil, err
	}
	token, err := jwt.Parse(tokenArr[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	if loginReq.Username == adminUsername && loginReq.Password == adminPassword {
		token, err := CreateToken(loginReq.Username)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"Token": token})
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Payload"})
}

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["admin"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := at.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
