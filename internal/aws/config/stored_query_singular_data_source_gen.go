// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_config_stored_query", storedQueryDataSourceType)
}

// storedQueryDataSourceType returns the Terraform awscc_config_stored_query data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Config::StoredQuery resource type.
func storedQueryDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"query_arn": {
			// Property: QueryArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 500,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"query_description": {
			// Property: QueryDescription
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "[\\s\\S]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"query_expression": {
			// Property: QueryExpression
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 4096,
			//   "minLength": 1,
			//   "pattern": "[\\s\\S]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"query_id": {
			// Property: QueryId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 1,
			//   "pattern": "^\\S+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"query_name": {
			// Property: QueryName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9-_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the stored query.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The tags for the stored query.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
		Description: "Data Source schema for AWS::Config::StoredQuery",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::StoredQuery").WithTerraformTypeName("awscc_config_stored_query")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":               "Key",
		"query_arn":         "QueryArn",
		"query_description": "QueryDescription",
		"query_expression":  "QueryExpression",
		"query_id":          "QueryId",
		"query_name":        "QueryName",
		"tags":              "Tags",
		"value":             "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
