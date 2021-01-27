
Prometheus
==

> about [prometheus](https://prometheus.io/)

Run
--

- `cd grapes/services/metric/exporter`

- `cp ./prometheus.example.yml prometheus.yml`
     
- `docker run -d -p 9090:9090 -v ${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus`
