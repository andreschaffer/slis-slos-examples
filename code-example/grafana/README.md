# grafana

## Deploying locally with Kubernetes
```
kubectl apply -f deployment/.
kubectl port-forward service/grafana-service 3000 -n grafana
curl localhost:3000
```

## Setup
- Login with admin:admin
- Add Prometheus Data Source with URL http://prometheus-service.prometheus.svc.cluster.local:9090
- Import Dashboard [SLOs Ping](dashboard/slos-ping.json)
