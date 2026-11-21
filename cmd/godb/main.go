package main

import "godb/internal/config"
import "fmt"

func main() {
	fmt.Printf("Linux log path: %s\n", config.DefaultLinuxLogPath)
	fmt.Printf("Linux cache path: %s\n", config.DefaultLinuxCachePath)
}
