// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ec2_transit_gateway_peering_attachment", transitGatewayPeeringAttachmentResourceType)
}

// transitGatewayPeeringAttachmentResourceType returns the Terraform awscc_ec2_transit_gateway_peering_attachment resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::TransitGatewayPeeringAttachment resource type.
func transitGatewayPeeringAttachmentResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The time the transit gateway peering attachment was created.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "The time the transit gateway peering attachment was created.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"peer_account_id": {
			// Property: PeerAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the peer account",
			//   "type": "string"
			// }
			Description: "The ID of the peer account",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"peer_region": {
			// Property: PeerRegion
			// CloudFormation resource type schema:
			// {
			//   "description": "Peer Region",
			//   "type": "string"
			// }
			Description: "Peer Region",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"peer_transit_gateway_id": {
			// Property: PeerTransitGatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the peer transit gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the peer transit gateway.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
			//   "type": "string"
			// }
			Description: "The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The status of the transit gateway peering attachment.",
			//   "properties": {
			//     "Code": {
			//       "description": "The status code.",
			//       "type": "string"
			//     },
			//     "Message": {
			//       "description": "The status message, if applicable.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The status of the transit gateway peering attachment.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"code": {
						// Property: Code
						Description: "The status code.",
						Type:        types.StringType,
						Optional:    true,
					},
					"message": {
						// Property: Message
						Description: "The status message, if applicable.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags for the transit gateway peering attachment.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The tags for the transit gateway peering attachment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
						Type:        types.StringType,
						Optional:    true,
					},
					"value": {
						// Property: Value
						Description: "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"transit_gateway_attachment_id": {
			// Property: TransitGatewayAttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway peering attachment.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway peering attachment.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
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
		Description: "The AWS::EC2::TransitGatewayPeeringAttachment type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayPeeringAttachment").WithTerraformTypeName("awscc_ec2_transit_gateway_peering_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"code":                          "Code",
		"creation_time":                 "CreationTime",
		"key":                           "Key",
		"message":                       "Message",
		"peer_account_id":               "PeerAccountId",
		"peer_region":                   "PeerRegion",
		"peer_transit_gateway_id":       "PeerTransitGatewayId",
		"state":                         "State",
		"status":                        "Status",
		"tags":                          "Tags",
		"transit_gateway_attachment_id": "TransitGatewayAttachmentId",
		"transit_gateway_id":            "TransitGatewayId",
		"value":                         "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
