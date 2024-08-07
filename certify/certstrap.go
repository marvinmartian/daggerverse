package main

import (
	"context"
	"dagger/certify/internal/dagger"
	"fmt"
)

type CertStrap struct {
	*dagger.Container
}

// A utility to bootstrap your own certificate authority and public key infrastructure
func (m *Certify) Ca(ctx context.Context) *CertStrap {
	ctr := m.Base(ctx)
	return &CertStrap{
		Container: ctr,
	}
}

// Create Certificate Authority, including certificate, key and extra information file.
func (cam *CertStrap) Init(
	ctx context.Context,
	// Passphrase to encrypt private key PEM block
	// +optional
	passphrase *dagger.Secret,
	// Size (in bits) of RSA keypair to generate (example: 4096) (default: 4096)
	// +optional
	// +default="4096"
	bits string,
	// Elliptic curve name. Must be one of P-521, Ed25519, P-224, P-256, P-384.
	// +optional
	curve string,
	// How long until the certificate expires (example: 1 year 2 days 3 months 4 hours) (default: "18 months")
	// +optional
	expires string,
	// Sets the Organization (O) field of the certificate
	// +optional
	// +default="SomeOrg"
	organization string,
	// Sets the Organizational Unit (OU) field of the certificate
	// +optional
	organizationalUnit string,
	// Sets the Country (C) field of the certificate
	// +optional
	country string,
	// Sets the Common Name (CN) field of the certificate
	// +optional
	// +default="SomeCert"
	commonName string,
	// Sets the State/Province (ST) field of the certificate
	// +optional
	province string,
	// Sets the Locality (L) field of the certificate
	// +optional
	locality string,
) (*dagger.Directory, error) {

	var flags []string

	ctr := cam.Container

	if passphrase != nil {
		passPhraseText, err := passphrase.Plaintext(ctx)
		if err != nil {
			return nil, err
		}
		flags = append(flags, fmt.Sprintf("--passphrase=%s", passPhraseText))
	} else {
		flags = append(flags, fmt.Sprintf("--passphrase=%s", ""))
	}
	if commonName != "" {
		flags = append(flags, fmt.Sprintf("--common-name=%s", commonName))
	}
	if curve != "" {
		flags = append(flags, fmt.Sprintf("--curve=%s", curve))
	}
	if bits != "" {
		flags = append(flags, fmt.Sprintf("--key-bits=%s", bits))
	}
	if expires != "" {
		flags = append(flags, fmt.Sprintf("--expires=%s", expires))
	}
	if organization != "" {
		flags = append(flags, fmt.Sprintf("--organization=%s", expires))
	}
	if organizationalUnit != "" {
		flags = append(flags, fmt.Sprintf("--organization-unit=%s", organizationalUnit))
	}
	if country != "" {
		flags = append(flags, fmt.Sprintf("--country=%s", country))
	}
	if province != "" {
		flags = append(flags, fmt.Sprintf("--province=%s", province))
	}
	if locality != "" {
		flags = append(flags, fmt.Sprintf("--locality=%s", locality))
	}

	// Prepend the base command to the flags slice
	cmd := append([]string{"certstrap", "init"}, flags...)
	// Run the connect command
	return ctr.WithExec(cmd).Directory("out"), nil
}

