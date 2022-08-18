// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

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
	registry.AddResourceTypeFactory("awscc_sagemaker_pipeline", pipelineResourceType)
}

// pipelineResourceType returns the Terraform awscc_sagemaker_pipeline resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::Pipeline resource type.
func pipelineResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"parallelism_configuration": {
			// Property: ParallelismConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "MaxParallelExecutionSteps": {
			//       "description": "Maximum parallel execution steps",
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "required": [
			//     "MaxParallelExecutionSteps"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"max_parallel_execution_steps": {
						// Property: MaxParallelExecutionSteps
						Description: "Maximum parallel execution steps",
						Type:        types.Int64Type,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(1),
						},
					},
				},
			),
			Optional: true,
		},
		"pipeline_definition": {
			// Property: PipelineDefinition
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "PipelineDefinitionBody": {
			//       "description": "A specification that defines the pipeline in JSON format.",
			//       "type": "string"
			//     },
			//     "PipelineDefinitionS3Location": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Bucket": {
			//           "description": "The name of the S3 bucket where the PipelineDefinition file is stored.",
			//           "type": "string"
			//         },
			//         "ETag": {
			//           "description": "The Amazon S3 ETag (a file checksum) of the PipelineDefinition file. If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.",
			//           "type": "string"
			//         },
			//         "Key": {
			//           "description": "The file name of the PipelineDefinition file (Amazon S3 object name).",
			//           "type": "string"
			//         },
			//         "Version": {
			//           "description": "For versioning-enabled buckets, a specific version of the PipelineDefinition file.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Bucket",
			//         "Key"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"pipeline_definition_body": {
						// Property: PipelineDefinitionBody
						Description: "A specification that defines the pipeline in JSON format.",
						Type:        types.StringType,
						Optional:    true,
					},
					"pipeline_definition_s3_location": {
						// Property: PipelineDefinitionS3Location
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"bucket": {
									// Property: Bucket
									Description: "The name of the S3 bucket where the PipelineDefinition file is stored.",
									Type:        types.StringType,
									Required:    true,
								},
								"e_tag": {
									// Property: ETag
									Description: "The Amazon S3 ETag (a file checksum) of the PipelineDefinition file. If you don't specify a value, SageMaker skips ETag validation of your PipelineDefinition file.",
									Type:        types.StringType,
									Optional:    true,
								},
								"key": {
									// Property: Key
									Description: "The file name of the PipelineDefinition file (Amazon S3 object name).",
									Type:        types.StringType,
									Required:    true,
								},
								"version": {
									// Property: Version
									Description: "For versioning-enabled buckets, a specific version of the PipelineDefinition file.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"pipeline_description": {
			// Property: PipelineDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the Pipeline.",
			//   "maxLength": 3072,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The description of the Pipeline.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 3072),
			},
		},
		"pipeline_display_name": {
			// Property: PipelineDisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "The display name of the Pipeline.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
			//   "type": "string"
			// }
			Description: "The display name of the Pipeline.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*"), ""),
			},
		},
		"pipeline_name": {
			// Property: PipelineName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Pipeline.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
			//   "type": "string"
			// }
			Description: "The name of the Pipeline.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Role Arn",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
			//   "type": "string"
			// }
			Description: "Role Arn",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
				validate.StringMatch(regexp.MustCompile("^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$"), ""),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
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
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
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
		Description: "Resource Type definition for AWS::SageMaker::Pipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Pipeline").WithTerraformTypeName("awscc_sagemaker_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket":                          "Bucket",
		"e_tag":                           "ETag",
		"key":                             "Key",
		"max_parallel_execution_steps":    "MaxParallelExecutionSteps",
		"parallelism_configuration":       "ParallelismConfiguration",
		"pipeline_definition":             "PipelineDefinition",
		"pipeline_definition_body":        "PipelineDefinitionBody",
		"pipeline_definition_s3_location": "PipelineDefinitionS3Location",
		"pipeline_description":            "PipelineDescription",
		"pipeline_display_name":           "PipelineDisplayName",
		"pipeline_name":                   "PipelineName",
		"role_arn":                        "RoleArn",
		"tags":                            "Tags",
		"value":                           "Value",
		"version":                         "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
