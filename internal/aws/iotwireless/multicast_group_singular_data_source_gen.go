// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotwireless_multicast_group", multicastGroupDataSourceType)
}

// multicastGroupDataSourceType returns the Terraform awscc_iotwireless_multicast_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTWireless::MulticastGroup resource type.
func multicastGroupDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Multicast group arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "Multicast group arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"associate_wireless_device": {
			// Property: AssociateWirelessDevice
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device to associate. Only for update request.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Wireless device to associate. Only for update request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Multicast group description",
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Description: "Multicast group description",
			Type:        types.StringType,
			Computed:    true,
		},
		"disassociate_wireless_device": {
			// Property: DisassociateWirelessDevice
			// CloudFormation resource type schema:
			// {
			//   "description": "Wireless device to disassociate. Only for update request.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Wireless device to disassociate. Only for update request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Multicast group id. Returned after successful create.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Multicast group id. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Multicast group LoRaWAN",
			//   "properties": {
			//     "DlClass": {
			//       "description": "Multicast group LoRaWAN DL Class",
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "NumberOfDevicesInGroup": {
			//       "description": "Multicast group number of devices in group. Returned after successful read.",
			//       "type": "integer"
			//     },
			//     "NumberOfDevicesRequested": {
			//       "description": "Multicast group number of devices requested. Returned after successful read.",
			//       "type": "integer"
			//     },
			//     "RfRegion": {
			//       "description": "Multicast group LoRaWAN RF region",
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "RfRegion",
			//     "DlClass"
			//   ],
			//   "type": "object"
			// }
			Description: "Multicast group LoRaWAN",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"dl_class": {
						// Property: DlClass
						Description: "Multicast group LoRaWAN DL Class",
						Type:        types.StringType,
						Computed:    true,
					},
					"number_of_devices_in_group": {
						// Property: NumberOfDevicesInGroup
						Description: "Multicast group number of devices in group. Returned after successful read.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"number_of_devices_requested": {
						// Property: NumberOfDevicesRequested
						Description: "Multicast group number of devices requested. Returned after successful read.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"rf_region": {
						// Property: RfRegion
						Description: "Multicast group LoRaWAN RF region",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Multicast group",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Name of Multicast group",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "Multicast group status. Returned after successful read.",
			//   "type": "string"
			// }
			Description: "Multicast group status. Returned after successful read.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the Multicast group.",
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
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the Multicast group.",
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
		Description: "Data Source schema for AWS::IoTWireless::MulticastGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::MulticastGroup").WithTerraformTypeName("awscc_iotwireless_multicast_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"associate_wireless_device":    "AssociateWirelessDevice",
		"description":                  "Description",
		"disassociate_wireless_device": "DisassociateWirelessDevice",
		"dl_class":                     "DlClass",
		"id":                           "Id",
		"key":                          "Key",
		"lo_ra_wan":                    "LoRaWAN",
		"name":                         "Name",
		"number_of_devices_in_group":   "NumberOfDevicesInGroup",
		"number_of_devices_requested":  "NumberOfDevicesRequested",
		"rf_region":                    "RfRegion",
		"status":                       "Status",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
