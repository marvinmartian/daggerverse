// A module to encode a message into a file using steganography
//
// Take a message and embed it into an image file using LSB steganography in order to produce a secret image file that will contain your message.
// Credit to https://github.com/auyer/steganography

package main

import (
	"context"
	"dagger/stego/internal/dagger"
	"fmt"
)

type Stego struct{}

// Base Image
func (m *Stego) Base() *dagger.Container {
	libFiles := dag.CurrentModule().Source().Directory("files")
	return dag.Container().From("golang:alpine").
		WithWorkdir("/go/src").
		WithMountedDirectory("/go/src", libFiles).
		WithExec([]string{"go", "mod", "tidy"}).
		WithExec([]string{"go", "build"})
}

// Encode any file into an image
func (m *Stego) Encode(
	ctx context.Context,
	// Source file
	sourceImage *dagger.File,
	// Any file you want to encode into the source image
	encodeFile *dagger.File,
) (*dagger.File, error) {
	sourceName, err := sourceImage.Name(ctx)
	if err != nil {
		return nil, err
	}

	encodeName, err := encodeFile.Name(ctx)
	if err != nil {
		return nil, err
	}

	inputImagePath := fmt.Sprintf("/input/%s", sourceName)
	inputEncodePath := fmt.Sprintf("/input/%s", encodeName)

	outImagePath := fmt.Sprintf("/input/encoded_%s", sourceName)
	ctr := m.Base().
		WithMountedFile(inputImagePath, sourceImage).
		WithMountedFile(inputEncodePath, encodeFile).
		WithExec([]string{"./stego", "-e", "-i", inputImagePath, "-mi", inputEncodePath, "-o", outImagePath})

	return ctr.File(outImagePath), nil
}

// Decode a message/file from an image
func (m *Stego) Decode(
	ctx context.Context,
	// Source file
	sourceImage *dagger.File,

) (*dagger.File, error) {
	sourceName, err := sourceImage.Name(ctx)
	if err != nil {
		return nil, err
	}

	inputImagePath := fmt.Sprintf("/input/%s", sourceName)

	outImagePath := fmt.Sprintf("/input/decoded_%s", sourceName)
	ctr := m.Base().
		WithMountedFile(inputImagePath, sourceImage).
		WithExec([]string{"./stego", "-d", "-i", inputImagePath, "-mo", outImagePath})

	return ctr.File(outImagePath), nil
}
