### Code Generation
* create server.cfg.yaml
* create generate.go file
* go get github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen
* mkdir api
* go generate
* go mod tidy

### Wiring the Code
* create api_func.go
* go get github.com/oapi-codegen/nethttp-middleware
* create main.go
