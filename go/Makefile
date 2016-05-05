.PHONY: cover

cover:
	go test -coverprofile=/tmp/cove.out "github.com/yubo/hw/go"
	go tool cover -html=/tmp/cove.out 
