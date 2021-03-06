// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type DescribeGameSessionQueuesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return. Use this parameter with NextToken
	// to get results as a set of sequential pages.
	Limit *int64 `min:"1" type:"integer"`

	// A list of queue names to retrieve information for. You can use either the
	// queue ID or ARN value. To request settings for all queues, leave this parameter
	// empty.
	Names []string `type:"list"`

	// A token that indicates the start of the next sequential page of results.
	// Use the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeGameSessionQueuesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGameSessionQueuesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGameSessionQueuesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type DescribeGameSessionQueuesOutput struct {
	_ struct{} `type:"structure"`

	// A collection of objects that describe the requested game session queues.
	GameSessionQueues []GameSessionQueue `type:"list"`

	// A token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeGameSessionQueuesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeGameSessionQueues = "DescribeGameSessionQueues"

// DescribeGameSessionQueuesRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the properties for one or more game session queues. When requesting
// multiple queues, use the pagination parameters to retrieve results as a set
// of sequential pages. If successful, a GameSessionQueue object is returned
// for each requested queue. When specifying a list of queues, objects are returned
// only for queues that currently exist in the Region.
//
// Learn more
//
//  View Your Queues (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-console.html)
//
// Related operations
//
//    * CreateGameSessionQueue
//
//    * DescribeGameSessionQueues
//
//    * UpdateGameSessionQueue
//
//    * DeleteGameSessionQueue
//
//    // Example sending a request using DescribeGameSessionQueuesRequest.
//    req := client.DescribeGameSessionQueuesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeGameSessionQueues
func (c *Client) DescribeGameSessionQueuesRequest(input *DescribeGameSessionQueuesInput) DescribeGameSessionQueuesRequest {
	op := &aws.Operation{
		Name:       opDescribeGameSessionQueues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeGameSessionQueuesInput{}
	}

	req := c.newRequest(op, input, &DescribeGameSessionQueuesOutput{})

	return DescribeGameSessionQueuesRequest{Request: req, Input: input, Copy: c.DescribeGameSessionQueuesRequest}
}

// DescribeGameSessionQueuesRequest is the request type for the
// DescribeGameSessionQueues API operation.
type DescribeGameSessionQueuesRequest struct {
	*aws.Request
	Input *DescribeGameSessionQueuesInput
	Copy  func(*DescribeGameSessionQueuesInput) DescribeGameSessionQueuesRequest
}

// Send marshals and sends the DescribeGameSessionQueues API request.
func (r DescribeGameSessionQueuesRequest) Send(ctx context.Context) (*DescribeGameSessionQueuesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGameSessionQueuesResponse{
		DescribeGameSessionQueuesOutput: r.Request.Data.(*DescribeGameSessionQueuesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGameSessionQueuesResponse is the response type for the
// DescribeGameSessionQueues API operation.
type DescribeGameSessionQueuesResponse struct {
	*DescribeGameSessionQueuesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGameSessionQueues request.
func (r *DescribeGameSessionQueuesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
