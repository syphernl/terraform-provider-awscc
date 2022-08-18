// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_flow_log", flowLogDataSourceType)
}

// flowLogDataSourceType returns the Terraform awscc_ec2_flow_log data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::FlowLog resource type.
func flowLogDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"deliver_logs_permission_arn": {
			// Property: DeliverLogsPermissionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
			//   "type": "string"
			// }
			Description: "The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
			Type:        types.StringType,
			Computed:    true,
		},
		"destination_options": {
			// Property: DestinationOptions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "FileFormat": {
			//       "enum": [
			//         "plain-text",
			//         "parquet"
			//       ],
			//       "type": "string"
			//     },
			//     "HiveCompatiblePartitions": {
			//       "type": "boolean"
			//     },
			//     "PerHourPartition": {
			//       "type": "boolean"
			//     }
			//   },
			//   "required": [
			//     "FileFormat",
			//     "HiveCompatiblePartitions",
			//     "PerHourPartition"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"file_format": {
						// Property: FileFormat
						Type:     types.StringType,
						Computed: true,
					},
					"hive_compatible_partitions": {
						// Property: HiveCompatiblePartitions
						Type:     types.BoolType,
						Computed: true,
					},
					"per_hour_partition": {
						// Property: PerHourPartition
						Type:     types.BoolType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The Flow Log ID",
			//   "type": "string"
			// }
			Description: "The Flow Log ID",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_destination": {
			// Property: LogDestination
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.",
			//   "type": "string"
			// }
			Description: "Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_destination_type": {
			// Property: LogDestinationType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.",
			//   "enum": [
			//     "cloud-watch-logs",
			//     "s3",
			//     "kinesis-data-firehose"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_format": {
			// Property: LogFormat
			// CloudFormation resource type schema:
			// {
			//   "description": "The fields to include in the flow log record, in the order in which they should appear.",
			//   "type": "string"
			// }
			Description: "The fields to include in the flow log record, in the order in which they should appear.",
			Type:        types.StringType,
			Computed:    true,
		},
		"log_group_name": {
			// Property: LogGroupName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
			//   "type": "string"
			// }
			Description: "The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.",
			Type:        types.StringType,
			Computed:    true,
		},
		"max_aggregation_interval": {
			// Property: MaxAggregationInterval
			// CloudFormation resource type schema:
			// {
			//   "description": "The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).",
			//   "type": "integer"
			// }
			Description: "The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the subnet, network interface, or VPC for which you want to create a flow log.",
			//   "type": "string"
			// }
			Description: "The ID of the subnet, network interface, or VPC for which you want to create a flow log.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.",
			//   "enum": [
			//     "NetworkInterface",
			//     "Subnet",
			//     "VPC"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags to apply to the flow logs.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The tags to apply to the flow logs.",
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
		"traffic_type": {
			// Property: TrafficType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.",
			//   "enum": [
			//     "ACCEPT",
			//     "ALL",
			//     "REJECT"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::FlowLog",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::FlowLog").WithTerraformTypeName("awscc_ec2_flow_log")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"deliver_logs_permission_arn": "DeliverLogsPermissionArn",
		"destination_options":         "DestinationOptions",
		"file_format":                 "FileFormat",
		"hive_compatible_partitions":  "HiveCompatiblePartitions",
		"id":                          "Id",
		"key":                         "Key",
		"log_destination":             "LogDestination",
		"log_destination_type":        "LogDestinationType",
		"log_format":                  "LogFormat",
		"log_group_name":              "LogGroupName",
		"max_aggregation_interval":    "MaxAggregationInterval",
		"per_hour_partition":          "PerHourPartition",
		"resource_id":                 "ResourceId",
		"resource_type":               "ResourceType",
		"tags":                        "Tags",
		"traffic_type":                "TrafficType",
		"value":                       "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
