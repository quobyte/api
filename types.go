package quobyte

// CreateVolumeRequest represents a CreateVolumeRequest
type CreateVolumeRequest struct {
	Name              string   `json:"name,omitempty"`
	RootUserID        string   `json:"root_user_id,omitempty"`
	RootGroupID       string   `json:"root_group_id,omitempty"`
	ReplicaDeviceIDS  []uint64 `json:"replica_device_ids,string,omitempty"`
	ConfigurationName string   `json:"configuration_name,omitempty"`
	AccessMode        uint32   `json:"access_mode,string,omitempty"`
	TenantID          string   `json:"tenant_id,omitempty"`
}

type resolveVolumeNameRequest struct {
	VolumeName   string `json:"volume_name,omitempty"`
	TenantDomain string `json:"tenant_domain,omitempty"`
}

type volumeUUID struct {
	VolumeUUID string `json:"volume_uuid,omitempty"`
}

type getClientListRequest struct {
	TenantDomain string `json:"tenant_domain,omitempty"`
}

type GetClientListResponse struct {
	Clients []Client `json:"client,omitempty"`
}

type Client struct {
	MountedUserName   string `json:"mount_user_name,omitempty"`
	MountedVolumeUUID string `json:"mounted_volume_uuid,omitempty"`
}

type GetDeviceNetworkEndpointsRequest struct {
	DeviceID uint64 `json:"device_id,omitempty"`
}

type DeviceNetworkEndpoint struct {
	DeviceType string `json:"device_type,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	Port       int    `json:"devicporte_id,omitempty"`
}

type GetDeviceNetworkEndpointsResponse struct {
	Endpoints []DeviceNetworkEndpoint `json:"endpoints,omitempty"`
}
