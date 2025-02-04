build:
	go build -o gwc wc.go

configure:
	go get github.com/tdewolff/argp

clean:
	rm -rf gwc

test:
	go test 
