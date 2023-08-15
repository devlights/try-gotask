# Looping over a static list

v3.28.0 からループ機能が追加された。

基本形は以下のようになる

```yaml
task-name:
  cmds:
    - for: loop-source
      cmd: command {{ .ITEM }}
```

ループ中の項目は ```{{ .ITEM }}``` で参照できる。

固定の情報をループする場合は以下のように出来る。

## サンプル

```yaml
# https://taskfile.dev

version: "3"

tasks:
  default:
    cmds:
      #
      # 固定の情報をループ
      #
      - for: ["one", "two", "three"]
        # ループ中の項目は .ITEM で参照できる
        cmd: echo {{ .ITEM }}

```

実行すると以下のようになります。

```sh
$ task --version
Task version: v3.28.0 (h1:PGYGwevlGQdYrqhO6lLCYylC7YuGoQLlVwHkO42gf0I=)

$ task
task: [default] echo one
one
task: [default] echo two
two
task: [default] echo three
three
```

## 参考情報

- [Looping over a static list](https://taskfile.dev/usage/#looping-over-a-static-list)
