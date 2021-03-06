// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a metric filter and associates it with the specified log
// group. Metric filters allow you to configure rules to extract metric data from
// log events ingested through PutLogEvents
// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutLogEvents.html).
// The maximum number of metric filters that can be associated with a log group is
// 100.
func (c *Client) PutMetricFilter(ctx context.Context, params *PutMetricFilterInput, optFns ...func(*Options)) (*PutMetricFilterOutput, error) {
	if params == nil {
		params = &PutMetricFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetricFilter", params, optFns, addOperationPutMetricFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetricFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetricFilterInput struct {

	// A name for the metric filter.
	//
	// This member is required.
	FilterName *string

	// A filter pattern for extracting metric data out of ingested log events.
	//
	// This member is required.
	FilterPattern *string

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// A collection of information that defines how metric data gets emitted.
	//
	// This member is required.
	MetricTransformations []types.MetricTransformation
}

type PutMetricFilterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutMetricFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutMetricFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutMetricFilter{}, middleware.After)
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
	if err = addOpPutMetricFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetricFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMetricFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutMetricFilter",
	}
}
