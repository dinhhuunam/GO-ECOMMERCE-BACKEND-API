package controller

import (
	"net/http"

	"example.com/ecommerce/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service
// controlle-->service-->repo-->models-->dbs
func (uc *UserController) GetUserById(c *gin.Context) {
	// name := c.Param("name")
	// name := c.DefaultQuery("name", "anonystick")
	// uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
		// "uid":     uid,
		"users": []string{"cr7", "m10", "nam"},
	})
}
