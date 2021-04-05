package client

import (
	"errors"
	"net/http"
	"testing"
)

func TestBadJson(t *testing.T) {
	err := postJson(http.DefaultClient, "asdasd", "{")
	if !errors.Is(err, ErrBadJson) {
		t.Errorf("expected bad json, got %#v", err)
	}
}

func TestNetworkErr(t *testing.T) {
	err := postJson(http.DefaultClient, "invalid_url", "{}")
	if !errors.Is(err, ErrNetwork) {
		t.Errorf("expected network err, got %#v", err)
	}
}

func TestBadCode(t *testing.T) {
	err := postJson(http.DefaultClient, "http://httpstat.us/500", "{}")
	httpBadCode, ok := err.(*ErrHTTPBadCode)
	if !ok {
		t.Errorf("expected ErrHTTPBadCode, got %#v", err)
		return
	}
	if httpBadCode.Code() != 500 {
		t.Errorf("expected ErrHTTPBadCode(500), got ErrHTTPBadCode(%d)", httpBadCode.Code())
	}
}

func TestOk(t *testing.T) {
	err := postJson(http.DefaultClient, "http://httpstat.us/200", "{}")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
