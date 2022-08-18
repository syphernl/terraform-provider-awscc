// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_lambda_code_signing_config", codeSigningConfigDataSourceType)
}

// codeSigningConfigDataSourceType returns the Terraform awscc_lambda_code_signing_config data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Lambda::CodeSigningConfig resource type.
func codeSigningConfigDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"allowed_publishers": {
			// Property: AllowedPublishers
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			//   "properties": {
			//     "SigningProfileVersionArns": {
			//       "description": "List of Signing profile version Arns",
			//       "items": {
			//         "maxLength": 1024,
			//         "minLength": 12,
			//         "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
			//         "type": "string"
			//       },
			//       "maxItems": 20,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "SigningProfileVersionArns"
			//   ],
			//   "type": "object"
			// }
			Description: "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"signing_profile_version_arns": {
						// Property: SigningProfileVersionArns
						Description: "List of Signing profile version Arns",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"code_signing_config_arn": {
			// Property: CodeSigningConfigArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique Arn for CodeSigningConfig resource",
			//   "pattern": "arn:(aws[a-zA-Z-]*)?:lambda:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\\d{1}:\\d{12}:code-signing-config:csc-[a-z0-9]{17}",
			//   "type": "string"
			// }
			Description: "A unique Arn for CodeSigningConfig resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"code_signing_config_id": {
			// Property: CodeSigningConfigId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for CodeSigningConfig resource",
			//   "pattern": "csc-[a-zA-Z0-9-_\\.]{17}",
			//   "type": "string"
			// }
			Description: "A unique identifier for CodeSigningConfig resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"code_signing_policies": {
			// Property: CodeSigningPolicies
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Policies to control how to act if a signature is invalid",
			//   "properties": {
			//     "UntrustedArtifactOnDeployment": {
			//       "default": "Warn",
			//       "description": "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
			//       "enum": [
			//         "Warn",
			//         "Enforce"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "UntrustedArtifactOnDeployment"
			//   ],
			//   "type": "object"
			// }
			Description: "Policies to control how to act if a signature is invalid",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"untrusted_artifact_on_deployment": {
						// Property: UntrustedArtifactOnDeployment
						Description: "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the CodeSigningConfig",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the CodeSigningConfig",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Lambda::CodeSigningConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::CodeSigningConfig").WithTerraformTypeName("awscc_lambda_code_signing_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allowed_publishers":               "AllowedPublishers",
		"code_signing_config_arn":          "CodeSigningConfigArn",
		"code_signing_config_id":           "CodeSigningConfigId",
		"code_signing_policies":            "CodeSigningPolicies",
		"description":                      "Description",
		"signing_profile_version_arns":     "SigningProfileVersionArns",
		"untrusted_artifact_on_deployment": "UntrustedArtifactOnDeployment",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
