// Package quobyte represents a golang API for the Quobyte Storage System
package quobyte

import "net/http"

type QuobyteClient struct {
	client   *http.Client
	url      string
	username string
	password string
}

// NewQuobyteClient creates a new Quobyte API client
func NewQuobyteClient(url string, username string, password string) *QuobyteClient {
	return &QuobyteClient{
		client:   &http.Client{},
		url:      url,
		username: username,
		password: password,
	}
}

// Create a new Quobyte volume. Its root directory will be owned by given user and group
func (client QuobyteClient) CreateVolume(request *CreateVolumeRequest) (string, error) {
	var response createVolumeResponse
	if err := client.sendRequest("createVolume", request, &response); err != nil {
		return "", err
	}

	return response.VolumeUUID, nil
}

// ResolveVolumeNameToUUID resolves a volume name to a UUID
func (client QuobyteClient) ResolveVolumeNameToUUID(volumeName string) (string, error) {
	request := &resolveVolumeNameRequest{
		VolumeName: volumeName,
	}
	var response resolveVolumeNameResponse
	if err := client.sendRequest("resolveVolumeName", request, &response); err != nil {
		return "", err
	}

	return response.VolumeUUID, nil
}

// Delete a Quobyte volume. Its root directory will be owned by given user and group and have access 700.
func (client QuobyteClient) DeleteVolume(volumeUUID string) error {
	request := &deleteVolumeRequest{
		VolumeUUID: volumeUUID,
	}

	return client.sendRequest("deleteVolume", request, nil)
}

func (client QuobyteClient) DeleteVolumeByName(volumeName string) error {
	uuid, err := client.ResolveVolumeNameToUUID(volumeName)
	if err != nil {
		return err
	}

	return client.DeleteVolume(uuid)
}
