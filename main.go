/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"embed"
	"go-embedded-svelte/cmd"
	"go-embedded-svelte/server"
)

// must use all: so we include _ files from sveltekit
//go:embed all:frontend/build
var assets embed.FS

func main() {
	server.SetAssets(assets)
	cmd.Execute()
}
