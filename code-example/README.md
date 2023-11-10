# code-example

## Step-by-step setup
- Deploy [prometheus](prometheus/README.md)
- Deploy [grafana](grafana/README.md)
- Deploy [traefik](traefik/README.md)
- Deploy [simple-service](simple-service/README.md)
- Simulate traffic with [k6](k6/README.md)
- Observe the simple-service SLOs in Grafana
- Clean up each resource with `kubectl delete -f deployment/.` when satisfied

## More automated setup
- Run `make all`
- Port-forward to Traefik web service `kubectl port-forward service/traefik-web-service 8081 -n traefik`
- Port-forward to Grafana `kubectl port-forward service/grafana-service 3000 -n grafana`
- Setup [Grafana Dashboard](grafana/README.md#Setup)
- Simulate traffic with [k6](k6/README.md)
- Observe the simple-service SLOs in Grafana
- Clean resources with `make clean` when satisfied
