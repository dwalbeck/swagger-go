## go-swagger

Supporting Swagger 2.0, this will generate client and server code, as well as validate and generate the swagger spec.
Code generation includes the ability for authentication and OAuth2


#### Serve specification UI
````bash
swagger serve https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
````

#### Validate a specification
````bash
swagger validate https://raw.githubusercontent.com/swagger-api/swagger-spec/master/examples/v2.0/json/petstore-expanded.json
````

#### Generate an API server
````bash
swagger generate server [-f ./swagger.json] -A [application-name [--principal [principal-name]]
````

#### Generate an API client
````bash
swagger generate client [-f ./swagger.json] -A [application-name [--principal [principal-name]]
````

#### Generate an CLI swagger Spec
````bash
swagger generate cli [-f ./swagger.json] -A [application-name [--principal [principal-name]]
````

#### Generate a spec from source
````bash
swagger generate spec -o ./swagger.json
````

#### Generate a data model
````bash
swagger generate model --spec={spec}
````
