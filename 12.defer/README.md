# deferの指定

Task では、Goのように defer を指定することが出来る。

Go の defer とは以下のような機能。

- https://go.dev/tour/flowcontrol/12


```yaml
tasks:
  helloworld:
    cmds:
      - echo hello
      - defer: echo hehehe
      - echo world
```

上記を実行すると以下のように出力される。

```sh
$ task -d 12.defer/ helloworld
task: [helloworld] echo hello
hello
task: [helloworld] echo world
world
task: [helloworld] echo hehehe
hehehe
```

クリーンアップ処理などを行う際に有用。

## 参考情報

- https://taskfile.dev/usage/#doing-task-cleanup-with-defer
