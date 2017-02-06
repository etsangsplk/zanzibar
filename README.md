## Documentation
To learn how to document your project on [EngDocs](https://engdocs.uberinternal.com),
see this [guide](https://engdocs.uberinternal.com/engdocs/index.html).

## Some make targets:

 - `make install`; install dependencies
 - `make test`; run tests
 - `make bins`; generates binaries for local dev
 - `make examples/example-gateway/example-gateway`; generates example binary

 ## Building a service:

 - Create an endpoint definition in idl/github.com/uber/zanzibar/endpoints
 - `make generate`; build definitions and marshalling code
 - Register the endpoint in service/endpoints/register.go
 - Add example inputs and outputs for the endpoints and downstream services in test_goldens/
 - `make test`; verify new endpoint behaviour against test_goldens
