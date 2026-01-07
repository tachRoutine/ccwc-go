build:
	go build -o ccwc

run:
	go run main.go test.txt

test:
	go test -v ./...

clean:
	rm -f ccwc
