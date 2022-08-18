// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

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
	registry.AddDataSourceTypeFactory("awscc_datasync_location_object_storages", locationObjectStoragesDataSourceType)
}

// locationObjectStoragesDataSourceType returns the Terraform awscc_datasync_location_object_storages data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::DataSync::LocationObjectStorage resource type.
func locationObjectStoragesDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
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
		Description: "Plural Data Source schema for AWS::DataSync::LocationObjectStorage",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationObjectStorage").WithTerraformTypeName("awscc_datasync_location_object_storages")
	opts = opts.WithTerraformSchema(schema)

	pluralDataSourceType, err := NewPluralDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return pluralDataSourceType, nil
}
