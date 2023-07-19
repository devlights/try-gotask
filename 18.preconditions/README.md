# precondition

```preconditions``` を定義することで、タスクの事前条件を指定することが出来る

## サンプル

```yaml
# https://taskfile.dev

version: '3'

vars:
  KENALL: utf_all.csv

tasks:
  default:
    preconditions:
      - (! test -f {{.KENALL}})
    cmds:
      - wget https://www.post.japanpost.jp/zipcode/{{.KENALL}}

```

実行すると以下のようになります。

```sh
$ task
task: [default] wget https://www.post.japanpost.jp/zipcode/utf_all.csv
--2023-07-19 07:51:38--  https://www.post.japanpost.jp/zipcode/utf_all.csv
Resolving www.post.japanpost.jp (www.post.japanpost.jp)... 43.253.48.17
Connecting to www.post.japanpost.jp (www.post.japanpost.jp)|43.253.48.17|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 18305986 (17M) [text/csv]
Saving to: ‘utf_all.csv’

utf_all.csv                                            100%[==========================================================================================================================>]  17.46M  1.78MB/s    in 10s     

2023-07-19 07:51:49 (1.66 MB/s) - ‘utf_all.csv’ saved [18305986/18305986]



$ task
task: `(! test -f utf_all.csv)` failed
task: precondition not met
```

## 参考情報

- [Using programmatic checks to cancel the execution of a task and its dependencies](https://taskfile.dev/usage/#using-programmatic-checks-to-cancel-the-execution-of-a-task-and-its-dependencies)
