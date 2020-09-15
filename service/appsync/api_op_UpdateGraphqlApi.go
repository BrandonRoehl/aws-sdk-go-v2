// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a GraphqlApi object.
func (c *Client) UpdateGraphqlApi(ctx context.Context, params *UpdateGraphqlApiInput, optFns ...func(*Options)) (*UpdateGraphqlApiOutput, error) {
	stack := middleware.NewStack("UpdateGraphqlApi", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateGraphqlApiMiddlewares(stack)
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
	addOpUpdateGraphqlApiValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGraphqlApi(options.Region), middleware.Before)

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
			OperationName: "UpdateGraphqlApi",
			Err:           err,
		}
	}
	out := result.(*UpdateGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGraphqlApiInput struct {
	// The Amazon CloudWatch Logs configuration for the GraphqlApi object.
	LogConfig *types.LogConfig
	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []*types.AdditionalAuthenticationProvider
	// A flag indicating whether to enable X-Ray tracing for the GraphqlApi.
	XrayEnabled *bool
	// The new Amazon Cognito user pool configuration for the GraphqlApi object.
	UserPoolConfig *types.UserPoolConfig
	// The new authentication type for the GraphqlApi object.
	AuthenticationType types.AuthenticationType
	// The OpenID Connect configuration for the GraphqlApi object.
	OpenIDConnectConfig *types.OpenIDConnectConfig
	// The new name for the GraphqlApi object.
	Name *string
	// The API ID.
	ApiId *string
}

type UpdateGraphqlApiOutput struct {
	// The updated GraphqlApi object.
	GraphqlApi *types.GraphqlApi

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateGraphqlApiMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGraphqlApi{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGraphqlApi{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGraphqlApi(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "UpdateGraphqlApi",
	}
}
