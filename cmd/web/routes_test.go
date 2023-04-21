package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/bertoxic/bert/internal/config"
)

func Test_Routes(t *testing.T) {

var app config.AppConfig
mux := routes(&app)

switch v := mux.(type) {
case http.Handler:
    default : t.Errorf(fmt.Sprintf("Routes test failed not a type http,Handler %T",v))
}

}