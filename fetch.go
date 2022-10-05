package goodle

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

type (
	Fetcher interface {
		Fetch(url string) ([]byte, int, string, error)
	}
	DefaultFetcher struct {
		client *http.Client
	}
)

// Fetch sends a GET request to the given URL and returns the response body, status, contentType and error.
func (f *DefaultFetcher) Fetch(url string) ([]byte, int, string, error) {
	if f.client == nil {
		transport := http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 20 * time.Second,
		}

		// All users of cookiejar should import "golang.org/x/net/publicsuffix"
		jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		if err != nil {
			return nil, 0, "", err
		}

		f.client = &http.Client{
			Timeout:   time.Second * 20,
			Transport: &transport,
			// net/http provides a default cookie jar implementation which implements the CookieJar interface.
			Jar: jar,
		}
	}

	resp, err := f.client.Get(url)
	if err != nil {
		return nil, 0, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, "", err
	}

	// check for non textual media types

	contentType := resp.Header.Get("Content-Type")
	if contentType != "" {
		if isNonTextualContentType(contentType) {
			return nil, 0, "", errors.New("Content-Type is not textual")
		}
	} else {
		return nil, 0, "", errors.New("no content type")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, "", err
	}
	return body, resp.StatusCode, contentType, nil

}
