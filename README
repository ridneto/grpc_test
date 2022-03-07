# fc2-grpc

Teste básico aplicando grpc com golang;

Para rodar local: go run cmd/server/server.go

Debug com evans (https://github.com/ktr0731/evans):
- Rodar ->  ```bash
evans -r --host localhost --port 50051
``` 

Garantir que o modo reflection esteja setado:
    reflection.Register(grpcServer)

- ou Rodar client com:
    go run cmd/client/client.go

--------------------------------------------------------------------------

Compilar alterações classes proto

```bash
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```