default:
	@echo "Please specify a target. Available targets are:"

build:
	go build .

install: build
	sudo cp aicoder /usr/local/bin/aicoder
	sudo cp aicoder.json /usr/local/bin/aicoder.json
	rm -rf aicoder

dist: build
	mkdir -p dist
	cp aicoder dist/aicoder
	cp aicoder.json dist/aicoder.json
	zip -r dist/aicoder.zip dist/aicoder dist/aicoder.json
	rm -f aicoder
	rm -f dist/aicoder
	rm -f dist/aicoder.json

build-windows:
	GOOS=windows GOARCH=amd64 go build -o aicoder.exe .

install-linux: build-windows	
	rm -rf aicoder.exe
