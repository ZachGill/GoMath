# GoMath

### GoMath is a Web API for answering simple math problems.

It takes two input numbers (integer OR decimal) and performs one of the following operations on them:

* Addition
* Subtraction
* Multiplication
* Division

### Running Locally:

1. Navigate to the root folder of the repository
2. Initialize submodules:
```
git submodule update --init
```
3. Build the binary:
```
go build ./cmd/math/
```
4. Run the binary:
```
./math
```

The service will be accessible at `localhost:8080`

See the API documentation (api.apib or the root web page of the service) for usage information