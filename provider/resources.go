// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pnap

import (
	"unicode"

	// "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	prcko "github.com/phoenixnap/terraform-provider-pnap/pnap"
	"github.com/pulumi/pulumi-terraform-bridge/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/go/common/resource"
	"github.com/pulumi/pulumi/sdk/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "pnap"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := prcko.Provider()

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "pnap",
		Description:          "A PNAP package for creating and managing X cloud resources.",
		Keywords:             []string{"pulumi", "pnap"},
		License:              "Apache-2.0",
		Homepage:             "https://pulumi.io",
		Repository:           "https://github.com/phoenixnap/pulumi-pnap",
		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. An example
			// is below.

			"pnap_server": {
				Tok: makeResource(mainMod, "Server"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"status": {
						Type: "string",
					},
					"hostname": {
						Type: "string",
					},
					"description": {
						Type: "string",
					},
					"public": {
						Type: "boolean",
					},
					"private_ip_addresses": {
						Elem: &tfbridge.SchemaInfo{Type: "string"},
					},
					"public_ip_addresses": {
						Elem: &tfbridge.SchemaInfo{Type: "string"},
					},
					"os": {
						Type: "string",
					},
					"type": {
						Type: "string",
					},
					"ssh_keys": {
						Elem: &tfbridge.SchemaInfo{Type: "string"},
					},
					"location": {
						Type: "string",
					},
					"cpu": {
						Type: "string",
					},
					"ram": {
						Type: "string",
					},
					"storage": {
						Type: "string",
					},
					"action": {
						Type: "string",
					},
				},
			},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^1.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "1.12.1-preview",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "PNAP",
			},
		},
	}

	return prov
}
