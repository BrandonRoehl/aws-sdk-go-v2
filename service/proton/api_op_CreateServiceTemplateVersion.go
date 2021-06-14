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

// Create a new major or minor version of a service template. A major version of a
// service template is a version that isn't backwards compatible. A minor version
// of a service template is a version that's backwards compatible within its major
// version.
func (c *Client) CreateServiceTemplateVersion(ctx context.Context, params *CreateServiceTemplateVersionInput, optFns ...func(*Options)) (*CreateServiceTemplateVersionOutput, error) {
	if params == nil {
		params = &CreateServiceTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceTemplateVersion", params, optFns, addOperationCreateServiceTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceTemplateVersionInput struct {

	// An array of compatible environment template objects for the new version of a
	// service template.
	//
	// This member is required.
	CompatibleEnvironmentTemplates []types.CompatibleEnvironmentTemplateInput

	// An object that includes the template bundle S3 bucket path and name for the new
	// version of a service template.
	//
	// This member is required.
	Source types.TemplateVersionSourceInput

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// When included, if two identicial requests are made with the same client token,
	// AWS Proton returns the service template version that the first request created.
	ClientToken *string

	// A description of the new version of a service template.
	Description *string

	// To create a new minor version of the service template, include a majorVersion.
	// To create a new major and minor version of the service template, exclude
	// majorVersion.
	MajorVersion *string

	// Create tags for a new version of a service template.
	Tags []types.Tag
}

type CreateServiceTemplateVersionOutput struct {

	// The service template version summary of detail data that's returned by AWS
	// Proton.
	//
	// This member is required.
	ServiceTemplateVersion *types.ServiceTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateServiceTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateServiceTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateServiceTemplateVersion{}, middleware.After)
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
	if err = addOpCreateServiceTemplateVersionValidationMiddleware(stack); err != nil {
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
