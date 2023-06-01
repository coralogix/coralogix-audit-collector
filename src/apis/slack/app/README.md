# Slack Util

This is a utility to install the Slack app.

## Installation

```bash
export SLACK_CLIENT_ID=""
export SLACK_CLIENT_SECRET=""
export LISTEN_ADDRESS=":8433"
export CERT_FILE=""
export CERT_KEY_FILE=""
go run main.go
```

You can use [mkcert](https://github.com/FiloSottile/mkcert) to generate a self-signed certificate.

```bash
$ mkcert -install
Created a new local CA 💥
The local CA is now installed in the system trust store! ⚡️
The local CA is now installed in the Firefox trust store (requires browser restart)! 🦊

$ mkcert localhost 127.0.0.1

Created a new certificate valid for the following names 📜
 - "localhost"
 - "127.0.0.1"

The certificate is at "./localhost+5.pem" and the key at "./localhost+5-key.pem" ✅

$ export CERT_FILE="./localhost+5.pem"
$ export CERT_KEY_FILE="./localhost+5-key.pem"
...
$ go run main.go
```

Your token will be printed to the console.
