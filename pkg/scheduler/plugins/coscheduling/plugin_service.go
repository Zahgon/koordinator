package coscheduling

import (
	"github.com/gin-gonic/gin"

	"github.com/koordinator-sh/koordinator/pkg/scheduler/frameworkext/services"
)

var _ services.APIServiceProvider = &Coscheduling{}

func (cs *Coscheduling) RegisterEndpoints(group *gin.RouterGroup) {
	_ = "STUB: not implemented"
	return
}
