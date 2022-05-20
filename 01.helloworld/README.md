# Overview

helloworld を出力するサンプルです。

```Taskfile.yml``` を初期生成するには

```sh
task --init
```

とすると、以下のファイルが生成されます。

```yaml
# https://taskfile.dev

version: '3'

vars:
  GREETING: Hello, World!

tasks:
  default:
    cmds:
      - echo "{{.GREETING}}"
    silent: true

```

## Run

```sh
task
```
