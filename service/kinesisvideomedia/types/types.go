// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Identifies the chunk on the Kinesis video stream where you want the GetMedia API
// to start returning media data. You have the following options to identify the
// starting chunk:
//
//     * Choose the latest (or oldest) chunk.
//
//     * Identify a
// specific chunk. You can identify a specific chunk either by providing a fragment
// number or timestamp (server or producer).
//
//     * Each chunk's metadata includes
// a continuation token as a Matroska (MKV) tag
// (AWS_KINESISVIDEO_CONTINUATION_TOKEN). If your previous GetMedia request
// terminated, you can use this tag value in your next GetMedia request. The API
// then starts returning chunks starting where the last API ended.
type StartSelector struct {
	// Continuation token that Kinesis Video Streams returned in the previous GetMedia
	// response. The GetMedia API then starts with the chunk identified by the
	// continuation token.
	ContinuationToken *string
	// A timestamp value. This value is required if you choose the PRODUCER_TIMESTAMP
	// or the SERVER_TIMESTAMP as the startSelectorType. The GetMedia API then starts
	// with the chunk containing the fragment that has the specified timestamp.
	StartTimestamp *time.Time
	// Specifies the fragment number from where you want the GetMedia API to start
	// returning the fragments.
	AfterFragmentNumber *string
	// Identifies the fragment on the Kinesis video stream where you want to start
	// getting the data from.
	//
	//     * NOW - Start with the latest chunk on the stream.
	//
	//
	// * EARLIEST - Start with earliest available chunk on the stream.
	//
	//     *
	// FRAGMENT_NUMBER - Start with the chunk after a specific fragment. You must also
	// specify the AfterFragmentNumber parameter.
	//
	//     * PRODUCER_TIMESTAMP or
	// SERVER_TIMESTAMP - Start with the chunk containing a fragment with the specified
	// producer or server timestamp. You specify the timestamp by adding
	// StartTimestamp.
	//
	//     * CONTINUATION_TOKEN - Read using the specified
	// continuation token.
	//
	// If you choose the NOW, EARLIEST, or CONTINUATION_TOKEN as
	// the startSelectorType, you don't provide any additional information in the
	// startSelector.
	StartSelectorType StartSelectorType
}
