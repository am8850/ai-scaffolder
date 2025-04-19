default:
	@echo "Please specify a target. Available targets are:"

build-linux:
	go build .

install: build-linux
	sudo cp aicoder /usr/local/bin/aicoder
	sudo cp aicoder.json /usr/local/bin/aicoder.json
	rm -rf aicoder

build-windows:
	GOOS=windows GOARCH=amd64 go build -o aicoder.exe .

install-linux: build-windows	
	rm -rf aicoder.exe
	