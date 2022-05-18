package main

import (
	// Modules of packwiz
	"github.com/wafflecoffee/packwiz/cmd"
	_ "github.com/wafflecoffee/packwiz/curseforge"
	_ "github.com/wafflecoffee/packwiz/modrinth"
	_ "github.com/wafflecoffee/packwiz/utils"
)

func main() {
	cmd.Execute()
}
