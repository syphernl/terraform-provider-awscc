// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package memorydb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_memorydb_subnet_group", subnetGroupDataSourceType)
}

// subnetGroupDataSourceType returns the Terraform awscc_memorydb_subnet_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::MemoryDB::SubnetGroup resource type.
func subnetGroupDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: ARN
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the subnet group.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the subnet group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional description of the subnet group.",
			//   "type": "string"
			// }
			Description: "An optional description of the subnet group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_group_name": {
			// Property: SubnetGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.",
			//   "pattern": "[a-z][a-z0-9\\-]*",
			//   "type": "string"
			// }
			Description: "The name of the subnet group. This value must be unique as it also serves as the subnet group identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of VPC subnet IDs for the subnet group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of VPC subnet IDs for the subnet group.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this subnet group.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for the tag. May not be null.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value. May be null.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this subnet group.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for the tag. May not be null.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value. May be null.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::MemoryDB::SubnetGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::MemoryDB::SubnetGroup").WithTerraformTypeName("awscc_memorydb_subnet_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":               "ARN",
		"description":       "Description",
		"key":               "Key",
		"subnet_group_name": "SubnetGroupName",
		"subnet_ids":        "SubnetIds",
		"tags":              "Tags",
		"value":             "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
