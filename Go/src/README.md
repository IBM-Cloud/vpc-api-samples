Installation instructions:
Follow these steps to get this Go workspace up and running.

Setup:

1. Follow the instruction to install Go - https://golang.org/doc/install

2. Set the GOPATH for this workspace
    export GOPATH={your_workspace}/Go
    export PATH=$PATH:$GOPATH/bin
    cd $GOPATH/src/main
    
3. Build and install
    go install

4. Execution
    go run main
