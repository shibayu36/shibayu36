setup:
	go get -v \
		github.com/Masterminds/glide \
		github.com/mitchellh/gox

deps: setup
	glide install

test: deps
	go test -v $$(glide novendor)

xbuild: deps
	gox -output "build/{{.Dir}}_{{.OS}}_{{.Arch}}/{{.Dir}}"

.PHONY: setup deps test xbuild
