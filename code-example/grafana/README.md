# grafana

## Deploying locally with Kubernetes
```
kubectl apply -f deployment/.
kubectl port-forward service/grafana-service 3000 -n grafana
curl localhost:3000
```

## Setup
- Add Prometheus Data Source
- Import Dashboard [Ping SLO](dashboard/ping-slo.json)
