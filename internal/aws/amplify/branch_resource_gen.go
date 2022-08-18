// Code generated by generators/resource/main.go; DO NOT EDIT.

package amplify

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_amplify_branch", branchResourceType)
}

// branchResourceType returns the Terraform awscc_amplify_branch resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Amplify::Branch resource type.
func branchResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_id": {
			// Property: AppId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "pattern": "d[a-z0-9]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 20),
				validate.StringMatch(regexp.MustCompile("d[a-z0-9]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "(?s).*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"basic_auth_config": {
			// Property: BasicAuthConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "EnableBasicAuth": {
			//       "type": "boolean"
			//     },
			//     "Password": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Username": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Username",
			//     "Password"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enable_basic_auth": {
						// Property: EnableBasicAuth
						Type:     types.BoolType,
						Optional: true,
					},
					"password": {
						// Property: Password
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
					"username": {
						// Property: Username
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			// BasicAuthConfig is a write-only property.
		},
		"branch_name": {
			// Property: BranchName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "(?s).+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
				validate.StringMatch(regexp.MustCompile("(?s).+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"build_spec": {
			// Property: BuildSpec
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 25000,
			//   "minLength": 1,
			//   "pattern": "(?s).+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 25000),
				validate.StringMatch(regexp.MustCompile("(?s).+"), ""),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "pattern": "(?s).*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1000),
				validate.StringMatch(regexp.MustCompile("(?s).*"), ""),
			},
		},
		"enable_auto_build": {
			// Property: EnableAutoBuild
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"enable_performance_mode": {
			// Property: EnablePerformanceMode
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"enable_pull_request_preview": {
			// Property: EnablePullRequestPreview
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"environment_variables": {
			// Property: EnvironmentVariables
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 255,
			//         "pattern": "(?s).*",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 5500,
			//         "pattern": "(?s).*",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Name",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(255),
							validate.StringMatch(regexp.MustCompile("(?s).*"), ""),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(5500),
							validate.StringMatch(regexp.MustCompile("(?s).*"), ""),
						},
					},
				},
			),
			Optional: true,
		},
		"pull_request_environment_name": {
			// Property: PullRequestEnvironmentName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 20,
			//   "pattern": "(?s).*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(20),
				validate.StringMatch(regexp.MustCompile("(?s).*"), ""),
			},
		},
		"stage": {
			// Property: Stage
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "EXPERIMENTAL",
			//     "BETA",
			//     "PULL_REQUEST",
			//     "PRODUCTION",
			//     "DEVELOPMENT"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"EXPERIMENTAL",
					"BETA",
					"PULL_REQUEST",
					"PRODUCTION",
					"DEVELOPMENT",
				}),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "insertionOrder": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
							validate.StringMatch(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						},
					},
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "The AWS::Amplify::Branch resource creates a new branch within an app.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Amplify::Branch").WithTerraformTypeName("awscc_amplify_branch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_id":                        "AppId",
		"arn":                           "Arn",
		"basic_auth_config":             "BasicAuthConfig",
		"branch_name":                   "BranchName",
		"build_spec":                    "BuildSpec",
		"description":                   "Description",
		"enable_auto_build":             "EnableAutoBuild",
		"enable_basic_auth":             "EnableBasicAuth",
		"enable_performance_mode":       "EnablePerformanceMode",
		"enable_pull_request_preview":   "EnablePullRequestPreview",
		"environment_variables":         "EnvironmentVariables",
		"key":                           "Key",
		"name":                          "Name",
		"password":                      "Password",
		"pull_request_environment_name": "PullRequestEnvironmentName",
		"stage":                         "Stage",
		"tags":                          "Tags",
		"username":                      "Username",
		"value":                         "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BasicAuthConfig",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
