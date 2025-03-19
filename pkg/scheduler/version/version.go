// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"fmt"
	"runtime"
)

var (
	buildDate    = "Not provided."
	gitCommit    = "Not provided."
	gitTreeState = "Not provided."
	gitVersion   = "Not provided."
)

// PrintVersion prints versions from the array returned by Info()
func PrintVersion() {
	for _, i := range Info() {
		fmt.Printf("%v\n", i)
	}
}

// Info returns an array of various service versions
func Info() []string {
	return []string{
		fmt.Sprintf("Go Version: %s", runtime.Version()),
		fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH),

		fmt.Sprintf("buildDate: %s", buildDate),
		fmt.Sprintf("gitCommit: %s", gitCommit),
		fmt.Sprintf("gitTreeState: %s", gitTreeState),
		fmt.Sprintf("gitVersion: %s", gitVersion),
	}
}