func (cam *CertStrap) Request(
	ctx context.Context,
	// Directory containing any previously generated CA,csr,crl,etc files
	// +optional
	fileDir *dagger.Directory,
	// Passphrase to encrypt private key PEM block
	// +optional
	passphrase *dagger.Secret,
	// Size (in bits) of RSA keypair to generate (example: 4096) (default: 4096)
	// +optional
	// +default="4096"
	bits string,
	// Elliptic curve name. Must be one of P-521, Ed25519, P-224, P-256, P-384.
	// +optional
	curve string,
	// How long until the certificate expires (example: 1 year 2 days 3 months 4 hours) (default: "18 months")
	// +optional
	expires string,
	// Sets the Organization (O) field of the certificate
	// +optional
	organization string,
	// Sets the Organizational Unit (OU) field of the certificate
	// +optional
	organizationalUnit string,
	// Sets the Country (C) field of the certificate
	// +optional
	country string,
	// Sets the Common Name (CN) field of the certificate
	commonName string,
	// Sets the State/Province (ST) field of the certificate
	// +optional
	province string,
	// Sets the Locality (L) field of the certificate
	// +optional
	locality string,
	// IP addresses to add as subject alt name (comma separated)
	// +optional
	ip string,
	// DNS entries to add as subject alt name (comma separated)
	// +optional
	domain string,
	// URI values to add as subject alt name (comma separated)
	// +optional
	uri string,
) (*dagger.Directory, error) {

	var flags []string

	ctr := cam.Container

	if passphrase != nil {
		passPhraseText, err := passphrase.Plaintext(ctx)
		if err != nil {
			return nil, err
		}
		flags = append(flags, fmt.Sprintf("--passphrase=%s", passPhraseText))
	} else {
		flags = append(flags, fmt.Sprintf("--passphrase=%s", ""))
	}

	if fileDir != nil {
		ctr = ctr.WithMountedDirectory(fmt.Sprintf("/project/out"), fileDir)
	}

	if commonName != "" {
		flags = append(flags, fmt.Sprintf("--common-name=%s", commonName))
	}
	if curve != "" {
		flags = append(flags, fmt.Sprintf("--curve=%s", curve))
	}
	if bits != "" {
		flags = append(flags, fmt.Sprintf("--key-bits=%s", bits))
	}
	if expires != "" {
		flags = append(flags, fmt.Sprintf("--expires=%s", expires))
	}
	if organization != "" {
		flags = append(flags, fmt.Sprintf("--organization=%s", expires))
	}
	if organizationalUnit != "" {
		flags = append(flags, fmt.Sprintf("--organization-unit=%s", organizationalUnit))
	}
	if country != "" {
		flags = append(flags, fmt.Sprintf("--country=%s", country))
	}
	if province != "" {
		flags = append(flags, fmt.Sprintf("--province=%s", province))
	}
	if locality != "" {
		flags = append(flags, fmt.Sprintf("--locality=%s", locality))
	}
	if ip != "" {
		flags = append(flags, fmt.Sprintf("--ip=%s", ip))
	}
	if domain != "" {
		flags = append(flags, fmt.Sprintf("--domain=%s", domain))
	}
	if uri != "" {
		flags = append(flags, fmt.Sprintf("--uri=%s", uri))
	}

	// Prepend the base command to the flags slice
	cmd := append([]string{"certstrap", "request-cert"}, flags...)
	// Run the connect command
	return ctr.WithExec(cmd).Directory("out"), nil
}

/*
--passphrase value   Passphrase to decrypt private-key PEM block of CA
--expires value      How long until the certificate expires (example: 1 year 2 days 3 months 4 hours) (default: "2 years")
--CA value           Name of CA to issue cert with
--csr value          Path to certificate request PEM file (if blank, will use --depot-path and default name)
--cert value         Path to certificate output PEM file (if blank, will use --depot-path and default name)
--stdout             Print certificate to stdout in addition to saving file
--intermediate       Whether generated certificate should be a intermediate
--path-length value  Maximum number of non-self-issued intermediate certificates that may follow this CA certificate in a valid certification path (default: 0)
*/

// Sign certificate request with CA, and generate certificate for the host.
func (cam *CertStrap) Sign(
	ctx context.Context,
	// Host name for certificate
	hostName string,
	// Directory containing any previously generated CA,csr,crl,etc files
	// +optional
	fileDir *dagger.Directory,
	// Passphrase to encrypt private key PEM block
	// +optional
	passphrase *dagger.Secret,
	// How long until the certificate expires (example: 1 year 2 days 3 months 4 hours) (default: "18 months")
	// +optional
	// +default="2 years"
	expires string,
	// Name of CA to issue cert with
	ca string,
	// Path to certificate request PEM file
	// +optional
	csr string,
	// Path to certificate output PEM file
	// +optional
	cert string,
	// Whether generated certificate should be a intermediate
	// +optional
	intermediate bool,

) (*dagger.Directory, error) {

	var flags []string

	ctr := cam.Container

	// Set certificate passphrase
	if passphrase != nil {
		passPhraseText, err := passphrase.Plaintext(ctx)
		if err != nil {
			return nil, err
		}
		flags = append(flags, fmt.Sprintf("--passphrase=%s", passPhraseText))
	} else {
		flags = append(flags, fmt.Sprintf("--passphrase=%s", ""))
	}

	// Mount the file directory for CA,CSR,CRL files
	if fileDir != nil {
		ctr = ctr.WithMountedDirectory(fmt.Sprintf("/project/out"), fileDir)
	}

	if expires != "" {
		flags = append(flags, fmt.Sprintf("--expires=%s", expires))
	}

	if ca != "" {
		flags = append(flags, fmt.Sprintf("--CA=out/%s", ca))
	}

	if csr != "" {
		flags = append(flags, fmt.Sprintf("--csr=out/%s", csr))
	}

	if cert != "" {
		flags = append(flags, fmt.Sprintf("--cert=out/%s", cert))
	}

	if intermediate {
		flags = append(flags, "--intermediate")
	}

	// Prepend the base command to the flags slice
	cmd := append([]string{"certstrap", "sign", hostName}, flags...)
	// Run the connect command
	return ctr.WithExec(cmd).Directory("out"), nil
}
