// Code generated by generators/resource/main.go; DO NOT EDIT.

package rum

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
	registry.AddResourceTypeFactory("awscc_rum_app_monitor", appMonitorResourceType)
}

// appMonitorResourceType returns the Terraform awscc_rum_app_monitor resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RUM::AppMonitor resource type.
func appMonitorResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_monitor_configuration": {
			// Property: AppMonitorConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "AppMonitor configuration",
			//   "properties": {
			//     "AllowCookies": {
			//       "description": "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
			//       "type": "boolean"
			//     },
			//     "EnableXRay": {
			//       "description": "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
			//       "type": "boolean"
			//     },
			//     "ExcludedPages": {
			//       "description": "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "Page Url",
			//         "maxLength": 1260,
			//         "minLength": 1,
			//         "pattern": "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?\u0026//=]*)",
			//         "type": "string"
			//       },
			//       "maxItems": 50,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "FavoritePages": {
			//       "description": "A list of pages in the RUM console that are to be displayed with a favorite icon.",
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 50,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "GuestRoleArn": {
			//       "description": "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
			//       "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
			//       "type": "string"
			//     },
			//     "IdentityPoolId": {
			//       "description": "The ID of the identity pool that is used to authorize the sending of data to RUM.",
			//       "maxLength": 55,
			//       "minLength": 1,
			//       "pattern": "[\\w-]+:[0-9a-f-]+",
			//       "type": "string"
			//     },
			//     "IncludedPages": {
			//       "description": "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "Page Url",
			//         "maxLength": 1260,
			//         "minLength": 1,
			//         "pattern": "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?\u0026//=]*)",
			//         "type": "string"
			//       },
			//       "maxItems": 50,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "SessionSampleRate": {
			//       "description": "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
			//       "maximum": 1,
			//       "minimum": 0,
			//       "type": "number"
			//     },
			//     "Telemetries": {
			//       "description": "An array that lists the types of telemetry data that this app monitor is to collect.",
			//       "insertionOrder": false,
			//       "items": {
			//         "enum": [
			//           "errors",
			//           "performance",
			//           "http"
			//         ],
			//         "type": "string"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "AppMonitor configuration",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allow_cookies": {
						// Property: AllowCookies
						Description: "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"enable_x_ray": {
						// Property: EnableXRay
						Description: "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"excluded_pages": {
						// Property: ExcludedPages
						Description: "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 50),
							validate.ArrayForEach(validate.StringLenBetween(1, 1260)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"), "")),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"favorite_pages": {
						// Property: FavoritePages
						Description: "A list of pages in the RUM console that are to be displayed with a favorite icon.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 50),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"guest_role_arn": {
						// Property: GuestRoleArn
						Description: "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringMatch(regexp.MustCompile("arn:[^:]*:[^:]*:[^:]*:[^:]*:.*"), ""),
						},
					},
					"identity_pool_id": {
						// Property: IdentityPoolId
						Description: "The ID of the identity pool that is used to authorize the sending of data to RUM.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 55),
							validate.StringMatch(regexp.MustCompile("[\\w-]+:[0-9a-f-]+"), ""),
						},
					},
					"included_pages": {
						// Property: IncludedPages
						Description: "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 50),
							validate.ArrayForEach(validate.StringLenBetween(1, 1260)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"), "")),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"session_sample_rate": {
						// Property: SessionSampleRate
						Description: "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
						Type:        types.Float64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatBetween(0.000000, 1.000000),
						},
					},
					"telemetries": {
						// Property: Telemetries
						Description: "An array that lists the types of telemetry data that this app monitor is to collect.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayForEach(validate.StringInSlice([]string{
								"errors",
								"performance",
								"http",
							})),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
				},
			),
			Optional: true,
		},
		"cw_log_enabled": {
			// Property: CwLogEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
			//   "type": "boolean"
			// }
			Description: "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
			Type:        types.BoolType,
			Optional:    true,
		},
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			// {
			//   "description": "The top-level internet domain name for which your application has administrative authority.",
			//   "maxLength": 253,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The top-level internet domain name for which your application has administrative authority.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 253),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the app monitor",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "[\\.\\-_/#A-Za-z0-9]+",
			//   "type": "string"
			// }
			Description: "A name for the app monitor",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
				validate.StringMatch(regexp.MustCompile("[\\.\\-_/#A-Za-z0-9]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
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
		Description: "Resource Type definition for AWS::RUM::AppMonitor",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RUM::AppMonitor").WithTerraformTypeName("awscc_rum_app_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_cookies":             "AllowCookies",
		"app_monitor_configuration": "AppMonitorConfiguration",
		"cw_log_enabled":            "CwLogEnabled",
		"domain":                    "Domain",
		"enable_x_ray":              "EnableXRay",
		"excluded_pages":            "ExcludedPages",
		"favorite_pages":            "FavoritePages",
		"guest_role_arn":            "GuestRoleArn",
		"identity_pool_id":          "IdentityPoolId",
		"included_pages":            "IncludedPages",
		"key":                       "Key",
		"name":                      "Name",
		"session_sample_rate":       "SessionSampleRate",
		"tags":                      "Tags",
		"telemetries":               "Telemetries",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
