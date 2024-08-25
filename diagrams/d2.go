package main

import (
	"context"
	"dagger/diagrams/internal/dagger"
	"fmt"
)


type D2 struct {
	*dagger.Container
}

// Generate diagram using D2
func (m *Diagrams) D2(
	ctx context.Context,
) *D2 {
	ctr := m.Base(ctx)
	return &D2{
		ctr,
	}
}

// Render diagram text to image
func (d2 *D2) Render(
	ctx context.Context,
	// File to render with D2
	file *dagger.File,
	// Theme to use for render
	// +default="100"
	theme string,
	// Export format
	// +default="svg"
	format string,
	// Animation interval timing for animated SVG exports
	// +optional
	animateInterval string,
) *dagger.File {
	// ctr := d2
	// return d2.Terminal()
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
	return d2.WithMountedFile(filePath, file).WithExec(cmd).File(outPath)
}

// List available themes
func (d2 *D2) Themes(
	ctx context.Context,
) (string, error) {
	// Run the command command
	return d2.WithExec([]string{"d2", "themes"}).Stdout(ctx)
}
