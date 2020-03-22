# Backedn API

Exposed HTTP JSON api.

## Build

```bash
make build   # builds the service
make install # builds and installs the binary into $GOROOT/bin 
```


## Run

### With Docker

```bash
make docker-build
docker run -ti --rm -p 8080:8080 api
```

### With Golang installed

```bash
go run main.go
```
