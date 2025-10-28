package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"errors"

	"api/config"
	"api/internal/server/client"
)

var action *http.Client = config.Http2Config()

func GetCinema[T any](endpoint string, res T) (T, error) {
	var res T

	resp, err := action.Get(client.Cinema + endpoint)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyByte, _ := io.ReadAll(resp.Body)
		return res, errors.New(string(bodyByte))
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, fmt.Errorf("failed decode data: %w", err)
	}

	return res, nil
}

func DeleteCinema[T any](endpoint string, body T) error {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewBuffer(jsonData)

	url := client.Cinema + endpoint
	req, err := http.NewRequest(http.MethodDelete, url, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := action.Do(req)
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

func PostCinema[Req any, Res any](endpoint string, body Req) (res Res, err error) {
	data, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, client.Cinema+endpoint, bytes.NewReader(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := action.Do(req)
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

func PutCinema[Req any, Res any](endpoint string, body Req) (res Res, err error) {
	data, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPut, client.Cinema+endpoint, bytes.NewReader(data))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := action.Do(req)
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
