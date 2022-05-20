# 概要

このリポジトリには、[Task](https://github.com/go-task/task) についての自分用のメモが置いてあります。

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

makeには ```Makefile``` のように、taskの場合は ```Taskfile.yml``` というファイルを使う。

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

```task --init``` とすることで、カレントディレクトリに Taskfile.yml を初期生成してくれる。

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
