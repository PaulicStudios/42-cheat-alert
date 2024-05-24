package requests

import (
	"encoding/json"
	"io"
)

func request(url string, target any) error {
	req, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	body, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		return readErr
	}
	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}
	return nil
}
