package NintendoSwitchEshop

import (
	"errors"
	"github.com/imroc/req"
	"net/http"
)

func sendNewRequest(u string) ([]byte, error) {
	r, err := req.Get(u, req.Header{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36",
	})

	if err != nil {
		return nil, err
	}

	if r.Response().StatusCode == http.StatusForbidden {
		return nil, errors.New("price api rate limit")
	}

	return r.ToBytes()
}
