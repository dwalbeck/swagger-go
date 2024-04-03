# Swaggo

Supports up to version Swagger 2.0.  The Golang chi support utilizes the package **http-swagger**.

### Overview

Swaggo uses annotations placed in the code that describe the various swagger definitions. Those annotations 
are then parsed and the swagger doc is generated from those annotations.  In other words, it is 100% up to 
the programmer to insure that any change in the code is also updated in the annotations.


## Declarative Comments Format

### General API Info
| Annotation                                                              | description                                                                                                                                                                         | example                                                 |
|-------------------------------------------------------------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:--------------------------------------------------------|
| title	                                                                  | Required. The title of the application.                                                                                                                                             | 	// @title Swagger Example API                          |
| version                                                                 | Required. Provides the version of the application API.                                                                                                                              | // @version 1.0                                         |
| description                                                             | 	A short description of the application.                                                                                                                                            | 	// @description This is a sample server celler server. |
| tag.name                                                                | 	Name of a tag.                                                                                                                                                                     | 	// @tag.name This is the name of the tag               |
| tag.description                                                         | 	Description of the tag                                                                                                                                                             | 	// @tag.description Cool Description                   |
| tag.docs.url                                                            | 	Url of the external Documentation of the tag                                                                                                                                       | 	// @tag.docs.url https://example.com                  |
| tag.docs.description                                                    | 	Description of the external Documentation of the tag                                                                                                                               | 	// @tag.docs.description Best example documentation   |
| termsOfService                                                          | 	The Terms of Service for the API.                                                                                                                                                  | 	// @termsOfService http://swagger.io/terms/           |
| contact.name                                                            | 	The contact information for the exposed API.                                                                                                                                       | 	// @contact.name API Support                          |
| contact.url                                                             | 	The URL pointing to the contact information. MUST be in the format of a URL.                                                                                                       | 	// @contact.url http://www.swagger.io/support         |
| contact.email                                                           | 	The email address of the contact person/organization. MUST be in the format of an email address.                                                                                   | 	// @contact.email support@swagger.io                                         |
| license.name                                                            | 	**Required**. The license name used for the API.                                                                                                                                   | 	// @license.name Apache 2.0                                                                      |
| license.url                                                             | 	A URL to the license used for the API. MUST be in the format of a URL.                                                                                                             | 	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html                                  |
| host                                                                    | 	The host (name or ip) serving the API.                                                                                                                                             | 	// @host localhost:8080                                                                          |
| BasePath                                                                | 	The base path on which the API is served.                                                                                                                                          | 	// @BasePath /api/v1                                                                             |
| accept                                                                  | 	A list of MIME types the APIs can consume. Note that Accept only affects operations with a request body, such as POST, PUT and PATCH. Value MUST be as described under Mime Types. | 	// @accept json                                                                                  |
| produce                                                                 | 	A list of MIME types the APIs can produce. Value MUST be as described under Mime Types.                                                                                            | 	// @produce json                                                                                                                                                                   |
| query.collection.format                                                 | 	The default collection(array) param format in query,enums:csv,multi,pipes,tsv,ssv. If not set, csv is the default.                                                                 | 	// @query.collection.format multi                                                                                                                                                  |
| schemes                                                                 | 	The transfer protocol for the operation that separated by spaces.                                                                                                                  | 	// @schemes http https                                                                                                                                                             |
| externalDocs.description                                                | 	Description of the external document.                                                                                                                                              | 	// @externalDocs.description OpenAPI                                                                                                                                               |
| externalDocs.url                                                        | 	URL of the external document.                                                                                                                                                      | 	// @externalDocs.url https://swagger.io/resources/open-api/                                                                                                                        |
| x-name | 	The extension key, must be start by x- and take only json value                                                                                                                    | 	// @x-example-key {"key": "value"}                                                                                                                                                 |

### API Operation

| Annotation                                                    | description                                       |
|---------------------------------------------------------------|:--------------------------------------------------|
| description | 	A verbose explanation of the operation behavior. |
| description.markdown | 	A short description of the application. The description will be read from a file. E.g. @description.markdown details will load details.md |
| id | 	A unique string used to identify the operation. Must be unique among all API operations. |
| tags | 	A list of tags to each API operation that separated by commas. |
| summary | 	A short summary of what the operation does. |
| accept | 	A list of MIME types the APIs can consume. Note that Accept only affects operations with a request body, such as POST, PUT and PATCH. Value MUST be as described under Mime Types. |
| produce | 	A list of MIME types the APIs can produce. Value MUST be as described under Mime Types. |
| param | 	Parameters that separated by spaces. param name,param type,data type,is mandatory?,comment attribute(optional) |
| security | 	Security to each API operation. |
| success | Success response that separated by spaces. return code or default,{param type},data type,comment |
| failure | 	Failure response that separated by spaces. return code or default,{param type},data type,comment |
| response | 	As same as success and failure |
| header | 	Header in response that separated by spaces. return code,{param type},data type,comment |
| router | 	Path definition that separated by spaces. path,[httpMethod] |
| deprecatedrouter | 	As same as router, but deprecated. |
| x-name | 	The extension key, must be start by x- and take only json value. |
| x-codeSample | 	Optional Markdown usage. take file as parameter. This will then search for a file named like the summary in the given folder. |
| deprecated | 	Mark endpoint as deprecated. |

