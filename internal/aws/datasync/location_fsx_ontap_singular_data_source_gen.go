// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_datasync_location_fsx_ontap", locationFSxONTAPDataSourceType)
}

// locationFSxONTAPDataSourceType returns the Terraform awscc_datasync_location_fsx_ontap data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataSync::LocationFSxONTAP resource type.
func locationFSxONTAPDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"fsx_filesystem_arn": {
			// Property: FsxFilesystemArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the FSx ONAP file system.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the FSx ONAP file system.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.",
			//   "maxLength": 128,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Amazon FSx ONTAP file system location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the FSx ONTAP file system that was described.",
			//   "maxLength": 4360,
			//   "pattern": "^(efs|nfs|s3|smb|hdfs|fsx[a-z0-9-]+)://[a-zA-Z0-9.:/\\-]+$",
			//   "type": "string"
			// }
			Description: "The URL of the FSx ONTAP file system that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration settings for NFS or SMB protocol.",
			//   "properties": {
			//     "NFS": {
			//       "additionalProperties": false,
			//       "description": "NFS protocol configuration for FSx ONTAP file system.",
			//       "properties": {
			//         "MountOptions": {
			//           "additionalProperties": false,
			//           "description": "The NFS mount options that DataSync can use to mount your NFS share.",
			//           "properties": {
			//             "Version": {
			//               "description": "The specific NFS version that you want DataSync to use to mount your NFS share.",
			//               "enum": [
			//                 "AUTOMATIC",
			//                 "NFS3",
			//                 "NFS4_0",
			//                 "NFS4_1"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "required": [
			//         "MountOptions"
			//       ],
			//       "type": "object"
			//     },
			//     "SMB": {
			//       "additionalProperties": false,
			//       "description": "SMB protocol configuration for FSx ONTAP file system.",
			//       "properties": {
			//         "Domain": {
			//           "description": "The name of the Windows domain that the SMB server belongs to.",
			//           "maxLength": 253,
			//           "pattern": "^([A-Za-z0-9]+[A-Za-z0-9-.]*)*[A-Za-z0-9-]*[A-Za-z0-9]$",
			//           "type": "string"
			//         },
			//         "MountOptions": {
			//           "additionalProperties": false,
			//           "description": "The mount options used by DataSync to access the SMB server.",
			//           "properties": {
			//             "Version": {
			//               "description": "The specific SMB version that you want DataSync to use to mount your SMB share.",
			//               "enum": [
			//                 "AUTOMATIC",
			//                 "SMB2",
			//                 "SMB3"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "Password": {
			//           "description": "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
			//           "maxLength": 104,
			//           "pattern": "^.{0,104}$",
			//           "type": "string"
			//         },
			//         "User": {
			//           "description": "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
			//           "maxLength": 104,
			//           "pattern": "^[^\\x5B\\x5D\\\\/:;|=,+*?]{1,104}$",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "User",
			//         "Password",
			//         "MountOptions"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Configuration settings for NFS or SMB protocol.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"nfs": {
						// Property: NFS
						Description: "NFS protocol configuration for FSx ONTAP file system.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"mount_options": {
									// Property: MountOptions
									Description: "The NFS mount options that DataSync can use to mount your NFS share.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"version": {
												// Property: Version
												Description: "The specific NFS version that you want DataSync to use to mount your NFS share.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"smb": {
						// Property: SMB
						Description: "SMB protocol configuration for FSx ONTAP file system.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"domain": {
									// Property: Domain
									Description: "The name of the Windows domain that the SMB server belongs to.",
									Type:        types.StringType,
									Computed:    true,
								},
								"mount_options": {
									// Property: MountOptions
									Description: "The mount options used by DataSync to access the SMB server.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"version": {
												// Property: Version
												Description: "The specific SMB version that you want DataSync to use to mount your SMB share.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"password": {
									// Property: Password
									Description: "The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.",
									Type:        types.StringType,
									Computed:    true,
								},
								"user": {
									// Property: User
									Description: "The user who can mount the share, has the permissions to access files and folders in the SMB share.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"security_group_arns": {
			// Property: SecurityGroupArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARNs of the security groups that are to use to configure the FSx ONTAP file system.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/sg-[a-f0-9]+$",
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "The ARNs of the security groups that are to use to configure the FSx ONTAP file system.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"storage_virtual_machine_arn": {
			// Property: StorageVirtualMachineArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) for the FSx ONTAP SVM.",
			//   "maxLength": 162,
			//   "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:storage-virtual-machine/fs-[0-9a-f]+/svm-[0-9a-f]{17,}$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) for the FSx ONTAP SVM.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "A subdirectory in the location's path.",
			//   "maxLength": 4096,
			//   "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
			//   "type": "string"
			// }
			Description: "A subdirectory in the location's path.",
			Type:        types.StringType,
			Computed:    true,
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
			//   "minItems": 0,
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
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::DataSync::LocationFSxONTAP",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationFSxONTAP").WithTerraformTypeName("awscc_datasync_location_fsx_ontap")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain":                      "Domain",
		"fsx_filesystem_arn":          "FsxFilesystemArn",
		"key":                         "Key",
		"location_arn":                "LocationArn",
		"location_uri":                "LocationUri",
		"mount_options":               "MountOptions",
		"nfs":                         "NFS",
		"password":                    "Password",
		"protocol":                    "Protocol",
		"security_group_arns":         "SecurityGroupArns",
		"smb":                         "SMB",
		"storage_virtual_machine_arn": "StorageVirtualMachineArn",
		"subdirectory":                "Subdirectory",
		"tags":                        "Tags",
		"user":                        "User",
		"value":                       "Value",
		"version":                     "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
