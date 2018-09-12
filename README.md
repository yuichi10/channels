# channels

lircdファイルを読み込んで自動でサーバーを構築。
apiを叩くと
```
irsend SEND_ONCE ...
```
コマンドが実行される。

とりあえず今読み込めるファイルとしては、赤外線情報が`raw code` か `begin code` のもの。 /lircd/conf.d に動かせるサンプルのlircd confファイルがある。
## build 方法
raspberry ようにbuild
```
env GOOS=linux GOARCH=arm GOARM=5 go build
```

## 実行方法
### option
```
-lircd : lircdファイルが存在するディレクトリを指定。デフォルトはカレントディレクトリのlircd.dディレクトリ
-port : 実行ポート番号。デフォルトは9999
```

### 実行コマンド例
```
./channels -lircd "/etc/lirc/lircd.conf.d"
```