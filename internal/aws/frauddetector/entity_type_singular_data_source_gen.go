// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_frauddetector_entity_type", entityTypeDataSourceType)
}

// entityTypeDataSourceType returns the Terraform awscc_frauddetector_entity_type data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::FraudDetector::EntityType resource type.
func entityTypeDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The entity type ARN.",
			//   "type": "string"
			// }
			Description: "The entity type ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The timestamp when the entity type was created.",
			//   "type": "string"
			// }
			Description: "The timestamp when the entity type was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The entity type description.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The entity type description.",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The timestamp when the entity type was last updated.",
			//   "type": "string"
			// }
			Description: "The timestamp when the entity type was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the entity type.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[0-9a-z_-]+$",
			//   "type": "string"
			// }
			Description: "The name of the entity type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags associated with this entity type.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "Tags associated with this entity type.",
			Attributes: tfsdk.ListNestedAttributes(
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
		Description: "Data Source schema for AWS::FraudDetector::EntityType",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FraudDetector::EntityType").WithTerraformTypeName("awscc_frauddetector_entity_type")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "Arn",
		"created_time":      "CreatedTime",
		"description":       "Description",
		"key":               "Key",
		"last_updated_time": "LastUpdatedTime",
		"name":              "Name",
		"tags":              "Tags",
		"value":             "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
