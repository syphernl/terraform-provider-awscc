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
	registry.AddDataSourceTypeFactory("awscc_imagebuilder_component", componentDataSourceType)
}

// componentDataSourceType returns the Terraform awscc_imagebuilder_component data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ImageBuilder::Component resource type.
func componentDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the component.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"change_description": {
			// Property: ChangeDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "The change description of the component.",
			//   "type": "string"
			// }
			Description: "The change description of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"data": {
			// Property: Data
			// CloudFormation resource type schema:
			// {
			//   "description": "The data of the component.",
			//   "maxLength": 16000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The data of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the component.",
			//   "type": "string"
			// }
			Description: "The description of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"encrypted": {
			// Property: Encrypted
			// CloudFormation resource type schema:
			// {
			//   "description": "The encryption status of the component.",
			//   "type": "boolean"
			// }
			Description: "The encryption status of the component.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The KMS key identifier used to encrypt the component.",
			//   "type": "string"
			// }
			Description: "The KMS key identifier used to encrypt the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the component.",
			//   "type": "string"
			// }
			Description: "The name of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"platform": {
			// Property: Platform
			// CloudFormation resource type schema:
			// {
			//   "description": "The platform of the component.",
			//   "enum": [
			//     "Windows",
			//     "Linux"
			//   ],
			//   "type": "string"
			// }
			Description: "The platform of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"supported_os_versions": {
			// Property: SupportedOsVersions
			// CloudFormation resource type schema:
			// {
			//   "description": "The operating system (OS) version supported by the component.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The operating system (OS) version supported by the component.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags associated with the component.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags associated with the component.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the component denotes whether the component is used to build the image or only to test it. ",
			//   "enum": [
			//     "BUILD",
			//     "TEST"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of the component denotes whether the component is used to build the image or only to test it. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"uri": {
			// Property: Uri
			// CloudFormation resource type schema:
			// {
			//   "description": "The uri of the component.",
			//   "type": "string"
			// }
			Description: "The uri of the component.",
			Type:        types.StringType,
			Computed:    true,
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "description": "The version of the component.",
			//   "type": "string"
			// }
			Description: "The version of the component.",
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
		Description: "Data Source schema for AWS::ImageBuilder::Component",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::Component").WithTerraformTypeName("awscc_imagebuilder_component")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"change_description":    "ChangeDescription",
		"data":                  "Data",
		"description":           "Description",
		"encrypted":             "Encrypted",
		"kms_key_id":            "KmsKeyId",
		"name":                  "Name",
		"platform":              "Platform",
		"supported_os_versions": "SupportedOsVersions",
		"tags":                  "Tags",
		"type":                  "Type",
		"uri":                   "Uri",
		"version":               "Version",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
