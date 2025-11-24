package bootstrap

import (
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
)

//go:embed banner.txt
var bannerString string

func StartUp(port string) {
	version := "v0.0.1"

	fmt.Println(bannerString)
	fmt.Println(version)

	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf(" Go Version    : %s\n", runtime.Version())
	fmt.Printf(" Listen Port   : %s\n", port)
	fmt.Printf(" OS/Arch       : %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf(" Start Time    : %s\n", time.Now().Format(time.RFC3339))
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
	fmt.Println("Started Server now")
}
