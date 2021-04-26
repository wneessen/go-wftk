package kratos

import (
	"bufio"
	"bytes"
	"fmt"
	go_wftk "github.com/wneessen/go-wftk"
	"log"
	"net/http"
	"strings"
)

// Perform a GET request
func HttpReqGet(u string, hc *http.Client, st string) ([]byte, error) {
	// Create a HTTP request
	httpReq, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("an error occured creating new HTTP GET request: %v", err)
	}

	// Set HTTP header
	setReqHeader(httpReq)

	// We have a session token
	if st != "" {
		httpReq.Header.Add("X-Session-Token", st)
	}

	serverResp, err := hc.Do(httpReq)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		err := serverResp.Body.Close()
		if err != nil {
			log.Printf("error while closing response body: %v", err)
		}
	}()
	if !strings.HasPrefix(serverResp.Header.Get("Content-Type"), "application/json") {
		return []byte{}, fmt.Errorf("kratos returned non-JSON response")
	}

	// Read the response body
	var respBody []byte
	buf := bufio.NewScanner(serverResp.Body)
	for buf.Scan() {
		respBody = buf.Bytes()
	}
	if err = buf.Err(); err != nil {
		return []byte{}, err
	}

	return respBody, nil
}

// Perform a POST request
func HttpReqPost(u string, pd []byte, hc *http.Client, st string) ([]byte, error) {
	httpReq, err := http.NewRequest("POST", u, bytes.NewBuffer(pd))
	if err != nil {
		return []byte{}, fmt.Errorf("an error occured creating new HTTP GET request: %v", err)
	}

	// Set HTTP header
	setReqHeader(httpReq)

	// We have a session token
	if st != "" {
		httpReq.Header.Add("X-Session-Token", st)
	}

	serverResp, err := hc.Do(httpReq)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		err := serverResp.Body.Close()
		if err != nil {
			log.Printf("error while closing response body: %v", err)
		}
	}()
	if serverResp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("kratos server responded with non-HTTP 200: %q", serverResp.Status)
	}
	if strings.HasSuffix("application/json", serverResp.Header.Get("Content-Type")) {
		return []byte{}, fmt.Errorf("kratos returned non-JSON response")
	}

	// Read the response body
	var respBody []byte
	buf := bufio.NewScanner(serverResp.Body)
	for buf.Scan() {
		respBody = buf.Bytes()
	}
	if err = buf.Err(); err != nil {
		return []byte{}, err
	}

	return respBody, nil
}

// Set package specific HTTP header
func setReqHeader(h *http.Request) {
	h.Header.Set("User-Agent", "wftk-go v"+go_wftk.Version)
	h.Header.Set("Content-Type", "application/json")
	h.Header.Set("Accept", "application/json")
}
