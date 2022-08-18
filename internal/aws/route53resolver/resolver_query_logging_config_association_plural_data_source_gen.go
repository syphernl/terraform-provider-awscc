// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_route53resolver_resolver_query_logging_config_associations", resolverQueryLoggingConfigAssociationsDataSourceType)
}

// resolverQueryLoggingConfigAssociationsDataSourceType returns the Terraform awscc_route53resolver_resolver_query_logging_config_associations data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation resource type.
func resolverQueryLoggingConfigAssociationsDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			Description: "Uniquely identifies the data source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ids": {
			Description: "Set of Resource Identifiers.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Plural Data Source schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation").WithTerraformTypeName("awscc_route53resolver_resolver_query_logging_config_associations")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
