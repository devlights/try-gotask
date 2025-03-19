# ```--init``` オプションにディレクトリ名を渡せるようになった

[v3.42.1](https://taskfile.dev/changelog/#v3420---2025-03-08) にて、以下の機能追加が行われた。

```
--init now accepts a file name or directory as an argument
```

```task --init directory_name``` とすると ```directory_name/Taskfile.yml``` が生成出来るようになった。

また、デフォルトのコマンドライン実行結果が控えめにもなった。(```Made --init less verbose by default and respect --silent and --verbose flags```)
