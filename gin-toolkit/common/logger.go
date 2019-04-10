package common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"math"
	"os"

	"strings"
	"time"
)

func InitLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.SetOutput(os.Stdout)
	if GetRunEnv() == EnvDev {
		logrus.SetLevel(logrus.DebugLevel)
	} else if GetRunEnv() == EnvTest {
		logrus.SetLevel(logrus.DebugLevel)
	} else if GetRunEnv() == EnvProduct {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func Logger(c *gin.Context) *logrus.Entry {
	path := c.Request.URL.RequestURI()
	requestId, _ := c.Get("requestId")

	entry := logrus.NewEntry(logrus.StandardLogger()).WithFields(logrus.Fields{
		"requestId": requestId,
		"path":      path,
		"method":    c.Request.Method,
	})

	return entry
}

func StatisLogger(value ...interface{}) {
	entry := logrus.NewEntry(logrus.StandardLogger()).WithFields(logrus.Fields{
		"module": "statistics",
	})
	entry.Info(value...)
}

func DebugLogger(c *gin.Context, value ...interface{}) {
	if c != nil {
		Logger(c).Debug(value...)
	} else {
		logrus.Debug(value...)
	}
}

func Debug4Logger(c *gin.Context, format string, value ...interface{}) {
	if c != nil {
		Logger(c).Debugf(format, value...)
	} else {
		logrus.Debugf(format, value...)
	}
}

func InfoLogger(c *gin.Context, value ...interface{}) {
	if c != nil {
		Logger(c).Info(value...)
	} else {
		logrus.Info(value...)
	}
}

func Info4Logger(c *gin.Context, format string, value ...interface{}) {
	if c != nil {
		Logger(c).Infof(format, value...)
	} else {
		logrus.Infof(format, value...)
	}
}

func WarnLogger(c *gin.Context, value ...interface{}) {
	if c != nil {
		Logger(c).Warn(value...)
	} else {
		logrus.Warn(value...)
	}
}

func Warn4Logger(c *gin.Context, format string, value ...interface{}) {
	if c != nil {
		Logger(c).Warnf(format, value...)
	} else {
		logrus.Warnf(format, value...)
	}
}

func ErrorLogger(c *gin.Context, value ...interface{}) {
	if c != nil {
		Logger(c).Error(value...)
	} else {
		logrus.Error(value...)
	}
}

func Error4Logger(c *gin.Context, format string, value ...interface{}) {
	if c != nil {
		Logger(c).Errorf(format, value...)
	} else {
		logrus.Errorf(format, value...)
	}
}

func PanicLogger(c *gin.Context, value ...interface{}) {
	if c != nil {
		Logger(c).Panic(value...)
	} else {
		logrus.Panic(value...)
	}
}

func Panic4Logger(c *gin.Context, format string, value ...interface{}) {
	if c != nil {
		Logger(c).Panicf(format, value...)
	} else {
		logrus.Panicf(format, value...)
	}
}

func StackLogger(c *gin.Context, err error) {
	if !strings.Contains(fmt.Sprintf("%+v", err), "\n") {
		return
	}

	var info []byte
	if c != nil {
		requestId, _ := c.Get("requestId")
		path := c.Request.URL.RequestURI()
		info, _ = json.Marshal(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": "error", "module": "errorStack", "requestId": requestId, "path": path})
	} else {
		info, _ = json.Marshal(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": "error", "module": "errorStack"})
	}

	log.Printf("%s\n-------------------stack-start-------------------\n%+v\n-------------------stack-end-------------------\n", string(info), err)
}

func LoggerTemplate() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.RequestURI()
		start := time.Now()
		requestId := c.Request.Header.Get("X-RequestId")
		if requestId == "" {
			requestId = GetUuidString()
		}
		c.Set("requestId", requestId)

		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		httpCode := c.Writer.Status()
		clientIp := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()

		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		_ = c.Request.ParseForm()
		requestBody := c.Request.PostForm.Encode()

		render, exist := c.Get("render")
		var retCode int
		if exist == true {
			retCode = render.(TRenderJson).ReturnCode
		} else {
			retCode = 0
		}

		response, _ := json.Marshal(render)

		logrus.NewEntry(logrus.StandardLogger()).WithFields(logrus.Fields{
			"requestId":   requestId,
			"requestBody": requestBody,
			"requestPath": path,
			"httpCode":    httpCode,
			"retCode":     retCode,
			"latency":     latency, // time to process
			"clientIp":    clientIp,
			"method":      c.Request.Method,
			"referer":     referer,
			"dataLength":  dataLength,
			"userAgent":   clientUserAgent,
		}).Info(string(response))
	}
}
