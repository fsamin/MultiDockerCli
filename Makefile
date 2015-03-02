all: multidocker


multidocker:
	go build src/desc/*.go
	go build src/cli/*.go
	go build src/*.go

clean:
	rm -f multidocker

clusterclean:
	rm -rf test/uploads*

run: multidocker
	./multidocker ps

fmt:
	go fmt src/*.go
	go fmt src/desc/*.go
	go fmt src/cli/*.go

install: multidocker
	cp -f multidocker /usr/local/bin/multidocker

test: multidocker
	go test src/

coverage: multidocker
	go test src/ -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html


install_deps:
	go get -u github.com/codegangsta/cli
	go get -u github.com/samalba/dockerclient
	go get -u github.com/crackcomm/go-clitable
