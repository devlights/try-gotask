# Renaming variables

v3.28.0 からループ機能が追加された。

基本形は以下のようになる

```yaml
task-name:
  cmds:
    - for: loop-source
      cmd: command {{ .ITEM }}
```

ループ中の項目は ```{{ .ITEM }}``` で参照できる。

ループ内で別のタスクを呼び出すことも出来る。その場合にループ項目をサブタスク側に渡すことも可能。

## サンプル

```yaml
# https://taskfile.dev

version: "3"

tasks:
  default:
    vars:
      FILES:
        sh: ls -1
    cmds:
      - for: { var: FILES }
        task: proc
        vars:
          FILE: "{{ .ITEM }}"
  proc:
    internal: true
    cmds:
      - task: echo
        vars:
          TARGET: "{{ .FILE }}"
      - task: wc
        vars:
          TARGET: "{{ .FILE }}"
      - task: hr
  echo:
    internal: true
    cmds:
      - echo '{{ .TARGET }}'
  wc:
    internal: true
    cmds:
      - wc '{{ .TARGET }}'
  hr:
    internal: true
    silent: true
    cmds:
      - echo '---------------------------------'

```

実行すると以下のようになります。

```sh
$ task
task: [echo] echo 'README.md'
README.md
task: [wc] wc 'README.md'
  70  113 1207 README.md
---------------------------------
task: [echo] echo 'Taskfile.yml'
Taskfile.yml
task: [wc] wc 'Taskfile.yml'
 37  75 609 Taskfile.yml
---------------------------------
```

## 参考情報

- [Looping over tasks](https://taskfile.dev/usage/#looping-over-tasks)
