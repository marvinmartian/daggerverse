// Generate diagrams from code using popular tools (D2,Mermaid)
//
// This module provides you the ability to pass in your diagram script language and get an image back.
//
// The first language supported is D2 (https://d2lang.com/). Mermaid support is in progress.

package main

import (
	"context"
	"dagger/diagrams/internal/dagger"
	"strings"
)

type Diagrams struct{}

const projectRoot string = "/project/"

// Build the base images
func (m *Diagrams) Build(ctx context.Context) *dagger.Container {
	d2_dir := dag.Git("https://github.com/terrastruct/d2").Tag("v0.6.6").Tree()

	d2 := dag.Container().
		From("golang:alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-121")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedDirectory("/go/src", d2_dir).
		WithWorkdir("/go/src").
		// Terminal().
		WithExec([]string{"go", "build"})

	return dag.Container().From("alpine").
		WithWorkdir(projectRoot).
		WithFile("/usr/bin/d2", d2.File("/go/src/d2"))
}

// Setup base image
func (m *Diagrams) Base(ctx context.Context) *dagger.Container {
	return m.Build(ctx)
}

// StripExtension removes the file extension from the filename
func StripExtension(filename string) string {
	if dot := strings.LastIndex(filename, "."); dot != -1 {
		return filename[:dot]
	}
	return filename
}
