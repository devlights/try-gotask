# dry-run モード

実行時に ```--dry``` オプションを付与することにより dry-run モードで動作する

## サンプル

```yaml
# https://taskfile.dev

version: "3"

vars:
  GREETING: Hello, World!

tasks:
  default:
    cmds:
      - echo "{{.GREETING}}"
      - echo "{{.GREETING}}"
      - echo "{{.GREETING}}"
      - echo "{{.GREETING}}"
      - echo "{{.GREETING}}"

```

実行すると以下のようになります。

```sh
$ task --dry
task: [default] echo "Hello, World!"
task: [default] echo "Hello, World!"
task: [default] echo "Hello, World!"
task: [default] echo "Hello, World!"
task: [default] echo "Hello, World!"
```

## 参考情報

- [Dry run mode](https://taskfile.dev/usage/#dry-run-mode)
