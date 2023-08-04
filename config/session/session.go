package session

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

// sessionStore menyimpan instance dari penyimpanan sesi dengan cookie.
var sessionStore = sessions.NewCookieStore([]byte("secret"))

// SetSession menetapkan nilai sesi untuk kunci yang diberikan.
func SetSession(c echo.Context, key string, value string) {
	// Dapatkan atau buat sesi baru dengan nama "session-name"
	session, _ := sessionStore.Get(c.Request(), "session-name")
	// Tetapkan nilai untuk kunci di sesi
	session.Values[key] = value
	// Simpan sesi
	session.Save(c.Request(), c.Response())
}

// GetSession mengembalikan nilai sesi untuk kunci yang diberikan.
func GetSession(c echo.Context, key string) string {
	// Dapatkan atau buat sesi baru dengan nama "session-name"
	session, _ := sessionStore.Get(c.Request(), "session-name")
	// Cek apakah nilai sesi dengan kunci yang diberikan adalah string, jika ya, kembalikan nilai tersebut.
	if val, ok := session.Values[key].(string); ok {
		return val
	}
	// Jika tidak ditemukan nilai sesi atau nilainya bukan string, kembalikan string kosong.
	return ""
}

// ClearSession menghapus data sesi yang terkait dengan pengguna.
func ClearSession(c echo.Context) {
	// Dapatkan atau buat sesi baru dengan nama "session-name"
	sess, _ := sessionStore.Get(c.Request(), "session-name")
	// Hapus data sesi untuk kunci "user"
	delete(sess.Values, "user")
	// Simpan sesi
	sess.Save(c.Request(), c.Response())
}