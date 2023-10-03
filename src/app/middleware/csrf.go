/*
Package middleware implements the middleware functions for the application.
*/
package middleware

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kallepan/go-backend/app/internal/uniuri"
)

const (
	secretKey = "csrfSecret"
	saltKey   = "csrfSalt"
	tokenKey  = "csrfToken"
)

// Methods to ignore
var ignoreMethods = []string{"OPTIONS", "GET", "HEAD"}

func tokenGetter(ctx *gin.Context) string {
	r := ctx.Request

	if t := r.FormValue("_csrf"); len(t) > 0 {
		return t
	} else if t := r.URL.Query().Get("_csrf"); len(t) > 0 {
		return t
	} else if t := r.Header.Get("X-CSRF-TOKEN"); len(t) > 0 {
		return t
	} else if t := r.Header.Get("X-XSRF-TOKEN"); len(t) > 0 {
		return t
	}

	return ""
}

func errorFunc(ctx *gin.Context) {
	/* Error function for the middleware */
	ctx.AbortWithStatus(403)
}

func inArray(arr []string, value string) bool {
	inarr := false

	for _, v := range arr {
		if v == value {
			inarr = true
			break
		}
	}

	return inarr
}

func generateToken(secret, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt+"-"+secret)
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return hash
}

func getToken(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	secret := ctx.MustGet(secretKey).(string)

	if t, ok := ctx.Get(tokenKey); ok {
		return t.(string)
	}

	salt, ok := session.Get(saltKey).(string)
	if !ok {
		salt = uniuri.New()
		session.Set(saltKey, salt)
		session.Save()
	}

	token := generateToken(secret, salt)
	ctx.Set(tokenKey, token)

	return token
}

func CSRF() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secret := os.Getenv("CSRF_SECRET")

		session := sessions.Default(ctx)
		ctx.Set(secretKey, secret)

		if inArray(ignoreMethods, ctx.Request.Method) {
			ctx.Next()
			return
		}

		salt, ok := session.Get(saltKey).(string)

		if !ok || len(salt) == 0 {
			errorFunc(ctx)
			return
		}

		token := tokenGetter(ctx)

		if getToken(ctx) != token {
			errorFunc(ctx)
			return
		}

		ctx.Next()
	}
}
