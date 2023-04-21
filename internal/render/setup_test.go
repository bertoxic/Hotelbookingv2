package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bertoxic/bert/internal/config"
	"github.com/bertoxic/bert/models"
)

var sessions *scs.SessionManager
var testApp config.AppConfig


func TestMain(m *testing.M){
    gob.Register(models.Reservation{})
	testApp.InProduction = false

	sessions = scs.New()
	sessions.Cookie.Persist = true
	sessions.Lifetime = 24 * time.Hour
	sessions.Cookie.Secure = false
	sessions.Cookie.SameSite = http.SameSiteLaxMode
	testApp.Sessions = sessions

    app = &testApp

    os.Exit(m.Run())
}

type myWriter struct {}

func (mw *myWriter) Header() http.Header{

	var h http.Header
	return h
}

func (mw *myWriter) Write(b []byte) (int, error){
	v:=len(b)
	return v,nil
}

func (mw *myWriter)WriteHeader(statuscode int ){

}