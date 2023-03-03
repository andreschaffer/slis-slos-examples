# simple-service

## Deploying locally with Kubernetes
```
docker build -t simple-service .
kubectl apply -f deployment/deployment.yaml
kubectl port-forward service/simple-service 8888 8889 -n simple-service
curl localhost:8888/ping
```