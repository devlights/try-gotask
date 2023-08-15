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

このループ変数名を変更することも出来る。

## サンプル

```yaml
# https://taskfile.dev

version: "3"

tasks:
  default:
    vars:
      ymlFiles:
        sh: find ../ -type f -name 'Taskfile.yml' | sed -e 's;../;;'
    cmds:
      - for: { var: ymlFiles, as: YML_FILE }
        cmd: echo {{ .YML_FILE }}
    silent: true

```

実行すると以下のようになります。

```sh
$ task
01.helloworld/Taskfile.yml
03.default-task/Taskfile.yml
04.env/Taskfile.yml
05.dotenv/Taskfile.yml
06.include/Taskfile.yml
06.include/other/Taskfile.yml
07.no-color/Taskfile.yml
08.silent/Taskfile.yml
09.os-specific/Taskfile.yml
10.variable/Taskfile.yml
11.cliargs/Taskfile.yml
12.defer/Taskfile.yml
13.internal-task/Taskfile.yml
14.shopt/Taskfile.yml
15.single-command-task/Taskfile.yml
16.prompt/Taskfile.yml
17.dryrun/Taskfile.yml
18.preconditions/Taskfile.yml
19.loop-over-static-list/Taskfile.yml
20.loop-over-task-source/Taskfile.yml
21.loop-over-variables/Taskfile.yml
22.loop-rename-variables/Taskfile.yml
```

## 参考情報

- [Renaming variables](https://taskfile.dev/usage/#renaming-variables)
