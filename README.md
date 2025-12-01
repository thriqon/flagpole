# Flagpole

A lightweight, dependency-free HTTP server written in Go that serves responses entirely from a YAML configuration file.
Define routes quickly, mock APIs instantly, and prototype without writing any extra code.

## ✨ Features

📝 YAML-based routing — define endpoints using simple declarative config

⚡ Fast and lightweight — built with Go’s standard library

📦 Serve plain text, JSON, HTML, or any MIME type

🚀 Perfect for mocks, demos, and rapid prototyping

🔒 Predictable behavior — always responds with HTTP 200 OK

📄 Example Configuration

```yaml
- route: GET /
    contentType: text/plain
    contents: |
      Hello, World!
```

## 🛠️ Installation

Clone the repository:

```
git clone https://github.com/thriqon/flagpole.git
cd flagpole
```

Build the binary:

```
go build -o flagpole
```

## ▶️ Usage

Run the server with a YAML config:

./flagpole -config-file routes.yaml

Default address is:

<http://localhost:8080/>

To specify a custom port:

```
./flagpole -config-file routes.yaml -listen-address :9000
```

## 📚 Configuration Reference

Each route entry supports:

| Field | Description | Required |
|-------|-------------|----------|
| route|HTTP method and path (e.g., GET /hello)|Yes |
| contentType|MIME type returned in the Content-Type header|Yes|
| contents|Response body returned to the client|Yes|

Note:

All responses always return status code 200

No dynamic request handling — responses are literal from YAML.

Example with multiple routes:

```yaml
- route: GET /
    contentType: text/plain
    contents: |
      Welcome!
- route: GET /api/info
    contentType: application/json
    contents: |
      { "name": "example", "version": 1 }
- route: POST /submit
    contentType: text/plain
    contents: OK
```

## 🧪 Testing

Use curl:

curl -i <http://localhost:8080/>

Check JSON route:

curl -i <http://localhost:8080/api/info>

## 🤝 Contributing

Pull requests are welcome!
For significant changes, please open an issue to discuss what you’d like to improve.

## 📜 License

MIT. See LICENSE for details.
