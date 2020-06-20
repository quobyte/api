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
    client := quobyte_api.NewQuobyteClient("http://apiserver:7860", "user", "password")
    client.SetAPIRetryPolicy(quobyte_api.RetryInfinitely) // Default quobyte_api.RetryInteractive
    req := &quobyte_api.CreateVolumeRequest{
        Name:              "MyVolume",
        RootUserId:        "root",
        RootGroupId:       "root",
        ConfigurationName: "BASE",
        Labels: []quobyte_api.Label{
            {Name: "label1", Value: "value1"},
            {Name: "label2", Value: "value2"},
        },
    }

    response, err := client.CreateVolume(req)
    if err != nil {
        log.Fatalf("Error:", err)
    }

    log.Printf("Created volume %s", response.volumeUuid)
}
```
