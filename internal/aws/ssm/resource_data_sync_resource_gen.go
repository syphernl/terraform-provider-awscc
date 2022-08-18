// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ssm_resource_data_sync", resourceDataSyncResourceType)
}

// resourceDataSyncResourceType returns the Terraform awscc_ssm_resource_data_sync resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSM::ResourceDataSync resource type.
func resourceDataSyncResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"bucket_prefix": {
			// Property: BucketPrefix
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"bucket_region": {
			// Property: BucketRegion
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"kms_key_arn": {
			// Property: KMSKeyArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 512),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"s3_destination": {
			// Property: S3Destination
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BucketName": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketPrefix": {
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketRegion": {
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "KMSKeyArn": {
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SyncFormat": {
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "BucketName",
			//     "BucketRegion",
			//     "SyncFormat"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"bucket_name": {
						// Property: BucketName
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
					"bucket_prefix": {
						// Property: BucketPrefix
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
					"bucket_region": {
						// Property: BucketRegion
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 64),
						},
					},
					"kms_key_arn": {
						// Property: KMSKeyArn
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 512),
						},
					},
					"sync_format": {
						// Property: SyncFormat
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1024),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"sync_format": {
			// Property: SyncFormat
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1024),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"sync_name": {
			// Property: SyncName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"sync_source": {
			// Property: SyncSource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AwsOrganizationsSource": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "OrganizationSourceType": {
			//           "maxLength": 64,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "OrganizationalUnits": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "OrganizationSourceType"
			//       ],
			//       "type": "object"
			//     },
			//     "IncludeFutureRegions": {
			//       "type": "boolean"
			//     },
			//     "SourceRegions": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SourceType": {
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SourceType",
			//     "SourceRegions"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"aws_organizations_source": {
						// Property: AwsOrganizationsSource
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"organization_source_type": {
									// Property: OrganizationSourceType
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 64),
									},
								},
								"organizational_units": {
									// Property: OrganizationalUnits
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"include_future_regions": {
						// Property: IncludeFutureRegions
						Type:     types.BoolType,
						Optional: true,
					},
					"source_regions": {
						// Property: SourceRegions
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
					"source_type": {
						// Property: SourceType
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 64),
						},
					},
				},
			),
			Optional: true,
		},
		"sync_type": {
			// Property: SyncType
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
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
		Description: "Resource Type definition for AWS::SSM::ResourceDataSync",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::ResourceDataSync").WithTerraformTypeName("awscc_ssm_resource_data_sync")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aws_organizations_source": "AwsOrganizationsSource",
		"bucket_name":              "BucketName",
		"bucket_prefix":            "BucketPrefix",
		"bucket_region":            "BucketRegion",
		"include_future_regions":   "IncludeFutureRegions",
		"kms_key_arn":              "KMSKeyArn",
		"organization_source_type": "OrganizationSourceType",
		"organizational_units":     "OrganizationalUnits",
		"s3_destination":           "S3Destination",
		"source_regions":           "SourceRegions",
		"source_type":              "SourceType",
		"sync_format":              "SyncFormat",
		"sync_name":                "SyncName",
		"sync_source":              "SyncSource",
		"sync_type":                "SyncType",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
