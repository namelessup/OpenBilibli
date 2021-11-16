#! /bin/sh
# proto.sh
gopath=$GOPATH/src
gogopath=$GOPATH/src/github.com/namelessup/bilibili/vendor/github.com/gogo/protobuf
protoc --gofast_out=. --proto_path=$gopath:$gogopath:. *.proto


