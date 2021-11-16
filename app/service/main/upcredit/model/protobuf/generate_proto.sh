gopath="$GOPATH/src"
protobuf_path="$gopath/github.com/namelessup/bilibili/vendor/github.com/gogo/protobuf"
echo $protobuf_path
ls "$gopath/github.com/namelessup/bilibili/vendor/github.com/gogo/protobuf/gogoproto/"
protoc --gofast_out=".." -I"../" -I"$gopath/github.com/namelessup/bilibili/vendor/" ../*.proto
