// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new job for an existing AWS Glue DataBrew recipe in the current AWS
// account. You can create a standalone job using either a project, or a
// combination of a recipe and a dataset.
func (c *Client) CreateRecipeJob(ctx context.Context, params *CreateRecipeJobInput, optFns ...func(*Options)) (*CreateRecipeJobOutput, error) {
	if params == nil {
		params = &CreateRecipeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRecipeJob", params, optFns, addOperationCreateRecipeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRecipeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRecipeJobInput struct {

	// A unique name for the job.
	//
	// This member is required.
	Name *string

	// One or more artifacts that represent the output from running the job.
	//
	// This member is required.
	Outputs []types.Output

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role to be assumed for this request.
	//
	// This member is required.
	RoleArn *string

	// The name of the dataset that this job processes.
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

	// A value that enables or disables Amazon CloudWatch logging for the current AWS
	// account. If logging is enabled, CloudWatch writes one log stream for each job
	// run.
	LogSubscription types.LogSubscription

	// The maximum number of nodes that DataBrew can consume when the job processes
	// data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// Either the name of an existing project, or a combination of a recipe and a
	// dataset to associate with the recipe.
	ProjectName *string

	// Represents all of the attributes of an AWS Glue DataBrew recipe.
	RecipeReference *types.RecipeReference

	// Metadata tags to apply to this job dataset.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT.
	Timeout int32
}

type CreateRecipeJobOutput struct {

	// The name of the job that you created.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRecipeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRecipeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRecipeJob{}, middleware.After)
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
	if err = addOpCreateRecipeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRecipeJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRecipeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "CreateRecipeJob",
	}
}
