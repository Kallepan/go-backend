package middleware

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/kallepan/go-backend/test"
)

type compareOriginTest struct {
	origin         string
	allowedOrigins string
	expected       string
}

var compareOriginTests = []compareOriginTest{
	{"http://localhost:3000", "http://localhost:3000", "http://localhost:3000"},
	{"http://localhost:3000", "http://localhost:3000,http://localhost:3001", "http://localhost:3000"},
	{"http://localhost:3000", "http://localhost:3001", ""},
	{"http://localhost:3000", "", "*"},
	{"", "http://localhost:3000", ""},
	{"", "", "*"},
}

func TestCompareOrigin(t *testing.T) {
	for _, testStep := range compareOriginTests {
		// prepare the necessary variables
		w := httptest.NewRecorder()
		ctx := test.GetGinTestCtx(w)

		// set the origin
		ctx.Request.Header.Set("Origin", testStep.origin)

		// set the allowed origins
		os.Setenv("ALLOWED_ORIGINS", testStep.allowedOrigins)

		// compare the origin
		origin := compareOrigin(ctx)
		if origin != testStep.expected {
			t.Errorf("Origin should be %s, got %s", testStep.expected, origin)
		}
	}
}
