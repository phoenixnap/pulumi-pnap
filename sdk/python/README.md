
# phoenixNAP provider

The phoenixNAP resource provider for Pulumi lets you use [phoenixNAP](https://phoenixnap.com//) resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).


## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install pulumi-pnap

or `yarn`:

    $ yarn add pulumi-pnap

### Python

To use from Python, install using `pip`:

    $ pip install pulumi-pnap

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/phoenixnap/pulumi-pnap/sdk/go/...
    
## Configuration

Configuration file config.yaml is required for authentication. File should be located in user home directory.
File path for Linux is /.pnap/config.yaml and file path for Windows is \\AppData\\Roaming\\pnap\\config.yaml

The following shows a sample config file: 

				# ===================================================== 
				#Sample yaml config file 
				# =====================================================
				# Authentication
				clientId: <enter your client id>
				clientSecret: <enter your client secret>


## Reference

For detailed reference documentation, please visit [the API docs][1].
