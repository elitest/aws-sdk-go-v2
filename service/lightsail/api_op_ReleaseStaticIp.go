// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ReleaseStaticIpInput struct {
	_ struct{} `type:"structure"`

	// The name of the static IP to delete.
	//
	// StaticIpName is a required field
	StaticIpName *string `locationName:"staticIpName" type:"string" required:"true"`
}

// String returns the string representation
func (s ReleaseStaticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseStaticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReleaseStaticIpInput"}

	if s.StaticIpName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StaticIpName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReleaseStaticIpOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected
	// by the request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s ReleaseStaticIpOutput) String() string {
	return awsutil.Prettify(s)
}

const opReleaseStaticIp = "ReleaseStaticIp"

// ReleaseStaticIpRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a specific static IP from your account.
//
//    // Example sending a request using ReleaseStaticIpRequest.
//    req := client.ReleaseStaticIpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/ReleaseStaticIp
func (c *Client) ReleaseStaticIpRequest(input *ReleaseStaticIpInput) ReleaseStaticIpRequest {
	op := &aws.Operation{
		Name:       opReleaseStaticIp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseStaticIpInput{}
	}

	req := c.newRequest(op, input, &ReleaseStaticIpOutput{})

	return ReleaseStaticIpRequest{Request: req, Input: input, Copy: c.ReleaseStaticIpRequest}
}

// ReleaseStaticIpRequest is the request type for the
// ReleaseStaticIp API operation.
type ReleaseStaticIpRequest struct {
	*aws.Request
	Input *ReleaseStaticIpInput
	Copy  func(*ReleaseStaticIpInput) ReleaseStaticIpRequest
}

// Send marshals and sends the ReleaseStaticIp API request.
func (r ReleaseStaticIpRequest) Send(ctx context.Context) (*ReleaseStaticIpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReleaseStaticIpResponse{
		ReleaseStaticIpOutput: r.Request.Data.(*ReleaseStaticIpOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReleaseStaticIpResponse is the response type for the
// ReleaseStaticIp API operation.
type ReleaseStaticIpResponse struct {
	*ReleaseStaticIpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReleaseStaticIp request.
func (r *ReleaseStaticIpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
