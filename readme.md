# Metrics, Traces and Logging using OpenTelemetry

## Metrics using Prometheus

```
sudo docker run \
  --rm \
  --name prometheus \
  -d \
  -v "${PWD}/prometheus.yml:/prometheus.yml" \
  -p 9090:9090 \
  prom/prometheus
```

Then open http://localhost:9090/

## Tracing using Jaeger

```
sudo docker run --rm --name jaeger -p 16686:16686 -p 14268:14268 jaegertracing/all-in-one
```

Then open http://localhost:14268/api/traces

Then open http://localhost:16686/search


## Runtime app

Runtime home        http://localhost:1323/

Runtime child span  http://localhost:1323/child-span

Runtime metrics     http://localhost:1323/metrics
