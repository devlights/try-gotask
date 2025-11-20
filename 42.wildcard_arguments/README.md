# これは何？

[Task](https://taskfile.dev/) には [Wildcard arguments](https://taskfile.dev/docs/guide#wildcard-arguments) という機能があります。

これはタスク名の部分に `*` をつけることでこの部分にマッチする値を変数として受け取ることが出来る機能です。

> Another way to parse arguments into a task is to use a wildcard in your task's name. Wildcards are denoted by an asterisk (*) and can be used multiple times in a task's name to pass in multiple arguments.

> 引数をタスクに解析する別の方法は、タスク名にワイルドカードを使用することです。ワイルドカードはアスタリスク（*）で示され、タスク名内で複数回使用することで複数の引数を渡すことができます。

例を上げると以下のような感じです。

```yaml
  mytask:*:
    vars:
      VALUE1: '{{index .MATCH 0}}'
    cmds:
      - echo "value is '{{.VALUE1}}'"
```

上の `mytask:*:` としている部分がワイルドカードです。

このようにタスク名を設定した状態でコマンドラインより

```sh
$ task mytask:helloworld
```

と実行すると内部で `VALUE1` にアスタリスクの部分の値が設定され、以下のように出力されます。

```sh
$ task mytask:helloworld
value is 'helloworld'
```

ワイルドカードは複数指定することが可能です。

```yaml
  mytask:*:*:
    vars:
      VALUE1: '{{index .MATCH 0}}'
      VALUE2: '{{index .MATCH 1}}''
    cmds:
      - echo "value1 is '{{.VALUE1}}'"
      - echo "value2 is '{{.VALUE2}}'"
```

以下のように指定します。

```sh
$ task mytask:foo:bar
value1 is 'foo'
value2 is 'bar'
```
