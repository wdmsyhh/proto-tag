# proto-tag

a plugin for protoc-gen-go to reset struct tags.

## why

golang protobuf doesn't support custom tags to generated structs. this plugin help to set custom tags to generated protobuf file.

## install

```shell
go get github.com/golang/protobuf
cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go
git clone https://github.com/PilockHulmes/proto_tag tag

然后在 $GOPATH/src/github.com/golang/protobuf/protoc-gen-go/link_grpc.go 文件中加一行 import _ "github.com/golang/protobuf/protoc-gen-go/tag"
最后执行 go install github.com/golang/protobuf/protoc-gen-go 并使用 $GOPATH/bin/protoc-gen-go
```

## usage

Add a comment with syntax `//｀custom_tag1:custom_value1 custom_tag2:custom_value2｀` after fields.

Example:

```proto
syntax = "proto3";

package staff;

message Staff {
    string ID = 1;    //`json:"id,omitempty"   xml:"id,omitempty"`
    string Name = 2;  //`json:"name,omitempty" xml:"name,omitempty"`
    int64 Age = 3;    //`json:"age,omitempty"  xml:"age,omitempty"`
}
```

generate `.pb.go` with command `protoc` as:

```shell
protoc --go_out=plugins=grpc+tag:. example.proto
```

the custom tag will set to strcut:

```golang
type Staff struct {
    ID   string `protobuf:"bytes,1,opt,name=ID"     json:"id,omitempty"    xml:"id,omitempty"`
    Name string `protobuf:"bytes,2,opt,name=Name"   json:"name,omitempty"  xml:"name,omitempty"`
    Age  int64  `protobuf:"varint,3,opt,name=Age"   json:"age,omitempty"   xml:"age,omitempty"`
}
```


message in message:

```protobuf

package staff;

message Staff {
    string ID = 1;    //`json:"id,omitempty"   xml:"id,omitempty"`
    string Name = 2;  //`json:"name,omitempty" xml:"name,omitempty"`
    int64 Age = 3;    //`json:"age,omitempty"  xml:"age,omitempty"`
    message Class {
        string ID = 1;      //`json:"id,omitempty"    xml:"id,omitempty"`
        string Type = 2;    //`json:"type,omitempty"  xml:"type,omitempty"`
    };

    Class MyClass = 4;    //`json:"class,omitempty"     xml:"class,omitempty"`
}
```

```golang
type Staff struct {
	ID      string       `protobuf:"bytes,1,opt,name=ID"        json:"id,omitempty"     xml:"id,omitempty"`
	Name    string       `protobuf:"bytes,2,opt,name=Name"      json:"name,omitempty"   xml:"name,omitempty"`
	Age     int64        `protobuf:"varint,3,opt,name=Age"      json:"age,omitempty"    xml:"age,omitempty"`
	MyClass *Staff_Class `protobuf:"bytes,4,opt,name=MyClass"   json:"class,omitempty"  xml:"class,omitempty"`
}

type Staff_Class struct {
	ID   string `protobuf:"bytes,1,opt,name=ID"        json:"id,omitempty"     xml:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type"      json:"type,omitempty"   xml:"type,omitempty"`
}
```