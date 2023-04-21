package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNosurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf(fmt.Sprintf("this is not a handler %T", v))
	}
}

func TestSessionsLoad(t *testing.T) {
    var mh myHandler
	h := SessionsLoad(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf(fmt.Sprintf("this is not a handler %T", v))
	}
}

func TestWriteToConsole(t *testing.T) {
    var mh myHandler
	h := WriteToConsole(&mh)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf(fmt.Sprintf("this is not a handler %T", v))
	}
}