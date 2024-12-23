# My Go Web Server

This is a simple web server application built using Go (Golang). The application serves as a basic template for creating web applications with Go.

## Project Structure

```
my-go-web-server
├── src
│   ├── main.go          # Entry point of the application
│   ├── handlers         # Contains HTTP request handlers
│   │   └── handlers.go  # Defines handler functions
│   ├── routes           # Contains route definitions
│   │   └── routes.go    # Sets up application routes
│   └── models           # Contains data models
│       └── models.go    # Defines data structures
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Getting Started

To run the application, follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   cd my-go-web-server
   ```

2. Install the necessary dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run src/main.go
   ```

4. Open your web browser and navigate to `http://localhost:8080` to see the application in action.

## Features

- Basic web server setup
- Route handling with dedicated handler functions
- Modular structure for easy maintenance and scalability

## License

This project is licensed under the MIT License. See the LICENSE file for more details.