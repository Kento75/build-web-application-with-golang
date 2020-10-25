# Mac で Go を動作させるための環境を構築する

バージョン管理ツールを使用しない場合、Homebrew でのインストールが簡単。

## インストール

```
# インストール(gitのインストールがまだの場合は、インストールしておく)
brew install go

# バージョン確認
go version
```

## パスの設定

Homebrew でインストール後、 `GOPATH`の設定を行う。

```
# 環境変数の追加
$ echo 'export GOPATH=$HOME/go' >> ~/.bash_profile
$ echo 'export GOROOT=/usr/local/Cellar/go/1.13.4/libexec' >> ~/.bash_profile
$ echo 'export GOTOOLDIR=/usr/local/Cellar/go/1.13.4/libexec/pkg/tool/darwin_amd64' >> ~/.bash_profile
$ echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile

# bash_profileを更新
$ source ~/.bash_profile
```

変更の反映は `go env` コマンドで確認できる。

```
go env

GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/username/Library/Caches/go-build"
GOENV="/Users/username/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/username/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/Cellar/go/1.13.4/libexec"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.13.4/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
...
```
