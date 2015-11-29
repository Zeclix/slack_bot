[![Build Status](https://travis-ci.org/PoolC/slack_bot.svg)](https://travis-ci.org/PoolC/slack_bot)
[![Coverage Status](https://coveralls.io/repos/PoolC/slack_bot/badge.svg?branch=master&service=github)](https://coveralls.io/github/PoolC/slack_bot?branch=master)
# slack_bot
우리들의 장난감 메우와 안즈

# Prerequirement for develop
- go
- goenv (recommended)
  - `go get github.com/crsmithdev/goenv`
- godep (recommended)
  - `go get github.com/tools/godep`

# Build instruction
```bash
# init goenv
goenv init github.com/PoolC/slack_bot
# activate goenv
source goenv/activate
# instll dependencies
godep get github.com/PoolC/slack_bot
# build it
go build
```

# command
slash command서버.  
sample로 echo server가 구현되어있음.

# bot
RTM bot.  
anzu, meu 두개의 구현체가 존재. 
