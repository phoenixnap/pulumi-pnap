<h1 align="center">
  <br>
  <a href="https://phoenixnap.com/bare-metal-cloud"><img src="https://user-images.githubusercontent.com/78744488/109779287-16da8600-7c06-11eb-81a1-97bf44983d33.png" alt="phoenixnap Bare Metal Cloud" width="300"></a>
  <br>
  Pulumi phoenixNAP Provider
  <br>
</h1>

<p align="center">
The <a href="https://pulumi.com">Pulumi</a> phoenixNAP provider makes it easy to deploy and manage your <a href="https://phoenixnap.com/bare-metal-cloud">Bare Metal Cloud</a> infrastructure with familiar programming languages such as JavaScript or Python.
</p>

<p align="center">
  <a href="https://phoenixnap.com/bare-metal-cloud">Bare Metal Cloud</a> •
  <a href="https://www.pulumi.com/">Pulumi</a> •
  <a href="https://developers.phoenixnap.com/">Developers Portal</a> •
  <a href="http://phoenixnap.com/kb">Knowledge Base</a> •
  <a href="https://developers.phoenixnap.com/support">Support</a>
</p>

<p align="center">
  <a href="https://badge.fury.io/js/pulumi-pnap"><img src="https://badge.fury.io/js/pulumi-pnap.png" alt="npm version" height="18"></a>
  <a href="https://badge.fury.io/py/pulumi-pnap"><img src="https://badge.fury.io/py/pulumi-pnap.svg" alt="PyPI version" height="18"></a>
  <a href="https://badge.fury.io/nu/pulumi-pnap"><img src="https://badge.fury.io/nu/pulumi-pnap.svg" alt="NuGet version" height="18"></a>
</p>


## Requirements
-	[Bare Metal Cloud](https://bmc.phoenixnap.com) account
-	Pulumi CLI

## Creating a Bare Metal Cloud account
You need to have a Bare Metal Cloud account in order to use the ***pnap*** Pulumi provider with Bare Metal Cloud. 

1. Go to the [Bare Metal Cloud signup page](https://support.phoenixnap.com/wap-jpost3/bmcSignup).
2. Follow the prompts to set up your account.
3. Use your credentials to [log in to Bare Metal Cloud portal](https://bmc.phoenixnap.com).

:arrow_forward: **Video tutorial:** [How to Create a Bare Metal Cloud Account](https://www.youtube.com/watch?v=RLRQOisEB-k)
<br>

:arrow_forward: **Video tutorial:** [Introduction to Bare Metal Cloud](https://www.youtube.com/watch?v=8TLsqgLDMN4)

## Installing Pulumi locally
Check out this official tutorial to learn how to install Pulumi on your local machine. 

-   [How to Download and Install Pulumi](https://www.pulumi.com/docs/get-started/install/)

## Installing the provider

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm` or `yarn`:

    $ npm install pulumi-pnap

    $ yarn add pulumi-pnap

### Python

To use from Python, install using `pip`:

    $ pip install pulumi-pnap

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/phoenixnap/pulumi-pnap/sdk/go/...
    
## Authentication

You need to create a configuration file called `config.yaml` and save it in the user home directory. This file is used to authenticate access to your Bare Metal Cloud resources.

File path for Linux: `/.pnap/config.yaml`. 

File path for Windows: `\\AppData\\Roaming\\pnap\\config.yaml`.

The configuration file needs to include only two lines: 

				clientId: <enter your client id>
				clientSecret: <enter your client secret>

To get the values for the clientId and clientSecret, follow these steps:

1. [Log in to the Bare Metal Cloud portal](https://bmc.phoenixnap.com). 
2. On the left side menu, click on API Credentials. 
3. Click the Create Credentials button. 
4. Fill in the Name and Description fields, select the permissions scope and click Create. 
5. In the table, click on Actions and select View Credentials from the dropdown.  
6. Copy the values from the Client ID and Client Secret fields into your `config.yaml` file. 

## Reference

For detailed reference documentation, please visit [the API docs](https://developers.phoenixnap.com/apis).

## Bare Metal Cloud community
Become part of the Bare Metal Cloud community to get updates on new features, help us improve the platform, and engage with developers and other users. 

-   Follow [@phoenixNAP on Twitter](https://twitter.com/phoenixnap)
-   Join the [official Slack channel](https://phoenixnap.slack.com)
-   Sign up for our [Developers Monthly newsletter](https://phoenixnap.com/developers-monthly-newsletter)


### Resources
-	[Product page](https://phoenixnap.com/bare-metal-cloud)
-	[Instance pricing](https://phoenixnap.com/bare-metal-cloud/instances)
-	[YouTube tutorials](https://www.youtube.com/watch?v=8TLsqgLDMN4&list=PLWcrQnFWd54WwkHM0oPpR1BrAhxlsy1Rc&ab_channel=PhoenixNAPGlobalITServices)
-	[Developers Portal](https://developers.phoenixnap.com)
-	[Knowledge Base](https://phoenixnap.com/kb)
-	[Blog](https:/phoenixnap.com/blog)

### Documentation
-	[Official Pulumi documentation](https://www.pulumi.com/docs/)
-	[API documentation](https://developers.phoenixnap.com/docs/bmc/1/overview)

### Contact phoenixNAP
Get in touch with us if you have questions or need help with Bare Metal Cloud. 

<p align="left">
  <a href="https://twitter.com/phoenixNAP">Twitter</a> •
  <a href="https://www.facebook.com/phoenixnap">Facebook</a> •
  <a href="https://www.linkedin.com/company/phoenix-nap">LinkedIn</a> •
  <a href="https://www.instagram.com/phoenixnap">Instagram</a> •
  <a href="https://www.youtube.com/user/PhoenixNAPdatacenter">YouTube</a> •
  <a href="https://developers.phoenixnap.com/support">Email</a> 
</p>

<p align="center">
  <br>
  <a href="https://phoenixnap.com/bare-metal-cloud"><img src="https://user-images.githubusercontent.com/81640346/115243282-0c773b80-a123-11eb-9de7-59e3934a5712.jpg" alt="phoenixnap Bare Metal Cloud"></a>
</p>
