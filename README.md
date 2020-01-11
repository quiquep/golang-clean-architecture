## Golang / Go: Implementing Clean Architecture in our REST API

## Install Golang - Linux

```bash
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz

vim ~/.profile

(append :/usr/local/go/bin to PATH)
```

## Install Mux library

```bash
go get github.com/gorilla/mux
```

## Install Chi library
```bash
go get github.com/go-chi/chi
```

## Install Firestore library

```bash
go get cloud.google.com/go/firestore
```

## Export Environment variable 

```bash
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private-key.json'
```

## How to get the private key JSON file:
## From the Firebase Console: Project Overview -> Project Settings -> Service Accounts -> Generate new private key

## Build

```bash
go build
```

## Run

```bash
go run .
```

```bash
go run *.go
```