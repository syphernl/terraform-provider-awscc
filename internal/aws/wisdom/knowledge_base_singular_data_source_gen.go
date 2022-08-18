// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package wisdom

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_wisdom_knowledge_base", knowledgeBaseDataSourceType)
}

// knowledgeBaseDataSourceType returns the Terraform awscc_wisdom_knowledge_base data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Wisdom::KnowledgeBase resource type.
func knowledgeBaseDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"knowledge_base_arn": {
			// Property: KnowledgeBaseArn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^arn:[a-z-]*?:wisdom:[a-z0-9-]*?:[0-9]{12}:[a-z-]*?/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}(?:/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12})?$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"knowledge_base_id": {
			// Property: KnowledgeBaseId
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"knowledge_base_type": {
			// Property: KnowledgeBaseType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "EXTERNAL",
			//     "CUSTOM"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"rendering_configuration": {
			// Property: RenderingConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "TemplateUri": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"template_uri": {
						// Property: TemplateUri
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"server_side_encryption_configuration": {
			// Property: ServerSideEncryptionConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "KmsKeyId": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"kms_key_id": {
						// Property: KmsKeyId
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"source_configuration": {
			// Property: SourceConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "oneOf": [
			//     {
			//       "required": [
			//         "AppIntegrations"
			//       ]
			//     }
			//   ],
			//   "properties": {
			//     "AppIntegrations": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "AppIntegrationArn": {
			//           "maxLength": 2048,
			//           "minLength": 1,
			//           "pattern": "^arn:[a-z-]+?:[a-z-]+?:[a-z0-9-]*?:([0-9]{12})?:[a-zA-Z0-9-:/]+$",
			//           "type": "string"
			//         },
			//         "ObjectFields": {
			//           "insertionOrder": false,
			//           "items": {
			//             "maxLength": 4096,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "AppIntegrationArn",
			//         "ObjectFields"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"app_integrations": {
						// Property: AppIntegrations
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"app_integration_arn": {
									// Property: AppIntegrationArn
									Type:     types.StringType,
									Computed: true,
								},
								"object_fields": {
									// Property: ObjectFields
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
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
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Wisdom::KnowledgeBase",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Wisdom::KnowledgeBase").WithTerraformTypeName("awscc_wisdom_knowledge_base")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_integration_arn":                  "AppIntegrationArn",
		"app_integrations":                     "AppIntegrations",
		"description":                          "Description",
		"key":                                  "Key",
		"kms_key_id":                           "KmsKeyId",
		"knowledge_base_arn":                   "KnowledgeBaseArn",
		"knowledge_base_id":                    "KnowledgeBaseId",
		"knowledge_base_type":                  "KnowledgeBaseType",
		"name":                                 "Name",
		"object_fields":                        "ObjectFields",
		"rendering_configuration":              "RenderingConfiguration",
		"server_side_encryption_configuration": "ServerSideEncryptionConfiguration",
		"source_configuration":                 "SourceConfiguration",
		"tags":                                 "Tags",
		"template_uri":                         "TemplateUri",
		"value":                                "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
