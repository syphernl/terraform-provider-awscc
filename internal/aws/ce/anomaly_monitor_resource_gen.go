// Code generated by generators/resource/main.go; DO NOT EDIT.

package ce

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ce_anomaly_monitor", anomalyMonitorResourceType)
}

// anomalyMonitorResourceType returns the Terraform awscc_ce_anomaly_monitor resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CE::AnomalyMonitor resource type.
func anomalyMonitorResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was created. ",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
			//   "type": "string"
			// }
			Description: "The date when the monitor was created. ",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			Type:        types.Int64Type,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"last_evaluated_date": {
			// Property: LastEvaluatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor last evaluated for anomalies.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
			//   "type": "string"
			// }
			Description: "The date when the monitor last evaluated for anomalies.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"last_updated_date": {
			// Property: LastUpdatedDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date when the monitor was last updated.",
			//   "maxLength": 40,
			//   "minLength": 0,
			//   "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
			//   "type": "string"
			// }
			Description: "The date when the monitor was last updated.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"monitor_arn": {
			// Property: MonitorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Monitor ARN",
			//   "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
			//   "type": "string"
			// }
			Description: "Monitor ARN",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"SERVICE",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"monitor_name": {
			// Property: MonitorName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the monitor.",
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "pattern": "[\\S\\s]*",
			//   "type": "string"
			// }
			Description: "The name of the monitor.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
				validate.StringMatch(regexp.MustCompile("[\\S\\s]*"), ""),
			},
		},
		"monitor_specification": {
			// Property: MonitorSpecification
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
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
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"DIMENSIONAL",
					"CUSTOM",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"resource_tags": {
			// Property: ResourceTags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tags to assign to monitor.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name for the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "Tags to assign to monitor.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name for the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 200),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalyMonitor").WithTerraformTypeName("awscc_ce_anomaly_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_date":           "CreationDate",
		"dimensional_value_count": "DimensionalValueCount",
		"key":                     "Key",
		"last_evaluated_date":     "LastEvaluatedDate",
		"last_updated_date":       "LastUpdatedDate",
		"monitor_arn":             "MonitorArn",
		"monitor_dimension":       "MonitorDimension",
		"monitor_name":            "MonitorName",
		"monitor_specification":   "MonitorSpecification",
		"monitor_type":            "MonitorType",
		"resource_tags":           "ResourceTags",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
