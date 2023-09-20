/*
* Package auth provides authentication and authorization services.
* using session cookies
 */
package auth

import (
	"os"

	"github.com/gin-contrib/sessions/cookie"
)

var store cookie.Store

func init() {
	secret := os.Getenv("COOKIE_SECRET")
	store = cookie.NewStore([]byte(secret))
}
