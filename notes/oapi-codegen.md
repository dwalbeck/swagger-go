## OApi-Codegen

https://github.com/deepmap/oapi-codegen

### Installation

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

### Usage
```bash
# server.cfg.yaml
package: main
output: api/main.go
output-options:
    skip-prune: true
generate:
    embedded-spec: true
    strict-server: true
    models: true
    chi-server: true # compatible with net/http

oapi-codegen -package order-upload-oapi_codegen swagger3.yaml > order-upload.gen.go

oapi-codegen --config server.cfg.yaml swagger3.yaml > order-upload.gen.go

oapi-codegen --config server.cfg.yaml swagger3.yaml 

oapi-codegen -generate chi-server,types,spec -config server.cfg.yaml -o api swagger3.yaml
```














