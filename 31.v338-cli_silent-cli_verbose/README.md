# v3.38で追加された特殊変数 (CLI_SILENT, CLI_VERBOSE)

v.3.38で以下の特殊変数が追加された。

- CLI_SILENT
  - 実行時に ```--silent``` が付与されていればtrueとなる
- CLI_VERBOSE
  - 実行時に ```--verbose``` が付与されていればtrueとなる

実行するアプリケーションに ```-verbose``` オプションや ```-silent``` オプションのようなものが存在する場合は、今回追加された特殊変数の値によって切り替えて実行したり出来る。

## タスクファイル

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - go build -o app .
      - ./app {{- if .CLI_SILENT}} -s {{end -}}{{- if .CLI_VERBOSE}} -v {{end -}}
```

## 実行

```sh
$ task 
task: [default] go build -o app .
task: [default] ./app
first
second

$ task --silent
first

$ task --verbose
task: [/workspace/try-gotask/31.v338-special-variable] Not found - Using alternative (Taskfile.yml)
task: "default" started
task: [default] go build -o app .
task: [default] ./app -v 
zero
first
second
fourth
task: "default" finished

$ task --silent --verbose
task: [/workspace/try-gotask/31.v338-special-variable] Not found - Using alternative (Taskfile.yml)
task: "default" started
task: [default] go build -o app .
task: [default] ./app -s  -v 
zero
first
fourth
task: "default" finished
```

## REFERENCES

- https://taskfile.dev/changelog/#v3380---2024-06-30
- https://github.com/go-task/task/issues/1616
- https://github.com/go-task/task/issues/1480
