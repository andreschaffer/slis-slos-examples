# code-example

This is a code example of implementing SLIs and SLOs on a simple application.

## SLIs and SLOs definition
Imagine that after good discussions with the different functions in your team, you agree that the "ping" functionality your service provides is the one that is important to your users, and come up with these SLIs and SLOs:
- Availability: 95% of the "ping" requests succeed.
- Latency: 95% of the "ping" requests complete in at most 100ms.

You also define that the compliance period will be 28 rolling days, and that the measurements will be taken from the load balancer.

## System Overview

How the system looks and which tools are involved in this example are described below:

             │                             │
             │                             │
             │                             │
             ▼                             ▼
      ┌──────────────┐              ┌───────────┐
      │              │              │           │
      │      k6      │              │  Grafana  │
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

- k6: tool used to generate traffic towards our application
- Traefik: used as load balancer for our application
- simple-service: our application providing the "ping" functionality
- Prometheus: used for recording our load balancer and application metrics
- Grafana: used for visualising our metrics, and how we are doing with our SLOs

The metrics used for our SLIs and SLOs are the ones recorded from Traefik. Out-of-scope, but good to mention, 
additional metrics recorded directly from the application are usually still useful for observability and troubleshooting.

## Prerequisites for getting it up and running
- Docker
- Kubernetes

There are two types of setup provided below, just pick the one you feel most comfortable with.

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
