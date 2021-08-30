package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {

	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Error(fmt.Sprintf("type is no http.handler but is %T", v))
	}
}
