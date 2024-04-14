# 依存関係タスク実行時にループ処理

v3.36.0 にて、[Looping over dependencies](https://taskfile.dev/usage/#looping-over-dependencies) という機能が追加されました。

これは、```deps``` にて依存タスクを指定時にループ処理が指定出来るというものです。


## サンプル

若干無理やり感が漂いますがサンプルってことで。。。

以下のようなタスクファイルを用意します。

```yaml
# https://taskfile.dev

version: "3"

tasks:
  default:
    deps:
      - for: ["32", "64"]
        task: build
        vars:
          OPTION_CPU: '{{.ITEM}}'
    cmds:
      - task: show
  build:
    cmds:
      - gcc -m{{.OPTION_CPU}} -g -Wall -o main{{.OPTION_CPU}}.o -c main.c
      - gcc -m{{.OPTION_CPU}} -g -Wall -o app{{.OPTION_CPU}} main{{.OPTION_CPU}}.o
  show:
    cmds:
      - file ./app32
      - file ./app64
  install-libc32:
    cmds:
      - sudo apt-get install gcc-multilib

```

実行すると以下のように、パラメータが渡ってるのが確認できます。


```sh
$ task
task: [build] gcc -m64 -g -Wall -o main64.o -c main.c
task: [build] gcc -m32 -g -Wall -o main32.o -c main.c
task: [build] gcc -m64 -g -Wall -o app64 main64.o
task: [build] gcc -m32 -g -Wall -o app32 main32.o
task: [show] file ./app32
./app32: ELF 32-bit LSB pie executable, Intel 80386, version 1 (SYSV), dynamically linked, interpreter /lib/ld-linux.so.2, BuildID[sha1]=b2f47ccd92c87fc9ad086099543c4583b80f124a, for GNU/Linux 3.2.0, with debug_info, not stripped
task: [show] file ./app64
./app64: ELF 64-bit LSB pie executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, BuildID[sha1]=5b0532275712cb255ae1fd1b5ae75f7a57b537ba, for GNU/Linux 3.2.0, with debug_info, not stripped
```
