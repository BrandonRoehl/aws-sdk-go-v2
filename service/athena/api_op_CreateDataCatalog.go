// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates (registers) a data catalog with the specified name and properties.
// Catalogs created are visible to all users of the same AWS account.
func (c *Client) CreateDataCatalog(ctx context.Context, params *CreateDataCatalogInput, optFns ...func(*Options)) (*CreateDataCatalogOutput, error) {
	stack := middleware.NewStack("CreateDataCatalog", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateDataCatalogMiddlewares(stack)
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
	addOpCreateDataCatalogValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataCatalog(options.Region), middleware.Before)

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
			OperationName: "CreateDataCatalog",
			Err:           err,
		}
	}
	out := result.(*CreateDataCatalogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataCatalogInput struct {
	// The name of the data catalog to create. The catalog name must be unique for the
	// AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or
	// hyphen characters.
	Name *string
	// The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS
	// Glue Catalog, or HIVE for an external hive metastore.
	Type types.DataCatalogType
	// Specifies the Lambda function or functions to use for creating the data catalog.
	// This is a mapping whose values depend on the catalog type.
	//
	//     * For the HIVE
	// data catalog type, use the following syntax. The metadata-function parameter is
	// required. The sdk-version parameter is optional and defaults to the currently
	// supported version. metadata-function=lambda_arn, sdk-version=version_number
	//
	//
	// * For the LAMBDA data catalog type, use one of the following sets of required
	// parameters, but not both.
	//
	//         * If you have one Lambda function that
	// processes metadata and another for reading the actual data, use the following
	// syntax. Both parameters are required. metadata-function=lambda_arn,
	// record-function=lambda_arn
	//
	//         * If you have a composite Lambda function
	// that processes both metadata and data, use the following syntax to specify your
	// Lambda function. function=lambda_arn
	//
	//     * The GLUE type has no parameters.
	Parameters map[string]*string
	// A list of comma separated tags to add to the data catalog that is created.
	Tags []*types.Tag
	// A description of the data catalog to be created.
	Description *string
}

type CreateDataCatalogOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateDataCatalogMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataCatalog{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataCatalog{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataCatalog(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "CreateDataCatalog",
	}
}
