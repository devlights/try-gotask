# task のコマンドラインオプション

```sh
$ task --version
Task version: v3.12.1 (h1:SmYDqaRgt9NOm/ogPaMviwckYCqCjbGek9DKF//5A2g=)
```

```sh
$ task -h
Usage: task [-ilfwvsd] [--init] [--list] [--force] [--watch] [--verbose] [--silent] [--dir] [--taskfile] [--dry] [--summary] [task...]

Runs the specified task(s). Falls back to the "default" task if no task name
was specified, or lists all tasks if an unknown task name was specified.

Example: 'task hello' with the following 'Taskfile.yml' file will generate an
'output.txt' file with the content "hello".

'''
version: '3'
tasks:
  hello:
    cmds:
      - echo "I am going to write a file named 'output.txt' now."
      - echo "hello" > output.txt
    generates:
      - output.txt
'''

Options:
  -c, --color                       colored output. Enabled by default. Set flag to false or use NO_COLOR=1 to disable (default true)
  -C, --concurrency int             limit number tasks to run concurrently
  -d, --dir string                  sets directory of execution
      --dry                         compiles and prints tasks in the order that they would be run, without executing them
  -f, --force                       forces execution even when the task is up-to-date
  -h, --help                        shows Task usage
  -i, --init                        creates a new Taskfile.yaml in the current folder
  -l, --list                        lists tasks with description of current Taskfile
  -a, --list-all                    lists tasks with or without a description
  -o, --output string               sets output style: [interleaved|group|prefixed]
      --output-group-begin string   message template to print before a task's grouped output
      --output-group-end string     message template to print after a task's grouped output
  -p, --parallel                    executes tasks provided on command line in parallel
  -s, --silent                      disables echoing
      --status                      exits with non-zero exit code if any of the given tasks is not up-to-date
      --summary                     show summary about a task
  -t, --taskfile string             choose which Taskfile to run. Defaults to "Taskfile.yml"
  -v, --verbose                     enables verbose mode
      --version                     show Task version
  -w, --watch                       enables watch of the given task
```

## -c, --color

色付きで出力を行う。デフォルトで有効。このフラグで false を設定するか、環境変数 ```NO_COLOR=1``` を指定することで無効にできる。

## -C, --concurrency int

同時実行するタスクの数を制限する場合に利用する。値に同時実行最大数を指定する。

## -d, --dir string

実行ディレクトリを指定する。```Taskfile.yml``` がカレントディレクトリとは異なる場所にある場合に利用する。

```make -C directory``` と同じ。

## --dry

dry-run で実行する。どのようなコマンドが実行されるのかを確認する際に利用できる。

## -f, --force

実行するタスクが既に最新状態で実行する必要がない場合でも、強制的に実行する。

## -h, --help

ヘルプを表示する。

## -i, --init

カレントディレクトリに ```Taskfile.yaml``` を初期生成する。

ただ、ドキュメントには ```Taskfile.yml``` という名前を推奨しているのに、```Taskfile.yaml``` という名前で生成される。

## -l, --list

定義しているタスクを一覧表示する。どのようなタスクが定義されているのか確認する際に便利。

```desc:``` に、説明を記載しているタスクのみが表示される。

## -a, --list-all

定義しているタスクを一覧表示する。どのようなタスクが定義されているのか確認する際に便利。

```desc:``` に、説明を記載していないタスクも含めて全て表示される。

## -o, --output string

出力スタイルを指定する。 以下から選択して値を指定する。

- interleaved
- group
- prefixed

## -p, --parallel

タスクを並列実行する。

## -s, --silent

実行するタスクをエコー表示しないようにする。つまり、コマンドの結果のみが表示される。

## --summary

```Taskfile.yml``` の内容をサマリー表示する。

## -t, --taskfile string

指定したファイルを実行する。ファイル名をデフォルトの ```Taskfile.yml``` から異なる名前にしている場合に利用できる。

## -v, --verbose

出力を冗長にする。

## --version

バージョンを表示する。

## -w, --watch

タスクの監視を有効にする。
