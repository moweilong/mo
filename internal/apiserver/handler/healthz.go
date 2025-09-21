package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/moweilong/mo/pkg/core"

	apiv1 "github.com/moweilong/mo/pkg/api/gen/go/v1"
	"github.com/moweilong/mo/pkg/mlog"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(c *gin.Context) {
	mlog.Infow("Healthz handler is called", "method", "Healthz", "status", "healthy")
	core.WriteResponse(c, apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil)
}
