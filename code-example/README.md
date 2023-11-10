# code-example

## Overview

             │                             │
          1  │                          2  │
             │                             │
             ▼                             ▼
      ┌──────────────┐              ┌───────────┐
      │              │              │           │
      │  k6 traffic  │              │  Grafana  │
      │              │              │           │
      └──────┬───────┘              └──────┬────┘
             │                             │
             │                             │
             ▼                             ▼
      ┌─────────────┐               ┌────────────┐
      │             │               │            │
      │   Traefik   │◄──────────────┤ Prometheus │
      │             │               │            │
      └──────┬──────┘               └──────┬─────┘
             │                             │
             │                             │
             ▼                             │
    ┌──────────────────┐                   │
    │                  │                   │
    │  simple-service  │◄──────────────────┘
    │                  │
    └──────────────────┘

The resources involved in this code example are:
- k6: used to generate traffic towards our application
- Traefik: used as load balancer for our application
- simple-service: a pretty simple application with a ping endpoint
- Prometheus: for recording our load balancer and application metrics
- Grafana: for visualising our application metrics

The metrics used for our SLIs and SLOs are the ones recorded from Traefik. Nevertheless, 
additional metrics recorded directly from the application are usually useful for troubleshooting.

There are two types of setup below, just pick the one you feel most comfortable with.

## Step-by-step setup
- Deploy [prometheus](prometheus/README.md)
- Deploy [traefik](traefik/README.md)
- Deploy [simple-service](simple-service/README.md)
- Simulate traffic with [k6](k6/README.md)
- Deploy [grafana](grafana/README.md)
- Observe the simple-service SLOs in Grafana
- Clean up each resource with `kubectl delete -f deployment/.` when satisfied

## More automated setup
- Run `make all` to deploy the resources
- Port-forward to Traefik web service `kubectl port-forward service/traefik-web-service 8081 -n traefik`
- Simulate traffic with [k6](k6/README.md)
- Port-forward to Grafana `kubectl port-forward service/grafana-service 3000 -n grafana`
- Setup [Grafana Dashboard](grafana/README.md#Setup)
- Observe the simple-service SLOs in Grafana
- Clean resources with `make clean` when satisfied
