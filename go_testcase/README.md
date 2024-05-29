# go_testcaseについて

Go言語でのテストケースのサンプルです。

# 実行方法

## mainの実行

~~~go
go run main.go
~~~

## テストケースの実行

~~~go
// gotestディレクトリへ移動
cd gotest

// 全てのテストケースを実行
go test -v ./...

// 個別のテストケースを実行
go test -v add_sub_test.go
go test -v multi_div_test.go
~~~