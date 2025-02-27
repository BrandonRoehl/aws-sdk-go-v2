// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an environment template.
func (c *Client) UpdateEnvironmentTemplate(ctx context.Context, params *UpdateEnvironmentTemplateInput, optFns ...func(*Options)) (*UpdateEnvironmentTemplateOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironmentTemplate", params, optFns, addOperationUpdateEnvironmentTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentTemplateInput struct {

	// The name of the environment template to update.
	//
	// This member is required.
	Name *string

	// A description of the environment template update.
	Description *string

	// The name of the environment template to update as displayed in the developer
	// interface.
	DisplayName *string
}

type UpdateEnvironmentTemplateOutput struct {

	// The environment template detail data that's returned by AWS Proton.
	//
	// This member is required.
	EnvironmentTemplate *types.EnvironmentTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateEnvironmentTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEnvironmentTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEnvironmentTemplate{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpUpdateEnvironmentTemplateValidationMiddleware(stack); err != nil {
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
