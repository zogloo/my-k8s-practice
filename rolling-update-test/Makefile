all: build push  clean
.PHONY: build push  clean

TAG = v1

# Build for linux amd64
build:
	GOOS=linux GOARCH=amd64 go build -o hello main.go
	docker build -t  zogloo/hello:${TAG} .

# Push to tenxcloud
push:
	docker push zogloo/hello:${TAG}

# Clean 
clean:
	rm -f hello
