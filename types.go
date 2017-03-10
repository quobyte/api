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

type GetDeviceListRequest struct {
	DeviceID   []uint64 `json:"device_id,omitempty"`
	DeviceType []string `json:"device_type,omitempty"`
}

type GetDeviceListResponse struct {
	DeviceList DeviceList `json:"device_list,omitempty"`
}

type DeviceList struct {
	Devices []Device `json:"devices,omitempty"`
}

type Device struct {
	DeviceStatus            string              `json:"device_status,omitempty"`
	DeviceID                uint64              `json:"device_id,omitempty"`
	DeviceLabel             string              `json:"device_label,omitempty"`
	Content                 []DeviceContent     `json:"content,omitempty"`
	HostName                string              `json:"host_name,omitempty"`
	TotalDiskSpaceBytes     uint64              `json:"total_disk_space_bytes,omitempty"`
	UsedDiskSpaceBytes      uint64              `json:"used_disk_space_bytes,omitempty"`
	DeviceTags              []string            `json:"device_tags,omitempty"`
	FailureDomainInfos      []FailureDomainInfo `json:"failure_domain_infos,omitempty"`
	IsEmpty                 bool                `json:"is_empty,omitempty"`
	Draining                bool                `json:"draining,omitempty"`
	DeviceSerialNumber      string              `json:"device_serial_number,omitempty"`
	DeviceModel             string              `json:"device_model,omitempty"`
	DetectedDiskType        string              `json:"detected_disk_type,omitempty"`
	CurrentMountPathstring  string              `json:"current_mount_pathstring,omitempty"`
	FileCount               int64               `json:"file_count,omitempty"`
	VolumeDatabaseCount     int64               `json:"volume_database_count,omitempty"`
	RegistryDatabaseCount   int64               `json:"registry_database_countcontent,omitempty"`
	IOErrorCount            int64               `json:"io_error_count,omitempty"`
	CRCErrorCount           int64               `json:"crc_error_count,omitempty"`
	LastSuccessfulCleanupMS int64               `json:"last_successful_cleanup_ms,omitempty"`
	CurrentUtilization      float64             `json:"current_utilization,omitempty"`
	ReallocatedSectorCt     int64               `json:"reallocated_sector_ct,omitempty"`
	ReportedUncorrect       int64               `json:"reported_uncorrect,omitempty"`
	CommandTimeout          int64               `json:"command_timeout,omitempty"`
	CurrentPendingSector    int64               `json:"current_pending_sector,omitempty"`
	OfflineUncorrectable    int64               `json:"offline_uncorrectable,omitempty"`
	IsPrimary               bool                `json:"is_primary,omitempty"`
}

type DeviceContent struct {
	ContentType         string `json:"content_type,omitempty"`
	ServiceUUID         string `json:"service_uuid,omitempty"`
	LastSeenTimestampMS uint64 `json:"last_seen_timestamp_ms,omitempty"`
	Available           bool   `json:"available,omitempty"`
	LastSeenServiceUUID string `json:"last_seen_service_uuid,omitempty"`
	LastSeenServiceName string `json:"last_seen_service_name,omitempty"`
	LastSeenMountPath   string `json:"last_seen_mount_path,omitempty"`
}

type FailureDomainInfo struct {
	Name       string `json:"name,omitempty"`
	DomainType string `json:"domain_type,omitempty"`
}
