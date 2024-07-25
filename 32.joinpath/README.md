# joinPath 関数

[go-task](https://taskfile.dev/) の [ビルドイン関数](https://taskfile.dev/next/reference/templating/#built-in-functions) には

いろいろと便利な関数が公開されている。

```joinPath``` 関数は、その名前の通り パスを結合 する仕事を行ってくれる。

以下のようにして利用できる。

```yaml
for: { var: ITEMS }
cmd: echo {{joinPath {{.BASE_DIR}} {{.ITEM}}}}
```
