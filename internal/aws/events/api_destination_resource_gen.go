// Code generated by generators/resource/main.go; DO NOT EDIT.

package events

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
	registry.AddResourceTypeFactory("awscc_events_api_destination", apiDestinationResourceType)
}

// apiDestinationResourceType returns the Terraform awscc_events_api_destination resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Events::ApiDestination resource type.
func apiDestinationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The arn of the api destination.",
			//   "type": "string"
			// }
			Description: "The arn of the api destination.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The arn of the connection.",
			//   "type": "string"
			// }
			Description: "The arn of the connection.",
			Type:        types.StringType,
			Required:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(512),
			},
		},
		"http_method": {
			// Property: HttpMethod
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "GET",
			//     "HEAD",
			//     "POST",
			//     "OPTIONS",
			//     "PUT",
			//     "DELETE",
			//     "PATCH"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"GET",
					"HEAD",
					"POST",
					"OPTIONS",
					"PUT",
					"DELETE",
					"PATCH",
				}),
			},
		},
		"invocation_endpoint": {
			// Property: InvocationEndpoint
			// CloudFormation resource type schema:
			// {
			//   "description": "Url endpoint to invoke.",
			//   "type": "string"
			// }
			Description: "Url endpoint to invoke.",
			Type:        types.StringType,
			Required:    true,
		},
		"invocation_rate_limit_per_second": {
			// Property: InvocationRateLimitPerSecond
			// CloudFormation resource type schema:
			// {
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntAtLeast(1),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the apiDestination.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the apiDestination.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Events::ApiDestination.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::ApiDestination").WithTerraformTypeName("awscc_events_api_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                              "Arn",
		"connection_arn":                   "ConnectionArn",
		"description":                      "Description",
		"http_method":                      "HttpMethod",
		"invocation_endpoint":              "InvocationEndpoint",
		"invocation_rate_limit_per_second": "InvocationRateLimitPerSecond",
		"name":                             "Name",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
