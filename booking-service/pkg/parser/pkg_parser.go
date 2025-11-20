package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ParamsInt(r *http.Request, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	val, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %v", err)
	}
	return val, nil
}

func ParamsUUID(r *http.Request, prefix string) (uuid.UUID, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid uuid: %v", err)
	}
	return id, nil
}

func Body[T any](r io.Reader, body T) (T, error) {
	// setting decoding
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()

	//run decoding
	if err := decoder.Decode(&body); err != nil {
		var zero T
		return zero, err
	}

	return body, nil

}
