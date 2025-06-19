# これは何？

[v3.44](https://taskfile.dev/changelog/#v3440---2025-06-08)にて ```uuid``` 関数が追加されました。

文字通り、uuidを出力してくれる関数です。

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - echo "{{uuid}}"
    silent: true

```

```sh
$ task
2d548ad7-9ce6-46c0-af29-bb446c699dc5
```
