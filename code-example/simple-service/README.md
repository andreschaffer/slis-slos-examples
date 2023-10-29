# simple-service

## Deploying locally with Kubernetes
```
docker build -t simple-service -f build/Dockerfile .
kubectl apply -f deployment/.
kubectl port-forward service/simple-service 8888 8889 -n simple-service
curl "localhost:8888/ping?fail=false&sleep=0ms"
```

## Accessing through Traefik
```
curl "localhost:8081/simple-service/ping?fail=false&sleep=0ms"
```
