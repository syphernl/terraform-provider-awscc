// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ecs_cluster_capacity_provider_associations_plural", clusterCapacityProviderAssociationsPluralDataSourceType)
}

// clusterCapacityProviderAssociationsPluralDataSourceType returns the Terraform awscc_ecs_cluster_capacity_provider_associations_plural data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECS::ClusterCapacityProviderAssociations resource type.
func clusterCapacityProviderAssociationsPluralDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
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
		Description: "Plural Data Source schema for AWS::ECS::ClusterCapacityProviderAssociations",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::ClusterCapacityProviderAssociations").WithTerraformTypeName("awscc_ecs_cluster_capacity_provider_associations_plural")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
