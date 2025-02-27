/*
Copyright The ORAS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

var (
	// Version is the current version of the oras.
	Version = "1.3.0-beta.2"
	// BuildMetadata is the extra build time data
	BuildMetadata = "unreleased"
	// GitCommit is the git sha1
	GitCommit = ""
	// GitTreeState is the state of the git tree
	GitTreeState = ""
)

// GetVersion returns the semver string of the version
func GetVersion() string {
	if BuildMetadata == "" {
		return Version
	}
	return Version + "+" + BuildMetadata
}
