package pkg

import "fmt"

// Logo print banner to console ...
func Logo() {
	// Generate with `ansi-shadow` font at `https://www.bootschool.net/ascii`
	banner :=
		`██████╗  ██████╗  ██████╗ ██╗  ██╗     █████╗ ██████╗ ██╗
██╔══██╗██╔═══██╗██╔═══██╗██║ ██╔╝    ██╔══██╗██╔══██╗██║
██████╔╝██║   ██║██║   ██║█████╔╝     ███████║██████╔╝██║
██╔══██╗██║   ██║██║   ██║██╔═██╗     ██╔══██║██╔═══╝ ██║
██████╔╝╚██████╔╝╚██████╔╝██║  ██╗    ██║  ██║██║     ██║
╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═╝    ╚═╝  ╚═╝╚═╝     ╚═╝`
	fmt.Println(banner)
	fmt.Println("┌───────────────────────────────────────────────────────┐")
	fmt.Println("│  Author: weifeng                                      │")
	fmt.Println("│  Email: 610025829@qq.com                              │")
	fmt.Println("│  Github: https://github.com/weifeng295/jikeshijian    │")
	fmt.Println("└───────────────────────────────────────────────────────┘")
}
