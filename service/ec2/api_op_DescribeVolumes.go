// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified EBS volumes or all of your EBS volumes. If you are
// describing a long list of volumes, we recommend that you paginate the output to
// make the list more manageable. The MaxResults parameter sets the maximum number
// of results returned in a single page. If the list of results exceeds your
// MaxResults value, then that number of results is returned along with a NextToken
// value that can be passed to a subsequent DescribeVolumes request to retrieve the
// remaining results. For more information about EBS volumes, see Amazon EBS
// Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumes.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVolumes(ctx context.Context, params *DescribeVolumesInput, optFns ...func(*Options)) (*DescribeVolumesOutput, error) {
	stack := middleware.NewStack("DescribeVolumes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeVolumesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVolumes(options.Region), middleware.Before)

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
			OperationName: "DescribeVolumes",
			Err:           err,
		}
	}
	out := result.(*DescribeVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVolumesInput struct {
	// The volume IDs.
	VolumeIds []*string
	// The filters.
	//
	//     * attachment.attach-time - The time stamp when the attachment
	// initiated.
	//
	//     * attachment.delete-on-termination - Whether the volume is
	// deleted on instance termination.
	//
	//     * attachment.device - The device name
	// specified in the block device mapping (for example, /dev/sda1).
	//
	//     *
	// attachment.instance-id - The ID of the instance the volume is attached to.
	//
	//
	// * attachment.status - The attachment state (attaching | attached | detaching).
	//
	//
	// * availability-zone - The Availability Zone in which the volume was created.
	//
	//
	// * create-time - The time stamp when the volume was created.
	//
	//     * encrypted -
	// Indicates whether the volume is encrypted (true | false)
	//
	//     *
	// multi-attach-enabled - Indicates whether the volume is enabled for Multi-Attach
	// (true | false)
	//
	//     * fast-restored - Indicates whether the volume was created
	// from a snapshot that is enabled for fast snapshot restore (true | false).
	//
	//     *
	// size - The size of the volume, in GiB.
	//
	//     * snapshot-id - The snapshot from
	// which the volume was created.
	//
	//     * status - The status of the volume (creating
	// | available | in-use | deleting | deleted | error).
	//
	//     * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	//     * tag-key - The key of a
	// tag assigned to the resource. Use this filter to find all resources assigned a
	// tag with a specific key, regardless of the tag value.
	//
	//     * volume-id - The
	// volume ID.
	//
	//     * volume-type - The Amazon EBS volume type. This can be gp2 for
	// General Purpose SSD, io1 for Provisioned IOPS SSD, st1 for Throughput Optimized
	// HDD, sc1 for Cold HDD, or standard for Magnetic volumes.
	Filters []*types.Filter
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The NextToken value returned from a previous paginated DescribeVolumes request
	// where MaxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// NextToken value. This value is null when there are no more results to return.
	NextToken *string
	// The maximum number of volume results returned by DescribeVolumes in paginated
	// output. When this parameter is used, DescribeVolumes only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeVolumes
	// request with the returned NextToken value. This value can be between 5 and 500;
	// if MaxResults is given a value larger than 500, only 500 results are returned.
	// If this parameter is not used, then DescribeVolumes returns all results. You
	// cannot specify this parameter and the volume IDs parameter in the same request.
	MaxResults *int32
}

type DescribeVolumesOutput struct {
	// Information about the volumes.
	Volumes []*types.Volume
	// The NextToken value to include in a future DescribeVolumes request. When the
	// results of a DescribeVolumes request exceed MaxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeVolumesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeVolumes{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVolumes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVolumes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVolumes",
	}
}
