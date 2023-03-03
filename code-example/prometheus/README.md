# prometheus

## Deploying locally with Kubernetes
```
kubectl apply -f .
kubectl port-forward service/prometheus-service 9090 -n prometheus
curl localhost:9090
```

## Exposed url for connecting to Prometheus from within the Kubernetes cluster
http://prometheus-service.prometheus.svc.cluster.local:9090
