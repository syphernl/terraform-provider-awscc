// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

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
	registry.AddDataSourceTypeFactory("awscc_ecs_primary_task_set", primaryTaskSetDataSourceType)
}

// primaryTaskSetDataSourceType returns the Terraform awscc_ecs_primary_task_set data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::ECS::PrimaryTaskSet resource type.
func primaryTaskSetDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"cluster": {
			// Property: Cluster
			// CloudFormation resource type schema:
			// {
			//   "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			//   "type": "string"
			// }
			Description: "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service": {
			// Property: Service
			// CloudFormation resource type schema:
			// {
			//   "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			//   "type": "string"
			// }
			Description: "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			Type:        types.StringType,
			Computed:    true,
		},
		"task_set_id": {
			// Property: TaskSetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID or full Amazon Resource Name (ARN) of the task set.",
			//   "type": "string"
			// }
			Description: "The ID or full Amazon Resource Name (ARN) of the task set.",
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
		Description: "Data Source schema for AWS::ECS::PrimaryTaskSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::PrimaryTaskSet").WithTerraformTypeName("awscc_ecs_primary_task_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cluster":     "Cluster",
		"service":     "Service",
		"task_set_id": "TaskSetId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
