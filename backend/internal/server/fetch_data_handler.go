package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) fetchDataHandler(w http.ResponseWriter, r *http.Request) {
	db := s.DB
	muxVars := mux.Vars(r)
	sdkKey := muxVars["sdk-key"]

	// Query the database for user data
	const getUserData = `
        SELECT 
            sdk_key, request_id, raw_req_id, raw_resp_id, 
            req_in, req_out, endpoint, status_code, latency, method 
        FROM api_submissions 
        WHERE sdk_key = $1
    `

	rows, err := db.Query(getUserData, sdkKey)
	if err != nil {
		log.Println("Database query error:", err)
		s.writeErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}
	defer rows.Close()

	// Prepare the response
	response := make([]map[string]interface{}, 0)

	for rows.Next() {
		var sdkKey, requestID, rawReqID, rawRespID, endpoint, method string
		var reqIn, reqOut time.Time
		var statusCode, latency int

		// Scan the row into variables
		err = rows.Scan(&sdkKey, &requestID, &rawReqID, &rawRespID, &reqIn, &reqOut, &endpoint, &statusCode, &latency, &method)
		if err != nil {
			log.Println("Row scan error:", err)
			s.writeErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
			return
		}

		// Convert row data into a map
		rowMap := map[string]interface{}{
			"sdk_key":     sdkKey,
			"request_id":  requestID,
			"raw_req_id":  rawReqID,
			"raw_resp_id": rawRespID,
			"req_in":      reqIn,
			"req_out":     reqOut,
			"endpoint":    endpoint,
			"status_code": statusCode,
			"latency":     latency,
			"method":      method,
		}
		response = append(response, rowMap)
	}

	// Check for errors after iterating through rows
	if err = rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
		s.writeErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Response encoding error:", err)
		s.writeErrorResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
