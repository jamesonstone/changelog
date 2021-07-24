package main

import (
	"github.com/jstone28/changelog/pkg/changelog"
	"github.com/jstone28/changelog/pkg/git"
)

func main() {
	changelog.Build()
	git.Calculate()
}
