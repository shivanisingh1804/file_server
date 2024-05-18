# static_server
# Simple Go HTTP File Server

This project demonstrates a simple HTTP file server written in Go that serves files from the current working directory.

## Project Structure

- `main.go`: The Go source code for the HTTP server.

## Prerequisites

- Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

## Getting Started

1. **Clone the repository or create the project directory**:

    ```sh
    mkdir project-directory
    cd project-directory
    ```

2. **Create the Go file**:

    - `main.go`:

    ```go
    package main

    import (
        "net/http"
        "os"
    )

    func main() {
        dir, _ := os.Getwd()
        http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
    }
    ```

3. **Run the server**:

    ```sh
    go run main.go
    ```

4. **Access the server in a web browser**:

    - Open [http://localhost:3000/](http://localhost:3000/) to see the content of the current working directory being served.


