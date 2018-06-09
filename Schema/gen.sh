protoc --proto_path=proto/ -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:proto/. \
base.proto