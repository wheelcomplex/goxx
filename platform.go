package main

import (
	"fmt"
	"strings"
)

// Platform is a combination of OS/arch that can be built against.
type Platform struct {
	OS   string
	Arch string

	// Default, if true, will be included as a default build target
	// if no OS/arch is specified. We try to only set as a default popular
	// targets or targets that are generally useful. For example, Android
	// is not a default because it is quite rare that you're cross-compiling
	// something to Android AND something like Linux.
	Default bool
}

func (p *Platform) String() string {
	return fmt.Sprintf("%s/%s", p.OS, p.Arch)
}

var (
	// OsList is GOOS list
	OsList = []string{
		"darwin",
		"dragonfly",
		"freebsd",
		"linux",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"windows",
	}

	// ArchList is GOARCH list
	ArchList = []string{
		"386",
		"amd64",
		"arm",
		"arm64",
		"ppc64",
		"ppc64le",
	}

	// Platforms1_0 is the os/arch list for go 1.0
	Platforms1_0 = []Platform{
		{"darwin", "386", true},
		{"darwin", "amd64", true},
		{"linux", "386", true},
		{"linux", "amd64", true},
		{"linux", "arm", true},
		{"freebsd", "386", true},
		{"freebsd", "amd64", true},
		{"openbsd", "386", true},
		{"openbsd", "amd64", true},
		{"windows", "386", true},
		{"windows", "amd64", true},
	}

	// Platforms1_1 append the os/arch list from go 1.1
	Platforms1_1 = append(Platforms1_0, []Platform{
		{"freebsd", "arm", true},
		{"netbsd", "386", true},
		{"netbsd", "amd64", true},
		{"netbsd", "arm", true},
		{"plan9", "386", false},
	}...)

	// Platforms1_3 append the os/arch list from go 1.3
	Platforms1_3 = append(Platforms1_1, []Platform{
		{"dragonfly", "386", false},
		{"dragonfly", "amd64", false},
		{"nacl", "amd64", false},
		{"nacl", "amd64p32", false},
		{"nacl", "arm", false},
		{"solaris", "amd64", false},
	}...)

	// Platforms1_4 append the os/arch list from go 1.4
	Platforms1_4 = append(Platforms1_3, []Platform{
		{"android", "arm", false},
		{"plan9", "amd64", false},
	}...)

	// Platforms1_5 append the os/arch list from go 1.5
	Platforms1_5 = append(Platforms1_4, []Platform{
		{"darwin", "arm", false},
		{"darwin", "arm64", false},
		{"linux", "arm64", false},
		{"linux", "ppc64", false},
		{"linux", "ppc64le", false},
	}...)
)

// SupportedPlatforms returns the full list of supported platforms for
// the version of Go that is
func SupportedPlatforms(v string) []Platform {
	if strings.HasPrefix(v, "go1.0") {
		return Platforms1_0
	} else if strings.HasPrefix(v, "go1.1") {
		return Platforms1_1
	} else if strings.HasPrefix(v, "go1.3") {
		return Platforms1_3
	} else if strings.HasPrefix(v, "go1.4") {
		return Platforms1_4
	} else if strings.HasPrefix(v, "go1.5") {
		return Platforms1_5
	}

	// Assume latest
	return Platforms1_5
}
