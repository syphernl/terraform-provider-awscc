// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_personalize_solution", solutionDataSourceType)
}

// solutionDataSourceType returns the Terraform awscc_personalize_solution data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Personalize::Solution resource type.
func solutionDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"dataset_group_arn": {
			// Property: DatasetGroupArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the dataset group that provides the training data.",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The ARN of the dataset group that provides the training data.",
			Type:        types.StringType,
			Computed:    true,
		},
		"event_type": {
			// Property: EventType
			// CloudFormation resource type schema:
			// {
			//   "description": "When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the solution",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//   "type": "string"
			// }
			Description: "The name for the solution",
			Type:        types.StringType,
			Computed:    true,
		},
		"perform_auto_ml": {
			// Property: PerformAutoML
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.",
			//   "type": "boolean"
			// }
			Description: "Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"perform_hpo": {
			// Property: PerformHPO
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.",
			//   "type": "boolean"
			// }
			Description: "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"recipe_arn": {
			// Property: RecipeArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the recipe to use for model training. Only specified when performAutoML is false.",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The ARN of the recipe to use for model training. Only specified when performAutoML is false.",
			Type:        types.StringType,
			Computed:    true,
		},
		"solution_arn": {
			// Property: SolutionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the solution",
			//   "maxLength": 256,
			//   "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//   "type": "string"
			// }
			Description: "The ARN of the solution",
			Type:        types.StringType,
			Computed:    true,
		},
		"solution_config": {
			// Property: SolutionConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The configuration to use with the solution. When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration.",
			//   "properties": {
			//     "AlgorithmHyperParameters": {
			//       "additionalProperties": false,
			//       "description": "Lists the hyperparameter names and ranges.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "AutoMLConfig": {
			//       "additionalProperties": false,
			//       "description": "The AutoMLConfig object containing a list of recipes to search when AutoML is performed.",
			//       "properties": {
			//         "MetricName": {
			//           "description": "The metric to optimize.",
			//           "maxLength": 256,
			//           "type": "string"
			//         },
			//         "RecipeList": {
			//           "description": "The list of candidate recipes.",
			//           "insertionOrder": true,
			//           "items": {
			//             "maxLength": 256,
			//             "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//             "type": "string"
			//           },
			//           "maxItems": 100,
			//           "type": "array"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "EventValueThreshold": {
			//       "description": "Only events with a value greater than or equal to this threshold are used for training a model.",
			//       "maxLength": 256,
			//       "type": "string"
			//     },
			//     "FeatureTransformationParameters": {
			//       "additionalProperties": false,
			//       "description": "Lists the feature transformation parameters.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "HpoConfig": {
			//       "additionalProperties": false,
			//       "description": "Describes the properties for hyperparameter optimization (HPO)",
			//       "properties": {
			//         "AlgorithmHyperParameterRanges": {
			//           "additionalProperties": false,
			//           "description": "The hyperparameters and their allowable ranges",
			//           "properties": {
			//             "CategoricalHyperParameterRanges": {
			//               "description": "The categorical hyperparameters and their ranges.",
			//               "insertionOrder": true,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "description": "Provides the name and values of a Categorical hyperparameter.",
			//                 "properties": {
			//                   "Name": {
			//                     "description": "The name of the hyperparameter.",
			//                     "maxLength": 256,
			//                     "type": "string"
			//                   },
			//                   "Values": {
			//                     "description": "A list of the categories for the hyperparameter.",
			//                     "insertionOrder": true,
			//                     "items": {
			//                       "maxLength": 1000,
			//                       "type": "string"
			//                     },
			//                     "maxItems": 100,
			//                     "type": "array"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "maxItems": 100,
			//               "type": "array"
			//             },
			//             "ContinuousHyperParameterRanges": {
			//               "description": "The continuous hyperparameters and their ranges.",
			//               "insertionOrder": true,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "description": "Provides the name and range of a continuous hyperparameter.",
			//                 "properties": {
			//                   "MaxValue": {
			//                     "description": "The maximum allowable value for the hyperparameter.",
			//                     "minimum": -1000000,
			//                     "type": "number"
			//                   },
			//                   "MinValue": {
			//                     "description": "The minimum allowable value for the hyperparameter.",
			//                     "minimum": -1000000,
			//                     "type": "number"
			//                   },
			//                   "Name": {
			//                     "description": "The name of the hyperparameter.",
			//                     "maxLength": 256,
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "maxItems": 100,
			//               "type": "array"
			//             },
			//             "IntegerHyperParameterRanges": {
			//               "description": "The integer hyperparameters and their ranges.",
			//               "insertionOrder": true,
			//               "items": {
			//                 "additionalProperties": false,
			//                 "description": "Provides the name and range of an integer-valued hyperparameter.",
			//                 "properties": {
			//                   "MaxValue": {
			//                     "description": "The maximum allowable value for the hyperparameter.",
			//                     "maximum": 1000000,
			//                     "type": "integer"
			//                   },
			//                   "MinValue": {
			//                     "description": "The minimum allowable value for the hyperparameter.",
			//                     "minimum": -1000000,
			//                     "type": "integer"
			//                   },
			//                   "Name": {
			//                     "description": "The name of the hyperparameter.",
			//                     "maxLength": 256,
			//                     "type": "string"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "maxItems": 100,
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "HpoObjective": {
			//           "additionalProperties": false,
			//           "description": "The metric to optimize during HPO.",
			//           "properties": {
			//             "MetricName": {
			//               "description": "The name of the metric",
			//               "maxLength": 256,
			//               "type": "string"
			//             },
			//             "MetricRegex": {
			//               "description": "A regular expression for finding the metric in the training job logs.",
			//               "maxLength": 256,
			//               "type": "string"
			//             },
			//             "Type": {
			//               "description": "The type of the metric. Valid values are Maximize and Minimize.",
			//               "enum": [
			//                 "Maximize",
			//                 "Minimize"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "HpoResourceConfig": {
			//           "additionalProperties": false,
			//           "description": "Describes the resource configuration for hyperparameter optimization (HPO).",
			//           "properties": {
			//             "MaxNumberOfTrainingJobs": {
			//               "description": "The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.",
			//               "maxLength": 256,
			//               "type": "string"
			//             },
			//             "MaxParallelTrainingJobs": {
			//               "description": "The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.",
			//               "maxLength": 256,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The configuration to use with the solution. When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"algorithm_hyper_parameters": {
						// Property: AlgorithmHyperParameters
						Description: "Lists the hyperparameter names and ranges.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Computed: true,
					},
					"auto_ml_config": {
						// Property: AutoMLConfig
						Description: "The AutoMLConfig object containing a list of recipes to search when AutoML is performed.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"metric_name": {
									// Property: MetricName
									Description: "The metric to optimize.",
									Type:        types.StringType,
									Computed:    true,
								},
								"recipe_list": {
									// Property: RecipeList
									Description: "The list of candidate recipes.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"event_value_threshold": {
						// Property: EventValueThreshold
						Description: "Only events with a value greater than or equal to this threshold are used for training a model.",
						Type:        types.StringType,
						Computed:    true,
					},
					"feature_transformation_parameters": {
						// Property: FeatureTransformationParameters
						Description: "Lists the feature transformation parameters.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Computed: true,
					},
					"hpo_config": {
						// Property: HpoConfig
						Description: "Describes the properties for hyperparameter optimization (HPO)",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"algorithm_hyper_parameter_ranges": {
									// Property: AlgorithmHyperParameterRanges
									Description: "The hyperparameters and their allowable ranges",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"categorical_hyper_parameter_ranges": {
												// Property: CategoricalHyperParameterRanges
												Description: "The categorical hyperparameters and their ranges.",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"name": {
															// Property: Name
															Description: "The name of the hyperparameter.",
															Type:        types.StringType,
															Computed:    true,
														},
														"values": {
															// Property: Values
															Description: "A list of the categories for the hyperparameter.",
															Type:        types.ListType{ElemType: types.StringType},
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
											"continuous_hyper_parameter_ranges": {
												// Property: ContinuousHyperParameterRanges
												Description: "The continuous hyperparameters and their ranges.",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"max_value": {
															// Property: MaxValue
															Description: "The maximum allowable value for the hyperparameter.",
															Type:        types.Float64Type,
															Computed:    true,
														},
														"min_value": {
															// Property: MinValue
															Description: "The minimum allowable value for the hyperparameter.",
															Type:        types.Float64Type,
															Computed:    true,
														},
														"name": {
															// Property: Name
															Description: "The name of the hyperparameter.",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
											"integer_hyper_parameter_ranges": {
												// Property: IntegerHyperParameterRanges
												Description: "The integer hyperparameters and their ranges.",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"max_value": {
															// Property: MaxValue
															Description: "The maximum allowable value for the hyperparameter.",
															Type:        types.Int64Type,
															Computed:    true,
														},
														"min_value": {
															// Property: MinValue
															Description: "The minimum allowable value for the hyperparameter.",
															Type:        types.Int64Type,
															Computed:    true,
														},
														"name": {
															// Property: Name
															Description: "The name of the hyperparameter.",
															Type:        types.StringType,
															Computed:    true,
														},
													},
												),
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"hpo_objective": {
									// Property: HpoObjective
									Description: "The metric to optimize during HPO.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"metric_name": {
												// Property: MetricName
												Description: "The name of the metric",
												Type:        types.StringType,
												Computed:    true,
											},
											"metric_regex": {
												// Property: MetricRegex
												Description: "A regular expression for finding the metric in the training job logs.",
												Type:        types.StringType,
												Computed:    true,
											},
											"type": {
												// Property: Type
												Description: "The type of the metric. Valid values are Maximize and Minimize.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"hpo_resource_config": {
									// Property: HpoResourceConfig
									Description: "Describes the resource configuration for hyperparameter optimization (HPO).",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"max_number_of_training_jobs": {
												// Property: MaxNumberOfTrainingJobs
												Description: "The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.",
												Type:        types.StringType,
												Computed:    true,
											},
											"max_parallel_training_jobs": {
												// Property: MaxParallelTrainingJobs
												Description: "The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Personalize::Solution",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Solution").WithTerraformTypeName("awscc_personalize_solution")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm_hyper_parameter_ranges":   "AlgorithmHyperParameterRanges",
		"algorithm_hyper_parameters":         "AlgorithmHyperParameters",
		"auto_ml_config":                     "AutoMLConfig",
		"categorical_hyper_parameter_ranges": "CategoricalHyperParameterRanges",
		"continuous_hyper_parameter_ranges":  "ContinuousHyperParameterRanges",
		"dataset_group_arn":                  "DatasetGroupArn",
		"event_type":                         "EventType",
		"event_value_threshold":              "EventValueThreshold",
		"feature_transformation_parameters":  "FeatureTransformationParameters",
		"hpo_config":                         "HpoConfig",
		"hpo_objective":                      "HpoObjective",
		"hpo_resource_config":                "HpoResourceConfig",
		"integer_hyper_parameter_ranges":     "IntegerHyperParameterRanges",
		"max_number_of_training_jobs":        "MaxNumberOfTrainingJobs",
		"max_parallel_training_jobs":         "MaxParallelTrainingJobs",
		"max_value":                          "MaxValue",
		"metric_name":                        "MetricName",
		"metric_regex":                       "MetricRegex",
		"min_value":                          "MinValue",
		"name":                               "Name",
		"perform_auto_ml":                    "PerformAutoML",
		"perform_hpo":                        "PerformHPO",
		"recipe_arn":                         "RecipeArn",
		"recipe_list":                        "RecipeList",
		"solution_arn":                       "SolutionArn",
		"solution_config":                    "SolutionConfig",
		"type":                               "Type",
		"values":                             "Values",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
