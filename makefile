all:
	@rm -rf ./dist
	@GOOS=windows GOARCH=386 go build -o ./dist/ServerStatus-windows.exe main.go
	@go build -o ./dist/ServerStatus-linux main.go
	@cp config.template.json ./dist/config.json