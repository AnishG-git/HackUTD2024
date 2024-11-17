package pineapplesdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/AnishG-git/HackUTD2024/backend/storage"
	"github.com/google/uuid"
)

// Middleware to emulate a round trip for requests and responses
func CreateMW(sdk_key string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		mwLog := log.Default()
		db, err := storage.LoadDB("hackutd2024", mwLog)
		if err != nil {
			mwLog.Fatalf("Failed to load database: %v", err)
		}

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Start: Before handling the request (Pre-processing)
			requestReceivedAt := time.Now()
			url := r.URL.String()
			method := r.Method
			headers := headersToString(r.Header)
			log.Println("Incoming Request:")
			log.Printf("Method: %s, URL: %s, Headers: %v\n", r.Method, r.URL, headers)

			// 2. Custom Response Writer to intercept outgoing response
			rawReq, err := httputil.DumpRequest(r, true)
			if err != nil {
				log.Println("Error dumping request:", err)
			}
			rw := &responseWriter{ResponseWriter: w}

			go func() {
				<-r.Context().Done()
				responseSentAt := time.Now()
				latency := int(responseSentAt.Sub(requestReceivedAt).Milliseconds())
				log.Printf("Status Code: %d, Duration: %v, Response Body: %s\n", rw.statusCode, latency, string(rw.body))

				rawResp := fmt.Sprintf("Headers:\n%s\nResponse body:\n%s\n", headersToString(rw.Header()), string(rw.body))
				pinataReqId, err := storeToPinata(string(rawReq), "request")
				if err != nil {
					log.Println("Error storing request to Pinata:", err)
				}
				pinataRespId, err := storeToPinata(rawResp, "response")
				if err != nil {
					log.Println("Error storing response to Pinata:", err)
				}

				fmt.Println("sdk key:", sdk_key)

				const insertQuery = "INSERT INTO api_submissions (sdk_key, request_id, raw_req_id, raw_resp_id, req_in, req_out, endpoint, status_code, method, latency) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
				_, err = db.Exec(insertQuery, sdk_key, uuid.New().String(), pinataReqId, pinataRespId, requestReceivedAt, responseSentAt, url, rw.statusCode, method, latency)
				if err != nil {
					log.Println("Error inserting into database:", err)
				}

				fmt.Println("Pinata request ID:", pinataReqId)
				fmt.Println("Pinata response ID:", pinataRespId)
			}()

			next.ServeHTTP(rw, r)
		})
	}
}

type pinataResponse struct {
	Data struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		CID           string `json:"cid"`
		CreatedAt     string `json:"created_at"`
		Size          int    `json:"size"`
		NumberOfFiles int    `json:"number_of_files"`
		MimeType      string `json:"mime_type"`
		UserID        string `json:"user_id"`
		GroupID       string `json:"group_id"`
		IsDuplicate   bool   `json:"is_duplicate"`
	} `json:"data"`
}

func storeToPinata(txtContent string, fileType string) (string, error) {
	const uploadUrl = "https://uploads.pinata.cloud/v3/files"

	// Create a buffer to hold the multipart data
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Create a file field for the txt file using CreateFormFile
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.txt", fileType))
	if err != nil {
		return "", fmt.Errorf("error creating form file: %w", err)
	}

	// Write the content of the txt file into the form field
	_, err = part.Write([]byte(txtContent))
	if err != nil {
		return "", fmt.Errorf("error writing to part: %w", err)
	}

	// Close the writer to finalize the multipart body
	writer.Close()

	// Create a new POST request
	req, err := http.NewRequest("POST", uploadUrl, &b)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set the Authorization header with the JWT token
	req.Header.Add("Authorization", "Bearer "+os.Getenv("PINATA_JWT"))

	// Set the Content-Type header with the boundary
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// Make the HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Print the response body for debugging
	fmt.Println(string(bodyBytes))

	// Parse the JSON response
	var respData pinataResponse
	err = json.Unmarshal(bodyBytes, &respData)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}

	// Return the ID from the response
	return respData.Data.CID, nil
}

func headersToString(headers http.Header) string {
	var builder strings.Builder

	for key, values := range headers {
		// Join all values for a header key with a comma
		builder.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", ")))
	}

	return builder.String()
}

// Custom ResponseWriter to capture response details
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// Override WriteHeader to capture status code
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

// Override Write to capture response body
func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body = append(rw.body, data...)
	return rw.ResponseWriter.Write(data)
}
