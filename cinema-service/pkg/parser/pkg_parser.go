package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ParamsInt(r *http.Request, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	val, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("ID tidak valid: %v", err)
	}
	return val, nil
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
