

bindata-dev:
	go-bindata --dev data/

bindata-build: 
	go-bindata data/

run:
	go run auth.go bindata.go commands.go dirjson.go dirsearch.go \
			dirzip.go gzip.go main.go serv_statics.go webcommand.go




