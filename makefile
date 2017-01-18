run:
	go run ./src/main.go

linux:
	go build -o ./dist/ServerStatus-linux ./src/main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/ServerStatus-mac ./src/main.go
	
windows:
	GOOS=windows GOARCH=386 go build -o ./dist/ServerStatus-windows.exe ./src/main.go

clean:
	rm -rf ./dist

copyfiles:
	cp config.template.json ./dist/config.json

all: linux mac windows copyfiles