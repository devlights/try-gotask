# 単一コマンドの場合はタスクに直接cmdと記載できる (v3.26より)

v3.26 より単一コマンドである場合は cmds -> cmd と書かなくても良くなった

## サンプル

```yaml
# https://taskfile.dev

version: '3'

vars:
  GREETING: Hello, World!

tasks:
  default:
    cmds:
      - task: new-style
      - task: old-style
  new-style:
    #
    # v3.26 より単一コマンドである場合は cmds -> cmd と書かなくても良くなった
    #
    cmd: echo "{{.GREETING}}"
    silent: true
  old-style:
    cmds:
      - cmd: echo "{{.GREETING}}"
        silent: true
```

実行すると以下のようになります。

```sh
$ task
Hello, World!
Hello, World!
```

## 参考情報

- [v3.26-ChangeLog](https://taskfile.dev/changelog/#v3260---2023-06-10)
