# ```.CHECKSUM``` と ```.TIMESTAMP```

v3.41.0にて、```cmds```と```status```の部分で以下の変数が使えるようになりました。

- ```.CHECKPOINT```
- ```.TIMESTAMP```

どちらも ```sources``` で対象にしているものに対しての値を出力します。

## サンプル

### ファイル構成

```sh
$ tree .
.
├── README.md
└── Taskfile.yml

0 directories, 2 files
```

### Taskfile.yml

```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    sources:
      - ./*.md
    method: timestamp
    cmds:
      - echo "CHECKSUM={{.CHECKSUM}}"
      - echo "TIMESTAMP={{.TIMESTAMP}}"
    silent: true
```

### 実行

```sh
$ task
task: [default] echo "CHECKSUM=4031f2f73b02c59ceb876c1d6eaf70e4"
CHECKSUM=4031f2f73b02c59ceb876c1d6eaf70e4
task: [default] echo "TIMESTAMP=2025-01-22 09:04:03.913762943 +0000 UTC"
TIMESTAMP=2025-01-22 09:04:03.913762943 +0000 UTC

$ task
task: Task "default" is up to date

$ task -f
task: [default] echo "CHECKSUM=4031f2f73b02c59ceb876c1d6eaf70e4"
CHECKSUM=4031f2f73b02c59ceb876c1d6eaf70e4
task: [default] echo "TIMESTAMP=2025-01-22 09:04:03.913762943 +0000 UTC"
TIMESTAMP=2025-01-22 09:04:03.913762943 +0000 UTC
```
