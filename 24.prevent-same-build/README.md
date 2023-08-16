# Prevent unnecessary work (By timestamping locally generated files and their sources)

Makeを利用しているときのように、何も変わっていないかどうかを判断してタスクを実行するように指示することも可能。

```sources```を用いて、ソース指定を行った場合、デフォルトではタスクファイルが存在するディレクトリに ```.task/``` というディレクトリが作成される。

このディレクトリにチェックサムやタイムスタンプの情報が保持されるため、```.gitignore```で除外対象としておく方が良い。

## サンプル

```yaml
# https://taskfile.dev

version: "3"

tasks:
  default:
    deps:
      - build
    cmds:
      - ./app.exe
  build:
    cmds:
      - go build -o app.exe main.go
    sources:
      - ./*.go
    generates:
      - app{{exeExt}}
    method: timestamp
  clean:
    cmds:
      - rm -rf .task/
      - rm -f app.exe

```

実行すると以下のようになります。

```sh
$ task
task: [build] go build -o app.exe main.go
task: [default] ./app.exe
HELLO WORLD

$ task
task: Task "build" is up to date
task: [default] ./app.exe
HELLO WORLD
```

初回はビルドタスクが実行されているが、２回目からは変更が無いので ```Task "build" is up to date``` と表示されてスキップされている。

## 参考情報

- [Prevent unnecessary work](https://taskfile.dev/usage/#prevent-unnecessary-work)
