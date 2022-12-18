# 概要

このリポジトリには、[Task](https://github.com/go-task/task) についての自分用のメモが置いてあります。

## ToC

- [Task (go-task) とは](#task-go-task-とは)
- [最も大事なルール](#最も大事なルール)
- [Taskfile.yml の作り方](#taskfileyml-の作り方)
- [実行方法](#実行方法)
- [環境変数 (env, dotenv)](#環境変数-env-dotenv)
- [他のタスクファイルを取込み](#他のタスクファイルを取込み)

## Task (go-task) とは

[Task(go-task)](https://github.com/go-task/task) とは、make のような タスクランナーでありビルドツール。

Goで作成されているので、シングルバイナリとなっており、どの環境であってもインストールが簡単。

Windowsの場合でも scoop で以下のようにするとインストールできる。

```sh
scoop bucket add extras
scoop install task
```

Linuxの場合も、パッケージマネージャからインストール出来るようになっている。

```sh
brew install go-task/tap/go-task
```

また、Goで作成されているので、Goが入っていれば以下でもインストールできる。

```sh
go install github.com/go-task/task/v3/cmd/task@latest
```

Windowsの場合、デフォルトではmakeが存在しないので、makeを個別で入れる代わりに[Task](https://github.com/go-task/task)を使ってもいいかもしれない。

以下にドキュメントを見ながら覚えていった内容をメモしておくことにする。

---

## 最も大事なルール

makeには ```Makefile``` のように、taskの場合は ```Taskfile.yml```(```Taskfile.yaml```) というファイルを使う。

```Taskfile.yml``` は、例えば以下のようになる。

```yaml
version: '3'

env:
  MESSAGE: helloworld

tasks:
  default:
    cmds:
      - echo $MESSAGE
```

## Taskfile.yml の作り方

```task --init``` とすることで、カレントディレクトリに ```Taskfile.yml``` を初期生成してくれる。

(*) 本家のドキュメントには ```Taskfile.yml``` って名前のファイルを使えとなっているがコマンドラインで生成されるものは ```Taskfile.yaml``` になる。

## 実行方法

対象となる ```Taskfile.yml``` が存在するディレクトリに移動して

```sh
# デフォルトのタスクが実行される
$ task

# タスク指定
$ task xxxx
```

とするか、ディレクトリは移動せずに

```sh
$ task -d /path/to/target
$ task -d /path/to/target task-name
```

としても良い。

## 環境変数 (env, dotenv)

タスク単位で環境変数を指定することが出来る。

```yaml
version: "3"

env:
  MYVARGLOBAL: myvar-global
  MYVARDUP: myvar-global

tasks:
  default:
    cmds:
      - echo $MYVAR
      - echo $MYVARGLOBAL
      - echo $MYVARDUP
    env:
      MYVAR: myvar-local
      MYVARDUP: myvar-local

```

```sh
$ task -d 04.env/
task: [default] echo $MYVAR
myvar-local
task: [default] echo $MYVARGLOBAL
myvar-global
task: [default] echo $MYVARDUP
myvar-local
```

また、```.env``` などを指定することも出来る。この場合は ```dotenv:``` を用いる。

### my.env

```yaml
MYVAR1=value1
MYVAR2=value2
```

### my2.env

```yaml
MYVAR2=value2-2
MYVAR3=value3
```

### Taskfile.yml

```yaml
version: "3"

dotenv: ["my.env", "my2.env"]

tasks:
  default:
    cmds:
      - echo $MYVAR1
      - echo $MYVAR2
      - echo $MYVAR3

```

```sh
$ task -d 05.dotenv/
task: [default] echo $MYVAR1
value1
task: [default] echo $MYVAR2
value2
task: [default] echo $MYVAR3
value3
```

## 他のタスクファイルを取込み

他のタスクファイルを取込み（include）することが可能。

取込むには、トップレベルで ```includes:``` を指定する。

存在しない場合でも処理を止めたくない場合は、```optional: true``` を指定する。

### ファイル階層

```sh
$ tree 06.include/
06.include/
├── other
│   └── Taskfile.yml
├── other2
│   └── othertaskfile.yml
└── Taskfile.yml

2 directories, 3 files
```


### other/Taskfile.yml

```yaml
version: "3"

tasks:
  task:
    cmds:
      - echo 'task1'

```

### other2/othertaskfile.yml

```yaml
version: "3"

tasks:
  task:
    cmds:
      - echo 'task2'

```

### Taskfile.yml

```yaml
version: "3"

includes:
  other1: ./other
  other2: ./other2/othertaskfile.yml
  other3:
    taskfile: ./other3/Taskfile.yml
    optional: true

tasks:
  default:
    cmds:
      - task: other1:task
      - task: other2:task

```

```sh
$ task -d 06.include/
task: [other1:task] echo 'task1'
task1
task: [other2:task] echo 'task2'
task2
```

### リスト

|directory|readme|taskfile|
|----|----|----|
|01.helloworld|[README](01.helloworld/README.md)|[Taskfile.yml](01.helloworld/Taskfile.yml)|
|02.cli-options|[README](02.cli-options/README.md)||
|03.default-task|[README](03.default-task/README.md)|[Taskfile.yml](03.default-task/Taskfile.yml)|
|04.env||[Taskfile.yml](04.env/Taskfile.yml)|
|05.dotenv||[Taskfile.yml](05.dotenv/Taskfile.yml)|
|06.include||[Taskfile.yml](06.include/Taskfile.yml)|
|07.no-color||[Taskfile.yml](07.no-color/Taskfile.yml)|
|08.silent||[Taskfile.yml](08.silent/Taskfile.yml)|
