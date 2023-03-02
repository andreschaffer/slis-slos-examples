# simple-service

## Building and running the application
```
go build -o ./bin/simple-service ./cmd
./bin/simple-service
curl localhost:8888/ping
```
or
```
docker build -t simple-service .
kubectl apply -f deployment/deployment.yaml
kubectl -n simple-service port-forward service/simple-service 8888 8889
curl localhost:8888/ping
```