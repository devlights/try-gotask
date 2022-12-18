# コマンド出力行を出力しないようにする

Taskではデフォルトで実行されるコマンドの行が出力されます。

以下の ```task:``` で始まっている行のことです。

```sh
task: [no-silent] ls -l
total 4
-rw-r--r-- 1 gitpod gitpod 304 Dec 18 09:18 Taskfile.yml
```

コマンド毎に ```silent: true``` を付与することで制御することができます。

```sh
      - cmd: ls -l
        silent: true
```

```sh
total 4
-rw-r--r-- 1 gitpod gitpod 304 Dec 18 09:18 Taskfile.yml
```
