// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the template alias of a template.
func (c *Client) UpdateTemplateAlias(ctx context.Context, params *UpdateTemplateAliasInput, optFns ...func(*Options)) (*UpdateTemplateAliasOutput, error) {
	stack := middleware.NewStack("UpdateTemplateAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateTemplateAliasMiddlewares(stack)
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
	addOpUpdateTemplateAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTemplateAlias(options.Region), middleware.Before)

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
			OperationName: "UpdateTemplateAlias",
			Err:           err,
		}
	}
	out := result.(*UpdateTemplateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTemplateAliasInput struct {
	// The ID of the AWS account that contains the template alias that you're updating.
	AwsAccountId *string
	// The version number of the template.
	TemplateVersionNumber *int64
	// The alias of the template that you want to update. If you name a specific alias,
	// you update the version that the alias points to. You can specify the latest
	// version of the template by providing the keyword $LATEST in the AliasName
	// parameter. The keyword $PUBLISHED doesn't apply to templates.
	AliasName *string
	// The ID for the template.
	TemplateId *string
}

type UpdateTemplateAliasOutput struct {
	// The AWS request ID for this operation.
	RequestId *string
	// The template alias.
	TemplateAlias *types.TemplateAlias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateTemplateAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTemplateAlias{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTemplateAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTemplateAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateTemplateAlias",
	}
}
