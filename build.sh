mkdir bin
cd bin

set GOOS=windows
set GOARCH=386
go build -v -trimpath -ldflags "-s -w" -o log4j2-exp_386.exe ../cmd/main.go

set GOOS=windows
set GOARCH=amd64
go build -v -trimpath -ldflags "-s -w" -o log4j2-exp_amd64.exe ../cmd/main.go

set GOOS=linux
set GOARCH=386
go build -v -trimpath -ldflags "-s -w" -o log4j2-exp_386.elf ../cmd/main.go

set GOOS=linux
set GOARCH=amd64
go build -v -trimpath -ldflags "-s -w" -o log4j2-exp_amd64.elf ../cmd/main.go

cd ..