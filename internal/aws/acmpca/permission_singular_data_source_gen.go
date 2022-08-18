// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_acmpca_permission", permissionDataSourceType)
}

// permissionDataSourceType returns the Terraform awscc_acmpca_permission data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ACMPCA::Permission resource type.
func permissionDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"actions": {
			// Property: Actions
			// CloudFormation resource type schema:
			// {
			//   "description": "The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"certificate_authority_arn": {
			// Property: CertificateAuthorityArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.",
			Type:        types.StringType,
			Computed:    true,
		},
		"principal": {
			// Property: Principal
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.",
			//   "type": "string"
			// }
			Description: "The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.",
			Type:        types.StringType,
			Computed:    true,
		},
		"source_account": {
			// Property: SourceAccount
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the calling account.",
			//   "type": "string"
			// }
			Description: "The ID of the calling account.",
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
		Description: "Data Source schema for AWS::ACMPCA::Permission",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ACMPCA::Permission").WithTerraformTypeName("awscc_acmpca_permission")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                   "Actions",
		"certificate_authority_arn": "CertificateAuthorityArn",
		"principal":                 "Principal",
		"source_account":            "SourceAccount",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
