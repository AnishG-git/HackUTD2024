package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AnishG-git/HackUTD2024/backend/internal/util"
)

func JsonMiddleware[ReqModel any](next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var reqBodyObj ReqModel
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		reader := bytes.NewReader(bodyBytes)
		err = json.NewDecoder(reader).Decode(&reqBodyObj)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
			return
		}

		missingFields, missingFieldsJsonErrMsg := util.ValidateFields(reqBodyObj)
		if len(missingFields) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, missingFieldsJsonErrMsg)))
			return
		}

		// Pass the decoded object to the next handler (this can be done by setting it on the context)
		ctx := context.WithValue(r.Context(), "reqBodyObj", &reqBodyObj)
		next(w, r.WithContext(ctx))
	}
}
