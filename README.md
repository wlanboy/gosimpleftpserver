# gosimpleftpserver
Golang simple ftp server
- depends on github.com/goftp/server and github.com/syndtr/goleveldb

# build
* go get -d -v
* go clean
* go build

# run
* go run ftp.go

# debug
* go get -u github.com/go-delve/delve/cmd/dlv
* dlv debug ./ftp
