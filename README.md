# JasperServer Go

## Client

This repo contains a Go client for the JasperReports Server Community Edition's v2 REST API. Releases of JasperReports can be found [here](http://community.jaspersoft.com/project/jasperreports-server/releases) and documentation for the v2 REST API are [here](http://community.jaspersoft.com/documentation/tibco-jasperreports-server-web-services-guide/v610/introduction-0). Currently this Go client only supports version 6.1.0 of JasperReports. 

The reason for creating this JasperServer API client in Go is to use it from a JasperServer [Terraform Provider](https://www.terraform.io/docs/providers/) that we plan to build. 

This will allow someone to version control their JasperServer setup and quickly recreate it in the spirit of [Infrastructure as Code](https://www.thoughtworks.com/insights/blog/infrastructure-code-reason-smile).

This client was automatically generated using [go-swagger](https://github.com/go-swagger/go-swagger) version 0.5.0.

## Generation Process

The process to generate this client was to:

1. Get the WADL from a running JasperServer instance (e.g. Using [our jasperserver docker container](https://github.com/retrievercommunications/docker-jasperserver)).

Example URL: http://localhost:8080/jasperserver/rest_v2/application.wadl

This file is saved in this repo under {version}/application.wadl 

2. Convert from WADL to Swagger 2.0 JSON using [API Transformer website](https://apitransformer.com/). This file is saved in this repo under {version}/converted.json

3. Generate the Go client using [go-swagger](https://github.com/go-swagger/go-swagger)

```
swagger generate client [-f {version}/swagger.json] -A [application-name [--principal [principal-name]]
```

E.g. 
```
Run `swagger generate client -f 6.1.0/converted.json -A jasperserver`
```

Notes:
Unfortunately the swagger-go client generation process wasn't as easy as hoped for:
* Had to vendor the version 0.5.0 packages which are moving in the upcoming 0.6.0 version to github.com/go-openapi
* Had to remove unused imports, mainly in _responses.go files, including a reference to a models package in this repo
* Had to add extra backslash to \d other compilation failed with "unknown escape sequence: d"
* Had to rename a parameter called 'type' which is forbidden
* Had to ensure that the jasperserver_client.go has the right base path of "/jasperserver/rest_v2/" other you get an unhelpful "no consumer" error

## TODOs

* Support multiple versions of the API, currently it only supports JasperServer version 6.1.0
* Include some sample code of how to use this client. [See go-swagger example](https://goswagger.io/generate/client/)