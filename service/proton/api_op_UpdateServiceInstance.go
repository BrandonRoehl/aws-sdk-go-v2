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

// Update a service instance. There are four modes for updating a service instance
// as described in the following. The deploymentType field defines the mode. NONE
// In this mode, a deployment doesn't occur. Only the requested metadata parameters
// are updated. CURRENT_VERSION In this mode, the service instance is deployed and
// updated with the new spec that you provide. Only requested parameters are
// updated. Don’t include minor or major version parameters when you use this
// deployment-type. MINOR_VERSION In this mode, the service instance is deployed
// and updated with the published, recommended (latest) minor version of the
// current major version in use, by default. You can also specify a different minor
// version of the current major version in use. MAJOR_VERSION In this mode, the
// service instance is deployed and updated with the published, recommended
// (latest) major and minor version of the current template, by default. You can
// also specify a different major version that is higher than the major version in
// use and a minor version (optional).
func (c *Client) UpdateServiceInstance(ctx context.Context, params *UpdateServiceInstanceInput, optFns ...func(*Options)) (*UpdateServiceInstanceOutput, error) {
	if params == nil {
		params = &UpdateServiceInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceInstance", params, optFns, addOperationUpdateServiceInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceInstanceInput struct {

	// The deployment type. There are four modes for updating a service instance as
	// described in the following. The deploymentType field defines the mode. NONE In
	// this mode, a deployment doesn't occur. Only the requested metadata parameters
	// are updated. CURRENT_VERSION In this mode, the service instance is deployed and
	// updated with the new spec that you provide. Only requested parameters are
	// updated. Don’t include minor or major version parameters when you use this
	// deployment-type. MINOR_VERSION In this mode, the service instance is deployed
	// and updated with the published, recommended (latest) minor version of the
	// current major version in use, by default. You can also specify a different minor
	// version of the current major version in use. MAJOR_VERSION In this mode, the
	// service instance is deployed and updated with the published, recommended
	// (latest) major and minor version of the current template, by default. You can
	// also specify a different major version that is higher than the major version in
	// use and a minor version (optional).
	//
	// This member is required.
	DeploymentType types.DeploymentUpdateType

	// The name of the service instance to update.
	//
	// This member is required.
	Name *string

	// The name of the service that the service instance belongs to.
	//
	// This member is required.
	ServiceName *string

	// The formatted specification that defines the service instance update.
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	// The major version of the service template to update.
	TemplateMajorVersion *string

	// The minor version of the service template to update.
	TemplateMinorVersion *string
}

type UpdateServiceInstanceOutput struct {

	// The service instance summary data returned by AWS Proton.
	//
	// This member is required.
	ServiceInstance *types.ServiceInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateServiceInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateServiceInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateServiceInstance{}, middleware.After)
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
	if err = addOpUpdateServiceInstanceValidationMiddleware(stack); err != nil {
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
