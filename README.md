
# Device Manager

REST-API server to store devices data and them events in database.

## API Reference
[here](./docs/swagger.html)


## Appendix

Any additional information goes here


## Run Locally

Clone the project:

```bash
  git clone https://github.com/Rustam2202/DeviceManager.git
```

Config parameters in ```cmd/config.yaml``` and start server:

```bash
  make run
```
Build app:

```bash
  make build
```

## Up in docker

To deploy this project run

```bash
  make compose
```


## Running Tests

To run tests and generate HTML coverage report 

```bash
  make test-cover-report
```


## Documentation

[Documentation](https://linktodocumentation)


## Tech Stack

- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [MongoDB](https://github.com/mongodb/mongo-go-driver)
- [Kafka](https://github.com/segmentio/kafka-go)
- [Docker Compose](https://docs.docker.com/compose/)
- [Zap Logger](https://github.com/uber-go/zap)
- Mock Tests: The project includes mock tests using [mtest](https://go.mongodb.org/mongo-driver/mongo/integration/mtest) to simulate MongoDB.
- Graceful Shutdown: The application gracefully handles shutdown signals to ensure all pending requests are completed before shutting down.
- Panic Recovery: The application recovers from panics to prevent crashes and logs detailed error information.
- [Swagger](https://github.com/swaggo/swag): An annotation-based Go library for automatically generating Swagger documentation for API endpoints.

## Usage/Examples

```javascript
import Component from 'my-project'

function App() {
  return <Component />
}
```

