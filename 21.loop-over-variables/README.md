# Looping over variables

v3.28.0 からループ機能が追加された。

基本形は以下のようになる

```yaml
task-name:
  cmds:
    - for: loop-source
      cmd: command {{ .ITEM }}
```

ループ中の項目は ```{{ .ITEM }}``` で参照できる。

変数の値をループする場合は以下のように出来る。

## サンプル

```yaml
# https://taskfile.dev

version: "3"

silent: true

vars:
  VAL1: Hello World
  VAL2: HELLO,WORLD
  VAL3:
    sh: find ../ -name '*.yml' -type f

tasks:
  default:
    cmds:
      - task: use-val1
      - cmd: echo '--------------------------------------'
      - task: use-val2
      - cmd: echo '--------------------------------------'
      - task: use-val3
  use-val1:
    cmds:
      - for: { var: VAL1 }
        cmd: echo {{ .ITEM }}
  use-val2:
    cmds:
      - for: { var: VAL2, split: "," }
        cmd: echo {{ .ITEM }}
  use-val3:
    cmds:
      - for: { var: VAL3 }
        cmd: wc -l {{ .ITEM }}

```

実行すると以下のようになります。

```sh
$ task
Hello
World
--------------------------------------
HELLO
WORLD
--------------------------------------
11 ../.gitpod.yml
20 ../01.helloworld/Taskfile.yml
9 ../03.default-task/Taskfile.yml
15 ../04.env/Taskfile.yml
10 ../05.dotenv/Taskfile.yml
14 ../06.include/Taskfile.yml
6 ../06.include/other/Taskfile.yml
6 ../06.include/other2/othertaskfile.yml
6 ../07.no-color/Taskfile.yml
17 ../08.silent/Taskfile.yml
9 ../09.os-specific/Taskfile.yml
5 ../09.os-specific/Taskfile_linux.yml
5 ../09.os-specific/Taskfile_windows.yml
18 ../10.variable/Taskfile.yml
8 ../11.cliargs/Taskfile.yml
25 ../12.defer/Taskfile.yml
24 ../13.internal-task/Taskfile.yml
37 ../14.shopt/Taskfile.yml
22 ../15.single-command-task/Taskfile.yml
13 ../16.prompt/Taskfile.yml
15 ../17.dryrun/Taskfile.yml
13 ../18.preconditions/Taskfile.yml
13 ../19.loop-over-static-list/Taskfile.yml
16 ../20.loop-over-task-source/Taskfile.yml
33 ../21.loop-over-variables/Taskfile.yml
```

## 参考情報

- [Looping over variables](https://taskfile.dev/usage/#looping-over-variables)
