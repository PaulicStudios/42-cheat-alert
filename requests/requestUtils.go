package requests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func request(url string, target any) (*http.Response, error) {
	req, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	if req.StatusCode != 200 {
		retryAfter, err := strconv.Atoi(req.Header.Get("Retry-After"))
		if err != nil {
			return nil, err
		}
		// log.Println("Rate limited, retrying in", retryAfter, "seconds")
		time.Sleep(time.Duration(retryAfter) * time.Second)
		return request(url, target)
	}

	body, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		return nil, readErr
	}
	err = json.Unmarshal(body, target)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func hasNextPage(responce *http.Response) bool {
	links := parseLinkHeader(responce.Header.Get("Link"))
	_, ok := links["next"]
	return ok
}

func parseLinkHeader(header string) map[string]string {
	links := make(map[string]string)

	linkHeaders := strings.Split(header, ", ")
	for _, linkHeader := range linkHeaders {
		parts := strings.Split(linkHeader, "; ")

		link, err := url.Parse(strings.Trim(parts[0], "<>"))
		if err != nil {
			continue
		}

		if len(parts) < 2 {
			continue
		}
		rel := strings.Trim(parts[1], `rel="`)
		links[rel] = link.String()
	}

	return links
}
