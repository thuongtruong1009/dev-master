package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleRequest(ctx *gin.Context) {
	path := ctx.Param("path")

	if path == "/" {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "CORS Proxy. Just go to /:url to use",
		})
		ctx.Done()
		return
	}

	if !strings.HasPrefix(path, "/http") {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid URL",
		})
		ctx.Done()
		return
	}

	requestedURL := path[1:]

	req, _ := http.NewRequest(ctx.Request.Method, requestedURL, ctx.Request.Body)

	req.Header = ctx.Request.Header.Clone()
	req.Header.Del("origin")
	req.Header.Del("referer")

	queryValues := req.URL.Query()
	for k, v := range ctx.Request.URL.Query() {
		queryValues.Add(k, v[0])
	}
	req.URL.RawQuery = queryValues.Encode()

	response, err1 := http.DefaultClient.Do(req)

	for k, v := range response.Header.Clone() {
		ctx.Header(k, v[0])
	}

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")

	responseBytes, err2 := ioutil.ReadAll(response.Body)

	if err1 != nil || err2 != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to request",
		})
		ctx.Done()
		return
	}

	ctx.Data(response.StatusCode, response.Header.Get("Content-Type"), responseBytes)
}

func main() {
	router := gin.Default()

	router.GET("*path", handleRequest)
	router.POST("*path", handleRequest)
	router.PUT("*path", handleRequest)
	router.PATCH("*path", handleRequest)
	router.DELETE("*path", handleRequest)
	router.OPTIONS("*path", handleRequest)
	router.HEAD("*path", handleRequest)

	router.Run(":5000")
}
