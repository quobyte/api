package quobyte

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"sync"
)

const (
	emptyResponse      string = "Empty result and no error occurred"
	errorMessageFormat string = "method: %s error message: %s"
)

var mux sync.Mutex

type request struct {
	ID      string      `json:"id"`
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type response struct {
	ID      string           `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

type rpcError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (err *rpcError) decodeErrorCode() string {
	switch err.Code {
	case -32600:
		return "ERROR_CODE_INVALID_REQUEST"
	case -32603:
		return "ERROR_CODE_JSON_ENCODING_FAILED"
	case -32601:
		return "ERROR_CODE_METHOD_NOT_FOUND"
	case -32700:
		return "ERROR_CODE_PARSE_ERROR"
	}

	return ""
}

func encodeRequest(method string, params interface{}) ([]byte, error) {
	return json.Marshal(&request{
		// Generate random ID and convert it to a string
		ID:      strconv.FormatInt(rand.Int63(), 10),
		Version: "2.0",
		Method:  method,
		Params:  params,
	})
}

func decodeResponse(method string, ioReader io.Reader, reply interface{}) error {
	var resp response
	if err := json.NewDecoder(ioReader).Decode(&resp); err != nil {
		return err
	}

	if resp.Error != nil {
		var rpcErr rpcError
		if err := json.Unmarshal(*resp.Error, &rpcErr); err != nil {
			return err
		}

		if rpcErr.Message != "" {
			return fmt.Errorf(errorMessageFormat, method, rpcErr.Message)
		}

		respError := rpcErr.decodeErrorCode()
		if respError != "" {
			return fmt.Errorf(errorMessageFormat, method, respError)
		}
	}

	if resp.Result != nil && reply != nil {
		return json.Unmarshal(*resp.Result, reply)
	}

	return fmt.Errorf(errorMessageFormat, method, emptyResponse)
}

func (client QuobyteClient) sendRequest(method string, request interface{}, response interface{}) error {
	etype := reflect.ValueOf(request).Elem()
	field := etype.FieldByName("RetryPolicy")
	if field.IsValid() {
		field.SetString(client.GetAPIRetryPolicy())
	}
	message, err := encodeRequest(method, request)
	if err != nil {
		return err
	}
	// If no cookies, serialize requests such that first successful request sets the cookies
	for {
		req, err := http.NewRequest("POST", client.url.String(), bytes.NewBuffer(message))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		mux.Lock()
		hasCookies, err := client.hasCookies()
		if err != nil {
			return err
		}
		if !hasCookies {
			req.SetBasicAuth(client.username, client.password)
			// no cookies available, must hold lock until request is completed and
			// new cookies are created by server
			defer mux.Unlock()
		} else {
			// let every thread/routine send request using the cookie
			mux.Unlock()
		}
		resp, err := client.client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			if resp.StatusCode == 401 {
				_, ok := req.Header["Authorization"]
				if ok {
					return errors.New("Unable to authenticate with Quobyte API service")
				}
				// Session is not valid anymore (service restart, session invalidated etc)!!
				// resend basic auth and get new cookies
				// invalidate session cookies
				cookieJar := client.client.Jar
				if cookieJar != nil {
					cookies := cookieJar.Cookies(client.url)
					for _, cookie := range cookies {
						cookie.MaxAge = -1
					}
					cookieJar.SetCookies(client.url, cookies)
				}
				// retry request with authorization header
				continue
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return (err)
			}
			return fmt.Errorf("JsonRPC failed with error (error code: %d) %s",
				resp.StatusCode, string(body))
		}
		return decodeResponse(method, resp.Body, &response)
	}
}
