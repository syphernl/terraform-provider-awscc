// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iot_account_audit_configuration", accountAuditConfigurationResourceType)
}

// accountAuditConfigurationResourceType returns the Terraform awscc_iot_account_audit_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::AccountAuditConfiguration resource type.
func accountAuditConfigurationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_id": {
			// Property: AccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "type": "string"
			// }
			Description: "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(12, 12),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"audit_check_configurations": {
			// Property: AuditCheckConfigurations
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies which audit checks are enabled and disabled for this account.",
			//   "properties": {
			//     "AuthenticatedCognitoRoleOverlyPermissiveCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "CaCertificateExpiringCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "CaCertificateKeyQualityCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ConflictingClientIdsCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "DeviceCertificateExpiringCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "DeviceCertificateKeyQualityCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "DeviceCertificateSharedCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "IotPolicyOverlyPermissiveCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "IotRoleAliasAllowsAccessToUnusedServicesCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "IotRoleAliasOverlyPermissiveCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "LoggingDisabledCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "RevokedCaCertificateStillActiveCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "RevokedDeviceCertificateStillActiveCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UnauthenticatedCognitoRoleOverlyPermissiveCheck": {
			//       "additionalProperties": false,
			//       "description": "The configuration for a specific audit check.",
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if the check is enabled.",
			//           "type": "boolean"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies which audit checks are enabled and disabled for this account.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"authenticated_cognito_role_overly_permissive_check": {
						// Property: AuthenticatedCognitoRoleOverlyPermissiveCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"ca_certificate_expiring_check": {
						// Property: CaCertificateExpiringCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"ca_certificate_key_quality_check": {
						// Property: CaCertificateKeyQualityCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"conflicting_client_ids_check": {
						// Property: ConflictingClientIdsCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"device_certificate_expiring_check": {
						// Property: DeviceCertificateExpiringCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"device_certificate_key_quality_check": {
						// Property: DeviceCertificateKeyQualityCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"device_certificate_shared_check": {
						// Property: DeviceCertificateSharedCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"iot_policy_overly_permissive_check": {
						// Property: IotPolicyOverlyPermissiveCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"iot_role_alias_allows_access_to_unused_services_check": {
						// Property: IotRoleAliasAllowsAccessToUnusedServicesCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"iot_role_alias_overly_permissive_check": {
						// Property: IotRoleAliasOverlyPermissiveCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"logging_disabled_check": {
						// Property: LoggingDisabledCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"revoked_ca_certificate_still_active_check": {
						// Property: RevokedCaCertificateStillActiveCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"revoked_device_certificate_still_active_check": {
						// Property: RevokedDeviceCertificateStillActiveCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"unauthenticated_cognito_role_overly_permissive_check": {
						// Property: UnauthenticatedCognitoRoleOverlyPermissiveCheck
						Description: "The configuration for a specific audit check.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if the check is enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
		},
		"audit_notification_target_configurations": {
			// Property: AuditNotificationTargetConfigurations
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about the targets to which audit notifications are sent.",
			//   "properties": {
			//     "Sns": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Enabled": {
			//           "description": "True if notifications to the target are enabled.",
			//           "type": "boolean"
			//         },
			//         "RoleArn": {
			//           "description": "The ARN of the role that grants permission to send notifications to the target.",
			//           "maxLength": 2048,
			//           "minLength": 20,
			//           "type": "string"
			//         },
			//         "TargetArn": {
			//           "description": "The ARN of the target (SNS topic) to which audit notifications are sent.",
			//           "maxLength": 2048,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Information about the targets to which audit notifications are sent.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sns": {
						// Property: Sns
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Description: "True if notifications to the target are enabled.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"role_arn": {
									// Property: RoleArn
									Description: "The ARN of the role that grants permission to send notifications to the target.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(20, 2048),
									},
								},
								"target_arn": {
									// Property: TargetArn
									Description: "The ARN of the target (SNS topic) to which audit notifications are sent.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(2048),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
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
		Description: "Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::AccountAuditConfiguration").WithTerraformTypeName("awscc_iot_account_audit_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                                            "AccountId",
		"audit_check_configurations":                            "AuditCheckConfigurations",
		"audit_notification_target_configurations":              "AuditNotificationTargetConfigurations",
		"authenticated_cognito_role_overly_permissive_check":    "AuthenticatedCognitoRoleOverlyPermissiveCheck",
		"ca_certificate_expiring_check":                         "CaCertificateExpiringCheck",
		"ca_certificate_key_quality_check":                      "CaCertificateKeyQualityCheck",
		"conflicting_client_ids_check":                          "ConflictingClientIdsCheck",
		"device_certificate_expiring_check":                     "DeviceCertificateExpiringCheck",
		"device_certificate_key_quality_check":                  "DeviceCertificateKeyQualityCheck",
		"device_certificate_shared_check":                       "DeviceCertificateSharedCheck",
		"enabled":                                               "Enabled",
		"iot_policy_overly_permissive_check":                    "IotPolicyOverlyPermissiveCheck",
		"iot_role_alias_allows_access_to_unused_services_check": "IotRoleAliasAllowsAccessToUnusedServicesCheck",
		"iot_role_alias_overly_permissive_check":                "IotRoleAliasOverlyPermissiveCheck",
		"logging_disabled_check":                                "LoggingDisabledCheck",
		"revoked_ca_certificate_still_active_check":             "RevokedCaCertificateStillActiveCheck",
		"revoked_device_certificate_still_active_check":         "RevokedDeviceCertificateStillActiveCheck",
		"role_arn":   "RoleArn",
		"sns":        "Sns",
		"target_arn": "TargetArn",
		"unauthenticated_cognito_role_overly_permissive_check": "UnauthenticatedCognitoRoleOverlyPermissiveCheck",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
