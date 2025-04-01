package di

import "github.com/gin-gonic/gin"

type ModuleInjector interface {
	Inject(routerGroup *gin.RouterGroup)
}
