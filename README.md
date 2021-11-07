# cert-installer
This application downloads TLS certificates from AWS Secrets Manager and installs them based on user configuration.

At the time of writing this, the arguments required are as follows:

```
./bin/cert-installer \
  -download-dir [download directory] \
  -region [aws region] \
  -secret-name [secret name]
  ```
### Example:
```
./bin/cert-installer \
  -download-dir /tmp/certs \
  -region us-east-1 \
  -secret-name my-super-secret
```
## How to build
`./scripts/build.sh`

## Notes
At the time of writing this, this application expects the payload string to be a JSON object containing the following keys (and their string values):
- `private.key`
- `certificate.crt`
- `ca_bundle.crt`