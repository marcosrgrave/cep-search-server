# Go server that consumes external API to get CEP info (Brazilian Zip Code)

## Go fundamentals applied in this project

- HTTP Request
- Response with JSON Body and Status Code
- URL Parameters
- Server, Handlers and Controllers Folder Structure

## How to use it

1. Run the server:
    ```bash
    go run main.go
    ```
2. Do a CEP Search on your browser using URL parameter "cep":
    ```
    http://localhost:8000/?cep=88034500
    ```