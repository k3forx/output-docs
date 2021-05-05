# Prometheus

## Kubernetes monitoring with Prometheus

1. Pull target auto discovery from Kubernetes API
1. Pull autodiscovered microservice metrics
1. Push metrics alerts to alertmanager

## Metrics

- Counter: Start from 0, increased monotonically
- Gauge: A single numerical value that can arbitrarily go up and down
- Summary: Samples observations (usually things like request durations and response sizes). While it also provides a total count of observations and a sum of all observed values, it calculates configurable quantiles over a sliding time window.
- Histogram: Samples observations (usually things like request durations or response sizes) and counts them in configurable buckets. It also provides a sum of all observed values.

## PromQL

- `up`
- `sum(increase(fastapi_requests_total[5m]))`
- `histogram_quantile(0.99, sum(rate(http_request_duration_highr_seconds_bucket[10m])) by (le))`
- `sum(rate(container_network_receive_errors_total[5m]))`
- `sum(rate(container_network_transmit_errors_total[5m]))`

# Demo

You can refer to https://github.com/k3forx/fastapi-example/tree/master/k8s/prometheus
