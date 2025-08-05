package router

import (
	authservice "github.com/abdullahshafaqat/Chatify/api/auth_service"
	messageservice "github.com/abdullahshafaqat/Chatify/api/message_service"
	"github.com/gin-gonic/gin"
)

type Router interface {
	DefineRoutes(r *gin.Engine)
}

type routerImpl struct {
	authService    authservice.Service
	messageService messageservice.Service
}

func NewRouter(authService authservice.Service, messageService messageservice.Service) Router {
	return &routerImpl{
		authService:    authService,
		messageService: messageService,
	}
}
