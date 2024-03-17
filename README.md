# Projeto Clean Architecture
1. Execute o comando para subir o banco de dados: ```docker-compose up```

### Subir servidor REST e GraphQL:
- Suba o servidor utilizado para as 2 tecnologias: ```go run main.go restAndGraphql```

### Subir servidor gRPC:
- Primeiro suba o grpc com o comando: ```go run main.go grpc```

- Então suba o evans para fazer a conexão das buscas: ```evans -r repl```

&nbsp;
# REST API
- Os endpoints estão no entity ```"/orders"```
1. POST ```"/new"``` criar um novo order
```json
{
  "name": "order 1",
  "total": 120.00
}
```

2. GET ```"/all"``` listar orders

&nbsp;
# GraphQL
1. Novo Order:
```graphql
mutation createOrder {
  createOrder(input: {name: "order name", total: 12.00}) {
    id
  }
}
```

2. Listar orders
```graphql
query orders {
  orders {
    id
    name
    total
  }
}
```

&nbsp;
# gRPC
- Acesse o package: ```package pb```
- Acesse o service ```service OrderService```

1. Novo Order: 
```grpc
call NewOrder
```
OBS: Digite os campos que serão solicitados: nome e total.

2. Listar orders: 
```grpc
call ListOrders
```
