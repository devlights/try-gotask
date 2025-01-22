# ```TASK_DIR``` 変数

v3.41.0にて、```TASK_DIR```という変数が新たに追加されました。

この変数には、実行時のタスクファイルの場所が設定されます。

親のタスクファイルにて、子のタスクファイルをincludeしている場合

子のタスクファイル内の ```TASK_DIR``` の値は親と同じになります。

（子のタスクファイルを直接実行した場合は、当然子のタスクファイルの場所になります）

## 実行例

### ファイル構成

```sh
$ tree .
.
├── README.md
├── subdir
│   └── Taskfile.yml
└── Taskfile.yml

1 directory, 3 files
```

### 親タスクファイル

```yaml
# https://taskfile.dev

version: '3'

includes:
  subdir: ./subdir

tasks:
  default:
    cmds:
      - task: call-self-taskdir
      - task: subdir:taskdir
      - task: call-subdir-taskfile
  call-self-taskdir:
      - echo "{{.TASK_DIR}}"
  call-subdir-taskfile:
      - task -d "{{.TASK_DIR}}/subdir" taskdir

```

### 子のタスクファイル

```yaml
# https://taskfile.dev

version: '3'

tasks:
  taskdir:
    cmds:
      - echo "{{.TASK_DIR}}"
```

### 実行

```sh
$ task
task: [call-self-taskdir] echo "/workspace/try-gotask/34.v341-taskdir"
/workspace/try-gotask/34.v341-taskdir                                   # <-- 親タスクファイルの場所
task: [subdir:taskdir] echo "/workspace/try-gotask/34.v341-taskdir"
/workspace/try-gotask/34.v341-taskdir                                   # <-- 子のタスクファイルからの実行だが、場所は親タスクファイルの場所となる
task: [call-subdir-taskfile] task -d "/workspace/try-gotask/34.v341-taskdir/subdir" taskdir
task: [taskdir] echo "/workspace/try-gotask/34.v341-taskdir/subdir"
/workspace/try-gotask/34.v341-taskdir/subdir                            # <-- 子のタスクファイルを直接実行した場合は、場所は子のタスクファイルの場所
```