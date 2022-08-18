// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_ec2_transit_gateway_multicast_domain_association", transitGatewayMulticastDomainAssociationDataSourceType)
}

// transitGatewayMulticastDomainAssociationDataSourceType returns the Terraform awscc_ec2_transit_gateway_multicast_domain_association data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::EC2::TransitGatewayMulticastDomainAssociation resource type.
func transitGatewayMulticastDomainAssociationDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"resource_id": {
			// Property: ResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the resource.",
			//   "type": "string"
			// }
			Description: "The ID of the resource.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_type": {
			// Property: ResourceType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of resource, for example a VPC attachment.",
			//   "type": "string"
			// }
			Description: "The type of resource, for example a VPC attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the subnet association.",
			//   "type": "string"
			// }
			Description: "The state of the subnet association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The IDs of the subnets to associate with the transit gateway multicast domain.",
			//   "type": "string"
			// }
			Description: "The IDs of the subnets to associate with the transit gateway multicast domain.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_attachment_id": {
			// Property: TransitGatewayAttachmentId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway attachment.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway attachment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"transit_gateway_multicast_domain_id": {
			// Property: TransitGatewayMulticastDomainId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the transit gateway multicast domain.",
			//   "type": "string"
			// }
			Description: "The ID of the transit gateway multicast domain.",
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
		Description: "Data Source schema for AWS::EC2::TransitGatewayMulticastDomainAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayMulticastDomainAssociation").WithTerraformTypeName("awscc_ec2_transit_gateway_multicast_domain_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"resource_id":                         "ResourceId",
		"resource_type":                       "ResourceType",
		"state":                               "State",
		"subnet_id":                           "SubnetId",
		"transit_gateway_attachment_id":       "TransitGatewayAttachmentId",
		"transit_gateway_multicast_domain_id": "TransitGatewayMulticastDomainId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
