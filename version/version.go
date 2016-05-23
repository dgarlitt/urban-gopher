package version

// Version - App version during build
var Version string

// Commit - Git commit used for build
var Commit string

// Branch - Git branch used for build
var Branch string

// SetVersion - can only be set once
func SetVersion(v string) {
	if Version == "" {
		Version = v
	}
}

// SetCommit - can only be set once
func SetCommit(c string) {
	if Commit == "" {
		Commit = c
	}
}

// SetBranch - can only be set once
func SetBranch(b string) {
	if Branch == "" {
		Branch = b
	}
}
