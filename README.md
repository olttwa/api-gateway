# RGate

A configurable HTTP API Gateway for applications running on Docker containers.

### Features

- Configurable routes and backends
- Graceful Shutdown
- Traffic stats

### Rationale behind decisions

- Golang because of docker SDK, stats middleware, reverseproxy utility,

### Assumptions

- For multiple backends, any container is randomly chosen at runtime, not during initialization
- No health-check is performed when choosing a backend

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
