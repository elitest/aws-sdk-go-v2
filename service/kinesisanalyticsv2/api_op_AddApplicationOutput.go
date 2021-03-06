// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationOutputInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to which you want to add the output configuration.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The version of the application to which you want to add the output configuration.
	// You can use the DescribeApplication operation to get the current application
	// version. If the version specified is not the current version, the ConcurrentModificationException
	// is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// An array of objects, each describing one output configuration. In the output
	// configuration, you specify the name of an in-application stream, a destination
	// (that is, a Kinesis data stream, a Kinesis Data Firehose delivery stream,
	// or an AWS Lambda function), and record the formation to use when writing
	// to the destination.
	//
	// Output is a required field
	Output *Output `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationOutputInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.Output == nil {
		invalidParams.Add(aws.NewErrParamRequired("Output"))
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			invalidParams.AddNested("Output", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationOutputOutput struct {
	_ struct{} `type:"structure"`

	// The application Amazon Resource Name (ARN).
	ApplicationARN *string `min:"1" type:"string"`

	// The updated application version ID. Kinesis Data Analytics increments this
	// ID when the application is updated.
	ApplicationVersionId *int64 `min:"1" type:"long"`

	// Describes the application output configuration. For more information, see
	// Configuring Application Output (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-output.html).
	OutputDescriptions []OutputDescription `type:"list"`
}

// String returns the string representation
func (s AddApplicationOutputOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddApplicationOutput = "AddApplicationOutput"

// AddApplicationOutputRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Adds an external destination to your SQL-based Amazon Kinesis Data Analytics
// application.
//
// If you want Kinesis Data Analytics to deliver data from an in-application
// stream within your application to an external destination (such as an Kinesis
// data stream, a Kinesis Data Firehose delivery stream, or an AWS Lambda function),
// you add the relevant configuration to your application using this operation.
// You can configure one or more outputs for your application. Each output configuration
// maps an in-application stream and an external destination.
//
// You can use one of the output configurations to deliver data from your in-application
// error stream to an external destination so that you can analyze the errors.
//
// Any configuration update, including adding a streaming source using this
// operation, results in a new version of the application. You can use the DescribeApplication
// operation to find the current application version.
//
//    // Example sending a request using AddApplicationOutputRequest.
//    req := client.AddApplicationOutputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/AddApplicationOutput
func (c *Client) AddApplicationOutputRequest(input *AddApplicationOutputInput) AddApplicationOutputRequest {
	op := &aws.Operation{
		Name:       opAddApplicationOutput,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddApplicationOutputInput{}
	}

	req := c.newRequest(op, input, &AddApplicationOutputOutput{})

	return AddApplicationOutputRequest{Request: req, Input: input, Copy: c.AddApplicationOutputRequest}
}

// AddApplicationOutputRequest is the request type for the
// AddApplicationOutput API operation.
type AddApplicationOutputRequest struct {
	*aws.Request
	Input *AddApplicationOutputInput
	Copy  func(*AddApplicationOutputInput) AddApplicationOutputRequest
}

// Send marshals and sends the AddApplicationOutput API request.
func (r AddApplicationOutputRequest) Send(ctx context.Context) (*AddApplicationOutputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddApplicationOutputResponse{
		AddApplicationOutputOutput: r.Request.Data.(*AddApplicationOutputOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddApplicationOutputResponse is the response type for the
// AddApplicationOutput API operation.
type AddApplicationOutputResponse struct {
	*AddApplicationOutputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddApplicationOutput request.
func (r *AddApplicationOutputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
