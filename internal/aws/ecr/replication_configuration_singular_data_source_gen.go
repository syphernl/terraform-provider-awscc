// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ecr_replication_configuration", replicationConfigurationDataSourceType)
}

// replicationConfigurationDataSourceType returns the Terraform awscc_ecr_replication_configuration data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECR::ReplicationConfiguration resource type.
func replicationConfigurationDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"registry_id": {
			// Property: RegistryId
			// CloudFormation resource type schema:
			// {
			//   "description": "The RegistryId associated with the aws account.",
			//   "type": "string"
			// }
			Description: "The RegistryId associated with the aws account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"replication_configuration": {
			// Property: ReplicationConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object representing the replication configuration for a registry.",
			//   "properties": {
			//     "Rules": {
			//       "description": "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain a maximum of 10 rules.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "An array of objects representing the details of a replication destination.",
			//         "properties": {
			//           "Destinations": {
			//             "description": "An array of objects representing the details of a replication destination.",
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "An array of objects representing the details of a replication destination.",
			//               "properties": {
			//                 "Region": {
			//                   "description": "A Region to replicate to.",
			//                   "pattern": "[0-9a-z-]{2,25}",
			//                   "type": "string"
			//                 },
			//                 "RegistryId": {
			//                   "description": "The account ID of the destination registry to replicate to.",
			//                   "pattern": "^[0-9]{12}$",
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Region",
			//                 "RegistryId"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 25,
			//             "minItems": 1,
			//             "type": "array"
			//           },
			//           "RepositoryFilters": {
			//             "description": "An array of objects representing the details of a repository filter.",
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "An array of objects representing the details of a repository filter.",
			//               "properties": {
			//                 "Filter": {
			//                   "description": "The repository filter to be applied for replication.",
			//                   "pattern": "^(?:[a-z0-9]+(?:[._-][a-z0-9]*)*/)*[a-z0-9]*(?:[._-][a-z0-9]*)*$",
			//                   "type": "string"
			//                 },
			//                 "FilterType": {
			//                   "description": "Type of repository filter",
			//                   "enum": [
			//                     "PREFIX_MATCH"
			//                   ],
			//                   "type": "string"
			//                 }
			//               },
			//               "required": [
			//                 "Filter",
			//                 "FilterType"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 100,
			//             "minItems": 0,
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Destinations"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 10,
			//       "minItems": 0,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "Rules"
			//   ],
			//   "type": "object"
			// }
			Description: "An object representing the replication configuration for a registry.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"rules": {
						// Property: Rules
						Description: "An array of objects representing the replication rules for a replication configuration. A replication configuration may contain a maximum of 10 rules.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"destinations": {
									// Property: Destinations
									Description: "An array of objects representing the details of a replication destination.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"region": {
												// Property: Region
												Description: "A Region to replicate to.",
												Type:        types.StringType,
												Computed:    true,
											},
											"registry_id": {
												// Property: RegistryId
												Description: "The account ID of the destination registry to replicate to.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"repository_filters": {
									// Property: RepositoryFilters
									Description: "An array of objects representing the details of a repository filter.",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"filter": {
												// Property: Filter
												Description: "The repository filter to be applied for replication.",
												Type:        types.StringType,
												Computed:    true,
											},
											"filter_type": {
												// Property: FilterType
												Description: "Type of repository filter",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ECR::ReplicationConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::ReplicationConfiguration").WithTerraformTypeName("awscc_ecr_replication_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destinations":              "Destinations",
		"filter":                    "Filter",
		"filter_type":               "FilterType",
		"region":                    "Region",
		"registry_id":               "RegistryId",
		"replication_configuration": "ReplicationConfiguration",
		"repository_filters":        "RepositoryFilters",
		"rules":                     "Rules",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
