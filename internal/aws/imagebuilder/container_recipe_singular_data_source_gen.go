// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_imagebuilder_container_recipe", containerRecipeDataSourceType)
}

// containerRecipeDataSourceType returns the Terraform awscc_imagebuilder_container_recipe data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ImageBuilder::ContainerRecipe resource type.
func containerRecipeDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the container recipe.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the container recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"components": {
			// Property: Components
			// CloudFormation resource type schema:
			// {
			//   "description": "Components for build and test that are included in the container recipe.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Configuration details of the component.",
			//     "properties": {
			//       "ComponentArn": {
			//         "description": "The Amazon Resource Name (ARN) of the component.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Components for build and test that are included in the container recipe.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"component_arn": {
						// Property: ComponentArn
						Description: "The Amazon Resource Name (ARN) of the component.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"container_type": {
			// Property: ContainerType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the type of container, such as Docker.",
			//   "enum": [
			//     "DOCKER"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the type of container, such as Docker.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the container recipe.",
			//   "type": "string"
			// }
			Description: "The description of the container recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dockerfile_template_data": {
			// Property: DockerfileTemplateData
			// CloudFormation resource type schema:
			// {
			//   "description": "Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.",
			//   "type": "string"
			// }
			Description: "Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"dockerfile_template_uri": {
			// Property: DockerfileTemplateUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The S3 URI for the Dockerfile that will be used to build your container image.",
			//   "type": "string"
			// }
			Description: "The S3 URI for the Dockerfile that will be used to build your container image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_os_version_override": {
			// Property: ImageOsVersionOverride
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the operating system version for the source image.",
			//   "type": "string"
			// }
			Description: "Specifies the operating system version for the source image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_configuration": {
			// Property: InstanceConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A group of options that can be used to configure an instance for building and testing container images.",
			//   "properties": {
			//     "BlockDeviceMappings": {
			//       "description": "Defines the block devices to attach for building an instance from this Image Builder AMI.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Defines block device mappings for the instance used to configure your image. ",
			//         "properties": {
			//           "DeviceName": {
			//             "description": "The device to which these mappings apply.",
			//             "type": "string"
			//           },
			//           "Ebs": {
			//             "additionalProperties": false,
			//             "description": "Use to manage Amazon EBS-specific configuration for this mapping.",
			//             "properties": {
			//               "DeleteOnTermination": {
			//                 "description": "Use to configure delete on termination of the associated device.",
			//                 "type": "boolean"
			//               },
			//               "Encrypted": {
			//                 "description": "Use to configure device encryption.",
			//                 "type": "boolean"
			//               },
			//               "Iops": {
			//                 "description": "Use to configure device IOPS.",
			//                 "type": "integer"
			//               },
			//               "KmsKeyId": {
			//                 "description": "Use to configure the KMS key to use when encrypting the device.",
			//                 "type": "string"
			//               },
			//               "SnapshotId": {
			//                 "description": "The snapshot that defines the device contents.",
			//                 "type": "string"
			//               },
			//               "Throughput": {
			//                 "description": "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
			//                 "type": "integer"
			//               },
			//               "VolumeSize": {
			//                 "description": "Use to override the device's volume size.",
			//                 "type": "integer"
			//               },
			//               "VolumeType": {
			//                 "description": "Use to override the device's volume type.",
			//                 "enum": [
			//                   "standard",
			//                   "io1",
			//                   "io2",
			//                   "gp2",
			//                   "gp3",
			//                   "sc1",
			//                   "st1"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "NoDevice": {
			//             "description": "Use to remove a mapping from the parent image.",
			//             "type": "string"
			//           },
			//           "VirtualName": {
			//             "description": "Use to manage instance ephemeral devices.",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array"
			//     },
			//     "Image": {
			//       "description": "The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A group of options that can be used to configure an instance for building and testing container images.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"block_device_mappings": {
						// Property: BlockDeviceMappings
						Description: "Defines the block devices to attach for building an instance from this Image Builder AMI.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"device_name": {
									// Property: DeviceName
									Description: "The device to which these mappings apply.",
									Type:        types.StringType,
									Computed:    true,
								},
								"ebs": {
									// Property: Ebs
									Description: "Use to manage Amazon EBS-specific configuration for this mapping.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"delete_on_termination": {
												// Property: DeleteOnTermination
												Description: "Use to configure delete on termination of the associated device.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"encrypted": {
												// Property: Encrypted
												Description: "Use to configure device encryption.",
												Type:        types.BoolType,
												Computed:    true,
											},
											"iops": {
												// Property: Iops
												Description: "Use to configure device IOPS.",
												Type:        types.Int64Type,
												Computed:    true,
											},
											"kms_key_id": {
												// Property: KmsKeyId
												Description: "Use to configure the KMS key to use when encrypting the device.",
												Type:        types.StringType,
												Computed:    true,
											},
											"snapshot_id": {
												// Property: SnapshotId
												Description: "The snapshot that defines the device contents.",
												Type:        types.StringType,
												Computed:    true,
											},
											"throughput": {
												// Property: Throughput
												Description: "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
												Type:        types.Int64Type,
												Computed:    true,
											},
											"volume_size": {
												// Property: VolumeSize
												Description: "Use to override the device's volume size.",
												Type:        types.Int64Type,
												Computed:    true,
											},
											"volume_type": {
												// Property: VolumeType
												Description: "Use to override the device's volume type.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"no_device": {
									// Property: NoDevice
									Description: "Use to remove a mapping from the parent image.",
									Type:        types.StringType,
									Computed:    true,
								},
								"virtual_name": {
									// Property: VirtualName
									Description: "Use to manage instance ephemeral devices.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"image": {
						// Property: Image
						Description: "The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "Identifies which KMS key is used to encrypt the container image.",
			//   "type": "string"
			// }
			Description: "Identifies which KMS key is used to encrypt the container image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the container recipe.",
			//   "type": "string"
			// }
			Description: "The name of the container recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parent_image": {
			// Property: ParentImage
			// CloudFormation resource type schema:
			// {
			//   "description": "The source image for the container recipe.",
			//   "type": "string"
			// }
			Description: "The source image for the container recipe.",
			Type:        types.StringType,
			Computed:    true,
		},
		"platform_override": {
			// Property: PlatformOverride
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the operating system platform when you use a custom source image.",
			//   "enum": [
			//     "Windows",
			//     "Linux"
			//   ],
			//   "type": "string"
			// }
			Description: "Specifies the operating system platform when you use a custom source image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Tags that are attached to the container recipe.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Tags that are attached to the container recipe.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"target_repository": {
			// Property: TargetRepository
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The destination repository for the container image.",
			//   "properties": {
			//     "RepositoryName": {
			//       "description": "The name of the container repository where the output container image is stored. This name is prefixed by the repository location.",
			//       "type": "string"
			//     },
			//     "Service": {
			//       "description": "Specifies the service in which this image was registered.",
			//       "enum": [
			//         "ECR"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The destination repository for the container image.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"repository_name": {
						// Property: RepositoryName
						Description: "The name of the container repository where the output container image is stored. This name is prefixed by the repository location.",
						Type:        types.StringType,
						Computed:    true,
					},
					"service": {
						// Property: Service
						Description: "Specifies the service in which this image was registered.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The semantic version of the container recipe (\u003cmajor\u003e.\u003cminor\u003e.\u003cpatch\u003e).",
			//   "type": "string"
			// }
			Description: "The semantic version of the container recipe (<major>.<minor>.<patch>).",
			Type:        types.StringType,
			Computed:    true,
		},
		"working_directory": {
			// Property: WorkingDirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "The working directory to be used during build and test workflows.",
			//   "type": "string"
			// }
			Description: "The working directory to be used during build and test workflows.",
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
		Description: "Data Source schema for AWS::ImageBuilder::ContainerRecipe",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ContainerRecipe").WithTerraformTypeName("awscc_imagebuilder_container_recipe")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"block_device_mappings":     "BlockDeviceMappings",
		"component_arn":             "ComponentArn",
		"components":                "Components",
		"container_type":            "ContainerType",
		"delete_on_termination":     "DeleteOnTermination",
		"description":               "Description",
		"device_name":               "DeviceName",
		"dockerfile_template_data":  "DockerfileTemplateData",
		"dockerfile_template_uri":   "DockerfileTemplateUri",
		"ebs":                       "Ebs",
		"encrypted":                 "Encrypted",
		"image":                     "Image",
		"image_os_version_override": "ImageOsVersionOverride",
		"instance_configuration":    "InstanceConfiguration",
		"iops":                      "Iops",
		"kms_key_id":                "KmsKeyId",
		"name":                      "Name",
		"no_device":                 "NoDevice",
		"parent_image":              "ParentImage",
		"platform_override":         "PlatformOverride",
		"repository_name":           "RepositoryName",
		"service":                   "Service",
		"snapshot_id":               "SnapshotId",
		"tags":                      "Tags",
		"target_repository":         "TargetRepository",
		"throughput":                "Throughput",
		"version":                   "Version",
		"virtual_name":              "VirtualName",
		"volume_size":               "VolumeSize",
		"volume_type":               "VolumeType",
		"working_directory":         "WorkingDirectory",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
