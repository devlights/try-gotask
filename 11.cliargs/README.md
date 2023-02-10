# コマンドライン引数を使う

```{{ .CLI_ARGS }}``` という形でコマンドラインから指定された引数を利用することが出来る。

## コマンドライン引数の指定方法

```task``` コマンド呼び出し時に ```--``` の後に指定する。

```sh
$ task run -- hello world
```

上記の場合、```{{ .CLI_ARGS }}``` の値は ```hello world``` となる。

## 参考情報

- https://taskfile.dev/usage/#forwarding-cli-arguments-to-commands
