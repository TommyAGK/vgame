all: 
	test vet fmt lint build
test:
	go test ./tests/vgame_tests/
	go test ./tests/fsm_tests/


vet: 
	go vet ./...

fmt:
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l
	go list -f '{{.Dir}}' ./... | grep -v /vendor/ | xargs -L1 gofmt -l

lint:
	echo 'currently not available'
#go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
	go build -o bin/vgame ./main.go
