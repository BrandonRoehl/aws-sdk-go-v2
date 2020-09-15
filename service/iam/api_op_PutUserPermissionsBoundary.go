// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates the policy that is specified as the IAM user's permissions
// boundary. You can use an AWS managed policy or a customer managed policy to set
// the boundary for a user. Use the boundary to control the maximum permissions
// that the user can have. Setting a permissions boundary is an advanced feature
// that can affect the permissions for the user. Policies that are used as
// permissions boundaries do not provide permissions. You must also attach a
// permissions policy to the user. To learn how the effective permissions for a
// user are evaluated, see IAM JSON Policy Evaluation Logic
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
// in the IAM User Guide.
func (c *Client) PutUserPermissionsBoundary(ctx context.Context, params *PutUserPermissionsBoundaryInput, optFns ...func(*Options)) (*PutUserPermissionsBoundaryOutput, error) {
	stack := middleware.NewStack("PutUserPermissionsBoundary", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutUserPermissionsBoundaryMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutUserPermissionsBoundaryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutUserPermissionsBoundary(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "PutUserPermissionsBoundary",
			Err:           err,
		}
	}
	out := result.(*PutUserPermissionsBoundaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutUserPermissionsBoundaryInput struct {
	// The name (friendly name, not ARN) of the IAM user for which you want to set the
	// permissions boundary.
	UserName *string
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary *string
}

type PutUserPermissionsBoundaryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutUserPermissionsBoundaryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutUserPermissionsBoundary{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutUserPermissionsBoundary{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutUserPermissionsBoundary(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "PutUserPermissionsBoundary",
	}
}
