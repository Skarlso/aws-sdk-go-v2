// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Associates a Web ACL with a regional application resource, to protect the
// resource. A regional application can be an Application Load Balancer (ALB), an
// API Gateway REST API, or an AppSync GraphQL API. For AWS CloudFront, don't use
// this call. Instead, use your CloudFront distribution configuration. To associate
// a Web ACL, in the CloudFront call UpdateDistribution, set the web ACL ID to the
// Amazon Resource Name (ARN) of the Web ACL. For information, see
// UpdateDistribution
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html).
func (c *Client) AssociateWebACL(ctx context.Context, params *AssociateWebACLInput, optFns ...func(*Options)) (*AssociateWebACLOutput, error) {
	if params == nil {
		params = &AssociateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWebACL", params, optFns, addOperationAssociateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWebACLInput struct {

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL.
	// The ARN must be in one of the following formats:
	//
	// * For an Application Load
	// Balancer:
	// arn:aws:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	// *
	// For an API Gateway REST API:
	// arn:aws:apigateway:region::/restapis/api-id/stages/stage-name
	//
	// * For an AppSync
	// GraphQL API: arn:aws:appsync:region:account-id:apis/GraphQLApiId
	//
	// This member is required.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with
	// the resource.
	//
	// This member is required.
	WebACLArn *string
}

type AssociateWebACLOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpAssociateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWebACL(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opAssociateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "AssociateWebACL",
	}
}
