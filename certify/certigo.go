package main

import (
	"context"
	"dagger/certify/internal/dagger"
	"fmt"
)

type Certigo struct {
	*dagger.Container
}

// A utility to examine and validate certificates to help with debugging SSL/TLS issues
func (m *Certify) View(ctx context.Context) *Certigo {
	ctr := m.Base(ctx)
	return &Certigo{
		Container: ctr,
	}
}

// Display information about a certificate from a file
func (certigo *Certigo) Cert(
	ctx context.Context,
	// Certificate File
	cert *dagger.File,
	// Password for PKCS12/JCEKS key stores
	// +optional
	passphrase *dagger.Secret,
	// Write output as PEM blocks instead of human-readable format.
	// +optional
	pem bool,
	// Write output as machine-readable JSON format.
	// +optional
	jsonFormat bool,
	// Only display the first certificate. This flag can be paired with --json or --pem.
	// +optional
	first bool,
) (string, error) {
	// ) *dagger.Container {
	var flags []string

	certFileName, err := cert.Name(ctx)
	if err != nil {
		return "", err
	}
	certMountPath := fmt.Sprintf("/project/%s", certFileName)
	ctr := certigo.Container.
		WithMountedFile(certMountPath, cert)

	if pem {
		flags = append(flags, "--pem")
	}
	if jsonFormat {
		flags = append(flags, "--json")
	}
	if first {
		flags = append(flags, "--first")
	}

	// Prepend the base command to the flags slice
	cmd := append([]string{"certigo", "dump", certMountPath}, flags...)
	// Run the command command
	return ctr.WithExec(cmd).Stdout(ctx)
}

// Verify a certificate chain from file
func (certigo *Certigo) Verify(
	ctx context.Context,
	// Certificate File
	cert *dagger.File,
	// Password for PKCS12/JCEKS key stores
	// +optional
	passphrase *dagger.Secret,
	// Server name to verify certificate against
	name string,
	// Path to CA bundle (system default if unspecified).
	// +optional
	ca *dagger.File,
	// Format of given input (PEM, DER, JCEKS, PKCS12; heuristic if missing).
	// +optional
	format string,
	// Write output as machine-readable JSON format.
	// +optional
	jsonFormat bool,
) (string, error) {
	// ) *dagger.Container {
	var flags []string

	certFileName, err := cert.Name(ctx)
	if err != nil {
		return "", err
	}
	certMountPath := fmt.Sprintf("/project/%s", certFileName)
	ctr := certigo.Container.
		WithMountedFile(certMountPath, cert)

	if passphrase != nil {
		passPhraseText, err := passphrase.Plaintext(ctx)
		if err != nil {
			return "", err
		}
		flags = append(flags, fmt.Sprintf("--p=%s", passPhraseText))
	}
	// Mount the ca file if it exists
	if ca != nil {
		caFileName, err := ca.Name(ctx)
		if err != nil {
			return "", err
		}
		ctr.WithMountedFile(fmt.Sprintf("/project/%s", caFileName), ca)
		flags = append(flags, fmt.Sprintf("--ca=/project/%s", caFileName))
	}

	if name != "" {
		flags = append(flags, fmt.Sprintf("--name=%s", name))
	}
	if jsonFormat {
		flags = append(flags, "--json")
	}

	// Prepend the base command to the flags slice
	cmd := append([]string{"certigo", "verify", certMountPath}, flags...)
	// Run the command command
	return ctr.WithExec(cmd).Stdout(ctx)
}

// Connect to a server and print its certificate(s).
func (certigo *Certigo) Connect(
	ctx context.Context,
	// Server address to connect to
	server string,
	// Server port to connect to
	// +default="443"
	port string,
	// Server name to verify certificate against
	// +optional
	name string,
	// Path to CA bundle (system default if unspecified).
	// +optional
	ca *dagger.File,
	// Certificate File
	// +optional
	certFile *dagger.File,
	// Private key for client certificate, if not in same file (PEM).
	// +optional
	key *dagger.File,
	// Enable StartTLS protocol ('ldap', 'mysql', 'postgres', 'smtp' or 'ftp').
	// +optional
	startTls string,
	// With --start-tls, sets the DB user or SMTP EHLO name.
	// +optional
	identity string,
	// Optional URI for HTTP(s) CONNECT proxy to dial connections with.
	// +optional
	proxy string,
	// Timeout for connecting to remote server (can be '5m', '1s', etc).
	// +optional
	timeout string,
	// Write output as PEM blocks instead of human-readable format.
	// +optional
	pem bool,
	// Write output as machine-readable JSON format.
	// +optional
	jsonFormat bool,
	// Only display the first certificate. This flag can be paired with --json or --pem.
	// +optional
	first bool,
	// Verify certificate chain.
	// +optional
	verify bool,
	// Name expected in the server TLS certificate. Defaults to name from SNI or, if SNI not overridden, the hostname to connect to.
	// +optional
	expectedName string,

) (string, error) {
	// ) *dagger.Container {
	var flags []string

	// Base container
	ctr := certigo.Container

	// Mount the certFile if it exists
	if certFile != nil {
		certFileName, err := certFile.Name(ctx)
		if err != nil {
			return "", err
		}
		ctr.WithMountedFile(fmt.Sprintf("/project/%s", certFileName), certFile)
		flags = append(flags, fmt.Sprintf("--cert=/project/%s", certFileName))
	}

	// Mount the ca file if it exists
	if ca != nil {
		caFileName, err := ca.Name(ctx)
		if err != nil {
			return "", err
		}
		ctr.WithMountedFile(fmt.Sprintf("/project/%s", caFileName), ca)
		flags = append(flags, fmt.Sprintf("--ca=/project/%s", caFileName))
	}

	// Override the name if set
	if name != "" {
		flags = append(flags, fmt.Sprintf("--name=%s", name))
	}

	// Mount the ca file if it exists
	if key != nil {
		keyFileName, err := key.Name(ctx)
		if err != nil {
			return "", err
		}
		ctr.WithMountedFile(fmt.Sprintf("/project/%s", keyFileName), key)
		flags = append(flags, fmt.Sprintf("--key=/project/%s", keyFileName))
	}

	if startTls != "" {
		flags = append(flags, fmt.Sprintf("--start-tls=%s", startTls))
	}
	if identity != "" {
		flags = append(flags, fmt.Sprintf("--identity=%s", identity))
	}
	if proxy != "" {
		flags = append(flags, fmt.Sprintf("--proxy=%s", proxy))
	}
	if timeout != "" {
		flags = append(flags, fmt.Sprintf("--timeout=%s", timeout))
	}
	if pem {
		flags = append(flags, "--pem")
	}
	if jsonFormat {
		flags = append(flags, "--json")
	}
	if first {
		flags = append(flags, "--first")
	}
	if verify {
		flags = append(flags, "--verify")
	}
	if expectedName != "" {
		flags = append(flags, fmt.Sprintf("--expected-name=%s", expectedName))
	}

	// Prepend the base command to the flags slice
	cmd := append([]string{"certigo", "connect"}, flags...)
	cmd = append(cmd, fmt.Sprintf("%s:%s", server, port))
	// Run the connect command
	return ctr.WithExec(cmd).Stdout(ctx)

}
