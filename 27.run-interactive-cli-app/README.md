# インタラクティブなCLIアプリを動作させる

タスクの中でインタラクティブなCLIアプリ、例えば vim を起動すると

出力内容が、たまに少しおかしくなったりします。

そのような場合は ```interactive: true``` を付与することでうまくいったりします。

インタラクティブなCLIアプリを起動する場合は指定しておいた方が無難ですね。

### サンプル

以下のようなタスクファイルを用意します。

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - echo "{{.PWD}}"
    silent: true
  run-vim:
    cmds:
      - cmd: vim README.md
    interactive: true

```

一旦、実行すると以下のようになります。

```sh
$ task -d 27.run-interactive-cli-app/ run-vim
task: [run-vim] vim README.md

## vim が起動する ##
```
