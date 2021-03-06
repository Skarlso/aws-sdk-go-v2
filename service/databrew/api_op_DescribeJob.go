// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the definition of a specific AWS Glue DataBrew job that is in the
// current AWS account.
func (c *Client) DescribeJob(ctx context.Context, params *DescribeJobInput, optFns ...func(*Options)) (*DescribeJobOutput, error) {
	if params == nil {
		params = &DescribeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJob", params, optFns, addOperationDescribeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobInput struct {

	// The name of the job to be described.
	//
	// This member is required.
	Name *string
}

type DescribeJobOutput struct {

	// The name of the job.
	//
	// This member is required.
	Name *string

	// The date and time that the job was created.
	CreateDate *time.Time

	// The identifier (user name) of the user associated with the creation of the job.
	CreatedBy *string

	// The dataset that the job acts upon.
	DatasetName *string

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job.
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//
	// * SSE-KMS -
	// Server-side encryption with AWS KMS-managed keys.
	//
	// * SSE-S3 - Server-side
	// encryption with keys managed by Amazon S3.
	EncryptionMode types.EncryptionMode

	// The identifier (user name) of the user who last modified the job.
	LastModifiedBy *string

	// The date and time that the job was last modified.
	LastModifiedDate *time.Time

	// A value that indicates whether Amazon CloudWatch logging is enabled for this
	// job.
	LogSubscription types.LogSubscription

	// The maximum number of nodes that AWS Glue DataBrew can consume when the job
	// processes data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// One or more artifacts that represent the output from running the job.
	Outputs []types.Output

	// The DataBrew project associated with this job.
	ProjectName *string

	// Represents all of the attributes of an AWS Glue DataBrew recipe.
	RecipeReference *types.RecipeReference

	// The Amazon Resource Name (ARN) of the job.
	ResourceArn *string

	// The ARN of the AWS Identity and Access Management (IAM) role that was assumed
	// for this request.
	RoleArn *string

	// Metadata tags associated with this job.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT.
	Timeout int32

	// The job type, which must be one of the following:
	//
	// * PROFILE - The job analyzes
	// the dataset to determine its size, data types, data distribution, and more.
	//
	// *
	// RECIPE - The job applies one or more transformations to a dataset.
	Type types.JobType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJob{}, middleware.After)
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
	if err = addOpDescribeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "DescribeJob",
	}
}
