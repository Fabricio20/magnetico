.PHONY: test format vet magneticod image image-magneticod

all: test magneticod

magneticod:
	go mod download
	go install --tags fts5 "-ldflags=-s -w -X main.compiledOn=`date -u +%Y-%m-%dT%H:%M:%SZ`" ./cmd/magneticod

.PHONY: docker
docker: docker_up docker_logs

.PHONY: docker_up
docker_up:
	docker-compose up -d

.PHONY: docker_down
docker_down:
	docker-compose down

.PHONY: docker_logs
docker_logs:
	docker-compose logs -ft --tail=10

image-magneticod:
	docker build -t boramalper/magneticod -f Dockerfile .

image: image-magneticod

vet:
	go vet ./...

test:
	go test ./...

format:
	gofmt -w ./cmd/
	gofmt -w ./pkg/

# Formatting Errors
#     Since gofmt returns zero even if there are files to be formatted, we use:
#
#       ! gofmt -d ${GOPATH}/path/ 2>&1 | read
#
#     to return 1 if there are files to be formatted, and 0 if not.
#     https://groups.google.com/forum/#!topic/Golang-Nuts/pdrN4zleUio
#
# How can I use Bash syntax in Makefile targets?
#     Because `read` is a bash command.
#     https://stackoverflow.com/a/589300/4466589
#
# How to ignore bindata.go
#     Due to irrational insistence of some Go developers, gofmt, like many
#     other tools of Go ecosystem, does not have flags for common scenarios
#     such as ignoring certain files by pattern etc. Thus we use `go list`
#     and grep together to achieve the desired result.
#
#     The original query is this:
#         gofmt -l $(go list -f $'{{range .GoFiles}}{{$.Dir}}/{{.}}\n{{end}}' ./... | grep -v bindata.go)
#
#     The original query is then escaped for Makefile (by repeating dollar signs $ -> $$).
check-formatting: SHELL:=/bin/bash   # HERE: this is setting the shell for check-formatting only
check-formatting:
	! gofmt -l $$(go list -f $$'{{range .GoFiles}}{{$$.Dir}}/{{.}}\n{{end}}' ./... | grep -v bindata.go) 2>&1 | tee /dev/fd/2 | read
	! gofmt -l $$(go list -f $$'{{range .GoFiles}}{{$$.Dir}}/{{.}}\n{{end}}' ./... | grep -v bindata.go) 2>&1 | tee /dev/fd/2 | read
