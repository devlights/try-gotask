# グローバル タスクファイル

Taskには ```グローバルタスクファイル``` という概念があります。

ユーザのホームディレクトリに配置された タスクファイル がグローバルタスクファイルとなります。

グローバルタスクファイルに定義されているタスクを実行するには

```sh
$ task -g task-name
```

とします。グローバルタスクなので、どの場所からでも実行できるタスクファイルとなります。

## サンプル

以下のようなタスクファイルを、まず自分の作業ディレクトリに作ります。

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - echo "PWD=$(pwd)"
      - echo "USER_WORKING_DIR={{.USER_WORKING_DIR}}"
      - echo "HOME_DIR={{.HOME}}"
    silent: true
  to_global_taskfile:
    cmds:
      - cp ./Taskfile.yml {{.HOME}}
      - ls -l "{{.HOME}}/Taskfile.yml"
  rm_global_taskfile:
    dir: "{{.HOME}}"
    cmds:
      - rm Taskfile.yml
      - ls -l Taskfile.yml
    ignore_error: true


```

一旦、実行すると以下のようになります。

```sh
$ task -d 26.global-taskfile/
PWD=/workspace/try-gotask/26.global-taskfile
USER_WORKING_DIR=/workspace/try-gotask
HOME_DIR=/home/gitpod
```

次に、上のタスクファイルに定義されている ```to_global_taskfile``` タスクを実行してホームディレクトリに配置します。

```sh
$ task -d 26.global-taskfile/ to_global_taskfile
task: [to_global_taskfile] cp ./Taskfile.yml /home/gitpod
task: [to_global_taskfile] ls -l "/home/gitpod/Taskfile.yml"
-rw-r--r-- 1 gitpod gitpod 304 Sep 24 13:45 /home/gitpod/Taskfile.yml
```

これで、同じタスクファイルがホームディレクトリにコピーされて、グローバルタスクファイルになりました。

グローバルタスクとして実行すると以下のようになります。

```sh
$ task -g
PWD=/home/gitpod
USER_WORKING_DIR=/workspace/try-gotask
HOME_DIR=/home/gitpod
```

PWDの出力は変わりましたが、 ```{{.USER_WORKING_DIR}}``` の場所は変わっていません。

グローバルタスクを作成する際に「現在実際いるディレクトリ」の情報を利用することは多いので、これは便利です。

今回のサンプルで利用したタスクファイルはもういらないので削除。

```sh
$ task -g rm_global_taskfile
task: [rm_global_taskfile] rm Taskfile.yml
task: [rm_global_taskfile] ls -l Taskfile.yml
ls: cannot access 'Taskfile.yml': No such file or directory
```
