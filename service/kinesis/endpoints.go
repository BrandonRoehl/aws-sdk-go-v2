// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	internalendpoints "github.com/aws/aws-sdk-go-v2/service/kinesis/internal/endpoints"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/url"
)

// ResolverOptions is the service endpoint resolver options
type ResolverOptions = internalendpoints.Options

// EndpointResolver interface for resolving service endpoints.
type EndpointResolver interface {
	ResolveEndpoint(region string, options ResolverOptions) (aws.Endpoint, error)
}

var _ EndpointResolver = &internalendpoints.Resolver{}

// NewDefaultEndpointResolver constructs a new service endpoint resolver
func NewDefaultEndpointResolver() *internalendpoints.Resolver {
	return internalendpoints.New()
}

// EndpointResolverFunc is a helper utility that wraps a function so it satisfies
// the EndpointResolver interface. This is useful when you want to add additional
// endpoint resolving logic, or stub out specific endpoints with custom values.
type EndpointResolverFunc func(region string, options ResolverOptions) (aws.Endpoint, error)

func (fn EndpointResolverFunc) ResolveEndpoint(region string, options ResolverOptions) (aws.Endpoint, error) {
	return fn(region, options)
}

func resolveDefaultEndpointConfiguration(o *Options) {
	if o.EndpointResolver != nil {
		return
	}
	o.EndpointResolver = NewDefaultEndpointResolver()
}

type ResolveEndpoint struct {
	Resolver EndpointResolver
	Options  ResolverOptions
}

func (*ResolveEndpoint) ID() string {
	return "ResolveEndpoint"
}

func (m *ResolveEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.Resolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	var endpoint aws.Endpoint
	endpoint, err = m.Resolver.ResolveEndpoint(awsmiddleware.GetRegion(ctx), m.Options)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint")
	}

	req.URL, err = url.Parse(endpoint.URL)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}

	if len(awsmiddleware.GetSigningName(ctx)) == 0 {
		signingName := endpoint.SigningName
		if len(signingName) == 0 {
			signingName = "kinesis"
		}
		ctx = awsmiddleware.SetSigningName(ctx, signingName)
	}
	ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)

	return next.HandleSerialize(ctx, in)
}

type ResolveEndpointMiddlewareOptions interface {
	GetEndpointResolver() EndpointResolver
	GetEndpointOptions() ResolverOptions
}

func AddResolveEndpointMiddleware(stack *middleware.Stack, options ResolveEndpointMiddlewareOptions) {
	stack.Serialize.Insert(&ResolveEndpoint{
		Resolver: options.GetEndpointResolver(),
		Options:  options.GetEndpointOptions(),
	}, "OperationSerializer", middleware.Before)
}

func RemoveResolveEndpointMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Remove((&ResolveEndpoint{}).ID())
}
