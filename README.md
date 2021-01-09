# grpc-tutorial

- grpcサーバー作成する
  - evansからリクエスト投げる
- gorm使ってリクエスト→DB操作をやってみる

## protocコマンド

```sh
protoc --go_out=./pb/calc --go_opt=paths=source_relative \
--go-grpc_out=./pb/calc --go-grpc_opt=paths=source_relative \
proto/calc.proto
```

--go_out
calc.pb.goを生成。
Request, ResponseのMessage Typeを含む。

--go-grpc_out
calc_grpc.pb.goを生成。
client, serverのコードを含む。

## evans

```sh
evans repl --proto proto/calc.proto
# replモード
call Sum
```
