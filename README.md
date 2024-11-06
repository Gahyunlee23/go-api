# Main Admin Api

- Implemented clean architecture patterns (Handler/Service/Repository) in Go
- Designed and developed RESTful APIs using Gin framework
- Achieved 40% reduction in API response times through optimized routing and middleware
- Integrated Swagger documentation using go-swag for API standardization
- Implemented comprehensive error handling middleware and logging system


## Table of Contents
1. [Installation](#installation)
2. [Usage](#usage)
3. [Configuration](#configuration)
4. [Development](#development)
5. [Testing](#testing)
6. [Contributing](#contributing)
7. [License](#license)

## Installation

### Prerequisites
- Go version `1.xx` or higher
- Database (if applicable, e.g., MySQL, PostgreSQL, etc.)
- Other dependencies (e.g., Gin, GORM, etc.)

### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/newprintgit/main-admin-api
    ```
2. Navigate to the project directory:
    ```bash
    cd main-admin-api
    ```
3. Install Go dependencies:
    ```bash
    go mod tidy
    ```
4. Set up environment variables (it's located in config directory):
    ```bash
    cp config/config.example.yaml config/config.yaml
    ```

## Usage

## development

### Running the Project
To run the server locally:
```bash
go run main.go
```

To access to the Jenkins:

http://localhost:{your_port}/swagger/index.html
