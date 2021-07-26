# Quobyte API Clients

Quobyte golang API follows `go module` system for versioning.

Get the Quobyte 2.x api client

```bash
# Replace v2.x.x with a valid tag v2.x.x api tag (https://github.com/quobyte/api/tags)
go get github.com/quobyte/api@<v2.x.x>
```

## Usage

```go
package main

import (
  "log"
  // Get Quobyte 2.x api
  quobyte_v2_api "github.com/quobyte/api/v2"
)

func main() {
    url := flag.String("url", "", "URL of Quobyte API")
    username := flag.String("username", "", "username")
    password := flag.String("password", "", "password")
    flag.Parse()

    if *url == "" || *username == "" || *password == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

    client := quobyte_v2_api.NewQuobyteClient(*url, *username, *password)
    client.SetAPIRetryPolicy(quobyte_v2_api.RetryInfinitely) // Default quobyte_api.RetryInteractive
    req := &quobyte_v2_api.CreateVolumeRequest{
        Name:              "MyVolume",
        TenantId:          "32edb36d-badc-affe-b44a-4ab749af4d9a",
        RootUserId:        "root",
        RootGroupId:	   "root",
        ConfigurationName: "BASE",
        Label: []*quobyte_v2_api.Label{
            {Name: "label1", Value: "value1"},
            {Name: "label2", Value: "value2"},
        },
    }

    response, err := client.CreateVolume(req)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    capactiy := int64(1024 * 1024 * 1024)
    err = client.SetVolumeQuota(response.VolumeUuid, uint64(capactiy))
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    log.Printf("%s", response.VolumeUuid)
}
```
