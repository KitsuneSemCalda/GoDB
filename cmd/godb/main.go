package main

import "godb/internal/config"
import "fmt"

func main() {
	fmt.Printf("Linux log path: %s\n", config.DEFAULT_LINUX_LOG_PATH)
	fmt.Printf("Linux cache path: %s\n", config.DEFAULT_LINUX_CACHE_PATH)
}
