## Installation instructions for Go

Follow these steps to get this Go workspace up and running.

## Setup

1. Follow the instructions [here](https://golang.org/doc/install) to install Go.

2. Set the GOPATH for this workspace

    ```
    export GOPATH={your_workspace}/Go
    export PATH=$PATH:$GOPATH/bin
    cd $GOPATH/src/main
    ```

3. Build and install

    ```
    go install
    ```

4. Execution
    ```
    go run main
    ```
