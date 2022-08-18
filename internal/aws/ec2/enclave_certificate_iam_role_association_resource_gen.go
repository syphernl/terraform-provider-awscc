// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

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
	registry.AddResourceTypeFactory("awscc_ec2_enclave_certificate_iam_role_association", enclaveCertificateIamRoleAssociationResourceType)
}

// enclaveCertificateIamRoleAssociationResourceType returns the Terraform awscc_ec2_enclave_certificate_iam_role_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EC2::EnclaveCertificateIamRoleAssociation resource type.
func enclaveCertificateIamRoleAssociationResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"certificate_arn": {
			// Property: CertificateArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
			//   "maxLength": 1283,
			//   "minLength": 1,
			//   "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:acm:[A-Za-z0-9-]{1,64}:([0-9]{12})?:certificate/.+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1283),
				validate.StringMatch(regexp.MustCompile("^arn:aws[A-Za-z0-9-]{0,64}:acm:[A-Za-z0-9-]{1,64}:([0-9]{12})?:certificate/.+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"certificate_s3_bucket_name": {
			// Property: CertificateS3BucketName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the Amazon S3 bucket to which the certificate was uploaded.",
			//   "type": "string"
			// }
			Description: "The name of the Amazon S3 bucket to which the certificate was uploaded.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"certificate_s3_object_key": {
			// Property: CertificateS3ObjectKey
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
			//   "type": "string"
			// }
			Description: "The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"encryption_kms_key_id": {
			// Property: EncryptionKmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
			//   "type": "string"
			// }
			Description: "The ID of the AWS KMS CMK used to encrypt the private key of the certificate.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
			//   "maxLength": 1283,
			//   "minLength": 1,
			//   "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:iam:.*:([0-9]{12})?:role/.+$",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1283),
				validate.StringMatch(regexp.MustCompile("^arn:aws[A-Za-z0-9-]{0,64}:iam:.*:([0-9]{12})?:role/.+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
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
		Description: "Associates an AWS Identity and Access Management (IAM) role with an AWS Certificate Manager (ACM) certificate. This association is based on Amazon Resource Names and it enables the certificate to be used by the ACM for Nitro Enclaves application inside an enclave.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EnclaveCertificateIamRoleAssociation").WithTerraformTypeName("awscc_ec2_enclave_certificate_iam_role_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_arn":            "CertificateArn",
		"certificate_s3_bucket_name": "CertificateS3BucketName",
		"certificate_s3_object_key":  "CertificateS3ObjectKey",
		"encryption_kms_key_id":      "EncryptionKmsKeyId",
		"role_arn":                   "RoleArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
