# shopt (globstar) の設定

グローバルまたはコマンドレベルでshoptを設定できます。

Bashの全オプションが有効なわけではなく、[https://github.com/mvdan/sh](https://github.com/mvdan/sh) で有効なものとなります。

グローバルなshoptの場合、以下のものだけサポートされている模様。

```
Valid values: "expand_aliases", "globstar", "nullglob"
```

コマンドレベルなshoptの場合、サポートされていないオプションを指定すると以下のようなメッセージが表示される。

```
shopt: invalid option name "autocd" "off" ("on" not supported)
```

```yaml
version: '3'

shopt:
  - "globstar" # グローバルな設定

tasks:
  default:
    - task: clean
    - task: init
    - task: run-no-shopt
    - task: run-with-shopt
  clean:
    internal: true
    cmds:
      - rm -rf a
  init:
    internal: true
    cmds:
      - mkdir -p a/b/c
      - echo 'hello' > a/hello.txt
      - echo 'world' > a/b/c/world.txt
  run:
    internal: true
    cmds:
      - cmd: ls **/*.txt    
  run-no-shopt:
    internal: true
    cmds:
      - cmd: ls **/*.txt
        shopt:
          - "-u globstar" # コマンド毎の設定
  run-with-shopt:
    internal: true
    cmds:
      - cmd: ls **/*.txt
        shopt:
          - "-s globstar" # コマンド毎の設定

```

実行すると以下のようになります。

```sh
$ task
task: [clean] rm -rf a
task: [init] mkdir -p a/b/c
task: [init] echo 'hello' > a/hello.txt
task: [init] echo 'world' > a/b/c/world.txt
task: [run-no-shopt] ls **/*.txt
a/hello.txt
task: [run-with-shopt] ls **/*.txt
a/b/c/world.txt  a/hello.txt
```