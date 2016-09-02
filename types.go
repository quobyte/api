package quobyte

// CreateVolumeRequest represents a CreateVolumeRequest
type CreateVolumeRequest struct {
	Name              string  `json:"name,omitempty"`
	RootUserID        string  `json:"root_user_id,omitempty"`
	RootGroupID       string  `json:"root_group_id,omitempty"`
	ReplicaDeviceIDS  []int64 `json:"replica_device_ids,omitempty"`
	ConfigurationName string  `json:"configuration_name,omitempty"`
	AccessMode        int32   `json:"access_mode,omitempty"`
	TenantID          string  `json:"tenant_id,omitempty"`
}

type createVolumeResponse struct {
	VolumeUUID string `json:"volume_uuid"`
}

type deleteVolumeRequest struct {
	VolumeUUID string `json:"volume_uuid"`
}

type resolveVolumeNameRequest struct {
	VolumeName string `json:"volume_name,omitempty"`
}

type resolveVolumeNameResponse struct {
	VolumeUUID string `json:"volume_uuid,omitempty"`
}
