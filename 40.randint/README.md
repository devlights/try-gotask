# これは何？

[v3.44](https://taskfile.dev/changelog/#v3440---2025-06-08)にて ```randInt```, ```randIntN``` 関数が追加されました。

それぞれ、[math/rand.Int()](https://pkg.go.dev/math/rand/v2@go1.24.4#Int), [math/rand.IntN()](https://pkg.go.dev/math/rand/v2@go1.24.4#IntN)が呼び出されます。


```yaml
# https://taskfile.dev

version: '3'

tasks:
  default:
    cmds:
      - echo "{{randInt}}"      # math/rand/v2.Int()が呼び出される
      - echo "{{randIntN 10}}"  # math/rand/v2.IntN()が呼び出される
```

```sh
$ task
task: [default] echo "6197453053355761480"
6197453053355761480
task: [default] echo "0"
0
```
