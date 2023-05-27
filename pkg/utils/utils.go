package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody this function is used to parse the request body and map it to the object passed as the second parameter
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil { //
			return
		}
	}
}
