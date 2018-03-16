VERSION := $(shell git describe --tags)

run:
	go run ./src/main.go

linux:
	go build -o ./dist/ServerStatus-linux -ldflags="-X main.version=${VERSION}" ./main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/ServerStatus-mac -ldflags="-X main.version=${VERSION}" ./main.go
	
windows:
	GOOS=windows GOARCH=386 go build -o ./dist/ServerStatus-windows.exe -ldflags="-X main.version=${VERSION}" ./main.go

clean:
	rm -rf ./dist

copyfiles:
	cp config.template.json ./dist/config.json

zip:
	zip -r dist.zip dist

all: linux mac windows copyfiles
