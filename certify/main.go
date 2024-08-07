// A module for creation and viewing of certificates
//
// Simplified functions to create a CA, CSR, CRL, and certificates. Also easy viewing, verifying of existing certificates.

package main

import (
	"context"
	"dagger/certify/internal/dagger"
)

type Certify struct{}

// Build the base images
func (m *Certify) Base(ctx context.Context) *dagger.Container {
	certigo_dir := dag.Git("https://github.com/square/certigo").Branch("master").Tree()
	// get build context with dockerfile added
	certigo := dag.Container().
		From("golang:alpine").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod-121")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedDirectory("/go/src", certigo_dir).
		WithWorkdir("/go/src").
		WithExec([]string{"go", "build"})

	certstrap_dir := dag.Git("https://github.com/square/certstrap").Branch("master").Tree()
	// get build context with dockerfile added
	workspace := dag.Container().
		WithMountedDirectory("/src", certstrap_dir).
		Directory("/src")

	// build using Dockerfile and publish to registry
	certStrapCtr := dag.Container().
		Build(workspace, dagger.ContainerBuildOpts{
			Dockerfile: "Dockerfile",
		})
	return dag.Container().From("alpine").
		WithWorkdir("/project").
		WithMountedFile("/usr/bin/certstrap", certStrapCtr.File("/usr/bin/certstrap")).
		WithMountedFile("/usr/bin/certigo", certigo.File("/go/src/certigo"))

}
