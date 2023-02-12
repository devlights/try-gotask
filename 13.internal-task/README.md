# 内部タスクを定義

タスクに ```internal: true``` を設定することで内部タスクとして定義することが出来ます。

## サンプル

```yaml
# https://taskfile.dev/usage/#internal-tasks
version: "3"

tasks:
  default:
    desc: デフォルトタスク
    cmds:
      - task: run
  run:
    desc: 処理を実行
    cmds:
      - task: _hr
      - task: _run
      - task: _hr
  _run:
    desc: 処理を実行（内部）
    internal: true
    cmds:
      - lsb_release -a
  _hr:
    internal: true
    cmds:
      - cmd: echo '------------------------------------'
        silent: true

```

実行すると以下のようになります。

```sh
$ task
------------------------------------
task: [_run] lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 20.04.5 LTS
Release:        20.04
Codename:       focal
------------------------------------
```

内部タスクは、```--list``` や ```--list-all``` を指定した際にも表示されません。

```sh
$ task --list
task: Available tasks for this project:
* default:       デフォルトタスク
* run:           処理を実行

$ task --list-all
task: Available tasks for this project:
* default:       デフォルトタスク
* run:           処理を実行
```

また、内部タスクを指定して実行しようとしてもエラーとなります。

```sh
$ task _run
task: Task "_run" is internal
```

## 参考情報

- [internal-task](https://taskfile.dev/usage/#internal-tasks)
