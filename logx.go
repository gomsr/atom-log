package logx

import (
	"github.com/gin-gonic/gin"
	"github.com/gomsr/atom-log/rich"
	"go.uber.org/zap"
)

func GetTLogger(c *gin.Context, logger *zap.Logger) rich.Logger {

	//c.Request = c.Request.WithContext(spanCtx)
	//c.Set("tlog", logx.WithLogger(global.LOGGER).WithContext(spanCtx))

	if logger, ok := c.Value("tlog").(rich.Logger); ok {
		return logger
	}

	if logger, ok := c.Request.Context().Value("tlog").(rich.Logger); ok {
		return logger
	}

	return rich.WithLogger(logger)
}
