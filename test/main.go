package main

import (
	"fmt"
	proto_tag "proto-tag"
)

func main() {
	tag := new(proto_tag.Tag)
	tag.GetStructTags("/home/user/GolandProjects/proto-tag/test/example/example.proto")
	tag.Tag("/home/user/GolandProjects/proto-tag/test/example/example.pb.go")
	fmt.Println(tag)
}
