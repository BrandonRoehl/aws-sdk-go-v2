// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new identity pool. The identity pool is a store of user identity
// information that is specific to your AWS account. The keys for
// SupportedLoginProviders are as follows:  <ul> <li> <p>Facebook:
// <code>graph.facebook.com</code> </p> </li> <li> <p>Google:
// <code>accounts.google.com</code> </p> </li> <li> <p>Amazon:
// <code>www.amazon.com</code> </p> </li> <li> <p>Twitter:
// <code>api.twitter.com</code> </p> </li> <li> <p>Digits:
// <code>www.digits.com</code> </p> </li> </ul> <p>You must use AWS Developer
// credentials to call this API.</p>
func (c *Client) CreateIdentityPool(ctx context.Context, params *CreateIdentityPoolInput, optFns ...func(*Options)) (*CreateIdentityPoolOutput, error) {
	stack := middleware.NewStack("CreateIdentityPool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateIdentityPoolMiddlewares(stack)
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
	addOpCreateIdentityPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIdentityPool(options.Region), middleware.Before)

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
			OperationName: "CreateIdentityPool",
			Err:           err,
		}
	}
	out := result.(*CreateIdentityPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the CreateIdentityPool action.
type CreateIdentityPoolInput struct {
	// A string that you provide.
	IdentityPoolName *string
	// The "domain" by which Cognito will refer to your users. This name acts as a
	// placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use letters
	// as well as period (.), underscore (_), and dash (-). Once you have set a
	// developer provider name, you cannot change it. Please take care in setting this
	// parameter.
	DeveloperProviderName *string
	// An array of Amazon Cognito user pools and their client IDs.
	CognitoIdentityProviders []*types.CognitoIdentityProvider
	// Enables or disables the Basic (Classic) authentication flow. For more
	// information, see Identity Pools (Federated Identities) Authentication Flow
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flow.html)
	// in the Amazon Cognito Developer Guide.
	AllowClassicFlow *bool
	// A list of OpendID Connect provider ARNs.
	OpenIdConnectProviderARNs []*string
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []*string
	// Tags to assign to the identity pool. A tag is a label that you can apply to
	// identity pools to categorize and manage them in different ways, such as by
	// purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]*string
	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]*string
	// TRUE if the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities *bool
}

// An object representing an Amazon Cognito identity pool.
type CreateIdentityPoolOutput struct {
	// Enables or disables the Basic (Classic) authentication flow. For more
	// information, see Identity Pools (Federated Identities) Authentication Flow
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flow.html)
	// in the Amazon Cognito Developer Guide.
	AllowClassicFlow *bool
	// A list representing an Amazon Cognito user pool and its client ID.
	CognitoIdentityProviders []*types.CognitoIdentityProvider
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []*string
	// A list of OpendID Connect provider ARNs.
	OpenIdConnectProviderARNs []*string
	// The tags that are assigned to the identity pool. A tag is a label that you can
	// apply to identity pools to categorize and manage them in different ways, such as
	// by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]*string
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId *string
	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string
	// A string that you provide.
	IdentityPoolName *string
	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]*string
	// TRUE if the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIdentities *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateIdentityPoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIdentityPool{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIdentityPool{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateIdentityPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "CreateIdentityPool",
	}
}
