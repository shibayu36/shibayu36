GOLDFLAGS = -X main.Version=$$(git describe --tags HEAD)

setup:
	go get -v \
		github.com/golang/dep/cmd/dep \
		github.com/laher/goxc \
		github.com/tcnksm/ghr

deps: setup
	dep ensure

test: deps
	go test -v ./...

xbuild: deps
	gox \
		-output "build/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}" \
		-ldflags "$(GOLDFLAGS)"

.PHONY: setup deps test xbuild
