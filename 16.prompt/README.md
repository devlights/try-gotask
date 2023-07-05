# タスク実行前に確認プロンプトを表示 (v3.26より)

v3.26 よりタスク実行前に確認用のプロンプトを表示できるようになった

## サンプル

```yaml
# https://taskfile.dev

version: '3'

vars:
  GREETING: hello world

tasks:
  default:
    prompt: 実行します。よろしいですか？
    cmds:
      - echo "{{.GREETING}}"
    silent: true

```

実行すると以下のようになります。

```sh
$ task
task: "実行します。よろしいですか？" [y/N]: y
hello world
```

## 強制的に y で進めるには？

CI環境などで入力を得られない場合は ```--yes``` オプションを付与して実行する。

```sh
$ task --yes
hello world
```

## 参考情報

- [v3.26-ChangeLog](https://taskfile.dev/changelog/#v3260---2023-06-10)
