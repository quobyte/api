# Releasing new version

1. Fix bugs/update code in v2/v3/v.. folders
2. Update go.mod file in the corresponding branch and update the `module`
   to specific release version. For example, if `./v3` is changed and will be
   released as github tag v3.0.1, the entry in `./v3/go.mod` must be as following  

    ```bash
     module github.com/quobyte/api/v3.0.1
    ```
