// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_sagemaker_image_version", imageVersionDataSourceType)
}

// imageVersionDataSourceType returns the Terraform awscc_sagemaker_image_version data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SageMaker::ImageVersion resource type.
func imageVersionDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"base_image": {
			// Property: BaseImage
			// CloudFormation resource type schema:
			// {
			//   "description": "The registry path of the container image on which this image version is based.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": ".+",
			//   "type": "string"
			// }
			Description: "The registry path of the container image on which this image version is based.",
			Type:        types.StringType,
			Computed:    true,
		},
		"container_image": {
			// Property: ContainerImage
			// CloudFormation resource type schema:
			// {
			//   "description": "The registry path of the container image that contains this image version.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": ".+",
			//   "type": "string"
			// }
			Description: "The registry path of the container image that contains this image version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_arn": {
			// Property: ImageArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the parent image.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image\\/[a-z0-9]([-.]?[a-z0-9])*$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the parent image.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_name": {
			// Property: ImageName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the image this version belongs to.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[A-Za-z0-9]([-.]?[A-Za-z0-9])*$",
			//   "type": "string"
			// }
			Description: "The name of the image this version belongs to.",
			Type:        types.StringType,
			Computed:    true,
		},
		"image_version_arn": {
			// Property: ImageVersionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the image version.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "^arn:aws(-[\\w]+)*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:image-version\\/[a-z0-9]([-.]?[a-z0-9])*\\/[0-9]+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the image version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version number of the image version.",
			//   "minimum": 1,
			//   "type": "integer"
			// }
			Description: "The version number of the image version.",
			Type:        types.Int64Type,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SageMaker::ImageVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::ImageVersion").WithTerraformTypeName("awscc_sagemaker_image_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"base_image":        "BaseImage",
		"container_image":   "ContainerImage",
		"image_arn":         "ImageArn",
		"image_name":        "ImageName",
		"image_version_arn": "ImageVersionArn",
		"version":           "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
