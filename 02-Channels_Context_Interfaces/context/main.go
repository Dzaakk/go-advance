package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second*2)
		defer cancel()

		req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://www.google.com", nil)
		if err != nil {
			panic(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		ctx.Data(200, "text/html", data)

	})
}
