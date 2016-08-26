// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Copyright (c) 2012 Rodrigo Moraes. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

	 * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
	 * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
	 * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.
*/

package quobyte

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
)

// ----------------------------------------------------------------------------
// Request, Response and Error
// ----------------------------------------------------------------------------

// clientRequest represents a JSON-RPC request sent by a client.
type clientRequest struct {
	// JSON-RPC protocol.
	Version string `json:"jsonrpc"`

	// A String containing the name of the method to be invoked.
	Method string `json:"method"`

	// Object to pass as request parameter to the method.
	Params interface{} `json:"params"`

	// The request id. This can be of any type. It is used to match the
	// response with the request that it is replying to.
	ID uint64 `json:"id"`
}

// clientResponse represents a JSON-RPC response returned to a client.
type clientResponse struct {
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   *json.RawMessage `json:"error"`
}

type responseError struct {
	// A Number that indicates the error type that occurred.
	Code int `json:"code"` /* required */

	// A String providing a short description of the error.
	// The message SHOULD be limited to a concise single sentence.
	Message string `json:"message"` /* required */

	// A Primitive or Structured value that contains additional information about the error.
	Data interface{} `json:"data"` /* optional */
}

func (e *responseError) Error() string {
	return e.Message
}

// EncodeClientRequest encodes parameters for a JSON-RPC client request.
func encodeClientRequest(method string, args interface{}) ([]byte, error) {
	c := &clientRequest{
		Version: "2.0",
		Method:  method,
		Params:  args,
		ID:      uint64(rand.Int63()),
	}

	return json.Marshal(c)
}

// DecodeClientResponse decodes the response body of a client request into
// the interface reply.
func decodeClientResponse(r io.Reader, reply interface{}) error {
	var c clientResponse
	if err := json.NewDecoder(r).Decode(&c); err != nil {
		return err
	}

	if c.Error != nil {
		jsonErr := &responseError{}
		if err := json.Unmarshal(*c.Error, jsonErr); err != nil {
			return &responseError{
				Code:    -32000,
				Message: string(*c.Error),
			}
		}
		return jsonErr
	}

	if c.Result == nil {
		return errors.New("result is nil")
	}

	return json.Unmarshal(*c.Result, reply)
}
