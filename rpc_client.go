package quobyte

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
)

type request struct {
	ID      int64       `json:"id"`
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type response struct {
	ID      int64            `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

type rpcError struct {
	Code    int    `json:"code"`
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
		ID:      int64(rand.Int63()),
		Version: "2.0",
		Method:  method,
		Params:  params,
	})
}

func decodeResponse(ioReader io.Reader, reply interface{}) error {
	var resp response
	if err := json.NewDecoder(ioReader).Decode(&resp); err != nil {
		return err
	}

	if resp.Error != nil {
		var rpcErr rpcError
		if err := json.Unmarshal(*resp.Error, &rpcErr); err != nil {
			if rpcErr.Message != "" {
				return errors.New(rpcErr.Message)
			}
			respError := rpcErr.decodeErrorCode()
			if respError != "" {
				return errors.New(respError)
			}

			return err
		}

		return errors.New(rpcErr.Message)
	}

	if resp.Result != nil {
		return json.Unmarshal(*resp.Result, reply)
	}

	return errors.New("Empty result and no error occured")
}

func (client QuobyteClient) sendRequest(method string, request interface{}, response interface{}) error {
	message, err := encodeRequest(method, request)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", client.url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(client.username, client.password)
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeResponse(resp.Body, &response)
}
