package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Get a gin context for testing
func GetGinTestCtx(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

// Mock GET request with JSON
func GET(ctx *gin.Context, params gin.Params, u url.Values) {
	/*
		Call like this:
		GET(ctx, gin.Params{
			gin.Param{Key: "id", Value: "1"},
		}, url.Values{
			"page": []string{"1"},
		})

		Params and query params are optional
	*/
	ctx.Request.Method = "GET"
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Set the params
	ctx.Params = params

	// Set the query params
	ctx.Request.URL.RawQuery = u.Encode()
}

// Mock POST request with JSON
func POST(ctx *gin.Context, params gin.Params, u url.Values, content interface{}) {
	/*
		Post method with JSON
	*/
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Set the params
	ctx.Params = params

	// Set the query params
	ctx.Request.URL.RawQuery = u.Encode()

	// json
	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	ctx.Request.Body = io.NopCloser(bytes.NewReader(jsonbytes))
}

func DELETE(ctx *gin.Context, params gin.Params, u url.Values) {
	/*
		Delete method
	*/
	ctx.Request.Method = "DELETE"
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Set the params
	ctx.Params = params

	// Set the query params
	ctx.Request.URL.RawQuery = u.Encode()
}
