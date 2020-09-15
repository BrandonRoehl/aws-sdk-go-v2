// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a self-service action.
func (c *Client) CreateServiceAction(ctx context.Context, params *CreateServiceActionInput, optFns ...func(*Options)) (*CreateServiceActionOutput, error) {
	stack := middleware.NewStack("CreateServiceAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateServiceActionMiddlewares(stack)
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
	addIdempotencyToken_opCreateServiceActionMiddleware(stack, options)
	addOpCreateServiceActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceAction(options.Region), middleware.Before)

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
			OperationName: "CreateServiceAction",
			Err:           err,
		}
	}
	out := result.(*CreateServiceActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceActionInput struct {
	// The self-service action description.
	Description *string
	// The self-service action name.
	Name *string
	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	IdempotencyToken *string
	// The self-service action definition. Can be one of the following: Name The name
	// of the AWS Systems Manager document (SSM document). For example,
	// AWS-RestartEC2Instance. If you are using a shared SSM document, you must provide
	// the ARN instead of the name. Version The AWS Systems Manager automation document
	// version. For example, "Version": "1" AssumeRole The Amazon Resource Name (ARN)
	// of the role that performs the self-service actions on your behalf. For example,
	// "AssumeRole": "arn:aws:iam::12345678910:role/ActionRole". To reuse the
	// provisioned product launch role, set to "AssumeRole": "LAUNCH_ROLE". Parameters
	// The list of parameters in JSON format. For example:
	// [{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}] or
	// [{\"Name\":\"InstanceId\",\"Type\":\"TEXT_VALUE\"}].
	Definition map[string]*string
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The service action definition type. For example, SSM_AUTOMATION.
	DefinitionType types.ServiceActionDefinitionType
}

type CreateServiceActionOutput struct {
	// An object containing information about the self-service action.
	ServiceActionDetail *types.ServiceActionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateServiceActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateServiceAction{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateServiceAction{}, middleware.After)
}

type idempotencyToken_initializeOpCreateServiceAction struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServiceAction) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServiceAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceActionInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceActionMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateServiceAction{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServiceAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateServiceAction",
	}
}
