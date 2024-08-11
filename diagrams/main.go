// A generated module for Diagrams functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/diagrams/internal/dagger"
	"fmt"
	"strings"
)

type Diagrams struct{}

const projectRoot string = "/project/"

// Build the base images
func (m *Diagrams) Build(ctx context.Context) *dagger.Container {
	d2_dir := dag.Git("https://github.com/terrastruct/d2").Tag("v0.6.6").Tree()

	certigo := dag.Container().
		From("golang:alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-121")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedDirectory("/go/src", d2_dir).
		WithWorkdir("/go/src").
		// Terminal().
		WithExec([]string{"go", "build"})

	return dag.Container().From("alpine").
		WithWorkdir(projectRoot).
		WithFile("/usr/bin/d2", certigo.File("/go/src/d2"))
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

// Generate diagram using d2
func (m *Diagrams) D2(
	ctx context.Context,
	// File to render with D2
	file *dagger.File,
	// Theme to use for render
	theme string,
) *dagger.File {

	ctr := m.Base(ctx)
	var flags []string

	fileName, err := file.Name(ctx)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	if theme != "" {
		flags = append(flags, fmt.Sprintf("--theme=%s", theme))
	}

	filePath := fmt.Sprintf("%s%s", projectRoot, fileName)
	outFile := StripExtension(fileName)
	outPath := fmt.Sprintf("%s%s.svg", projectRoot, outFile)

	// Prepend the base command to the flags slice
	cmd := append([]string{"d2", filePath}, flags...)
	// Run the command command
	return ctr.WithMountedFile(filePath, file).WithExec(cmd).File(outPath)
}

// List available themes
func (m *Diagrams) Themes(
	ctx context.Context,
) (string, error) {

	ctr := m.Base(ctx)

	// Run the command command
	return ctr.WithExec([]string{"d2", "themes"}).Stdout(ctx)
}
