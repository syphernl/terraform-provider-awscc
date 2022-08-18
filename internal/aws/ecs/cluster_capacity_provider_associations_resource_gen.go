// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

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
	registry.AddResourceTypeFactory("awscc_ecs_cluster_capacity_provider_associations", clusterCapacityProviderAssociationsResourceType)
}

// clusterCapacityProviderAssociationsResourceType returns the Terraform awscc_ecs_cluster_capacity_provider_associations resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ECS::ClusterCapacityProviderAssociations resource type.
func clusterCapacityProviderAssociationsResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"capacity_providers": {
			// Property: CapacityProviders
			// CloudFormation resource type schema:
			// {
			//   "description": "List of capacity providers to associate with the cluster",
			//   "items": {
			//     "anyOf": [
			//       {},
			//       {}
			//     ],
			//     "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "List of capacity providers to associate with the cluster",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
		"cluster": {
			// Property: Cluster
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the cluster",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the cluster",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"default_capacity_provider_strategy": {
			// Property: DefaultCapacityProviderStrategy
			// CloudFormation resource type schema:
			// {
			//   "description": "List of capacity providers to associate with the cluster",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Base": {
			//         "maximum": 100000,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "CapacityProvider": {
			//         "anyOf": [
			//           {},
			//           {}
			//         ],
			//         "description": "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
			//         "type": "string"
			//       },
			//       "Weight": {
			//         "maximum": 1000,
			//         "minimum": 0,
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "CapacityProvider"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "List of capacity providers to associate with the cluster",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"base": {
						// Property: Base
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 100000),
						},
					},
					"capacity_provider": {
						// Property: CapacityProvider
						Description: "If using ec2 auto-scaling, the name of the associated capacity provider. Otherwise FARGATE, FARGATE_SPOT.",
						Type:        types.StringType,
						Required:    true,
					},
					"weight": {
						// Property: Weight
						Type:     types.Int64Type,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 1000),
						},
					},
				},
			),
			Required: true,
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
		Description: "Associate a set of ECS Capacity Providers with a specified ECS Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::ClusterCapacityProviderAssociations").WithTerraformTypeName("awscc_ecs_cluster_capacity_provider_associations")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"base":                               "Base",
		"capacity_provider":                  "CapacityProvider",
		"capacity_providers":                 "CapacityProviders",
		"cluster":                            "Cluster",
		"default_capacity_provider_strategy": "DefaultCapacityProviderStrategy",
		"weight":                             "Weight",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
