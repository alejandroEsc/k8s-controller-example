package version

import (
	"runtime"
	"time"

	"github.com/alejandroEsc/k8s-controller-example/pkg/util"
	"github.com/juju/loggo"
)

// Copyright © 2017 The Kubicorn Authors
//
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

var (
	logger = util.GetModuleLogger("pkg.util", loggo.INFO)
	// The GitSha of the current commit (automatically set at compile time)
	GitSha string

	// The Version of the program from the VERSION file (automatically set at compile time)
	// Assume version is master so we can fetch versions from tests.
	// ldflags will automatically override this string.
	appVersion = "master"
)

// Version represents Kubicorn version.
type Version struct {
	Version   string
	GitCommit string
	BuildDate string
	GoVersion string
	GOOS      string
	GOArch    string
}

// GetVersion returns Kubicorn version.
func GetVersion() *Version {
	return &Version{
		Version:   appVersion,
		GitCommit: GitSha,
		BuildDate: time.Now().UTC().String(),
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
		GOArch:    runtime.GOARCH,
	}
}
