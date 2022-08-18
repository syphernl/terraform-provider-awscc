// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

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
	registry.AddResourceTypeFactory("awscc_iotwireless_network_analyzer_configuration", networkAnalyzerConfigurationResourceType)
}

// networkAnalyzerConfigurationResourceType returns the Terraform awscc_iotwireless_network_analyzer_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::NetworkAnalyzerConfiguration resource type.
func networkAnalyzerConfigurationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn for network analyzer configuration, Returned upon successful create.",
			//   "type": "string"
			// }
			Description: "Arn for network analyzer configuration, Returned upon successful create.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the new resource",
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Description: "The description of the new resource",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(2048),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the network analyzer configuration",
			//   "maxLength": 1024,
			//   "pattern": "^[a-zA-Z0-9-_]+$",
			//   "type": "string"
			// }
			Description: "Name of the network analyzer configuration",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(1024),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9-_]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"trace_content": {
			// Property: TraceContent
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Trace content for your wireless gateway and wireless device resources",
			//   "properties": {
			//     "LogLevel": {
			//       "enum": [
			//         "INFO",
			//         "ERROR",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     },
			//     "WirelessDeviceFrameInfo": {
			//       "enum": [
			//         "ENABLED",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Trace content for your wireless gateway and wireless device resources",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"log_level": {
						// Property: LogLevel
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"INFO",
								"ERROR",
								"DISABLED",
							}),
						},
					},
					"wireless_device_frame_info": {
						// Property: WirelessDeviceFrameInfo
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ENABLED",
								"DISABLED",
							}),
						},
					},
				},
			),
			Optional: true,
		},
		"wireless_devices": {
			// Property: WirelessDevices
			// CloudFormation resource type schema:
			// {
			//   "description": "List of wireless gateway resources that have been added to the network analyzer configuration",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 250,
			//   "type": "array"
			// }
			Description: "List of wireless gateway resources that have been added to the network analyzer configuration",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(250),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"wireless_gateways": {
			// Property: WirelessGateways
			// CloudFormation resource type schema:
			// {
			//   "description": "List of wireless gateway resources that have been added to the network analyzer configuration",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 250,
			//   "type": "array"
			// }
			Description: "List of wireless gateway resources that have been added to the network analyzer configuration",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(250),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
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
		Description: "Create and manage NetworkAnalyzerConfiguration resource.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::NetworkAnalyzerConfiguration").WithTerraformTypeName("awscc_iotwireless_network_analyzer_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                        "Arn",
		"description":                "Description",
		"key":                        "Key",
		"log_level":                  "LogLevel",
		"name":                       "Name",
		"tags":                       "Tags",
		"trace_content":              "TraceContent",
		"value":                      "Value",
		"wireless_device_frame_info": "WirelessDeviceFrameInfo",
		"wireless_devices":           "WirelessDevices",
		"wireless_gateways":          "WirelessGateways",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
