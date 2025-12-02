# Flagpole

A lightweight, dependency-free HTTP server written in Go that serves responses entirely from a YAML configuration file.
Define routes quickly, mock APIs instantly, and prototype without writing any extra code.

## ✨ Features

📝 YAML-based routing — define endpoints using simple declarative config

⚡ Fast and lightweight — built with Go’s standard library

📦 Serve plain text, JSON, HTML, or any MIME type

🚀 Perfect for mocks, demos, and rapid prototyping

🔒 Predictable behavior

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

```
./flagpole -config-file routes.yaml
```

Default address is:

http://localhost:8080/

To specify a custom port:

```
./flagpole -config-file routes.yaml -listen-address :9000
```

## 📚 Configuration Reference

```
routes:
  - route: GET /hello
    headers:
      Content-type: text/plain
    body: |
      Hello, World!
  - route: GET /
    status: 307
    headers:
      Location: /hello
```

## 🧪 Testing

Use curl:

```
curl -i http://localhost:8080/
```
Check JSON route:

```
curl -i http://localhost:8080/api/info
```

## 🤝 Contributing

Pull requests are welcome!
For significant changes, please open an issue to discuss what you’d like to improve.

## 📜 License

MIT. See LICENSE for details.
