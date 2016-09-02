package quobyte

// CreateVolumeRequest represents a CreateVolumeRequest
type CreateVolumeRequest struct {
	Name              string  `json:"name"`
	RootUserID        string  `json:"root_user_id"`
	RootGroupID       string  `json:"root_group_id"`
	ReplicaDeviceIDS  []int64 `json:"replica_device_ids"`
	ConfigurationName string  `json:"configuration_name"`
	AccessMode        int32   `json:"access_mode"`
	TenantID          string  `json:"tenant_id"`
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
