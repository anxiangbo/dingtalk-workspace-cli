// Copyright 2026 Alibaba Group
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"os"
	"runtime"
	"testing"
)

func setTestHome(t *testing.T, home string) {
	t.Helper()
	t.Setenv("HOME", home)
	if runtime.GOOS == "windows" {
		// os.UserHomeDir uses USERPROFILE on Windows rather than HOME.
		t.Setenv("USERPROFILE", home)
	}
}

func TestConfigureLogLevelReleasesPreviousFileLogger(t *testing.T) {
	firstConfig := t.TempDir()
	secondConfig := t.TempDir()
	t.Cleanup(CloseFileLogger)

	t.Setenv("DWS_CONFIG_DIR", firstConfig)
	configureLogLevel(&GlobalFlags{})
	t.Setenv("DWS_CONFIG_DIR", secondConfig)
	configureLogLevel(&GlobalFlags{})

	if err := os.RemoveAll(firstConfig); err != nil {
		t.Fatalf("remove previous logger directory: %v", err)
	}
}
