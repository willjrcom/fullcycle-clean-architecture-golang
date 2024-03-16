# Clean Architecture
1. Execute os comandos:
- docker-compose up
- go run main.go server

### Endpoints REST
- "/new" criar um novo order
- "/orders" listar orders

### GraphQL
- mutation createOrder {
  createOrder(input: {name: "order name", total: 12.00}) {
    id
  }
}

- query orders {
  orders {
    id
    name
    total
  }
}

### gRPC
...