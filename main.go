// An example module building a binary and putting it in a container

package main

import (
	"context"
)

type Foo struct{}

// Builds go proj from Directory builds Container
func (m *Foo) Build(ctx context.Context, source *Directory) (*Container, error) {
	build := dag.Container().
		From("cgr.dev/chainguard/go:latest").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"build", "-o", "foo"})

	return dag.Container().
		From("cgr.dev/chainguard/wolfi-base:latest").
		WithFile("/bin/foo", build.File("/src/foo")).
		WithEntrypoint([]string{"/bin/foo"}).Sync(ctx)
}
