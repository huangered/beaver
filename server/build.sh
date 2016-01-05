path=`pwd`
export GOPATH=$path:$path/../lib

rm -rf bin/*
cd bin
go build server