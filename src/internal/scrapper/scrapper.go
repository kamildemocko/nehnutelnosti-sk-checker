package scrapper

import (
	"fmt"
	"io"
	"net/http"
)

type ErrStatusCodeNk struct {
	StatusCode int
}

func (e ErrStatusCodeNk) Error() string {
	return fmt.Sprintf("status code is not ok: %d", e.StatusCode)
}

// scraps web page and returns body as text
func ScrapWebPage(uri string) (string, error) {
	res, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", ErrStatusCodeNk{StatusCode: res.StatusCode}
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
