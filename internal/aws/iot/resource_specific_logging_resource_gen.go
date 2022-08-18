// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

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
	registry.AddResourceTypeFactory("awscc_iot_resource_specific_logging", resourceSpecificLoggingResourceType)
}

// resourceSpecificLoggingResourceType returns the Terraform awscc_iot_resource_specific_logging resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::ResourceSpecificLogging resource type.
func resourceSpecificLoggingResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"log_level": {
			// Property: LogLevel
			// CloudFormation resource type schema:
			// {
			//   "description": "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			//   "enum": [
			//     "ERROR",
			//     "WARN",
			//     "INFO",
			//     "DEBUG",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ERROR",
					"WARN",
					"INFO",
					"DEBUG",
					"DISABLED",
				}),
			},
		},
		"target_id": {
			// Property: TargetId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
			//   "maxLength": 140,
			//   "minLength": 13,
			//   "pattern": "[a-zA-Z0-9.:_-]+",
			//   "type": "string"
			// }
			Description: "Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"target_name": {
			// Property: TargetName
			// CloudFormation resource type schema:
			// {
			//   "description": "The target name.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "[a-zA-Z0-9.:_-]+",
			//   "type": "string"
			// }
			Description: "The target name.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 128),
				validate.StringMatch(regexp.MustCompile("[a-zA-Z0-9.:_-]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			// {
			//   "description": "The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID.",
			//   "enum": [
			//     "THING_GROUP",
			//     "CLIENT_ID",
			//     "SOURCE_IP",
			//     "PRINCIPAL_ID"
			//   ],
			//   "type": "string"
			// }
			Description: "The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"THING_GROUP",
					"CLIENT_ID",
					"SOURCE_IP",
					"PRINCIPAL_ID",
				}),
			},
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
		Description: "Resource-specific logging allows you to specify a logging level for a specific thing group.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ResourceSpecificLogging").WithTerraformTypeName("awscc_iot_resource_specific_logging")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"log_level":   "LogLevel",
		"target_id":   "TargetId",
		"target_name": "TargetName",
		"target_type": "TargetType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
