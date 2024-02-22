package main

import (
	"fmt"
	"time"
	"vuln/Plugins"
	"vuln/common"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
}
