// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ecs_capacity_provider", capacityProviderDataSourceType)
}

// capacityProviderDataSourceType returns the Terraform awscc_ecs_capacity_provider data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECS::CapacityProvider resource type.
func capacityProviderDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"auto_scaling_group_provider": {
			// Property: AutoScalingGroupProvider
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AutoScalingGroupArn": {
			//       "type": "string"
			//     },
			//     "ManagedScaling": {
			//       "additionalProperties": false,
			//       "description": "The managed scaling settings for the Auto Scaling group capacity provider.",
			//       "properties": {
			//         "InstanceWarmupPeriod": {
			//           "type": "integer"
			//         },
			//         "MaximumScalingStepSize": {
			//           "type": "integer"
			//         },
			//         "MinimumScalingStepSize": {
			//           "type": "integer"
			//         },
			//         "Status": {
			//           "enum": [
			//             "DISABLED",
			//             "ENABLED"
			//           ],
			//           "type": "string"
			//         },
			//         "TargetCapacity": {
			//           "type": "integer"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ManagedTerminationProtection": {
			//       "enum": [
			//         "DISABLED",
			//         "ENABLED"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AutoScalingGroupArn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"auto_scaling_group_arn": {
						// Property: AutoScalingGroupArn
						Type:     types.StringType,
						Computed: true,
					},
					"managed_scaling": {
						// Property: ManagedScaling
						Description: "The managed scaling settings for the Auto Scaling group capacity provider.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"instance_warmup_period": {
									// Property: InstanceWarmupPeriod
									Type:     types.Int64Type,
									Computed: true,
								},
								"maximum_scaling_step_size": {
									// Property: MaximumScalingStepSize
									Type:     types.Int64Type,
									Computed: true,
								},
								"minimum_scaling_step_size": {
									// Property: MinimumScalingStepSize
									Type:     types.Int64Type,
									Computed: true,
								},
								"status": {
									// Property: Status
									Type:     types.StringType,
									Computed: true,
								},
								"target_capacity": {
									// Property: TargetCapacity
									Type:     types.Int64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"managed_termination_protection": {
						// Property: ManagedTerminationProtection
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
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
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
		Description: "Data Source schema for AWS::ECS::CapacityProvider",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::CapacityProvider").WithTerraformTypeName("awscc_ecs_capacity_provider")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_scaling_group_arn":         "AutoScalingGroupArn",
		"auto_scaling_group_provider":    "AutoScalingGroupProvider",
		"instance_warmup_period":         "InstanceWarmupPeriod",
		"key":                            "Key",
		"managed_scaling":                "ManagedScaling",
		"managed_termination_protection": "ManagedTerminationProtection",
		"maximum_scaling_step_size":      "MaximumScalingStepSize",
		"minimum_scaling_step_size":      "MinimumScalingStepSize",
		"name":                           "Name",
		"status":                         "Status",
		"tags":                           "Tags",
		"target_capacity":                "TargetCapacity",
		"value":                          "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
