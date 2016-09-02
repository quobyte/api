package quobyte

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSuccesfullEncodeRequest(t *testing.T) {
	req := &CreateVolumeRequest{
		RootUserID:  "root",
		RootGroupID: "root",
		Name:        "test",
	}

	//Generate Params here
	var param map[string]interface{}
	byt, _ := json.Marshal(req)
	_ = json.Unmarshal(byt, &param)

	expectedRPCRequest := &request{
		ID:      0,
		Method:  "createVolume",
		Version: "2.0",
		Params:  param,
	}

	res, err := encodeRequest("createVolume", req)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var reqResult request
	if err := json.Unmarshal(res, &reqResult); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if expectedRPCRequest.Version != reqResult.Version {
		t.Logf("Expected Version: %s got %s\n", expectedRPCRequest.Version, reqResult.Version)
		t.Fail()
	}

	if expectedRPCRequest.Method != reqResult.Method {
		t.Logf("Expected Method: %s got %s\n", expectedRPCRequest.Method, reqResult.Method)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedRPCRequest.Params, reqResult.Params) {
		t.Logf("Expected Params: %v got %v\n", expectedRPCRequest.Params, reqResult.Params)
		t.Fail()
	}
}

//TODO
func TestSuccesfullDecodeResponse(t *testing.T) {}

//TODO
func TestSuccesfullDecodeResponseWithError(t *testing.T) {}

//TODO test -> decodeErrorCode
