regenerate:
	(cd validator && make regenerate)

nuke:
	(cd validator && make nuke)

regenerate-tests:
	(cd validator && make regenerate-tests)

build:
	go build ./...