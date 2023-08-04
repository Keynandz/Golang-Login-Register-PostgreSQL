package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var sessionStore = sessions.NewCookieStore([]byte("secret"))

func SetSession(c echo.Context, key string, value string) {
	session, _ := sessionStore.Get(c.Request(), "session-name")
	session.Values[key] = value
	session.Save(c.Request(), c.Response())
}

func GetSession(c echo.Context, key string) string {
	session, _ := sessionStore.Get(c.Request(), "session-name")
	if val, ok := session.Values[key].(string); ok {
		return val
	}
	return ""
}

func ClearSession(c echo.Context) {
	sess, _ := sessionStore.Get(c.Request(), "session-name")
	delete(sess.Values, "user")
	sess.Save(c.Request(), c.Response())
}