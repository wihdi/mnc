package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/wihdi/mnc/usecase"
	"github.com/wihdi/mnc/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var user domain.LoginUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInDb, err := h.userUsecase.FindByUsername(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	if userInDb == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}
	
	if userInDb.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}
	

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

