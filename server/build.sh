path=`pwd`
export GOPATH=$path:$path/../lib

rm -rf bin/*
cp test bin
cd bin
go build server