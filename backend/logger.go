package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = bodyBytes
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		fmt.Printf("[REQUEST] %s %s\nHeaders: %v\nBody: %s\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Header,
			string(requestBody),
		)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		fmt.Printf("[RESPONSE] %d\nBody: %s\n",
			c.Writer.Status(),
			blw.body.String(),
		)
	}
}