### Param Type
- query
- path
- header
- body
- formData

### Data Type
- string (string)
- integer (int, uint, uint32, uint64)
- number (float32)
- boolean (bool)
- file (param data type when uploading)
- user defined struct

### Security

| Annotation       | description    | parameters | example  |
|------------------|:---------------|:-----------|:---------|
| securitydefinitions.basic |	Basic auth. |	 |	// @securityDefinitions.basic BasicAuth |
| securitydefinitions.apikey |	API key auth. |	in, name, description |	// @securityDefinitions.apikey ApiKeyAuth |
| securitydefinitions.oauth2.application |	OAuth2 application auth. |	tokenUrl, scope, description |	// @securitydefinitions.oauth2.application OAuth2Application |
| securitydefinitions.oauth2.implicit |	OAuth2 implicit auth. |	authorizationUrl, scope, description |	// @securitydefinitions.oauth2.implicit OAuth2Implicit |
| securitydefinitions.oauth2.password |	OAuth2 password auth. |	tokenUrl, scope, description |	// @securitydefinitions.oauth2.password OAuth2Password |
| securitydefinitions.oauth2.accessCode |	OAuth2 access code auth. |	tokenUrl, authorizationUrl, scope, description |	// @securitydefinitions.oauth2.accessCode OAuth2AccessCode |

| parameters annotation | description                                       |
|-----------------------|:--------------------------------------------------|
| in |	// @in header |
| name |	// @name Authorization |
| tokenUrl |	// @tokenUrl https://example.com/oauth/token |
| authorizationurl |	// @authorizationurl https://example.com/oauth/authorize |
| scope.hoge |	// @scope.write Grants write access |
| description |	// @description OAuth protects our entity endpoints |

### Available

| Field Name       | Type | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
|------------------|:----:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| validate         |	string | 	Determines the validation for the parameter. Possible values are: required,optional.                                                                                                                                                                                                                                                                                                                                                                                     |
| default          |	* | 	Declares the value of the parameter that the server will use if none is provided, for example a "count" to control the number of results per page might default to 100 if not supplied by the client in the request. (Note: "default" has no meaning for required parameters.) See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-6.2. Unlike JSON Schema this value MUST conform to the defined type for this parameter.                       |
| maximum          |	number | 	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.1.2.                                                                                                                                                                                                                                                                                                                                                                                       |
| minimum          |	number | 	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.1.3.                                                                                                                                                                                                                                                                                                                                                                                       |
| multipleOf       |	number | 	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.1.1.                                                                                                                                                                                                                                                                                                                                                                                       |
| maxLength        |	integer | 	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.1.                                                                                                                                                                                                                                                                                                                                                                                       |
| minLength        |	integer | 	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.2.                                                                                                                                                                                                                                                                                                                                                                                       |
| enums            |	[*] | 	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.5.1.                                                                                                                                                                                                                                                                                                                                                                                       |
| format           |	string | 	The extending format for the previously mentioned type. See Data Type Formats for further details.                                                                                                                                                                                                                                                                                                                                                                       |
| collectionFormat |	string | 	Determines the format of the array if type array is used. Possible values are:  **csv** - comma separated values foo,bar.  **ssv** - space separated values foo bar.  **tsv** - tab separated values foo\tbar.  **pipes** - pipe separated values foo bar.  **multi** - corresponds to multiple parameter instances instead of multiple values for a single instance foo=bar&foo=baz. This is valid only for parameters in "query" or "formData".  Default value is csv. |
| example          |	* | 	Declares the example for the parameter value                                                                                                                                                                                                                                                                                                                                                                                                                             |
| extensions       |	string | 	Add extension to parameters.                                                                                                                                                                                                                                                                                                                                                                                                                                             |

### Future

| Field Name       | Type | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
|------------------|:----:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| pattern |	string |	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.3. |
| maxItems |	integer |	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.3.2. |
| minItems |	integer |	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.3.3. |
| uniqueItems |	boolean |	See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.3.4. |







