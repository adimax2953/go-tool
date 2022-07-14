# go-tool#

Packet go-tool implements a way to use go more easily

## Install

```console
go get -u -v github.com/adimax2953/go-tool
```

## Usage

Let's start with a trivial example:

```go
package main

import (
	"github.com/adimax2953/go-tool"
	"fmt"
)

func main() {
    TgbotChatID := chatid
	TgbotToken :=  Telegram Bot Token 
	
	msg := fmt.Sprintf("\n事件：" + "山豬想睡覺了阿")
	gotool.SendToTG(TgbotChatID,TgbotToken,msg)
}
```

----------

## TODO

1. [X] Add TG Us test method.
2. [X] Improve or remove useless code.
3. [ ] Check code formatting.
