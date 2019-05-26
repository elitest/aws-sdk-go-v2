// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// The name of the authorizer to delete.
	//
	// AuthorizerName is a required field
	AuthorizerName *string `location:"uri" locationName:"authorizerName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAuthorizerInput"}

	if s.AuthorizerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthorizerName"))
	}
	if s.AuthorizerName != nil && len(*s.AuthorizerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthorizerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AuthorizerName != nil {
		v := *s.AuthorizerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "authorizerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteAuthorizerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAuthorizer = "DeleteAuthorizer"

// DeleteAuthorizerRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes an authorizer.
//
//    // Example sending a request using DeleteAuthorizerRequest.
//    req := client.DeleteAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteAuthorizerRequest(input *DeleteAuthorizerInput) DeleteAuthorizerRequest {
	op := &aws.Operation{
		Name:       opDeleteAuthorizer,
		HTTPMethod: "DELETE",
		HTTPPath:   "/authorizer/{authorizerName}",
	}

	if input == nil {
		input = &DeleteAuthorizerInput{}
	}

	req := c.newRequest(op, input, &DeleteAuthorizerOutput{})
	return DeleteAuthorizerRequest{Request: req, Input: input, Copy: c.DeleteAuthorizerRequest}
}

// DeleteAuthorizerRequest is the request type for the
// DeleteAuthorizer API operation.
type DeleteAuthorizerRequest struct {
	*aws.Request
	Input *DeleteAuthorizerInput
	Copy  func(*DeleteAuthorizerInput) DeleteAuthorizerRequest
}

// Send marshals and sends the DeleteAuthorizer API request.
func (r DeleteAuthorizerRequest) Send(ctx context.Context) (*DeleteAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAuthorizerResponse{
		DeleteAuthorizerOutput: r.Request.Data.(*DeleteAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAuthorizerResponse is the response type for the
// DeleteAuthorizer API operation.
type DeleteAuthorizerResponse struct {
	*DeleteAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAuthorizer request.
func (r *DeleteAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}