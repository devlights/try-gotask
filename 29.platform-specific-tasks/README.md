# プラットフォームごとに実行するタスクを定義

特定のプラットフォームでのみ実行するタスクを定義することが出来ます。

利用する場合は ```platforms: [ OS/ARCH ]``` という形で指定します。

例： ```platforms: [ linux/amd64 ]```

また、OSのみ、ARCHのみの指定でも可能です。

例： ```platforms: [ linux ]```

## サンプル

以下のようなタスクファイルを用意します。

```yaml
# https://taskfile.dev

version: "3"

tasks:
  default:
    cmds:
      - task: 32bit
      - task: 64bit
      - task: windowsonly
      - task: linuxonly
  32bit:
    platforms: [linux/386]
    cmds:
      - echo 32bit
  64bit:
    platforms: [linux/amd64]
    cmds:
      - echo 64bit
  windowsonly:
    platforms: [ windows ]
    cmds:
      - echo windows
  linuxonly:
    platforms: [ linux ]
    cmds:
      - echo linux
```

64ビット環境のLinux上で実行すると以下のようになります。

```sh
$ task
task: [64bit] echo 64bit
64bit
task: [linuxonly] echo linux
linux
```
