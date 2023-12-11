# リモートのTaskfileを実行 (v3.32.0の段階では実験的機能)

実験的機能として、「リモートのタスクファイルの実行」がサポートされています。(v3.32.0時点)

やり方としては、[06.include](../06.include/Taskfile.yml) と同様で、取込先がリモートとなるだけです。

### サンプル

以下のようなタスクファイルを用意します。

```yaml
# https://taskfile.dev

version: '3'

#
# リモートのタスクファイルを指定
#   - 以下は my-remote という prefix でリモートタスクファイルを設定
#
includes:
  my-remote:
    taskfile: https://raw.githubusercontent.com/devlights/try-gotask/main/03.default-task/Taskfile.yml

tasks:
  default:
    cmds:
      #
      # リモートタスクファイルに定義されているタスクを実行
      #
      - task: my-remote:hello

```

実験的機能であるため、環境変数 ```TASK_X_REMOTE_TASKFILES``` を指定していない場合はエラーとなります。

```sh
$ task
task: Remote taskfiles are not enabled. You can read more about this experiment and how to enable it at https://taskfile.dev/experiments/remote-taskfiles
```

環境変数を指定して実行します。

```sh
$ TASK_X_REMOTE_TASKFILES=1 task
The task you are attempting to run depends on the remote Taskfile at "https://raw.githubusercontent.com/devlights/try-gotask/main/03.default-task/Taskfile.yml".
--- Make sure you trust the source of this Taskfile before continuing ---
Continue? [y/N]
y
task: [my-remote:hello] echo 'hello world'
hello world
```

リモートタスクファイルの場合、本当に取得して続行して良いかを確認するプロンプトが表示されます。

ここで、```y``` と入力すると、処理が続行されリモートのタスクファイルに定義されているタスクが実行されます。

一度、許可した後の再実行時はプロンプトは表示されません。

## 生成されるファイル

リモートタスクファイルを実行すると、タスクファイルが存在するディレクトリに ```.task``` ディレクトリが生成され取得したファイルのチェックサムが記録されます。
