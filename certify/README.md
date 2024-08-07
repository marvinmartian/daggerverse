# Certify Dagger Module

Simple creation,signing,viewing,verifying of Certificates.

## Description

Built on top of [certigo](https://github.com/square/certigo) and [certstrap](https://github.com/square/certstrap).

The `ca` function provides a certificate manager that allows you to bootstrap your own certificate authority and public key infrastructure. A very convenient app if you don't feel like dealing with openssl, its myriad of options or config files.

The `view` function examine and validate certificates to help with debugging SSL/TLS issues.

### Usage

#### Create & Sign Certificates

Initialize Certificate Authority

```bash
dagger call ca init --common-name SomeCert \
    --country CA \
    --locality Toronto \
    --organization SomeOrg \
    export --path out
```
> You can now view the exported CA in the `out/` directory

View CA Certificate information
```bash
dagger call view cert --cert out/SomeCert.crt
```

## TODO:

* Add documentation for CSR function
* Add documentation for sign function
* Add code/documentation for revoke function
* Add view documentation for connect and verify functions

## License

This project is licensed under the Apache-2.0 License - see the LICENSE file for details

## Acknowledgments

Inspired and using:
* [certigo](https://github.com/square/certigo)
* [certstrap](https://github.com/square/certigo)
