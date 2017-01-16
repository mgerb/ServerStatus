run:
	@go run ./src/main.go

build:
	@rm -rf ./dist
	@GOOS=windows GOARCH=386 go build -o ./dist/ServerStatus-windows.exe ./src/main.go
	@go build -o ./dist/ServerStatus-linux ./src/main.go
	@cp config.template.json ./dist/config.json