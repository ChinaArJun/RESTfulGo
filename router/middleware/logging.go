package middleware

import (
	"RESTfulGo/pkg/response"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"regexp"
	"time"

	"RESTfulGo/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("(/v1/user|/login)")
		if !reg.MatchString(path) {
			return
		}

		// Skip for the health check requests.
		if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
			return
		}

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		// 客户端IP
		ip := c.ClientIP()
		// 客户端设备
		phone := c.GetHeader("User-Agent")

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// get code and message
		var result handler.Result
		if err := json.Unmarshal(blw.body.Bytes(), &result); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = response.InternalServerError.Code
			message = err.Error()
		} else {
			code = result.Code
			message = result.Message
		}

		log.Infof("%-13s | %-12s | %s %s %s | {code: %d, message: %s}", latency, ip, phone, pad.Right(method, 5, ""), path, code, message)
	}
}
