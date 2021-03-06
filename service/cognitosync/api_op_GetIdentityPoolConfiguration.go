// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the configuration settings of an identity pool. This API can only be called
// with developer credentials. You cannot call this API with the temporary user
// credentials provided by Cognito Identity.
func (c *Client) GetIdentityPoolConfiguration(ctx context.Context, params *GetIdentityPoolConfigurationInput, optFns ...func(*Options)) (*GetIdentityPoolConfigurationOutput, error) {
	if params == nil {
		params = &GetIdentityPoolConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityPoolConfiguration", params, optFns, addOperationGetIdentityPoolConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityPoolConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetIdentityPoolConfiguration operation.
type GetIdentityPoolConfigurationInput struct {

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. This is the ID of the pool for which to return a
	// configuration.
	//
	// This member is required.
	IdentityPoolId *string
}

// The output for the GetIdentityPoolConfiguration operation.
type GetIdentityPoolConfigurationOutput struct {

	// Options to apply to this identity pool for Amazon Cognito streams.
	CognitoStreams *types.CognitoStreams

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito.
	IdentityPoolId *string

	// Options to apply to this identity pool for push synchronization.
	PushSync *types.PushSync

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIdentityPoolConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIdentityPoolConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIdentityPoolConfiguration{}, middleware.After)
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
	if err = addOpGetIdentityPoolConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityPoolConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdentityPoolConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "GetIdentityPoolConfiguration",
	}
}
