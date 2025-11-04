# httpshare

**httpshare** is a simple web-based file explorer written in Go. It allows you to quickly share and browse files over HTTP with minimal setup.

## Features

- Lightweight and easy to use
- Simple web interface for browsing files
- Configurable port

## Installation

Make sure you have [Go](https://golang.org/dl/) installed. Then, run:

```bash
go install github.com/isa0-gh/httpshare@latest
````

This will install the `httpshare` binary in your Go `bin` directory (usually `$GOPATH/bin`).

## Usage

Run the command below to start the file explorer on a specific port (default: 8080):

```bash
httpshare --port 8080
```

Then, open your web browser and navigate to:

```
http://localhost:8080
```

You will see a simple web interface to explore and download files in the current directory.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License.


