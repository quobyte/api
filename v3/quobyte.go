package quobyte

import (
	"net/http"
	"regexp"
)

// retry policy codes
const (
	RetryInteractive string = "INTERACTIVE"
	RetryInfinitely  string = "INFINITELY"
)

var UUIDValidator = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

type QuobyteClient struct {
	client         *http.Client
	url            string
	username       string
	password       string
	apiRetryPolicy string
}

func (client *QuobyteClient) SetAPIRetryPolicy(retry string) {
	client.apiRetryPolicy = retry
}

func (client *QuobyteClient) GetAPIRetryPolicy() string {
	return client.apiRetryPolicy
}

func (client *QuobyteClient) SetTransport(t http.RoundTripper) {
	client.client.Transport = t
}

// NewQuobyteClient creates a new Quobyte API client
func NewQuobyteClient(url string, username string, password string) *QuobyteClient {
	return &QuobyteClient{
		client:         &http.Client{},
		url:            url,
		username:       username,
		password:       password,
		apiRetryPolicy: RetryInteractive,
	}
}

// GetVolumeUUID resolves the volumeUUID for the given volume and tenant name.
// This method should be used when it is not clear if the given string is volume UUID or Name.
func (client *QuobyteClient) GetVolumeUUID(volume, tenant string) (string, error) {
	if len(volume) != 0 && !IsValidUUID(volume) {
		tenantUUID, err := client.GetTenantUUID(tenant)
		if err != nil {
			return "", err
		}
		volUUID, err := client.ResolveVolumeNameToUUID(volume, tenantUUID)
		if err != nil {
			return "", err
		}

		return volUUID, nil
	}
	return volume, nil
}

// GetTenantUUID resolves the tenatnUUID for the given name
// This method should be used when it is not clear if the given string is Tenant UUID or Name.
func (client *QuobyteClient) GetTenantUUID(tenant string) (string, error) {
	if len(tenant) != 0 && !IsValidUUID(tenant) {
		tenantUUID, err := client.ResolveTenantNameToUUID(tenant)
		if err != nil {
			return "", err
		}
		return tenantUUID, nil
	}
	return tenant, nil
}

// ResolveVolumeNameToUUID resolves a volume name to a UUID
func (client *QuobyteClient) ResolveVolumeNameToUUID(volumeName, tenant string) (string, error) {
	request := &ResolveVolumeNameRequest{
		VolumeName:   volumeName,
		TenantDomain: tenant,
	}
	var response ResolveVolumeNameResponse
	if err := client.SendRequest("resolveVolumeName", request, &response); err != nil {
		return "", err
	}

	return response.VolumeUuid, nil
}

// DeleteVolumeByResolvingNamesToUUID deletes the volume by resolving the volume name and tenant name to
// respective UUID if required.
// This method should be used if the given volume, tenant information is name or UUID.
func (client *QuobyteClient) DeleteVolumeByResolvingNamesToUUID(volume, tenant string) error {
	volumeUUID, err := client.GetVolumeUUID(volume, tenant)
	if err != nil {
		return err
	}

	_, err = client.DeleteVolume(&DeleteVolumeRequest{VolumeUuid: volumeUUID})
	return err
}

// DeleteVolumeByName deletes a volume by a given name
func (client *QuobyteClient) DeleteVolumeByName(volumeName, tenant string) error {
	uuid, err := client.ResolveVolumeNameToUUID(volumeName, tenant)
	if err != nil {
		return err
	}

	_, err = client.DeleteVolume(&DeleteVolumeRequest{VolumeUuid: uuid})
	return err
}

// SetVolumeQuota sets a Quota to the specified Volume
func (client *QuobyteClient) SetVolumeQuota(volumeUUID string, quotaSize int64) error {
	request := &SetQuotaRequest{
		Quotas: []*Quota{
			&Quota{
				Consumer: []*ConsumingEntity{
					&ConsumingEntity{
						Type:       ConsumingEntity_Type_VOLUME,
						Identifier: volumeUUID,
					},
				},
				Limits: []*Resource{
					&Resource{
						Type:  Resource_Type_LOGICAL_DISK_SPACE,
						Value: quotaSize,
					},
				},
			},
		},
	}

	return client.SendRequest("setQuota", request, nil)
}

// GetTenantMap returns a map that contains all tenant names and there ID's
func (client *QuobyteClient) GetTenantMap() (map[string]string, error) {
	result := map[string]string{}
	response, err := client.GetTenant(&GetTenantRequest{})

	if err != nil {
		return result, err
	}

	for _, tenant := range response.Tenant {
		result[tenant.Name] = tenant.TenantId
	}

	return result, nil
}

// IsValidUUID Validates the given uuid
func IsValidUUID(uuid string) bool {
	return UUIDValidator.MatchString(uuid)
}

// ResolveTenantNameToUUID Returns UUID for given name, error if not found.
func (client *QuobyteClient) ResolveTenantNameToUUID(name string) (string, error) {
	request := &ResolveTenantNameRequest{
		TenantName: name,
	}

	var response ResolveTenantNameResponse
	err := client.SendRequest("resolveTenantName", request, &response)
	if err != nil {
		return "", err
	}
	return response.TenantId, nil
}
