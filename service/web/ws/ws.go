package ws

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func Respond(w http.ResponseWriter, code int, data interface{}) {
	switch v := data.(type) {
	case error:
		if v == io.EOF {
			Respond(w, http.StatusBadRequest, errors.New("Received empty body"))
			return
		}
		data = errorResponse{Error: v.Error()}
	}

	var buf []byte
	var err error

	if data == nil {
		if code >= 200 && code < 300 {
			data = statusResponse{http.StatusText(code)}
		} else {
			data = errorResponse{http.StatusText(code)}
		}
	}

	buf, err = json.Marshal(data)
	if err != nil {
		log.Printf("json.Marshal", err)
		Respond(w, 500, errors.New(http.StatusText(http.StatusInternalServerError)))
		return
	}

	w.WriteHeader(code)
	w.Write(buf)
}

func Bind(r *http.Request, dest interface{}) error {
	j := json.NewDecoder(r.Body)
	err := j.Decode(dest)
	if err != nil {
		return errors.New("Could not decode JSON")
	}
	return nil
}
