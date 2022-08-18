// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ecr_registry_policy", registryPolicyDataSourceType)
}

// registryPolicyDataSourceType returns the Terraform awscc_ecr_registry_policy data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECR::RegistryPolicy resource type.
func registryPolicyDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"policy_text": {
			// Property: PolicyText
			// CloudFormation resource type schema:
			// {
			//   "description": "The JSON policy text to apply to your registry. The policy text follows the same format as IAM policy text. For more information, see Registry permissions (https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the Amazon Elastic Container Registry User Guide.",
			//   "type": "object"
			// }
			Description: "The JSON policy text to apply to your registry. The policy text follows the same format as IAM policy text. For more information, see Registry permissions (https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the Amazon Elastic Container Registry User Guide.",
			Type:        types.MapType{ElemType: types.StringType},
			Computed:    true,
		},
		"registry_id": {
			// Property: RegistryId
			// CloudFormation resource type schema:
			// {
			//   "description": "The registry id.",
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "pattern": "^[0-9]{12}$",
			//   "type": "string"
			// }
			Description: "The registry id.",
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
		Description: "Data Source schema for AWS::ECR::RegistryPolicy",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::RegistryPolicy").WithTerraformTypeName("awscc_ecr_registry_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"policy_text": "PolicyText",
		"registry_id": "RegistryId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
