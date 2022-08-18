// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_globalaccelerator_endpoint_group", endpointGroupDataSourceType)
}

// endpointGroupDataSourceType returns the Terraform awscc_globalaccelerator_endpoint_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::GlobalAccelerator::EndpointGroup resource type.
func endpointGroupDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"endpoint_configurations": {
			// Property: EndpointConfigurations
			// CloudFormation resource type schema:
			// {
			//   "description": "The list of endpoint objects.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The configuration for a given endpoint",
			//     "properties": {
			//       "ClientIPPreservationEnabled": {
			//         "default": true,
			//         "description": "true if client ip should be preserved",
			//         "type": "boolean"
			//       },
			//       "EndpointId": {
			//         "description": "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
			//         "type": "string"
			//       },
			//       "Weight": {
			//         "default": 100,
			//         "description": "The weight for the endpoint.",
			//         "maximum": 255,
			//         "minimum": 0,
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "EndpointId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "The list of endpoint objects.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"client_ip_preservation_enabled": {
						// Property: ClientIPPreservationEnabled
						Description: "true if client ip should be preserved",
						Type:        types.BoolType,
						Computed:    true,
					},
					"endpoint_id": {
						// Property: EndpointId
						Description: "Id of the endpoint. For Network/Application Load Balancer this value is the ARN.  For EIP, this value is the allocation ID.  For EC2 instances, this is the EC2 instance ID",
						Type:        types.StringType,
						Computed:    true,
					},
					"weight": {
						// Property: Weight
						Description: "The weight for the endpoint.",
						Type:        types.Int64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"endpoint_group_arn": {
			// Property: EndpointGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the endpoint group",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the endpoint group",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint_group_region": {
			// Property: EndpointGroupRegion
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the AWS Region where the endpoint group is located",
			//   "type": "string"
			// }
			Description: "The name of the AWS Region where the endpoint group is located",
			Type:        types.StringType,
			Computed:    true,
		},
		"health_check_interval_seconds": {
			// Property: HealthCheckIntervalSeconds
			// CloudFormation resource type schema:
			// {
			//   "default": 30,
			//   "description": "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
			//   "type": "integer"
			// }
			Description: "The time in seconds between each health check for an endpoint. Must be a value of 10 or 30",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"health_check_path": {
			// Property: HealthCheckPath
			// CloudFormation resource type schema:
			// {
			//   "default": "/",
			//   "description": "",
			//   "type": "string"
			// }
			Description: "",
			Type:        types.StringType,
			Computed:    true,
		},
		"health_check_port": {
			// Property: HealthCheckPort
			// CloudFormation resource type schema:
			// {
			//   "default": -1,
			//   "description": "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			//   "maximum": 65535,
			//   "minimum": -1,
			//   "type": "integer"
			// }
			Description: "The port that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"health_check_protocol": {
			// Property: HealthCheckProtocol
			// CloudFormation resource type schema:
			// {
			//   "default": "TCP",
			//   "description": "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			//   "enum": [
			//     "TCP",
			//     "HTTP",
			//     "HTTPS"
			//   ],
			//   "type": "string"
			// }
			Description: "The protocol that AWS Global Accelerator uses to check the health of endpoints in this endpoint group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"listener_arn": {
			// Property: ListenerArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the listener",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the listener",
			Type:        types.StringType,
			Computed:    true,
		},
		"port_overrides": {
			// Property: PortOverrides
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "listener to endpoint port mapping.",
			//     "properties": {
			//       "EndpointPort": {
			//         "description": "A network port number",
			//         "maximum": 65535,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "ListenerPort": {
			//         "description": "A network port number",
			//         "maximum": 65535,
			//         "minimum": 0,
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "ListenerPort",
			//       "EndpointPort"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoint_port": {
						// Property: EndpointPort
						Description: "A network port number",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"listener_port": {
						// Property: ListenerPort
						Description: "A network port number",
						Type:        types.Int64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"threshold_count": {
			// Property: ThresholdCount
			// CloudFormation resource type schema:
			// {
			//   "default": 3,
			//   "description": "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
			//   "type": "integer"
			// }
			Description: "The number of consecutive health checks required to set the state of the endpoint to unhealthy.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"traffic_dial_percentage": {
			// Property: TrafficDialPercentage
			// CloudFormation resource type schema:
			// {
			//   "default": 100,
			//   "description": "The percentage of traffic to sent to an AWS Region",
			//   "maximum": 100,
			//   "minimum": 0,
			//   "type": "number"
			// }
			Description: "The percentage of traffic to sent to an AWS Region",
			Type:        types.Float64Type,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::GlobalAccelerator::EndpointGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::EndpointGroup").WithTerraformTypeName("awscc_globalaccelerator_endpoint_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"client_ip_preservation_enabled": "ClientIPPreservationEnabled",
		"endpoint_configurations":        "EndpointConfigurations",
		"endpoint_group_arn":             "EndpointGroupArn",
		"endpoint_group_region":          "EndpointGroupRegion",
		"endpoint_id":                    "EndpointId",
		"endpoint_port":                  "EndpointPort",
		"health_check_interval_seconds":  "HealthCheckIntervalSeconds",
		"health_check_path":              "HealthCheckPath",
		"health_check_port":              "HealthCheckPort",
		"health_check_protocol":          "HealthCheckProtocol",
		"listener_arn":                   "ListenerArn",
		"listener_port":                  "ListenerPort",
		"port_overrides":                 "PortOverrides",
		"threshold_count":                "ThresholdCount",
		"traffic_dial_percentage":        "TrafficDialPercentage",
		"weight":                         "Weight",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
