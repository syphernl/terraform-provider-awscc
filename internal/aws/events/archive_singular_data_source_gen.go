// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package events

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_events_archive", archiveDataSourceType)
}

// archiveDataSourceType returns the Terraform awscc_events_archive data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Events::Archive resource type.
func archiveDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"archive_name": {
			// Property: ArchiveName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 48,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"event_pattern": {
			// Property: EventPattern
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"retention_days": {
			// Property: RetentionDays
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"source_arn": {
			// Property: SourceArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Events::Archive",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::Archive").WithTerraformTypeName("awscc_events_archive")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"archive_name":   "ArchiveName",
		"arn":            "Arn",
		"description":    "Description",
		"event_pattern":  "EventPattern",
		"retention_days": "RetentionDays",
		"source_arn":     "SourceArn",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
