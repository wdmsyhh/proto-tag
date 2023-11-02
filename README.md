# proto-tag

a plugin for protoc-gen-go to add struct tags.

## install

```shell
#go get github.com/golang/protobuf
#git clone https://github.com/wdmsyhh/proto_tag
#cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go

git clone git@github.com:golang/protobuf.git
cd protoc-gen-go
mkdir tag
cp ~/proto_tag/tag.go ./tag/tag.go

protobuf/protoc-gen-go/link_grpc.go
import _ "github.com/golang/protobuf/protoc-gen-go/tag"
go install
```