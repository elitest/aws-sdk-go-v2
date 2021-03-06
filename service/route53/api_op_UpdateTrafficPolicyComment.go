// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the traffic policy that you
// want to update the comment for.
type UpdateTrafficPolicyCommentInput struct {
	_ struct{} `locationName:"UpdateTrafficPolicyCommentRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// The new comment for the specified traffic policy and version.
	//
	// Comment is a required field
	Comment *string `type:"string" required:"true"`

	// The value of Id for the traffic policy that you want to update the comment
	// for.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`

	// The value of Version for the traffic policy that you want to update the comment
	// for.
	//
	// Version is a required field
	Version *int64 `location:"uri" locationName:"Version" min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s UpdateTrafficPolicyCommentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTrafficPolicyCommentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTrafficPolicyCommentInput"}

	if s.Comment == nil {
		invalidParams.Add(aws.NewErrParamRequired("Comment"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && *s.Version < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTrafficPolicyCommentInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "UpdateTrafficPolicyCommentRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.Comment != nil {
			v := *s.Comment

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "Comment", protocol.StringValue(v), metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Version", protocol.Int64Value(v), metadata)
	}
	return nil
}

// A complex type that contains the response information for the traffic policy.
type UpdateTrafficPolicyCommentOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains settings for the specified traffic policy.
	//
	// TrafficPolicy is a required field
	TrafficPolicy *TrafficPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTrafficPolicyCommentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTrafficPolicyCommentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TrafficPolicy != nil {
		v := s.TrafficPolicy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrafficPolicy", v, metadata)
	}
	return nil
}

const opUpdateTrafficPolicyComment = "UpdateTrafficPolicyComment"

// UpdateTrafficPolicyCommentRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Updates the comment for a specified traffic policy version.
//
//    // Example sending a request using UpdateTrafficPolicyCommentRequest.
//    req := client.UpdateTrafficPolicyCommentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateTrafficPolicyComment
func (c *Client) UpdateTrafficPolicyCommentRequest(input *UpdateTrafficPolicyCommentInput) UpdateTrafficPolicyCommentRequest {
	op := &aws.Operation{
		Name:       opUpdateTrafficPolicyComment,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/trafficpolicy/{Id}/{Version}",
	}

	if input == nil {
		input = &UpdateTrafficPolicyCommentInput{}
	}

	req := c.newRequest(op, input, &UpdateTrafficPolicyCommentOutput{})

	return UpdateTrafficPolicyCommentRequest{Request: req, Input: input, Copy: c.UpdateTrafficPolicyCommentRequest}
}

// UpdateTrafficPolicyCommentRequest is the request type for the
// UpdateTrafficPolicyComment API operation.
type UpdateTrafficPolicyCommentRequest struct {
	*aws.Request
	Input *UpdateTrafficPolicyCommentInput
	Copy  func(*UpdateTrafficPolicyCommentInput) UpdateTrafficPolicyCommentRequest
}

// Send marshals and sends the UpdateTrafficPolicyComment API request.
func (r UpdateTrafficPolicyCommentRequest) Send(ctx context.Context) (*UpdateTrafficPolicyCommentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrafficPolicyCommentResponse{
		UpdateTrafficPolicyCommentOutput: r.Request.Data.(*UpdateTrafficPolicyCommentOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrafficPolicyCommentResponse is the response type for the
// UpdateTrafficPolicyComment API operation.
type UpdateTrafficPolicyCommentResponse struct {
	*UpdateTrafficPolicyCommentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrafficPolicyComment request.
func (r *UpdateTrafficPolicyCommentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
