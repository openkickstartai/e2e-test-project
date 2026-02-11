# E2E Test Project

A comprehensive end-to-end testing framework built with Go and Vue.js. This project demonstrates automated testing practices and CI/CD integration.

## Features

- ✅ HTTP API endpoints with health checks
- ✅ Comprehensive unit tests with realistic scenarios
- ✅ Ready for containerization and deployment
- ✅ Clean architecture following Go best practices

## Getting Started

### Prerequisites
- Go 1.19+

### Installation
```bash
git clone https://github.com/openkickstartai/e2e-test-project.git
cd e2e-test-project
go mod init e2e-test
go run src/main.go
```

### Running Tests
```bash
go test ./test/...
```

### API Endpoints

#### GET /
Returns a greeting message.

#### GET /health
Returns health status and version information.

```json
{
  "status": "healthy",
  "version": "1.0.0"
}
```

## Project Structure

```
.
├── src/
│   └── main.go          # HTTP server and handlers
├── test/
│   └── main_test.go       # Unit tests
└── README.md              # This file
```

## Usage Examples

### Start the server
```bash
go run src/main.go
# Server running on http://localhost:8080
```

### Test the API
```bash
# Test the main endpoint
curl http://localhost:8080/
# Response: Hello, E2E Test!

# Test health check
curl http://localhost:8080/health
# Response: {"status":"healthy","version":"1.0.0"}
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Run the test suite
5. Submit a pull request

## License

MIT License - see LICENSE file for details.