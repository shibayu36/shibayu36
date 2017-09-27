setup:
	go get github.com/Masterminds/glide

deps: setup
	glide install

test: deps
	go test -v $$(glide novendor)

.PHONY: setup deps test
