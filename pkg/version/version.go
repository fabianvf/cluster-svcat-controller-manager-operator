package version

import (
	"fmt"
	"github.com/openshift/cluster-svcat-controller-manager-operator/pkg/metrics"
	"k8s.io/apimachinery/pkg/version"
)

var (
	// commitFromGit is a constant representing the source version that
	// generated this build. It should be set during build via -ldflags.
	commitFromGit string
	// versionFromGit is a constant representing the version tag that
	// generated this build. It should be set during build via -ldflags.
	versionFromGit string
	// major version
	majorFromGit string
	// minor version
	minorFromGit string
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate string
	// SourceGitCommit indicates which git commit the binary was built from
	SourceGitCommit string
)

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() version.Info {
	return version.Info{
		Major:      majorFromGit,
		Minor:      minorFromGit,
		GitCommit:  commitFromGit,
		GitVersion: versionFromGit,
		BuildDate:  buildDate,
	}
}

// Commit returns a pretty string concatenation of SourceGitCommit
func Commit() string {
	return fmt.Sprintf("cluster-svcat-controller-manager-operator source git commit: %s\n", SourceGitCommit)
}

func init() {
	metrics.RegisterVersion(majorFromGit, minorFromGit, commitFromGit, versionFromGit)
}
