# Lillybox Backend Repository

<p align="center">
    <a href="https://go.dev/"><img alt=golang height="100" src="https://raw.githubusercontent.com/rfyiamcool/golang_logo/master/svg/golang_3.svg"/></a>
    <a href="https://gofiber.io">
    <img alt="Fiber" height="100" src="https://raw.githubusercontent.com/gofiber/docs/master/static/fiber_v2_logo.svg">
  </a>
    <a href="https://42seoul.kr/seoul42/main/view">
    <img src="https://ipfs.io/ipfs/QmNxXEhanbadVVaELQbhXCokDAEPQZF6e2AuYKQvTHoe83?filename=42.svg" alt="42logo" height="100" />
  </a>
</p>


## Domains

### 1. HTTP server for API

### 2. Batch server for background process

### 3. Chat server for users in live streaming video

## 🚀 How to Run

```
make
```



### Run http

```
build/httpd
```

### Run batch

```
build/batchd
```

### Run Chat

#### Chat is not implemented yet :(

## Project Layout

### [Inspired from...](https://github.com/golang-standards/project-layout)

## Directories

### `/cmd`

Main applications for this project.

> Chat application is under development.

### `/internal`

Private application and library code. 

> All directories in `/internal` except `/internal/database/` are used by applications that match the name.
> `/internal/database` directory is used by all applications that have to access the database.


### `/log`

Directory for log files. Each application put their access logs and error logs in this directory.

### `/docs`

Documentations for Swagger are in here.

### `/build`

Execution files that made by Makefile will be placed in this directory. 

### `/contracts`

The solidity code deployed in Baobab testnet and the go file that compiles the solidity file to Golang are located in this directory.
