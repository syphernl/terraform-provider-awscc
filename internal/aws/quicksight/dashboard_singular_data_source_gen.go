// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_quicksight_dashboard", dashboardDataSourceType)
}

// dashboardDataSourceType returns the Terraform awscc_quicksight_dashboard data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::QuickSight::Dashboard resource type.
func dashboardDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"aws_account_id": {
			// Property: AwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "pattern": "^[0-9]{12}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe time that this dataset was created.\u003c/p\u003e",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "<p>The time that this dataset was created.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"dashboard_id": {
			// Property: DashboardId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "[\\w\\-]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"dashboard_publish_options": {
			// Property: DashboardPublishOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eDashboard publish options.\u003c/p\u003e",
			//   "properties": {
			//     "AdHocFilteringOption": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eAd hoc (one-time) filtering option.\u003c/p\u003e",
			//       "properties": {
			//         "AvailabilityStatus": {
			//           "enum": [
			//             "ENABLED",
			//             "DISABLED"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ExportToCSVOption": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eExport to .csv option.\u003c/p\u003e",
			//       "properties": {
			//         "AvailabilityStatus": {
			//           "enum": [
			//             "ENABLED",
			//             "DISABLED"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "SheetControlsOption": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eSheet controls option.\u003c/p\u003e",
			//       "properties": {
			//         "VisibilityState": {
			//           "enum": [
			//             "EXPANDED",
			//             "COLLAPSED"
			//           ],
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>Dashboard publish options.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"ad_hoc_filtering_option": {
						// Property: AdHocFilteringOption
						Description: "<p>Ad hoc (one-time) filtering option.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"availability_status": {
									// Property: AvailabilityStatus
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"export_to_csv_option": {
						// Property: ExportToCSVOption
						Description: "<p>Export to .csv option.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"availability_status": {
									// Property: AvailabilityStatus
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"sheet_controls_option": {
						// Property: SheetControlsOption
						Description: "<p>Sheet controls option.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"visibility_state": {
									// Property: VisibilityState
									Type:     types.StringType,
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
		"last_published_time": {
			// Property: LastPublishedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe last time that this dataset was published.\u003c/p\u003e",
			//   "format": "string",
			//   "type": "string"
			// }
			Description: "<p>The last time that this dataset was published.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe last time that this dataset was updated.\u003c/p\u003e",
			//   "format": "string",
			//   "type": "string"
			// }
			Description: "<p>The last time that this dataset was updated.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe display name of the dashboard.\u003c/p\u003e",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "<p>The display name of the dashboard.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eA list of QuickSight parameters and the list's override values.\u003c/p\u003e",
			//   "properties": {
			//     "DateTimeParameters": {
			//       "description": "\u003cp\u003eDate-time parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA date-time parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eA display name for the date-time parameter.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values for the date-time parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "DecimalParameters": {
			//       "description": "\u003cp\u003eDecimal parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA decimal parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eA display name for the decimal parameter.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values for the decimal parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "number"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "IntegerParameters": {
			//       "description": "\u003cp\u003eInteger parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eAn integer parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eThe name of the integer parameter.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values for the integer parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "number"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "StringParameters": {
			//       "description": "\u003cp\u003eString parameters.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA string parameter.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eA display name for a string parameter.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "Values": {
			//             "description": "\u003cp\u003eThe values of a string parameter.\u003c/p\u003e",
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "required": [
			//           "Name",
			//           "Values"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>A list of QuickSight parameters and the list's override values.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"date_time_parameters": {
						// Property: DateTimeParameters
						Description: "<p>Date-time parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>A display name for the date-time parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values for the date-time parameter.</p>",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"decimal_parameters": {
						// Property: DecimalParameters
						Description: "<p>Decimal parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>A display name for the decimal parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values for the decimal parameter.</p>",
									Type:        types.ListType{ElemType: types.Float64Type},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"integer_parameters": {
						// Property: IntegerParameters
						Description: "<p>Integer parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>The name of the integer parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values for the integer parameter.</p>",
									Type:        types.ListType{ElemType: types.Float64Type},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"string_parameters": {
						// Property: StringParameters
						Description: "<p>String parameters.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>A display name for a string parameter.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"values": {
									// Property: Values
									Description: "<p>The values of a string parameter.</p>",
									Type:        types.ListType{ElemType: types.StringType},
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
		"permissions": {
			// Property: Permissions
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA structure that contains the permissions of the dashboard. You can use this structure\n            for granting permissions by providing a list of IAM action information for each\n            principal ARN. \u003c/p\u003e\n\n        \u003cp\u003eTo specify no permissions, omit the permissions list.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003ePermission for the resource.\u003c/p\u003e",
			//     "properties": {
			//       "Actions": {
			//         "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
			//         "items": {
			//           "type": "string"
			//         },
			//         "maxItems": 16,
			//         "minItems": 1,
			//         "type": "array"
			//       },
			//       "Principal": {
			//         "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Actions",
			//       "Principal"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 64,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>A structure that contains the permissions of the dashboard. You can use this structure\n            for granting permissions by providing a list of IAM action information for each\n            principal ARN. </p>\n\n        <p>To specify no permissions, omit the permissions list.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"actions": {
						// Property: Actions
						Description: "<p>The IAM action to grant or revoke permissions on.</p>",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"principal": {
						// Property: Principal
						Description: "<p>The Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:</p>\n        <ul>\n            <li>\n                <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) </p>\n            </li>\n         </ul>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"source_entity": {
			// Property: SourceEntity
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eDashboard source entity.\u003c/p\u003e",
			//   "properties": {
			//     "SourceTemplate": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eDashboard source template.\u003c/p\u003e",
			//       "properties": {
			//         "Arn": {
			//           "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
			//           "type": "string"
			//         },
			//         "DataSetReferences": {
			//           "description": "\u003cp\u003eDataset references.\u003c/p\u003e",
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "\u003cp\u003eDataset reference.\u003c/p\u003e",
			//             "properties": {
			//               "DataSetArn": {
			//                 "description": "\u003cp\u003eDataset Amazon Resource Name (ARN).\u003c/p\u003e",
			//                 "type": "string"
			//               },
			//               "DataSetPlaceholder": {
			//                 "description": "\u003cp\u003eDataset placeholder.\u003c/p\u003e",
			//                 "pattern": ".*\\S.*",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "DataSetArn",
			//               "DataSetPlaceholder"
			//             ],
			//             "type": "object"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "Arn",
			//         "DataSetReferences"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>Dashboard source entity.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"source_template": {
						// Property: SourceTemplate
						Description: "<p>Dashboard source template.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"data_set_references": {
									// Property: DataSetReferences
									Description: "<p>Dataset references.</p>",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"data_set_arn": {
												// Property: DataSetArn
												Description: "<p>Dataset Amazon Resource Name (ARN).</p>",
												Type:        types.StringType,
												Computed:    true,
											},
											"data_set_placeholder": {
												// Property: DataSetPlaceholder
												Description: "<p>Dataset placeholder.</p>",
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
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the\n            dashboard.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003eThe key or keys of the key-value pairs for the resource tag or tags assigned to the\n            resource.\u003c/p\u003e",
			//     "properties": {
			//       "Key": {
			//         "description": "\u003cp\u003eTag key.\u003c/p\u003e",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "\u003cp\u003eTag value.\u003c/p\u003e",
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
			//   "maxItems": 200,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>Contains a map of the key-value pairs for the resource tag or tags assigned to the\n            dashboard.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "<p>Tag key.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "<p>Tag value.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"theme_arn": {
			// Property: ThemeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If\n            you add a value for this field, it overrides the value that is used in the source\n            entity. The theme ARN must exist in the same AWS account where you create the\n            dashboard.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If\n            you add a value for this field, it overrides the value that is used in the source\n            entity. The theme ARN must exist in the same AWS account where you create the\n            dashboard.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eDashboard version.\u003c/p\u003e",
			//   "properties": {
			//     "Arn": {
			//       "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
			//       "type": "string"
			//     },
			//     "CreatedTime": {
			//       "description": "\u003cp\u003eThe time that this dashboard version was created.\u003c/p\u003e",
			//       "format": "string",
			//       "type": "string"
			//     },
			//     "DataSetArns": {
			//       "description": "\u003cp\u003eThe Amazon Resource Numbers (ARNs) for the datasets that are associated with this\n            version of the dashboard.\u003c/p\u003e",
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 100,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "Description": {
			//       "description": "\u003cp\u003eDescription.\u003c/p\u003e",
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Errors": {
			//       "description": "\u003cp\u003eErrors associated with this dashboard version.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eDashboard error.\u003c/p\u003e",
			//         "properties": {
			//           "Message": {
			//             "description": "\u003cp\u003eMessage.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "Type": {
			//             "enum": [
			//               "ACCESS_DENIED",
			//               "SOURCE_NOT_FOUND",
			//               "DATA_SET_NOT_FOUND",
			//               "INTERNAL_FAILURE",
			//               "PARAMETER_VALUE_INCOMPATIBLE",
			//               "PARAMETER_TYPE_INVALID",
			//               "PARAMETER_NOT_FOUND",
			//               "COLUMN_TYPE_MISMATCH",
			//               "COLUMN_GEOGRAPHIC_ROLE_MISMATCH",
			//               "COLUMN_REPLACEMENT_MISSING"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "Sheets": {
			//       "description": "\u003cp\u003eA list of the associated sheets with the unique identifier and name of each sheet.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA \u003ci\u003esheet\u003c/i\u003e, which is an object that contains a set of visuals that\n            are viewed together on one page in the Amazon QuickSight console. Every analysis and dashboard\n            contains at least one sheet. Each sheet contains at least one visualization widget, for\n            example a chart, pivot table, or narrative insight. Sheets can be associated with other\n            components, such as controls, filters, and so on.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eThe name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "SheetId": {
			//             "description": "\u003cp\u003eThe unique identifier associated with a sheet.\u003c/p\u003e",
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "[\\w\\-]+",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 20,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "SourceEntityArn": {
			//       "description": "\u003cp\u003eSource entity ARN.\u003c/p\u003e",
			//       "type": "string"
			//     },
			//     "Status": {
			//       "enum": [
			//         "CREATION_IN_PROGRESS",
			//         "CREATION_SUCCESSFUL",
			//         "CREATION_FAILED",
			//         "UPDATE_IN_PROGRESS",
			//         "UPDATE_SUCCESSFUL",
			//         "UPDATE_FAILED",
			//         "DELETED"
			//       ],
			//       "type": "string"
			//     },
			//     "ThemeArn": {
			//       "description": "\u003cp\u003eThe ARN of the theme associated with a version of the dashboard.\u003c/p\u003e",
			//       "type": "string"
			//     },
			//     "VersionNumber": {
			//       "description": "\u003cp\u003eVersion number for this version of the dashboard.\u003c/p\u003e",
			//       "minimum": 1,
			//       "type": "number"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>Dashboard version.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"created_time": {
						// Property: CreatedTime
						Description: "<p>The time that this dashboard version was created.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"data_set_arns": {
						// Property: DataSetArns
						Description: "<p>The Amazon Resource Numbers (ARNs) for the datasets that are associated with this\n            version of the dashboard.</p>",
						Type:        types.ListType{ElemType: types.StringType},
						Computed:    true,
					},
					"description": {
						// Property: Description
						Description: "<p>Description.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"errors": {
						// Property: Errors
						Description: "<p>Errors associated with this dashboard version.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"message": {
									// Property: Message
									Description: "<p>Message.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"type": {
									// Property: Type
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"sheets": {
						// Property: Sheets
						Description: "<p>A list of the associated sheets with the unique identifier and name of each sheet.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>The name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
								"sheet_id": {
									// Property: SheetId
									Description: "<p>The unique identifier associated with a sheet.</p>",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"source_entity_arn": {
						// Property: SourceEntityArn
						Description: "<p>Source entity ARN.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"status": {
						// Property: Status
						Type:     types.StringType,
						Computed: true,
					},
					"theme_arn": {
						// Property: ThemeArn
						Description: "<p>The ARN of the theme associated with a version of the dashboard.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"version_number": {
						// Property: VersionNumber
						Description: "<p>Version number for this version of the dashboard.</p>",
						Type:        types.Float64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"version_description": {
			// Property: VersionDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA description for the first version of the dashboard being created.\u003c/p\u003e",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "<p>A description for the first version of the dashboard being created.</p>",
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
		Description: "Data Source schema for AWS::QuickSight::Dashboard",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::QuickSight::Dashboard").WithTerraformTypeName("awscc_quicksight_dashboard")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                   "Actions",
		"ad_hoc_filtering_option":   "AdHocFilteringOption",
		"arn":                       "Arn",
		"availability_status":       "AvailabilityStatus",
		"aws_account_id":            "AwsAccountId",
		"created_time":              "CreatedTime",
		"dashboard_id":              "DashboardId",
		"dashboard_publish_options": "DashboardPublishOptions",
		"data_set_arn":              "DataSetArn",
		"data_set_arns":             "DataSetArns",
		"data_set_placeholder":      "DataSetPlaceholder",
		"data_set_references":       "DataSetReferences",
		"date_time_parameters":      "DateTimeParameters",
		"decimal_parameters":        "DecimalParameters",
		"description":               "Description",
		"errors":                    "Errors",
		"export_to_csv_option":      "ExportToCSVOption",
		"integer_parameters":        "IntegerParameters",
		"key":                       "Key",
		"last_published_time":       "LastPublishedTime",
		"last_updated_time":         "LastUpdatedTime",
		"message":                   "Message",
		"name":                      "Name",
		"parameters":                "Parameters",
		"permissions":               "Permissions",
		"principal":                 "Principal",
		"sheet_controls_option":     "SheetControlsOption",
		"sheet_id":                  "SheetId",
		"sheets":                    "Sheets",
		"source_entity":             "SourceEntity",
		"source_entity_arn":         "SourceEntityArn",
		"source_template":           "SourceTemplate",
		"status":                    "Status",
		"string_parameters":         "StringParameters",
		"tags":                      "Tags",
		"theme_arn":                 "ThemeArn",
		"type":                      "Type",
		"value":                     "Value",
		"values":                    "Values",
		"version":                   "Version",
		"version_description":       "VersionDescription",
		"version_number":            "VersionNumber",
		"visibility_state":          "VisibilityState",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
