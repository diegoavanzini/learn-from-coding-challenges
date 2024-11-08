all: clean build test

build: 
	go build -o ./target/ccwc.exe

test:
	go test ./...

format:
	go fmt .

clean:
	rm -rf ./target/*