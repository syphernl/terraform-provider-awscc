// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_cloudtrail_event_data_store", eventDataStoreDataSourceType)
}

// eventDataStoreDataSourceType returns the Terraform awscc_cloudtrail_event_data_store data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CloudTrail::EventDataStore resource type.
func eventDataStoreDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"advanced_event_selectors": {
			// Property: AdvancedEventSelectors
			// CloudFormation resource type schema:
			// {
			//   "description": "The advanced event selectors that were used to select events for the data store.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Advanced event selectors let you create fine-grained selectors for the following AWS CloudTrail event record ?elds. They help you control costs by logging only those events that are important to you.",
			//     "properties": {
			//       "FieldSelectors": {
			//         "description": "Contains all selector statements in an advanced event selector.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "A single selector statement in an advanced event selector.",
			//           "properties": {
			//             "EndsWith": {
			//               "description": "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "maxLength": 2048,
			//                 "minLength": 1,
			//                 "pattern": "(.+)",
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "Equals": {
			//               "description": "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "maxLength": 2048,
			//                 "minLength": 1,
			//                 "pattern": "(.+)",
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "Field": {
			//               "description": "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
			//               "maxLength": 1000,
			//               "minLength": 1,
			//               "pattern": "([\\w|\\d|\\.|_]+)",
			//               "type": "string"
			//             },
			//             "NotEndsWith": {
			//               "description": "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "maxLength": 2048,
			//                 "minLength": 1,
			//                 "pattern": "(.+)",
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "NotEquals": {
			//               "description": "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "maxLength": 2048,
			//                 "minLength": 1,
			//                 "pattern": "(.+)",
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "NotStartsWith": {
			//               "description": "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "maxLength": 2048,
			//                 "minLength": 1,
			//                 "pattern": "(.+)",
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             },
			//             "StartsWith": {
			//               "description": "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
			//               "insertionOrder": false,
			//               "items": {
			//                 "maxLength": 2048,
			//                 "minLength": 1,
			//                 "pattern": "(.+)",
			//                 "type": "string"
			//               },
			//               "minItems": 1,
			//               "type": "array",
			//               "uniqueItems": true
			//             }
			//           },
			//           "required": [
			//             "Field"
			//           ],
			//           "type": "object"
			//         },
			//         "minItems": 1,
			//         "type": "array",
			//         "uniqueItems": true
			//       },
			//       "Name": {
			//         "description": "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "FieldSelectors"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The advanced event selectors that were used to select events for the data store.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"field_selectors": {
						// Property: FieldSelectors
						Description: "Contains all selector statements in an advanced event selector.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"ends_with": {
									// Property: EndsWith
									Description: "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
								"equals": {
									// Property: Equals
									Description: "An operator that includes events that match the exact value of the event record field specified as the value of Field. This is the only valid operator that you can use with the readOnly, eventCategory, and resources.type fields.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
								"field": {
									// Property: Field
									Description: "A field in an event record on which to filter events to be logged. Supported fields include readOnly, eventCategory, eventSource (for management events), eventName, resources.type, and resources.ARN.",
									Type:        types.StringType,
									Computed:    true,
								},
								"not_ends_with": {
									// Property: NotEndsWith
									Description: "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
								"not_equals": {
									// Property: NotEquals
									Description: "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
								"not_starts_with": {
									// Property: NotStartsWith
									Description: "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
								"starts_with": {
									// Property: StartsWith
									Description: "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
									Type:        types.SetType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"name": {
						// Property: Name
						Description: "An optional, descriptive name for an advanced event selector, such as \"Log data events for only two S3 buckets\".",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"created_timestamp": {
			// Property: CreatedTimestamp
			// CloudFormation resource type schema:
			// {
			//   "description": "The timestamp of the event data store's creation.",
			//   "type": "string"
			// }
			Description: "The timestamp of the event data store's creation.",
			Type:        types.StringType,
			Computed:    true,
		},
		"event_data_store_arn": {
			// Property: EventDataStoreArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the event data store.",
			//   "type": "string"
			// }
			Description: "The ARN of the event data store.",
			Type:        types.StringType,
			Computed:    true,
		},
		"multi_region_enabled": {
			// Property: MultiRegionEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the event data store includes events from all regions, or only from the region in which it was created.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the event data store.",
			//   "type": "string"
			// }
			Description: "The name of the event data store.",
			Type:        types.StringType,
			Computed:    true,
		},
		"organization_enabled": {
			// Property: OrganizationEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates that an event data store is collecting logged events for an organization.",
			//   "type": "boolean"
			// }
			Description: "Indicates that an event data store is collecting logged events for an organization.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"retention_period": {
			// Property: RetentionPeriod
			// CloudFormation resource type schema:
			// {
			//   "description": "The retention period, in days.",
			//   "type": "integer"
			// }
			Description: "The retention period, in days.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of an event data store. Values are ENABLED and PENDING_DELETION.",
			//   "type": "string"
			// }
			Description: "The status of an event data store. Values are ENABLED and PENDING_DELETION.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "An arbitrary set of tags (key-value pairs) for this event data store.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"termination_protection_enabled": {
			// Property: TerminationProtectionEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the event data store is protected from termination.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the event data store is protected from termination.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"updated_timestamp": {
			// Property: UpdatedTimestamp
			// CloudFormation resource type schema:
			// {
			//   "description": "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
			//   "type": "string"
			// }
			Description: "The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
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
		Description: "Data Source schema for AWS::CloudTrail::EventDataStore",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudTrail::EventDataStore").WithTerraformTypeName("awscc_cloudtrail_event_data_store")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"advanced_event_selectors":       "AdvancedEventSelectors",
		"created_timestamp":              "CreatedTimestamp",
		"ends_with":                      "EndsWith",
		"equals":                         "Equals",
		"event_data_store_arn":           "EventDataStoreArn",
		"field":                          "Field",
		"field_selectors":                "FieldSelectors",
		"key":                            "Key",
		"multi_region_enabled":           "MultiRegionEnabled",
		"name":                           "Name",
		"not_ends_with":                  "NotEndsWith",
		"not_equals":                     "NotEquals",
		"not_starts_with":                "NotStartsWith",
		"organization_enabled":           "OrganizationEnabled",
		"retention_period":               "RetentionPeriod",
		"starts_with":                    "StartsWith",
		"status":                         "Status",
		"tags":                           "Tags",
		"termination_protection_enabled": "TerminationProtectionEnabled",
		"updated_timestamp":              "UpdatedTimestamp",
		"value":                          "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
