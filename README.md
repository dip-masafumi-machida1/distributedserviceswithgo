# Distributed Services with Go
Go言語による分散サービスのハンズオンリポジトリ<br>

## 環境構築
- Goのインストール
  - https://go.dev/doc/install
- protobufのインストール
  - https://grpc.io/docs/protoc-installation/
  - `protoc --version`を実行してエラーにならなければOK
  - `protoc`コマンドでスキーマ生成した際に`Please specify a program using absolute path or make sure the program is available in your PATH system variable`が表示される場合
  - `protoc-gen-go`が認識されていない or インストールされていない可能性あり
      - 以下を参考にインストールする
      - https://grpc.io/docs/languages/go/quickstart/

## その他
パッケージのダウンロード
```
go mod tidy
```
