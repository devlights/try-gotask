# Looping over your task's sources

v3.28.0 からループ機能が追加された。

基本形は以下のようになる

```yaml
task-name:
  cmds:
    - for: loop-source
      cmd: command {{ .ITEM }}
```

ループ中の項目は ```{{ .ITEM }}``` で参照できる。

タスクソースをループする場合は以下のように出来る。

## サンプル

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    deps: [clear-checksums]
    sources:
      - ./*.txt
      - ./*.go
    cmds:
      - for: sources
        cmd: stat {{ .ITEM }} | head -n 3
  clear-checksums:
    cmds:
      - rm -rf ./.task

```

実行すると以下のようになります。

```sh
task: [clear-checksums] rm -rf ./.task
task: [default] stat hello.txt | head -n 3
  File: hello.txt
  Size: 6               Blocks: 8          IO Block: 4096   regular file
Device: fd01h/64769d    Inode: 2776629404  Links: 1
task: [default] stat main1.go | head -n 3
  File: main1.go
  Size: 85              Blocks: 8          IO Block: 4096   regular file
Device: fd01h/64769d    Inode: 2776629407  Links: 1
task: [default] stat main2.go | head -n 3
  File: main2.go
  Size: 85              Blocks: 8          IO Block: 4096   regular file
Device: fd01h/64769d    Inode: 2776629408  Links: 1
task: [default] stat world.txt | head -n 3
  File: world.txt
  Size: 6               Blocks: 8          IO Block: 4096   regular file
Device: fd01h/64769d    Inode: 2776629405  Links: 1
```

## 参考情報

- [Looping over your task's sources](https://taskfile.dev/usage/#looping-over-your-tasks-sources)
