# Health Checker CLI Tool

A simple command-line tool to check the health of a given URL.

## Usage

### Health Check Site
Run the tool with the following command:

```bash
go run main.go -url <URL>
```
Replace <URL> with the URL you want to check.

Healthy Example:
```bash
go run main.go -url https://www.example.com/
```

### Health Check Port
Run the tool with the following command:

```bash
go run main.go -port <PORT> -host <HOST>
```
Replace <PORT> with the port you want to check.
Replace <HOST> with the host you want to check.

Healthy Example:
```bash
go run main.go -port 80 -host example.com
```


## Test Health Checker
Navigate to inside of healthchecker directory and run:
```
go test -v
```
