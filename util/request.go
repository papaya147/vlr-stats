package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func Request(url string, out any) error {
	resp, err := http.Get(url)
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
