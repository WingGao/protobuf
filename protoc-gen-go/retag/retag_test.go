package retag

import (
	"testing"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func TestRetag_Import(t *testing.T) {
	rt := retag{}
	rt.Init(generator.New())
	rt.getStructTags("./example/exampleb.proto")
	t.Logf("%#v", rt.tags)
	rt.retag()
}
