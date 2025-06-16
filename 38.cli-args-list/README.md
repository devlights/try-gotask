# これは何？

[v3.44](https://taskfile.dev/changelog/#v3440---2025-06-08)で

- CLI_ARGS_LIST

が追加されました。元々存在している ```CLI_ARGS``` とほぼ同様であるが、文字列ではなくリストで取得出来る。

```sh
$ task -- hello world へろー わーるど
[hello world へろー わーるど]
hello
world
へろー
わーるど
```
