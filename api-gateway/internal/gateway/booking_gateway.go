package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"api/config"
	"api/internal/server/client"
)

var execution *http.Client = config.Http2Config()

func GetBooking[T any](endpoint string) (T, error) {
	var result T
	resp, err := execution.Get(client.Booking + endpoint)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyByte, _ := io.ReadAll(resp.Body)
		return result, errors.New(string(bodyByte))
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("failed decode data: %w", err)
	}

	return result, nil
}

func PostBooking[Req any, Res any](endpoint string, body Req) (res Res, err error) {
	data, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, client.Booking+endpoint, bytes.NewReader(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := execution.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		err = errors.New(string(b))
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, fmt.Errorf("failed decode data: %w", err)
	}
	return
}

func DeleteBooking(endpoint string) error {
	url := client.Booking + endpoint
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := execution.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		bodyByte, _ := io.ReadAll(resp.Body)
		return errors.New(string(bodyByte))
	}

	return nil
}
