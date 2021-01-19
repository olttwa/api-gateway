# RGate

A configurable HTTP API Gateway for applications running on Docker containers.

### Features

- Configurable routes and backends
- Graceful Shutdown
- Traffic stats

### Rationale behind decisions

- Golang because of docker SDK, stats middleware, reverseproxy utility,

### Assumptions

- A backend container is chosen during intialization, not during runtime.

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
