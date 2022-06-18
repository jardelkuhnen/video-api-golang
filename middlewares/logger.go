package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const logPath = "logs"
const fileName = "gin.log"
const mode = 0700

func SetupLogOutput() {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		_ = os.Mkdir(logPath, mode)
	}

	f, _ := os.Create(logPath + "/" + fileName)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Logger() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
