// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

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
	registry.AddResourceTypeFactory("awscc_datasync_task", taskResourceType)
}

// taskResourceType returns the Terraform awscc_datasync_task resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::Task resource type.
func taskResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cloudwatch_log_group_arn": {
			// Property: CloudWatchLogGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.",
			//   "maxLength": 562,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):logs:[a-z\\-0-9]*:[0-9]{12}:log-group:([^:\\*]*)(:\\*)?$",
			//   "type": "string"
			// }
			Description: "The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(562),
				validate.StringMatch(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):logs:[a-z\\-0-9]*:[0-9]{12}:log-group:([^:\\*]*)(:\\*)?$"), ""),
			},
		},
		"destination_location_arn": {
			// Property: DestinationLocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of an AWS storage resource's location.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//   "type": "string"
			// }
			Description: "The ARN of an AWS storage resource's location.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(128),
				validate.StringMatch(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"destination_network_interface_arns": {
			// Property: DestinationNetworkInterfaceArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Names (ARNs) of the destination ENIs (Elastic Network Interfaces) that were created for your subnet.",
			//   "insertionOrder": false,
			//   "items": {
			//     "pattern": "^arn:aws[\\-a-z]{0,}:ec2:[a-z\\-0-9]*:[0-9]{12}:network-interface/eni-[0-9a-f]+$",
			//     "type": "string"
			//   },
			//   "maxItems": 128,
			//   "type": "array"
			// }
			Description: "The Amazon Resource Names (ARNs) of the destination ENIs (Elastic Network Interfaces) that were created for your subnet.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"error_code": {
			// Property: ErrorCode
			// CloudFormation resource type schema:
			// {
			//   "description": "Errors that AWS DataSync encountered during execution of the task. You can use this error code to help troubleshoot issues.",
			//   "type": "string"
			// }
			Description: "Errors that AWS DataSync encountered during execution of the task. You can use this error code to help troubleshoot issues.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"error_detail": {
			// Property: ErrorDetail
			// CloudFormation resource type schema:
			// {
			//   "description": "Detailed description of an error that was encountered during the task execution.",
			//   "type": "string"
			// }
			Description: "Detailed description of an error that was encountered during the task execution.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"excludes": {
			// Property: Excludes
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Specifies which files folders and objects to include or exclude when transferring files from source to destination.",
			//     "properties": {
			//       "FilterType": {
			//         "description": "The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.",
			//         "enum": [
			//           "SIMPLE_PATTERN"
			//         ],
			//         "maxLength": 128,
			//         "pattern": "^[A-Z0-9_]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A single filter string that consists of the patterns to include or exclude. The patterns are delimited by \"|\".",
			//         "maxLength": 409600,
			//         "pattern": "^[^\\x00]+$",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"filter_type": {
						// Property: FilterType
						Description: "The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(128),
							validate.StringMatch(regexp.MustCompile("^[A-Z0-9_]+$"), ""),
							validate.StringInSlice([]string{
								"SIMPLE_PATTERN",
							}),
						},
					},
					"value": {
						// Property: Value
						Description: "A single filter string that consists of the patterns to include or exclude. The patterns are delimited by \"|\".",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(409600),
							validate.StringMatch(regexp.MustCompile("^[^\\x00]+$"), ""),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"includes": {
			// Property: Includes
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Specifies which files folders and objects to include or exclude when transferring files from source to destination.",
			//     "properties": {
			//       "FilterType": {
			//         "description": "The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.",
			//         "enum": [
			//           "SIMPLE_PATTERN"
			//         ],
			//         "maxLength": 128,
			//         "pattern": "^[A-Z0-9_]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A single filter string that consists of the patterns to include or exclude. The patterns are delimited by \"|\".",
			//         "maxLength": 409600,
			//         "pattern": "^[^\\x00]+$",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"filter_type": {
						// Property: FilterType
						Description: "The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(128),
							validate.StringMatch(regexp.MustCompile("^[A-Z0-9_]+$"), ""),
							validate.StringInSlice([]string{
								"SIMPLE_PATTERN",
							}),
						},
					},
					"value": {
						// Property: Value
						Description: "A single filter string that consists of the patterns to include or exclude. The patterns are delimited by \"|\".",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(409600),
							validate.StringMatch(regexp.MustCompile("^[^\\x00]+$"), ""),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of a task. This value is a text reference that is used to identify the task in the console.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
			//   "type": "string"
			// }
			Description: "The name of a task. This value is a text reference that is used to identify the task in the console.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
			},
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Represents the options that are available to control the behavior of a StartTaskExecution operation.",
			//   "properties": {
			//     "Atime": {
			//       "description": "A file metadata value that shows the last time a file was accessed (that is, when the file was read or written to).",
			//       "enum": [
			//         "NONE",
			//         "BEST_EFFORT"
			//       ],
			//       "type": "string"
			//     },
			//     "BytesPerSecond": {
			//       "description": "A value that limits the bandwidth used by AWS DataSync.",
			//       "format": "int64",
			//       "minimum": -1,
			//       "type": "integer"
			//     },
			//     "Gid": {
			//       "description": "The group ID (GID) of the file's owners.",
			//       "enum": [
			//         "NONE",
			//         "INT_VALUE",
			//         "NAME",
			//         "BOTH"
			//       ],
			//       "type": "string"
			//     },
			//     "LogLevel": {
			//       "description": "A value that determines the types of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.",
			//       "enum": [
			//         "OFF",
			//         "BASIC",
			//         "TRANSFER"
			//       ],
			//       "type": "string"
			//     },
			//     "Mtime": {
			//       "description": "A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.",
			//       "enum": [
			//         "NONE",
			//         "PRESERVE"
			//       ],
			//       "type": "string"
			//     },
			//     "ObjectTags": {
			//       "description": "A value that determines whether object tags should be read from the source object store and written to the destination object store.",
			//       "enum": [
			//         "PRESERVE",
			//         "NONE"
			//       ],
			//       "type": "string"
			//     },
			//     "OverwriteMode": {
			//       "description": "A value that determines whether files at the destination should be overwritten or preserved when copying files.",
			//       "enum": [
			//         "ALWAYS",
			//         "NEVER"
			//       ],
			//       "type": "string"
			//     },
			//     "PosixPermissions": {
			//       "description": "A value that determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file.",
			//       "enum": [
			//         "NONE",
			//         "PRESERVE"
			//       ],
			//       "type": "string"
			//     },
			//     "PreserveDeletedFiles": {
			//       "description": "A value that specifies whether files in the destination that don't exist in the source file system should be preserved.",
			//       "enum": [
			//         "PRESERVE",
			//         "REMOVE"
			//       ],
			//       "type": "string"
			//     },
			//     "PreserveDevices": {
			//       "description": "A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and recreate the files with that device name and metadata on the destination.",
			//       "enum": [
			//         "NONE",
			//         "PRESERVE"
			//       ],
			//       "type": "string"
			//     },
			//     "SecurityDescriptorCopyFlags": {
			//       "description": "A value that determines which components of the SMB security descriptor are copied during transfer.",
			//       "enum": [
			//         "NONE",
			//         "OWNER_DACL",
			//         "OWNER_DACL_SACL"
			//       ],
			//       "type": "string"
			//     },
			//     "TaskQueueing": {
			//       "description": "A value that determines whether tasks should be queued before executing the tasks.",
			//       "enum": [
			//         "ENABLED",
			//         "DISABLED"
			//       ],
			//       "type": "string"
			//     },
			//     "TransferMode": {
			//       "description": "A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location.",
			//       "enum": [
			//         "CHANGED",
			//         "ALL"
			//       ],
			//       "type": "string"
			//     },
			//     "Uid": {
			//       "description": "The user ID (UID) of the file's owner.",
			//       "enum": [
			//         "NONE",
			//         "INT_VALUE",
			//         "NAME",
			//         "BOTH"
			//       ],
			//       "type": "string"
			//     },
			//     "VerifyMode": {
			//       "description": "A value that determines whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred.",
			//       "enum": [
			//         "POINT_IN_TIME_CONSISTENT",
			//         "ONLY_FILES_TRANSFERRED",
			//         "NONE"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Represents the options that are available to control the behavior of a StartTaskExecution operation.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"atime": {
						// Property: Atime
						Description: "A file metadata value that shows the last time a file was accessed (that is, when the file was read or written to).",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"BEST_EFFORT",
							}),
						},
					},
					"bytes_per_second": {
						// Property: BytesPerSecond
						Description: "A value that limits the bandwidth used by AWS DataSync.",
						Type:        types.Int64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntAtLeast(-1),
						},
					},
					"gid": {
						// Property: Gid
						Description: "The group ID (GID) of the file's owners.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"INT_VALUE",
								"NAME",
								"BOTH",
							}),
						},
					},
					"log_level": {
						// Property: LogLevel
						Description: "A value that determines the types of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"OFF",
								"BASIC",
								"TRANSFER",
							}),
						},
					},
					"mtime": {
						// Property: Mtime
						Description: "A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"PRESERVE",
							}),
						},
					},
					"object_tags": {
						// Property: ObjectTags
						Description: "A value that determines whether object tags should be read from the source object store and written to the destination object store.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PRESERVE",
								"NONE",
							}),
						},
					},
					"overwrite_mode": {
						// Property: OverwriteMode
						Description: "A value that determines whether files at the destination should be overwritten or preserved when copying files.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ALWAYS",
								"NEVER",
							}),
						},
					},
					"posix_permissions": {
						// Property: PosixPermissions
						Description: "A value that determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"PRESERVE",
							}),
						},
					},
					"preserve_deleted_files": {
						// Property: PreserveDeletedFiles
						Description: "A value that specifies whether files in the destination that don't exist in the source file system should be preserved.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PRESERVE",
								"REMOVE",
							}),
						},
					},
					"preserve_devices": {
						// Property: PreserveDevices
						Description: "A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and recreate the files with that device name and metadata on the destination.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"PRESERVE",
							}),
						},
					},
					"security_descriptor_copy_flags": {
						// Property: SecurityDescriptorCopyFlags
						Description: "A value that determines which components of the SMB security descriptor are copied during transfer.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"OWNER_DACL",
								"OWNER_DACL_SACL",
							}),
						},
					},
					"task_queueing": {
						// Property: TaskQueueing
						Description: "A value that determines whether tasks should be queued before executing the tasks.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ENABLED",
								"DISABLED",
							}),
						},
					},
					"transfer_mode": {
						// Property: TransferMode
						Description: "A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"CHANGED",
								"ALL",
							}),
						},
					},
					"uid": {
						// Property: Uid
						Description: "The user ID (UID) of the file's owner.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"NONE",
								"INT_VALUE",
								"NAME",
								"BOTH",
							}),
						},
					},
					"verify_mode": {
						// Property: VerifyMode
						Description: "A value that determines whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"POINT_IN_TIME_CONSISTENT",
								"ONLY_FILES_TRANSFERRED",
								"NONE",
							}),
						},
					},
				},
			),
			Optional: true,
		},
		"schedule": {
			// Property: Schedule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies the schedule you want your task to use for repeated executions.",
			//   "properties": {
			//     "ScheduleExpression": {
			//       "description": "A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location",
			//       "maxLength": 256,
			//       "pattern": "^[a-zA-Z0-9\\ \\_\\*\\?\\,\\|\\^\\-\\/\\#\\s\\(\\)\\+]*$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "ScheduleExpression"
			//   ],
			//   "type": "object"
			// }
			Description: "Specifies the schedule you want your task to use for repeated executions.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"schedule_expression": {
						// Property: ScheduleExpression
						Description: "A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\ \\_\\*\\?\\,\\|\\^\\-\\/\\#\\s\\(\\)\\+]*$"), ""),
						},
					},
				},
			),
			Optional: true,
		},
		"source_location_arn": {
			// Property: SourceLocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the source location for the task.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//   "type": "string"
			// }
			Description: "The ARN of the source location for the task.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(128),
				validate.StringMatch(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"source_network_interface_arns": {
			// Property: SourceNetworkInterfaceArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Names (ARNs) of the source ENIs (Elastic Network Interfaces) that were created for your subnet.",
			//   "insertionOrder": false,
			//   "items": {
			//     "pattern": "^arn:aws[\\-a-z]{0,}:ec2:[a-z\\-0-9]*:[0-9]{12}:network-interface/eni-[0-9a-f]+$",
			//     "type": "string"
			//   },
			//   "maxItems": 128,
			//   "type": "array"
			// }
			Description: "The Amazon Resource Names (ARNs) of the source ENIs (Elastic Network Interfaces) that were created for your subnet.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the task that was described.",
			//   "enum": [
			//     "AVAILABLE",
			//     "CREATING",
			//     "QUEUED",
			//     "RUNNING",
			//     "UNAVAILABLE"
			//   ],
			//   "type": "string"
			// }
			Description: "The status of the task that was described.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
						},
					},
				},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(50),
			},
		},
		"task_arn": {
			// Property: TaskArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the task.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]*:[0-9]{12}:task/task-[0-9a-f]{17}$",
			//   "type": "string"
			// }
			Description: "The ARN of the task.",
			Type:        types.StringType,
			Computed:    true,
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
		Description: "Resource schema for AWS::DataSync::Task.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::Task").WithTerraformTypeName("awscc_datasync_task")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"atime":                              "Atime",
		"bytes_per_second":                   "BytesPerSecond",
		"cloudwatch_log_group_arn":           "CloudWatchLogGroupArn",
		"destination_location_arn":           "DestinationLocationArn",
		"destination_network_interface_arns": "DestinationNetworkInterfaceArns",
		"error_code":                         "ErrorCode",
		"error_detail":                       "ErrorDetail",
		"excludes":                           "Excludes",
		"filter_type":                        "FilterType",
		"gid":                                "Gid",
		"includes":                           "Includes",
		"key":                                "Key",
		"log_level":                          "LogLevel",
		"mtime":                              "Mtime",
		"name":                               "Name",
		"object_tags":                        "ObjectTags",
		"options":                            "Options",
		"overwrite_mode":                     "OverwriteMode",
		"posix_permissions":                  "PosixPermissions",
		"preserve_deleted_files":             "PreserveDeletedFiles",
		"preserve_devices":                   "PreserveDevices",
		"schedule":                           "Schedule",
		"schedule_expression":                "ScheduleExpression",
		"security_descriptor_copy_flags":     "SecurityDescriptorCopyFlags",
		"source_location_arn":                "SourceLocationArn",
		"source_network_interface_arns":      "SourceNetworkInterfaceArns",
		"status":                             "Status",
		"tags":                               "Tags",
		"task_arn":                           "TaskArn",
		"task_queueing":                      "TaskQueueing",
		"transfer_mode":                      "TransferMode",
		"uid":                                "Uid",
		"value":                              "Value",
		"verify_mode":                        "VerifyMode",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
