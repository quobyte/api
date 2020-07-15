# Quobyte API Clients

Get the quobyte api client

```bash
go get github.com/quobyte/api
```

## Usage

```go
package main

import (
  "log"
  quobyte_api "github.com/quobyte/api"
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

    client := quobyte_api.NewQuobyteClient(*url, *username, *password)
    client.SetAPIRetryPolicy(quobyte_api.RetryInfinitely) // Default quobyte_api.RetryInteractive
    req := &quobyte_api.CreateVolumeRequest{
        Name:              "MyVolume",
        TenantId:          "32edb36d-badc-affe-b44a-4ab749af4d9a",
        RootUserId:        "root",
        RootGroupId:	   "root",
        ConfigurationName: "BASE",
        Label: []*quobyte_api.Label{
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
