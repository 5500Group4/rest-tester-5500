### Go Server Setup using Gin Framework

This guide explains how to set up and run a simple Go server using the Gin framework for basic CRUD operations on users.

### Prerequisites

Ensure you have **Go** installed on your system.

#### Installation Instructions:

##### Windows

1. **Install Go**:
   - Download and install Go from [golang.org](https://golang.org/dl/).
   - Verify the installation by running:
     ```bash
     go version
     ```

2. **Set Up Environment Variables**:
   - Ensure the `GOPATH` and `GOROOT` are correctly set (this is typically handled by the Go installer).

##### macOS

1. **Install Go**:
   - Install Go using [Homebrew](https://brew.sh/):
     ```bash
     brew install go
     ```
   - Verify the installation with:
     ```bash
     go version
     ```

2. **Set Up Environment Variables**:
   - Ensure your `GOPATH` is set, typically `$HOME/go` by default.

### Installing Dependencies

1. Navigate to your project directory.

2. Initialize a Go module:
   This command will create a `go.mod` file in our project directory. This has already been included in this repository.

   ```bash
   go mod init rest-multiple-languages
    ```

3. Install Gin and CORS:
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/gin-contrib/cors
   ```

### Running the Server


Use the following command to start the Ruby server:

```bash
go run src/servers/go-server/server.go
```

The server will run on **port 5004** by default.

### To Stop the Server

Press `Ctrl + C` in the terminal to stop the server.