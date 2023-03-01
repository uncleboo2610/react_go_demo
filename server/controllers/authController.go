package controllers

import (
	config "first-app/configuration"
	"first-app/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var reqBody struct {
		Email    string
		Password string
	}
	c.Bind(&reqBody)

	//hashedPassword
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if error != nil {
		c.JSON(400, gin.H{
			"error": error,
		})
		return
	}

	auth := models.Auth{
		Email:    reqBody.Email,
		Password: string(hashedPassword),
	}
	config.DB.Create(&auth)
}

func Login(c *gin.Context) {
	var reqBody struct {
		Email    string
		Password string
	}
	c.Bind(&reqBody)

	var auth models.Auth
	config.DB.First(&auth, "email = ?", reqBody.Email)
	if auth.ID == 0 {
		c.JSON(400, gin.H{
			"message": "Invalid Email",
		})
		return
	}
	//check pass
	error := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(reqBody.Password))
	if error != nil {
		c.JSON(400, gin.H{
			"message": "Invalid email or password",
		})
		return
	}
	//get jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": auth.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

	c.JSON(200, gin.H{
		"message": tokenString,
	})
}

func Validate(c *gin.Context) {
	auth, _ := c.Get("auth")

	c.JSON(200, gin.H{
		"message": auth,
	})
}
