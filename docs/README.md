# Developer notes

## Releasing new version

* Compile Quobyte source and copy generated types.go from <QUOBYTE_SOURCE>/build/golang/api/types.go
  to <api-rep>/quobyte/types.go
* Run `gofmt -w <api-rep>/quobyte/types.go`
* `go.mod` files must be present at the root level of the project
* Run `go build ./...` in the project root (in the parent of go.mod file) to run compilation pass
  (binary is not produced but checks for possible compilation issues)
* If compilation is successful, run `go generate ./...` to generate mocks
* Each major release beyond V1 (such =v2[+].a.b) must provide unique import path such as `github.com/quobyte/api/vX`
  * To get around this issue, we always use v1.x.x (**NEVER** make v2 release)
  * Further, each `*.go` file must have a `package XYZ` statement as the first line and must be placed into `XZY`
    directory. (It seems go mod ignores .go file that is not placed in declared package directory!!)
* For local testing of edited module,
  * Create a standalone project with the `testing or main.go` and `go mod init` inside the project root.
      The `go mod init` fetches the dependencies required by code.
  * Replace Quobyte API with updated API `go mod edit -replace github.com/quobyte/api=</path/to/local/quobyte/api>`
* Publishing change must always have highest minor version of all the published tags (even if the tag is deleted,
 the new version must have the higher version than deleted tag).

Note: go mod updates dependency to the highest minor version available. Even if the version is deleted from the tags,
go-git gets it from the deleted references (Uhhh!!). The only way out is to create a new higher minor
version with the changes.
