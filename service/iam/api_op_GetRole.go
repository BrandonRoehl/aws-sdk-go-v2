// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the specified role, including the role's path, GUID,
// ARN, and the role's trust policy that grants permission to assume the role. For
// more information about roles, see Working with Roles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
// Policies returned by this API are URL-encoded compliant with RFC 3986
// (https://tools.ietf.org/html/rfc3986). You can use a URL decoding method to
// convert the policy back to plain JSON text. For example, if you use Java, you
// can use the decode method of the java.net.URLDecoder utility class in the Java
// SDK. Other languages and SDKs provide similar functionality.
func (c *Client) GetRole(ctx context.Context, params *GetRoleInput, optFns ...func(*Options)) (*GetRoleOutput, error) {
	stack := middleware.NewStack("GetRole", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpGetRoleMiddlewares(stack)
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
	addOpGetRoleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRole(options.Region), middleware.Before)

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
			OperationName: "GetRole",
			Err:           err,
		}
	}
	out := result.(*GetRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoleInput struct {
	// The name of the IAM role to get information about. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	RoleName *string
}

// Contains the response to a successful GetRole () request.
type GetRoleOutput struct {
	// A structure containing details about the IAM role.
	Role *types.Role

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpGetRoleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpGetRole{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpGetRole{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRole(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetRole",
	}
}
