default:
	@go run main.go test name="example"

linux-build:
	@echo "building grunt binary - linux"
	@go build -o ./bin/grunt
	@echo "grunt binary built - linux"

win-build:
	@echo "building grunt binary - windows"
	@go build -o ./bin/grunt.exe
	@echo "grunt binary built - windows"

build-all:
	@echo "building grunt binaries - both"
	@make linux-build
	@make win-build
	@echo "grunt binaries built - both"