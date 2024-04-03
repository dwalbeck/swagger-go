## OpenAPI-Generator

### Installation
```bash
npm install @openapitools/openapi-generator-cli -g

  or

wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.2.0/openapi-generator-cli-7.2.0.jar \
    -O openapi-generator-cli.jar
    
  or 

docker run --rm \
  -v ${PWD}:/local openapitools/openapi-generator-cli generate \
  -i /local/petstore.yaml \
  -g go \
  -o /local/out/go
```

### Generate Server Code

```bash
vi config.yaml

router: chi
outputAsLibrary: false
onlyInterfaces: false
serverPort: 8000
sourceFolder: go
packageName: openapi


openapi-generator-cli generate -g go-server -o ./api -i swagger.yaml -c config.yaml
  
  or

openapi-generator-cli generate -g nodejs-express-server -o ./api -i swagger.yaml -c config.yaml

  or

openapi-generator-cli generate -g php-laravel -o ./api -i swagger.yaml -c config.yaml
```

https://openapi-generator.tech/docs/generators















