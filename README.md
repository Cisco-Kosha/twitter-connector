# Kosha Twitter Connector

Use the Twitter API to listen to and analyze the public conversation, engage with people on Twitter, and innovate.


![Twitter](images/twitter.png)

This Connector API exposes REST API endpoints to perform any operations on Twitter v2 API in a simple, quick and intuitive fashion.

It describes various API operations, related request and response structures, and error codes.

## Build

To build the project binary, run
```
    go build -o main .

```

## Run locally

To run the project, simply provide env variables to supply the twitter api key, api key secret, access token and access token secret.


```bash
go build -o main .
API_KEY=<apikey> API_KEY_SECRET=<apikey-secret> ACCESS_TOKEN=<token> ACCESS_TOKEN_SECRET=<token-secret> ./main
```

This will start a worker and expose the API on port `8005` on the host machine

Swagger docs is available at `https://localhost:8005/docs`

## Generating Swagger Documentation

To generate `swagger.json` and `swagger.yaml` files based on the API documentation, simple run -

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g main.go --parseDependency --parseInternal
```

To generate OpenAPISpec version 3 from Swagger 2.0 specification, run - 

```bash
npm i api-spec-converter
npx api-spec-converter --from=swagger_2 --to=openapi_3 --syntax=json ./docs/swagger.json > openapi.json
```
