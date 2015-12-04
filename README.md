[![Build Status](https://travis-ci.org/PoolC/slack_bot.svg)](https://travis-ci.org/PoolC/slack_bot)
[![Coverage Status](https://coveralls.io/repos/PoolC/slack_bot/badge.svg?branch=master&service=github)](https://coveralls.io/github/PoolC/slack_bot?branch=master)
# slack_bot
우리들의 장난감 메우와 안즈

# Prerequirement for develop
- go
- godep (recommended)
  - `go get github.com/tools/godep`

# Prepare build environment for beginner(sample)
`go get github.com/Perlmint/goautoenv`

## Linux/OSX
```bash
git clone https://github.com/PoolC/slack_bot.git
cd slack_bot
goautoenv init
source .goenv/bin/activate
```
## Windows(powershell)
```
git clone https://github.com/PoolC/slack_bot.git
cd slack_bot
goautoenv init
.\.goenv\bin\activate.ps1
```

# Build instruction(sample)
```bash
mkdir -p .workspace/src/github.com/PoolC
ln -s `pwd` .workspace/src/github.com/PoolC/slack_bot
export GOPATH=`pwd`/.workspace
cd $GOPATH/src/github.com/PoolC/slack_bot
# instll dependencies
godep restore
# build it
go build
```

# Run testcode
```bash
go test -cover ./...
```

# Update dependencies
```bash
godep save -r
```

# command
slash command서버.  
sample로 echo server가 구현되어있음.

# bot
RTM bot.  
anzu, meu 두개의 구현체가 존재. 
