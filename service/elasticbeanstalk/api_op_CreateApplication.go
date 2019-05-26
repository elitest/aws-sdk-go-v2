// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to create an application.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateApplicationMessage
type CreateApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application.
	//
	// Constraint: This name must be unique within your account. If the specified
	// name already exists, the action returns an InvalidParameterValue error.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Describes the application.
	Description *string `type:"string"`

	// Specify an application resource lifecycle configuration to prevent your application
	// from accumulating too many versions.
	ResourceLifecycleConfig *ApplicationResourceLifecycleConfig `type:"structure"`

	// Specifies the tags applied to the application.
	//
	// Elastic Beanstalk applies these tags only to the application. Environments
	// that you create in the application don't inherit the tags.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.ResourceLifecycleConfig != nil {
		if err := s.ResourceLifecycleConfig.Validate(); err != nil {
			invalidParams.AddNested("ResourceLifecycleConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message containing a single description of an application.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ApplicationDescriptionMessage
type CreateApplicationOutput struct {
	_ struct{} `type:"structure"`

	// The ApplicationDescription of the application.
	Application *ApplicationDescription `type:"structure"`
}

// String returns the string representation
func (s CreateApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApplication = "CreateApplication"

// CreateApplicationRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Creates an application that has one configuration template named default
// and no application versions.
//
//    // Example sending a request using CreateApplicationRequest.
//    req := client.CreateApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/CreateApplication
func (c *Client) CreateApplicationRequest(input *CreateApplicationInput) CreateApplicationRequest {
	op := &aws.Operation{
		Name:       opCreateApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApplicationInput{}
	}

	req := c.newRequest(op, input, &CreateApplicationOutput{})
	return CreateApplicationRequest{Request: req, Input: input, Copy: c.CreateApplicationRequest}
}

// CreateApplicationRequest is the request type for the
// CreateApplication API operation.
type CreateApplicationRequest struct {
	*aws.Request
	Input *CreateApplicationInput
	Copy  func(*CreateApplicationInput) CreateApplicationRequest
}

// Send marshals and sends the CreateApplication API request.
func (r CreateApplicationRequest) Send(ctx context.Context) (*CreateApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApplicationResponse{
		CreateApplicationOutput: r.Request.Data.(*CreateApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApplicationResponse is the response type for the
// CreateApplication API operation.
type CreateApplicationResponse struct {
	*CreateApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApplication request.
func (r *CreateApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}