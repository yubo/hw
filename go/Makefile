.PHONY: doc docd cover test cover bench

doc:
	godoc github.com/yubo/hw/go

docd:
	godoc -http=:6060

test:
	go test

cover:
	go test -coverprofile=/tmp/cove.out "github.com/yubo/hw/go"
	go tool cover -html=/tmp/cove.out 

bench:
	go test -bench=.
