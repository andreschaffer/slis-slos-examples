# k6

## Simulating traffic to simple-service with k6
```
docker run --rm --add-host=host.docker.internal:host-gateway -i grafana/k6 \
    run --vus 10 --duration 10s - <traffic.js
```

Control Parameters:
- In the docker command:
  - vus: number of virtual users generating traffic concurrently
  - duration: for how long traffic will be generated
- In the traffic.js script:
  - sleep(x): sleep time (in seconds) between requests, for each virtual user
  - const failure: whether the simple-service ping request will return a failed response
  - const sleepMs: how much delay will be added to the simple-service ping response

Simulate traffic at will, tweak the control parameters above between the runs, 
and observe how the simple-service is doing with its SLOs.
