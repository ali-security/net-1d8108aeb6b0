// Helper: package a Go module directory into a proxy.golang.org-format module zip.
// Usage: create_mod <module-path> <version> <dir> <out-zip>
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("usage: create_mod <module-path> <version> <dir> <out-zip>")
	}
	modPath := os.Args[1]
	version := os.Args[2]
	dir := os.Args[3]
	out := os.Args[4]

	f, err := os.Create(out)
	if err != nil {
		log.Fatalf("create out: %v", err)
	}
	defer f.Close()

	m := module.Version{Path: modPath, Version: version}
	if err := zip.CreateFromDir(f, m, dir); err != nil {
		log.Fatalf("CreateFromDir: %v", err)
	}
}
