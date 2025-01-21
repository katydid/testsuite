gofmt:
	gofmt -l -s -w .

regenerate:
	(protoc --go_out=. --go_opt=paths=source_relative person.proto)
	(protoc --go_out=. --go_opt=paths=source_relative srctree.proto)
	(protoc --go_out=. --go_opt=paths=source_relative treeregister.proto)
	(protoc --go_out=. --go_opt=paths=source_relative typewriterprison.proto)
	(protoc --go_out=. --go_opt=paths=source_relative puddingmilkshake.proto)
	(protoc --go_out=. --go_opt=paths=source_relative topsyturvy.proto)
	(protoc --go_out=. --go_opt=paths=source_relative knot.proto)
	(protoc --go_out=. --go_opt=paths=source_relative banana.proto)
	find . -name "*.pb.go" | xargs gofmt -l -s -w
