// Code generated by generators/resource/main.go; DO NOT EDIT.

package events

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
	registry.AddResourceTypeFactory("awscc_events_endpoint", endpointResourceType)
}

// endpointResourceType returns the Terraform awscc_events_endpoint resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Events::Endpoint resource type.
func endpointResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "minLength": 1,
			//   "pattern": "^arn:aws([a-z]|\\-)*:events:([a-z]|\\d|\\-)*:([0-9]{12})?:endpoint\\/[/\\.\\-_A-Za-z0-9]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "pattern": ".*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(512),
				validate.StringMatch(regexp.MustCompile(".*"), ""),
			},
		},
		"endpoint_id": {
			// Property: EndpointId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 50,
			//   "minLength": 1,
			//   "pattern": "^[A-Za-z0-9\\-]+[\\.][A-Za-z0-9\\-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"endpoint_url": {
			// Property: EndpointUrl
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^(https://)?[\\.\\-a-z0-9]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"event_buses": {
			// Property: EventBuses
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "EventBusArn": {
			//         "maxLength": 512,
			//         "minLength": 1,
			//         "pattern": "^arn:aws[a-z-]*:events:[a-z]{2}-[a-z-]+-\\d+:\\d{12}:event-bus/[\\w.-]+$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "EventBusArn"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 2,
			//   "minItems": 2,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"event_bus_arn": {
						// Property: EventBusArn
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 512),
							validate.StringMatch(regexp.MustCompile("^arn:aws[a-z-]*:events:[a-z]{2}-[a-z-]+-\\d+:\\d{12}:event-bus/[\\w.-]+$"), ""),
						},
					},
				},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(2, 2),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[\\.\\-_A-Za-z0-9]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
				validate.StringMatch(regexp.MustCompile("^[\\.\\-_A-Za-z0-9]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"replication_config": {
			// Property: ReplicationConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "State": {
			//       "enum": [
			//         "ENABLED",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "State"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"state": {
						// Property: State
						Type:     types.StringType,
						Required: true,
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
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws[a-z-]*:iam::\\d{12}:role\\/[\\w+=,.@/-]+$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^arn:aws[a-z-]*:iam::\\d{12}:role\\/[\\w+=,.@/-]+$"), ""),
			},
		},
		"routing_config": {
			// Property: RoutingConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "FailoverConfig": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Primary": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "HealthCheck": {
			//               "maxLength": 1600,
			//               "minLength": 1,
			//               "pattern": "^arn:aws([a-z]|\\-)*:route53:::healthcheck/[\\-a-z0-9]+$",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "HealthCheck"
			//           ],
			//           "type": "object"
			//         },
			//         "Secondary": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Route": {
			//               "maxLength": 20,
			//               "minLength": 9,
			//               "pattern": "^[\\-a-z0-9]+$",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Route"
			//           ],
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "Primary",
			//         "Secondary"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "FailoverConfig"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"failover_config": {
						// Property: FailoverConfig
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"primary": {
									// Property: Primary
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"health_check": {
												// Property: HealthCheck
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 1600),
													validate.StringMatch(regexp.MustCompile("^arn:aws([a-z]|\\-)*:route53:::healthcheck/[\\-a-z0-9]+$"), ""),
												},
											},
										},
									),
									Required: true,
								},
								"secondary": {
									// Property: Secondary
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"route": {
												// Property: Route
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(9, 20),
													validate.StringMatch(regexp.MustCompile("^[\\-a-z0-9]+$"), ""),
												},
											},
										},
									),
									Required: true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Required: true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ACTIVE",
			//     "CREATING",
			//     "UPDATING",
			//     "DELETING",
			//     "CREATE_FAILED",
			//     "UPDATE_FAILED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"state_reason": {
			// Property: StateReason
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "pattern": "^.*$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::Events::Endpoint.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::Endpoint").WithTerraformTypeName("awscc_events_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"description":        "Description",
		"endpoint_id":        "EndpointId",
		"endpoint_url":       "EndpointUrl",
		"event_bus_arn":      "EventBusArn",
		"event_buses":        "EventBuses",
		"failover_config":    "FailoverConfig",
		"health_check":       "HealthCheck",
		"name":               "Name",
		"primary":            "Primary",
		"replication_config": "ReplicationConfig",
		"role_arn":           "RoleArn",
		"route":              "Route",
		"routing_config":     "RoutingConfig",
		"secondary":          "Secondary",
		"state":              "State",
		"state_reason":       "StateReason",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
