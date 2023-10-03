package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newServer() *gin.Engine {
	g := gin.New()

	store := cookie.NewStore([]byte("secret123"))

	g.Use(sessions.Sessions("my_session", store))
	g.Use(CSRF())

	return g
}

type requestOptions struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    io.Reader
}

func request(server *gin.Engine, options requestOptions) *httptest.ResponseRecorder {
	if options.Method == "" {
		options.Method = "GET"
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(options.Method, options.URL, options.Body)

	if options.Headers != nil {
		for key, value := range options.Headers {
			req.Header.Set(key, value)
		}
	}

	server.ServeHTTP(w, req)

	if err != nil {
		panic(err)
	}

	return w
}

func TestForm(t *testing.T) {
	var token string
	g := newServer()

	g.GET("/login", func(c *gin.Context) {
		token = getToken(c)
	})

	g.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r1 := request(g, requestOptions{URL: "/login"})
	r2 := request(g, requestOptions{
		Method: "POST",
		URL:    "/login",
		Headers: map[string]string{
			"Cookie":       r1.Header().Get("Set-Cookie"),
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: strings.NewReader("_csrf=" + token),
	})

	if body := r2.Body.String(); body != "OK" {
		t.Error("Response is not OK: ", body)
	}
}

func TestQueryString(t *testing.T) {
	var token string
	os.Setenv("CSRF_SECRET", "secret123")
	g := newServer()

	g.GET("/login", func(c *gin.Context) {
		token = getToken(c)
	})

	g.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r1 := request(g, requestOptions{URL: "/login"})
	r2 := request(g, requestOptions{
		Method: "POST",
		URL:    "/login?_csrf=" + token,
		Headers: map[string]string{
			"Cookie": r1.Header().Get("Set-Cookie"),
		},
	})

	if body := r2.Body.String(); body != "OK" {
		t.Error("Response is not OK: ", body)
	}
}

func TestQueryHeader1(t *testing.T) {
	var token string
	os.Setenv("CSRF_SECRET", "secret123")
	g := newServer()

	g.GET("/login", func(c *gin.Context) {
		token = getToken(c)
	})

	g.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r1 := request(g, requestOptions{URL: "/login"})
	r2 := request(g, requestOptions{
		Method: "POST",
		URL:    "/login",
		Headers: map[string]string{
			"Cookie":       r1.Header().Get("Set-Cookie"),
			"X-CSRF-Token": token,
		},
	})

	if body := r2.Body.String(); body != "OK" {
		t.Error("Response is not OK: ", body)
	}
}

func TestQueryHeader2(t *testing.T) {
	var token string
	os.Setenv("CSRF_SECRET", "secret123")
	g := newServer()

	g.GET("/login", func(c *gin.Context) {
		token = getToken(c)
	})

	g.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r1 := request(g, requestOptions{URL: "/login"})
	r2 := request(g, requestOptions{
		Method: "POST",
		URL:    "/login",
		Headers: map[string]string{
			"Cookie":       r1.Header().Get("Set-Cookie"),
			"X-XSRF-Token": token,
		},
	})

	if body := r2.Body.String(); body != "OK" {
		t.Error("Response is not OK: ", body)
	}
}
