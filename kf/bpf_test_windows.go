// Copyright Contributors to the L3AF Project.
// SPDX-License-Identifier: Apache-2.0
// +build WINDOWS

package kf

import (
	"fmt"
	"os"
	"strings"
)

func GetTestNonexecutablePathName() string {
    return "c:/windows/system32/drivers/etc/host"
}

func GetTestExecutablePathName() string {
    return "c:/windows/system32/net.exe"
}

func GetTestExecutablePath() string {
    return "c:/windows/system32"
}

func GetTestExecutableName() string {
	return "net.exe"
}

// assertExecutable checks for executable permissions
func assertExecutable(fPath string) error {
	_, err := os.Stat(fPath)
	if err != nil {
		return fmt.Errorf("Could not stat file: %s with error: %w", fPath, err)
	}

	// info.Mode() does not return the correct permissions on Windows,
	// it always has the 'x' permissions clear, so instead use the file suffix.
	if !strings.HasSuffix(fPath, ".exe") {
		return fmt.Errorf("File: %s, is not executable.", fPath)
	}

	return nil
}
