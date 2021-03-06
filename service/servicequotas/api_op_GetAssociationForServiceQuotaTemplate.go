// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the ServiceQuotaTemplateAssociationStatus value from the service. Use
// this action to determine if the Service Quota template is associated, or
// enabled.
func (c *Client) GetAssociationForServiceQuotaTemplate(ctx context.Context, params *GetAssociationForServiceQuotaTemplateInput, optFns ...func(*Options)) (*GetAssociationForServiceQuotaTemplateOutput, error) {
	if params == nil {
		params = &GetAssociationForServiceQuotaTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssociationForServiceQuotaTemplate", params, optFns, addOperationGetAssociationForServiceQuotaTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssociationForServiceQuotaTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociationForServiceQuotaTemplateInput struct {
}

type GetAssociationForServiceQuotaTemplateOutput struct {

	// Specifies whether the template is ASSOCIATED or DISASSOCIATED. If the template
	// is ASSOCIATED, then it requests service quota increases for all new accounts
	// created in your organization.
	ServiceQuotaTemplateAssociationStatus types.ServiceQuotaTemplateAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAssociationForServiceQuotaTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAssociationForServiceQuotaTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAssociationForServiceQuotaTemplate{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociationForServiceQuotaTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAssociationForServiceQuotaTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetAssociationForServiceQuotaTemplate",
	}
}
