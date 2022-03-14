## RGate

A configurable HTTP API Gateway for applications running on Docker containers.
Detailed architecture can be found in [api_gateway_docker.pdf](https://github.com/olttwa/api-gateway/blob/master/api_gateway_docker.pdf).

### Features

- Configurable routes and backends
- Graceful Shutdown
- Accurate traffic stats

### Rationale behind decisions

- Chose Golang as programming language because of:
  - Docker SDK
  - Simple routing and middleware usage
  - Reverse Proxy Utility
- Encapsulation is done at package level to avoid mutations

### Assumptions

- For multiple backends, any container is randomly chosen at runtime (instead of load time)
- HTTP Method isn't taken into consideration for routing
- Mutex is used to avoid race conditions when measuring stats
- HTTP Connections to backend aren't pooled/reused
- No health-check is performed when choosing a backend

### Development

#### Instalment

```bash
make install
```

#### Running

```bash
rgate --port 8080 --config config.yml
```

#### Unit Tests

```bash
make test
```

#### Linting

```bash
make lint
```
