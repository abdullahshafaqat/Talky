package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *routerImpl) GetUsersHandler(c *gin.Context) {
	users, err := r.authService.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
