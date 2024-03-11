package util

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func Request(ctx context.Context, url string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("got response from", url)

	// buf, err := io.ReadAll(resp.Body)
	// log.Println(string(buf))

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
