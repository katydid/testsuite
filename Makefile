dep:
	go install -v github.com/gogo/protobuf/protoc-gen-gogo
	go install -v github.com/awalterschulze/goderive

build:
	go build ./...

nuke:
	(cd validator && make nuke)

regenerate-tests:
	(cd validator && make regenerate-tests)

regenerate-all:
	(cd validator && make regenerate-all)
