run:
	go run .

watch:
	air

build:
	go build -o bin/storer-server

clean:
	rm -rf bin