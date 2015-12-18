

VERSION = 1.0

bindata-dev:
	go-bindata --dev data/

bindata-build: 
	go-bindata data/

run:
	FILESERVER_DIR=/tmp FILESERVER_PORT=:9000 \
	go run auth.go bindata.go commands.go dirjson.go dirsearch.go \
			dirzip.go gzip.go main.go serv_statics.go webcommand.go


build: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

build_image: 
	docker build -t jordic/file_server:$(VERSION) .

push: 
	docker push jordic/file_server:$(VERSION)

all: bindata-build build build_image
