# JasperServer Go

## Client

This repo contains a Go client for [JasperReports Server Community Edition](http://community.jaspersoft.com/project/jasperreports-server/releases). Currently it only supports version 6.1.0.

The reason for creating this JasperServer API client in Go is to use it from a JasperServer [Terraform Provider](https://www.terraform.io/docs/providers/) that we plan to build. 

This will allow someone to version control their JasperServer setup and quickly recreate it in the spirit of [Infrastructure as Code](https://www.thoughtworks.com/insights/blog/infrastructure-code-reason-smile).

This client was automatically generated using [go-swagger](https://github.com/go-swagger/go-swagger).

## Generation Process

The process to generate this client was to:

1. Get the WADL from a running JasperServer instance (e.g. Using [our jasperserver docker container](https://github.com/retrievercommunications/docker-jasperserver)). This file is saved in this repo under {version}/application.wadl 

2. Convert from WADL to Swagger 2.0 JSON using [API Transformer website](https://apitransformer.com/). This file is saved in this repo under {version}/converted.json

3. Generate the Go client using [go-swagger](https://github.com/go-swagger/go-swagger)

```
swagger generate client [-f {version}/swagger.json] -A [application-name [--principal [principal-name]]
```

E.g. 
```
Run `swagger generate client -f 6.1.0/converted.json -A jasperserver`
```


## TODOs

* Support multiple versions of the API, currently it only supports JasperServer version 6.1.0
* Include some sample code of how to use this client. [See goswagger example](https://goswagger.io/generate/client/)