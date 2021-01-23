# RGate

A configurable HTTP API Gateway for applications running on Docker containers.

### Features

- Configurable routes and backends
- Graceful Shutdown
- Traffic stats

### Rationale behind decisions

- Golang because of docker SDK, stats middleware, reverseproxy utility,
- encapsulation is performed at package level to avoid mutations

### Assumptions

- For multiple backends, any container is randomly chosen at runtime, not during initialization
- No health-check is performed when choosing a backend
- HTTP Method isn't taken into consideration for routing

### Development

##### Instalment
```
make install
```

##### Running
```
make run
```

##### Unit Tests
```
make test
```

##### Linting
```
make lint
```