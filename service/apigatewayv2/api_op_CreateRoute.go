// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Route for an API.
func (c *Client) CreateRoute(ctx context.Context, params *CreateRouteInput, optFns ...func(*Options)) (*CreateRouteOutput, error) {
	stack := middleware.NewStack("CreateRoute", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRouteMiddlewares(stack)
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
	addOpCreateRouteValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoute(options.Region), middleware.Before)

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
			OperationName: "CreateRoute",
			Err:           err,
		}
	}
	out := result.(*CreateRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new Route resource to represent a route.
type CreateRouteInput struct {
	// The route key for the route.
	RouteKey *string
	// The operation name for the route.
	OperationName *string
	// The API identifier.
	ApiId *string
	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string
	// The identifier of the Authorizer resource to be associated with this route. The
	// authorizer identifier is generated by API Gateway when you created the
	// authorizer.
	AuthorizerId *string
	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters map[string]*types.ParameterConstraints
	// The route response selection expression for the route. Supported only for
	// WebSocket APIs.
	RouteResponseSelectionExpression *string
	// The authorization type for the route. For WebSocket APIs, valid values are NONE
	// for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a
	// Lambda authorizer For HTTP APIs, valid values are NONE for open access, or JWT
	// for using JSON Web Tokens.
	AuthorizationType types.AuthorizationType
	// Specifies whether an API key is required for the route. Supported only for
	// WebSocket APIs.
	ApiKeyRequired *bool
	// The authorization scopes supported by this route.
	AuthorizationScopes []*string
	// The target for the route.
	Target *string
	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels map[string]*string
}

type CreateRouteOutput struct {
	// The target for the route.
	Target *string
	// The request parameters for the route. Supported only for WebSocket APIs.
	RequestParameters map[string]*types.ParameterConstraints
	// The identifier of the Authorizer resource to be associated with this route. The
	// authorizer identifier is generated by API Gateway when you created the
	// authorizer.
	AuthorizerId *string
	// The route ID.
	RouteId *string
	// The request models for the route. Supported only for WebSocket APIs.
	RequestModels map[string]*string
	// Specifies whether a route is managed by API Gateway. If you created an API using
	// quick create, the $default route is managed by API Gateway. You can't modify the
	// $default route key.
	ApiGatewayManaged *bool
	// The operation name for the route.
	OperationName *string
	// Specifies whether an API key is required for this route. Supported only for
	// WebSocket APIs.
	ApiKeyRequired *bool
	// A list of authorization scopes configured on a route. The scopes are used with a
	// JWT authorizer to authorize the method invocation. The authorization works by
	// matching the route scopes against the scopes parsed from the access token in the
	// incoming request. The method invocation is authorized if any route scope matches
	// a claimed scope in the access token. Otherwise, the invocation is not
	// authorized. When the route scope is configured, the client must provide an
	// access token instead of an identity token for authorization purposes.
	AuthorizationScopes []*string
	// The route key for the route.
	RouteKey *string
	// The route response selection expression for the route. Supported only for
	// WebSocket APIs.
	RouteResponseSelectionExpression *string
	// The model selection expression for the route. Supported only for WebSocket APIs.
	ModelSelectionExpression *string
	// The authorization type for the route. For WebSocket APIs, valid values are NONE
	// for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a
	// Lambda authorizer For HTTP APIs, valid values are NONE for open access, or JWT
	// for using JSON Web Tokens.
	AuthorizationType types.AuthorizationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRouteMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoute{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoute{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRoute(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateRoute",
	}
}
