# Windows上でコマンドを実行する場合のやり方

Windows上でコマンドを実行する場合は

```sh
cmd /c "コマンド"
```

という形で実行する。普通に ```rd /s /q .task``` とかで指定すると

rdコマンドが見つからないというエラーになる。

## サンプル

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - task: run
  build:
    cmds:
      - go build -o app{{exeExt}}
    sources:
      - ./*.go
    generates:
      - app{{exeExt}}
    method: timestamp
  run:
    deps:
      - build
    cmds:
      - ./app{{exeExt}}
  clean:
    cmds:
      - cmd /c "del /F /Q app{{exeExt}}"
      - cmd /c "rd /s /q .task"
    ignore_error: true

```

実行すると以下のようになります。

```sh
$ task
task: [build] go build -o app.exe
task: [run] ./app.exe
0:helloworld
1:helloworld
2:helloworld
3:helloworld
4:helloworld


$ task
task: Task "build" is up to date
task: [run] ./app.exe
0:helloworld
1:helloworld
2:helloworld
3:helloworld
4:helloworld

$ task clean
task: [clean] cmd /c "del /F /Q app.exe"
task: [clean] cmd /c "rd /s /q .task"


$ task
task: [build] go build -o app.exe
task: [run] ./app.exe
0:helloworld
1:helloworld
2:helloworld
3:helloworld
4:helloworld
```