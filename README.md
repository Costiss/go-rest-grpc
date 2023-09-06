# HTTP + gRPC Server

This is a server that can be called both via HTTP (using Gin Gonic Framework) or gRPC

    .
    ├── grpc                    # gRPC Server implementation
    ├── internal                # Common internal files
    │   └── service             # Service that can be called both by HTTP & gRPC server
    ├── proto                   # Protobuffs definitions
    ├── rest                    # REST server implementation (router, controllers...)
    ├── mian.go                 # App entrypoint
    └── README.md
