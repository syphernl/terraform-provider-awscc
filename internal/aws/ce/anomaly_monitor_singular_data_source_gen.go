// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ce

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ce_anomaly_monitor", anomalyMonitorDataSourceType)
}

// anomalyMonitorDataSourceType returns the Terraform awscc_ce_anomaly_monitor data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::CE::AnomalyMonitor resource type.
func anomalyMonitorDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was created. ",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date when the monitor was created. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"dimensional_value_count": {
			// Property: DimensionalValueCount
			// CloudFormation resource type schema:
			// {
			//   "description": "The value for evaluated dimensions.",
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "The value for evaluated dimensions.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"last_evaluated_date": {
			// Property: LastEvaluatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor last evaluated for anomalies.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date when the monitor last evaluated for anomalies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_updated_date": {
			// Property: LastUpdatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was last updated.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date when the monitor was last updated.",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_arn": {
			// Property: MonitorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Monitor ARN",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Monitor ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_dimension": {
			// Property: MonitorDimension
			// CloudFormation resource type schema:
			// {
			//   "description": "The dimensions to evaluate",
			//   "enum": [
			//     "SERVICE"
			//   ],
			//   "type": "string"
			// }
			Description: "The dimensions to evaluate",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_name": {
			// Property: MonitorName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the monitor.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the monitor.",
			Type:        types.StringType,
			Computed:    true,
		},
		"monitor_specification": {
			// Property: MonitorSpecification
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"monitor_type": {
			// Property: MonitorType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "DIMENSIONAL",
			//     "CUSTOM"
			//   ],
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
		Description: "Data Source schema for AWS::CE::AnomalyMonitor",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalyMonitor").WithTerraformTypeName("awscc_ce_anomaly_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_date":           "CreationDate",
		"dimensional_value_count": "DimensionalValueCount",
		"last_evaluated_date":     "LastEvaluatedDate",
		"last_updated_date":       "LastUpdatedDate",
		"monitor_arn":             "MonitorArn",
		"monitor_dimension":       "MonitorDimension",
		"monitor_name":            "MonitorName",
		"monitor_specification":   "MonitorSpecification",
		"monitor_type":            "MonitorType",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
