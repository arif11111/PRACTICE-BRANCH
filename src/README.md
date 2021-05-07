# go-calc

go-calc is a simple 2 operand calculator over http written in go-lang. Please fork this repo and add the solutions to your fork and raise a PR once you are done. We will not merge it but review it on github.

## Tasks

1. Build a docker image for the application and publish it
1. Deploy the application in a kubernetes cluster - minikube will be fine. Deployment technique should be extendible to multiple environments.
1. Bonus: Build the missing functionality
1. Anything extra is always welcome (If you could use Helm)


## How to

### Run Local

```
make run_local
```

### Run as a binary

```
make build
./build/go-calc-darwin ## For Mac
./build/go-calc ## For Linux
```

## Sample inputs for current implementation

```
$ curl -X POST http://localhost:8080/add -d '{"operand1": 1, "operand2": 2}'
$ 3
$ curl -X POST http://localhost:8080/subtract -d '{"operand1": 4, "operand2": 3}'
$ 1
$ curl -X POST http://localhost:8080/multiply -d '{"operand1": 2, "operand2": 3}'
$ 6
$ curl -X GET http://localhost:8080/status
$ ok
```
