// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListIAMPolicyAssignmentsInput struct {
	_ struct{} `type:"structure"`

	// The status of the assignments.
	AssignmentStatus AssignmentStatus `type:"string" enum:"true"`

	// The ID of the AWS account that contains these IAM policy assignments.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The maximum number of results to be returned per request.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The namespace for the assignments.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListIAMPolicyAssignmentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIAMPolicyAssignmentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIAMPolicyAssignmentsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListIAMPolicyAssignmentsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.AssignmentStatus) > 0 {
		v := s.AssignmentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AssignmentStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListIAMPolicyAssignmentsOutput struct {
	_ struct{} `type:"structure"`

	// Information describing the IAM policy assignments.
	IAMPolicyAssignments []IAMPolicyAssignmentSummary `type:"list"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s ListIAMPolicyAssignmentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListIAMPolicyAssignmentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.IAMPolicyAssignments != nil {
		v := s.IAMPolicyAssignments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "IAMPolicyAssignments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opListIAMPolicyAssignments = "ListIAMPolicyAssignments"

// ListIAMPolicyAssignmentsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists IAM policy assignments in the current Amazon QuickSight account.
//
//    // Example sending a request using ListIAMPolicyAssignmentsRequest.
//    req := client.ListIAMPolicyAssignmentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListIAMPolicyAssignments
func (c *Client) ListIAMPolicyAssignmentsRequest(input *ListIAMPolicyAssignmentsInput) ListIAMPolicyAssignmentsRequest {
	op := &aws.Operation{
		Name:       opListIAMPolicyAssignments,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/iam-policy-assignments",
	}

	if input == nil {
		input = &ListIAMPolicyAssignmentsInput{}
	}

	req := c.newRequest(op, input, &ListIAMPolicyAssignmentsOutput{})

	return ListIAMPolicyAssignmentsRequest{Request: req, Input: input, Copy: c.ListIAMPolicyAssignmentsRequest}
}

// ListIAMPolicyAssignmentsRequest is the request type for the
// ListIAMPolicyAssignments API operation.
type ListIAMPolicyAssignmentsRequest struct {
	*aws.Request
	Input *ListIAMPolicyAssignmentsInput
	Copy  func(*ListIAMPolicyAssignmentsInput) ListIAMPolicyAssignmentsRequest
}

// Send marshals and sends the ListIAMPolicyAssignments API request.
func (r ListIAMPolicyAssignmentsRequest) Send(ctx context.Context) (*ListIAMPolicyAssignmentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIAMPolicyAssignmentsResponse{
		ListIAMPolicyAssignmentsOutput: r.Request.Data.(*ListIAMPolicyAssignmentsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIAMPolicyAssignmentsResponse is the response type for the
// ListIAMPolicyAssignments API operation.
type ListIAMPolicyAssignmentsResponse struct {
	*ListIAMPolicyAssignmentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIAMPolicyAssignments request.
func (r *ListIAMPolicyAssignmentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
