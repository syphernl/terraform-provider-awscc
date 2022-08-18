// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmincidents

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_ssmincidents_response_plan", responsePlanResourceType)
}

// responsePlanResourceType returns the Terraform awscc_ssmincidents_response_plan resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSMIncidents::ResponsePlan resource type.
func responsePlanResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"actions": {
			// Property: Actions
			// CloudFormation resource type schema:
			// {
			//   "default": [],
			//   "description": "The list of actions.",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The automation configuration to launch.",
			//     "properties": {
			//       "SsmAutomation": {
			//         "additionalProperties": false,
			//         "description": "The configuration to use when starting the SSM automation document.",
			//         "properties": {
			//           "DocumentName": {
			//             "description": "The document name to use when starting the SSM automation document.",
			//             "maxLength": 128,
			//             "type": "string"
			//           },
			//           "DocumentVersion": {
			//             "description": "The version of the document to use when starting the SSM automation document.",
			//             "maxLength": 128,
			//             "type": "string"
			//           },
			//           "DynamicParameters": {
			//             "description": "The parameters with dynamic values to set when starting the SSM automation document.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "A parameter with a dynamic value to set when starting the SSM automation document.",
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 50,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Value": {
			//                   "additionalProperties": false,
			//                   "description": "Value of the dynamic parameter to set when starting the SSM automation document.",
			//                   "properties": {
			//                     "Variable": {
			//                       "description": "The variable types used as dynamic parameter value when starting the SSM automation document.",
			//                       "enum": [
			//                         "INCIDENT_RECORD_ARN",
			//                         "INVOLVED_RESOURCES"
			//                       ],
			//                       "type": "string"
			//                     }
			//                   },
			//                   "type": "object"
			//                 }
			//               },
			//               "required": [
			//                 "Value",
			//                 "Key"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 200,
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "Parameters": {
			//             "description": "The parameters to set when starting the SSM automation document.",
			//             "insertionOrder": false,
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "A parameter to set when starting the SSM automation document.",
			//               "properties": {
			//                 "Key": {
			//                   "maxLength": 50,
			//                   "minLength": 1,
			//                   "type": "string"
			//                 },
			//                 "Values": {
			//                   "insertionOrder": true,
			//                   "items": {
			//                     "description": "A value of the parameter to set when starting the SSM automation document.",
			//                     "maxLength": 10000,
			//                     "type": "string"
			//                   },
			//                   "maxItems": 10,
			//                   "minItems": 1,
			//                   "type": "array",
			//                   "uniqueItems": true
			//                 }
			//               },
			//               "required": [
			//                 "Values",
			//                 "Key"
			//               ],
			//               "type": "object"
			//             },
			//             "maxItems": 200,
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "RoleArn": {
			//             "description": "The role ARN to use when starting the SSM automation document.",
			//             "maxLength": 1000,
			//             "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//             "type": "string"
			//           },
			//           "TargetAccount": {
			//             "description": "The account type to use when starting the SSM automation document.",
			//             "enum": [
			//               "IMPACTED_ACCOUNT",
			//               "RESPONSE_PLAN_OWNER_ACCOUNT"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "RoleArn",
			//           "DocumentName"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The list of actions.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"ssm_automation": {
						// Property: SsmAutomation
						Description: "The configuration to use when starting the SSM automation document.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"document_name": {
									// Property: DocumentName
									Description: "The document name to use when starting the SSM automation document.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(128),
									},
								},
								"document_version": {
									// Property: DocumentVersion
									Description: "The version of the document to use when starting the SSM automation document.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(128),
									},
								},
								"dynamic_parameters": {
									// Property: DynamicParameters
									Description: "The parameters with dynamic values to set when starting the SSM automation document.",
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 50),
												},
											},
											"value": {
												// Property: Value
												Description: "Value of the dynamic parameter to set when starting the SSM automation document.",
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"variable": {
															// Property: Variable
															Description: "The variable types used as dynamic parameter value when starting the SSM automation document.",
															Type:        types.StringType,
															Optional:    true,
															Validators: []tfsdk.AttributeValidator{
																validate.StringInSlice([]string{
																	"INCIDENT_RECORD_ARN",
																	"INVOLVED_RESOURCES",
																}),
															},
														},
													},
												),
												Required: true,
											},
										},
									),
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtMost(200),
									},
								},
								"parameters": {
									// Property: Parameters
									Description: "The parameters to set when starting the SSM automation document.",
									Attributes: tfsdk.SetNestedAttributes(
										map[string]tfsdk.Attribute{
											"key": {
												// Property: Key
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 50),
												},
											},
											"values": {
												// Property: Values
												Type:     types.ListType{ElemType: types.StringType},
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenBetween(1, 10),
													validate.UniqueItems(),
													validate.ArrayForEach(validate.StringLenAtMost(10000)),
												},
											},
										},
									),
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtMost(200),
									},
								},
								"role_arn": {
									// Property: RoleArn
									Description: "The role ARN to use when starting the SSM automation document.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(1000),
										validate.StringMatch(regexp.MustCompile("^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$"), ""),
									},
								},
								"target_account": {
									// Property: TargetAccount
									Description: "The account type to use when starting the SSM automation document.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"IMPACTED_ACCOUNT",
											"RESPONSE_PLAN_OWNER_ACCOUNT",
										}),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtMost(1),
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.List{ElemType: types.StringType, Elems: []attr.Value{}}),
				resource.UseStateForUnknown(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the response plan.",
			//   "maxLength": 1000,
			//   "pattern": "^arn:aws(-(cn|us-gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//   "type": "string"
			// }
			Description: "The ARN of the response plan.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"chat_channel": {
			// Property: ChatChannel
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The chat channel configuration.",
			//   "properties": {
			//     "ChatbotSns": {
			//       "insertionOrder": true,
			//       "items": {
			//         "description": "The ARN of the Chatbot SNS topic.",
			//         "maxLength": 1000,
			//         "pattern": "^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The chat channel configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"chatbot_sns": {
						// Property: ChatbotSns
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.UniqueItems(),
							validate.ArrayForEach(validate.StringLenAtMost(1000)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$"), "")),
						},
					},
				},
			),
			Optional: true,
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			// {
			//   "description": "The display name of the response plan.",
			//   "maxLength": 200,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The display name of the response plan.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 200),
			},
		},
		"engagements": {
			// Property: Engagements
			// CloudFormation resource type schema:
			// {
			//   "default": [],
			//   "description": "The list of engagements to use.",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "The ARN of the contact.",
			//     "maxLength": 1000,
			//     "pattern": "^arn:aws(-(cn|us-gov))?:ssm-contacts:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//     "type": "string"
			//   },
			//   "maxItems": 5,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The list of engagements to use.",
			Type:        types.SetType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 5),
				validate.ArrayForEach(validate.StringLenAtMost(1000)),
				validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^arn:aws(-(cn|us-gov))?:ssm-contacts:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$"), "")),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Set{ElemType: types.StringType, Elems: []attr.Value{}}),
				resource.UseStateForUnknown(),
			},
		},
		"incident_template": {
			// Property: IncidentTemplate
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The incident template configuration.",
			//   "properties": {
			//     "DedupeString": {
			//       "description": "The deduplication string.",
			//       "maxLength": 1000,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Impact": {
			//       "description": "The impact value.",
			//       "maximum": 5,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "IncidentTags": {
			//       "default": [],
			//       "description": "Tags that get applied to incidents created by the StartIncident API action.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "A key-value pair to tag a resource.",
			//         "properties": {
			//           "Key": {
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Value": {
			//             "maxLength": 256,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Value",
			//           "Key"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 50,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "NotificationTargets": {
			//       "description": "The list of notification targets.",
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "A notification target.",
			//         "properties": {
			//           "SnsTopicArn": {
			//             "description": "The ARN of the Chatbot SNS topic.",
			//             "maxLength": 1000,
			//             "pattern": "^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 10,
			//       "type": "array"
			//     },
			//     "Summary": {
			//       "description": "The summary string.",
			//       "maxLength": 4000,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Title": {
			//       "description": "The title string.",
			//       "maxLength": 200,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Title",
			//     "Impact"
			//   ],
			//   "type": "object"
			// }
			Description: "The incident template configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"dedupe_string": {
						// Property: DedupeString
						Description: "The deduplication string.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 1000),
						},
					},
					"impact": {
						// Property: Impact
						Description: "The impact value.",
						Type:        types.Int64Type,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 5),
						},
					},
					"incident_tags": {
						// Property: IncidentTags
						Description: "Tags that get applied to incidents created by the StartIncident API action.",
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 128),
									},
								},
								"value": {
									// Property: Value
									Type:     types.StringType,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 256),
									},
								},
							},
						),
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(50),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Set{ElemType: types.StringType, Elems: []attr.Value{}}),
							resource.UseStateForUnknown(),
						},
					},
					"notification_targets": {
						// Property: NotificationTargets
						Description: "The list of notification targets.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"sns_topic_arn": {
									// Property: SnsTopicArn
									Description: "The ARN of the Chatbot SNS topic.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenAtMost(1000),
										validate.StringMatch(regexp.MustCompile("^arn:aws(-(cn|us-gov))?:sns:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$"), ""),
									},
								},
							},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtMost(10),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"summary": {
						// Property: Summary
						Description: "The summary string.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 4000),
						},
					},
					"title": {
						// Property: Title
						Description: "The title string.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(200),
						},
					},
				},
			),
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the response plan.",
			//   "maxLength": 200,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9_-]*$",
			//   "type": "string"
			// }
			Description: "The name of the response plan.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 200),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]*$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "default": [],
			//   "description": "The tags to apply to the response plan.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to tag a resource.",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The tags to apply to the response plan.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Set{ElemType: types.StringType, Elems: []attr.Value{}}),
				resource.UseStateForUnknown(),
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
		Description: "Resource type definition for AWS::SSMIncidents::ResponsePlan",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMIncidents::ResponsePlan").WithTerraformTypeName("awscc_ssmincidents_response_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":              "Actions",
		"arn":                  "Arn",
		"chat_channel":         "ChatChannel",
		"chatbot_sns":          "ChatbotSns",
		"dedupe_string":        "DedupeString",
		"display_name":         "DisplayName",
		"document_name":        "DocumentName",
		"document_version":     "DocumentVersion",
		"dynamic_parameters":   "DynamicParameters",
		"engagements":          "Engagements",
		"impact":               "Impact",
		"incident_tags":        "IncidentTags",
		"incident_template":    "IncidentTemplate",
		"key":                  "Key",
		"name":                 "Name",
		"notification_targets": "NotificationTargets",
		"parameters":           "Parameters",
		"role_arn":             "RoleArn",
		"sns_topic_arn":        "SnsTopicArn",
		"ssm_automation":       "SsmAutomation",
		"summary":              "Summary",
		"tags":                 "Tags",
		"target_account":       "TargetAccount",
		"title":                "Title",
		"value":                "Value",
		"values":               "Values",
		"variable":             "Variable",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
