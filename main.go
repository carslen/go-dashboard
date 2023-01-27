package main

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

type AppVersion struct {
	CurrentVersion semver.Version
}

func main() {
	fmt.Println("lala")
	semver.MustParse("v1.2.3")
}
