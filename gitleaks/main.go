// Protect and discover secrets using Gitleaks
//
// Gitleaks is a SAST tool for detecting and preventing hardcoded secrets like passwords, api keys, and tokens in git repos. Gitleaks is an easy-to-use, all-in-one solution for detecting secrets, past or present, in your code.

package main

import (
	"context"
	"dagger/gitleaks/internal/dagger"
)

type Gitleaks struct{}

const projectRoot string = "/project/"

// Returns a gitleaks container
func (m *Gitleaks) base(ctx context.Context) *dagger.Container {
	return dag.Container().From("zricethezav/gitleaks").
		WithWorkdir(projectRoot)
}

// Detect secrets in code
func (m *Gitleaks) Detect(
	ctx context.Context,
	// git directory to scan
	dir *dagger.Directory,
	// No git option
	// +optional
	// default=false
	nogit bool,
) (string, error) {
	ctr := m.base(ctx)

	var flags []string

	if nogit {
		flags = append(flags, "--no-git")
	}

	cmd := append([]string{"gitleaks", "detect"}, flags...)

	return ctr.
		WithMountedDirectory(projectRoot, dir).
		WithExec(cmd).Stdout(ctx)
}
