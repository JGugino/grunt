default:
	@go run main.go test name="example"

linux-build:
	@go build -o ./bin/grunt

win-build:
	@go build -o ./bin/grunt.exe

build-all:
	@make linux-build
	@make win-build