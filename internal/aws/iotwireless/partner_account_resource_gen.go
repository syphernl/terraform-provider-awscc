// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotwireless_partner_account", partnerAccountResourceType)
}

// partnerAccountResourceType returns the Terraform awscc_iotwireless_partner_account resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::PartnerAccount resource type.
func partnerAccountResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_linked": {
			// Property: AccountLinked
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the partner account is linked to the AWS account.",
			//   "type": "boolean"
			// }
			Description: "Whether the partner account is linked to the AWS account.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "PartnerAccount arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "PartnerAccount arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"fingerprint": {
			// Property: Fingerprint
			// CloudFormation resource type schema:
			// {
			//   "description": "The fingerprint of the Sidewalk application server private key.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The fingerprint of the Sidewalk application server private key.",
			Type:        types.StringType,
			Optional:    true,
		},
		"partner_account_id": {
			// Property: PartnerAccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "The partner account ID to disassociate from the AWS account",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "The partner account ID to disassociate from the AWS account",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"partner_type": {
			// Property: PartnerType
			// CloudFormation resource type schema:
			// {
			//   "description": "The partner type",
			//   "enum": [
			//     "Sidewalk"
			//   ],
			//   "type": "string"
			// }
			Description: "The partner type",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"Sidewalk",
				}),
			},
		},
		"sidewalk": {
			// Property: Sidewalk
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Sidewalk account credentials.",
			//   "properties": {
			//     "AppServerPrivateKey": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AppServerPrivateKey"
			//   ],
			//   "type": "object"
			// }
			Description: "The Sidewalk account credentials.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"app_server_private_key": {
						// Property: AppServerPrivateKey
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 4096),
						},
					},
				},
			),
			Optional: true,
		},
		"sidewalk_response": {
			// Property: SidewalkResponse
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Sidewalk account credentials.",
			//   "properties": {
			//     "AmazonId": {
			//       "maxLength": 2048,
			//       "type": "string"
			//     },
			//     "Arn": {
			//       "type": "string"
			//     },
			//     "Fingerprint": {
			//       "maxLength": 64,
			//       "minLength": 64,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Sidewalk account credentials.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"amazon_id": {
						// Property: AmazonId
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(2048),
						},
					},
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Optional: true,
					},
					"fingerprint": {
						// Property: Fingerprint
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(64, 64),
						},
					},
				},
			),
			Computed: true,
		},
		"sidewalk_update": {
			// Property: SidewalkUpdate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Sidewalk account credentials.",
			//   "properties": {
			//     "AppServerPrivateKey": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Sidewalk account credentials.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"app_server_private_key": {
						// Property: AppServerPrivateKey
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 4096),
						},
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the destination.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Create and manage partner account",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::PartnerAccount").WithTerraformTypeName("awscc_iotwireless_partner_account")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_linked":         "AccountLinked",
		"amazon_id":              "AmazonId",
		"app_server_private_key": "AppServerPrivateKey",
		"arn":                    "Arn",
		"fingerprint":            "Fingerprint",
		"key":                    "Key",
		"partner_account_id":     "PartnerAccountId",
		"partner_type":           "PartnerType",
		"sidewalk":               "Sidewalk",
		"sidewalk_response":      "SidewalkResponse",
		"sidewalk_update":        "SidewalkUpdate",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
