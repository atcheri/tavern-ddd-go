## Let's GO to a Tavern in DDD (golang based project)

### Overview

A small project in go that aims to use clean-architecture and domain-driven-design (DDD).

###

### Install dependencies

```shell
$ go mod tidy
```

### Run the app
```shell
$ go run ./cmd/main.go
```

### Run the tests
```shell
$ go test ./...
```

### Project structure
```shell
├── README.md
├── cmd
│   └── main.go
├── domain
│   ├── billing
│   ├── customer
│   │   ├── customer.go
│   │   ├── customer_test.go
│   │   └── repository.go
│   └── product
│       ├── product.go
│       ├── product_test.go
│       └── repository.go
├── entity
│   ├── item.go
│   └── person.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── db
│   │   └── memory
│   │       ├── customer
│   │       │   ├── memory.go
│   │       │   └── memory_test.go
│   │       └── product
│   │           ├── memory.go
│   │           └── memory_test.go
│   └── sender
│       └── log_sender.go
├── services
│   ├── billing_service.go
│   ├── billing_service_test.go
│   ├── order_service.go
│   ├── order_service_test.go
│   ├── tavern_service.go
│   └── tavern_service_test.go
└── valueobject
    └── transaction.go
```

### What is this project using
- [faker v3](github.com/bxcodec/faker/v3)
- [uuid](github.com/google/uuid)
- [ensure](github.com/iamkoch/ensure)
- [lo](github.com/samber/lo)
- [testify](github.com/stretchr/testify)