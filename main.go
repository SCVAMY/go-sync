package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("https://baidu.com", "", 800, 600, "--disable-sync", "--disable-translate")

	chSignal := make(chan os.Signal, 1)

	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	// select 会阻塞当前进程，只选择其中一个 case 执行
	select {
	case <-ui.Done():
	case <-chSignal:
	}

	ui.Close()
}
