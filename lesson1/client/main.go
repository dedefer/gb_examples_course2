package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrBadJson = errors.New("bad json")
	ErrNetwork = errors.New("network error")
)

type ErrHTTPBadCode struct {
	code uint16
}

func (e *ErrHTTPBadCode) Code() uint16 { return e.code }

func (e *ErrHTTPBadCode) Error() string {
	return fmt.Sprintf("bad code: %d", e.code)
}

func postJson(c *http.Client, url string, body string) error {
	var js map[string]interface{}
	if err := json.Unmarshal([]byte(body), &js); err != nil {
		return fmt.Errorf("%w: %s", ErrBadJson, err) // В функцию передали невалидный json
	}

	resp, err := c.Post(url, "application/json", strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("%w: %s", ErrNetwork, err) // Тупит сеть
	}

	if resp.StatusCode != http.StatusOK { // Код не ОК
		return &ErrHTTPBadCode{uint16(resp.StatusCode)}
	}

	return nil
}
