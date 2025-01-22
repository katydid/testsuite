gofmt:
	gofmt -l -s -w .

regenerate:
	goderive .
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf person.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf srctree.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf treeregister.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf typewriterprison.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf puddingmilkshake.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf topsyturvy.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf knot.proto)
	(protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf banana.proto)
	find . -name "*.pb.go" | xargs gofmt -l -s -w
