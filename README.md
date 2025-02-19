# TODO リスト - Go

## Description

Go の基本的な文法を学習するための TODO リストです。
e2e テストまで作りました。

## アーキテクチャ

controller - usecase - repository
のオーソドックスな構成
controller には usecase のインターフェースを渡す形。

## Get Started

```
docker-compose up -d
go mod tidy
go run main.go
```

## 参考

[API を作りながら進む Go 中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)
