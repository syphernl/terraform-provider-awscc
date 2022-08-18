// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package panorama

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_panorama_application_instance", applicationInstanceDataSourceType)
}

// applicationInstanceDataSourceType returns the Terraform awscc_panorama_application_instance data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Panorama::ApplicationInstance resource type.
func applicationInstanceDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_instance_id": {
			// Property: ApplicationInstanceId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"application_instance_id_to_replace": {
			// Property: ApplicationInstanceIdToReplace
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"default_runtime_context_device": {
			// Property: DefaultRuntimeContextDevice
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"default_runtime_context_device_name": {
			// Property: DefaultRuntimeContextDeviceName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 0,
			//   "pattern": "^.*$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"device_id": {
			// Property: DeviceId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"health_status": {
			// Property: HealthStatus
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "RUNNING",
			//     "ERROR",
			//     "NOT_AVAILABLE"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"manifest_overrides_payload": {
			// Property: ManifestOverridesPayload
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "PayloadData": {
			//       "maxLength": 51200,
			//       "minLength": 0,
			//       "pattern": "^.+$",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"payload_data": {
						// Property: PayloadData
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"manifest_payload": {
			// Property: ManifestPayload
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "PayloadData": {
			//       "maxLength": 51200,
			//       "minLength": 1,
			//       "pattern": "^.+$",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"payload_data": {
						// Property: PayloadData
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"runtime_role_arn": {
			// Property: RuntimeRoleArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "^arn:[a-z0-9][-.a-z0-9]{0,62}:iam::[0-9]{12}:role/.+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEPLOYMENT_PENDING",
			//     "DEPLOYMENT_REQUESTED",
			//     "DEPLOYMENT_IN_PROGRESS",
			//     "DEPLOYMENT_ERROR",
			//     "DEPLOYMENT_SUCCEEDED",
			//     "REMOVAL_PENDING",
			//     "REMOVAL_REQUESTED",
			//     "REMOVAL_IN_PROGRESS",
			//     "REMOVAL_FAILED",
			//     "REMOVAL_SUCCEEDED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"status_description": {
			// Property: StatusDescription
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"status_filter": {
			// Property: StatusFilter
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DEPLOYMENT_SUCCEEDED",
			//     "DEPLOYMENT_ERROR",
			//     "REMOVAL_SUCCEEDED",
			//     "REMOVAL_FAILED",
			//     "PROCESSING_DEPLOYMENT",
			//     "PROCESSING_REMOVAL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of tags",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^.+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "^.+$",
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
			Description: "List of tags",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
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
		Description: "Data Source schema for AWS::Panorama::ApplicationInstance",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Panorama::ApplicationInstance").WithTerraformTypeName("awscc_panorama_application_instance")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_instance_id":             "ApplicationInstanceId",
		"application_instance_id_to_replace":  "ApplicationInstanceIdToReplace",
		"arn":                                 "Arn",
		"created_time":                        "CreatedTime",
		"default_runtime_context_device":      "DefaultRuntimeContextDevice",
		"default_runtime_context_device_name": "DefaultRuntimeContextDeviceName",
		"description":                         "Description",
		"device_id":                           "DeviceId",
		"health_status":                       "HealthStatus",
		"key":                                 "Key",
		"last_updated_time":                   "LastUpdatedTime",
		"manifest_overrides_payload":          "ManifestOverridesPayload",
		"manifest_payload":                    "ManifestPayload",
		"name":                                "Name",
		"payload_data":                        "PayloadData",
		"runtime_role_arn":                    "RuntimeRoleArn",
		"status":                              "Status",
		"status_description":                  "StatusDescription",
		"status_filter":                       "StatusFilter",
		"tags":                                "Tags",
		"value":                               "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
