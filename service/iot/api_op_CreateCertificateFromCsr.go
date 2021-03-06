// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the CreateCertificateFromCsr operation.
type CreateCertificateFromCsrInput struct {
	_ struct{} `type:"structure"`

	// The certificate signing request (CSR).
	//
	// CertificateSigningRequest is a required field
	CertificateSigningRequest *string `locationName:"certificateSigningRequest" min:"1" type:"string" required:"true"`

	// Specifies whether the certificate is active.
	SetAsActive *bool `location:"querystring" locationName:"setAsActive" type:"boolean"`
}

// String returns the string representation
func (s CreateCertificateFromCsrInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCertificateFromCsrInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCertificateFromCsrInput"}

	if s.CertificateSigningRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateSigningRequest"))
	}
	if s.CertificateSigningRequest != nil && len(*s.CertificateSigningRequest) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateSigningRequest", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCertificateFromCsrInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CertificateSigningRequest != nil {
		v := *s.CertificateSigningRequest

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateSigningRequest", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SetAsActive != nil {
		v := *s.SetAsActive

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "setAsActive", protocol.BoolValue(v), metadata)
	}
	return nil
}

// The output from the CreateCertificateFromCsr operation.
type CreateCertificateFromCsrOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the certificate. You can use the ARN as
	// a principal for policy operations.
	CertificateArn *string `locationName:"certificateArn" type:"string"`

	// The ID of the certificate. Certificate management operations only take a
	// certificateId.
	CertificateId *string `locationName:"certificateId" min:"64" type:"string"`

	// The certificate data, in PEM format.
	CertificatePem *string `locationName:"certificatePem" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateCertificateFromCsrOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCertificateFromCsrOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateArn != nil {
		v := *s.CertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificatePem != nil {
		v := *s.CertificatePem

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificatePem", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateCertificateFromCsr = "CreateCertificateFromCsr"

// CreateCertificateFromCsrRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates an X.509 certificate using the specified certificate signing request.
//
// Note: The CSR must include a public key that is either an RSA key with a
// length of at least 2048 bits or an ECC key from NIST P-256 or NIST P-384
// curves.
//
// Note: Reusing the same certificate signing request (CSR) results in a distinct
// certificate.
//
// You can create multiple certificates in a batch by creating a directory,
// copying multiple .csr files into that directory, and then specifying that
// directory on the command line. The following commands show how to create
// a batch of certificates given a batch of CSRs.
//
// Assuming a set of CSRs are located inside of the directory my-csr-directory:
//
// On Linux and OS X, the command is:
//
// $ ls my-csr-directory/ | xargs -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{}
//
// This command lists all of the CSRs in my-csr-directory and pipes each CSR
// file name to the aws iot create-certificate-from-csr AWS CLI command to create
// a certificate for the corresponding CSR.
//
// The aws iot create-certificate-from-csr part of the command can also be run
// in parallel to speed up the certificate creation process:
//
// $ ls my-csr-directory/ | xargs -P 10 -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{}
//
// On Windows PowerShell, the command to create certificates for all CSRs in
// my-csr-directory is:
//
// > ls -Name my-csr-directory | %{aws iot create-certificate-from-csr --certificate-signing-request
// file://my-csr-directory/$_}
//
// On a Windows command prompt, the command to create certificates for all CSRs
// in my-csr-directory is:
//
// > forfiles /p my-csr-directory /c "cmd /c aws iot create-certificate-from-csr
// --certificate-signing-request file://@path"
//
//    // Example sending a request using CreateCertificateFromCsrRequest.
//    req := client.CreateCertificateFromCsrRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateCertificateFromCsrRequest(input *CreateCertificateFromCsrInput) CreateCertificateFromCsrRequest {
	op := &aws.Operation{
		Name:       opCreateCertificateFromCsr,
		HTTPMethod: "POST",
		HTTPPath:   "/certificates",
	}

	if input == nil {
		input = &CreateCertificateFromCsrInput{}
	}

	req := c.newRequest(op, input, &CreateCertificateFromCsrOutput{})

	return CreateCertificateFromCsrRequest{Request: req, Input: input, Copy: c.CreateCertificateFromCsrRequest}
}

// CreateCertificateFromCsrRequest is the request type for the
// CreateCertificateFromCsr API operation.
type CreateCertificateFromCsrRequest struct {
	*aws.Request
	Input *CreateCertificateFromCsrInput
	Copy  func(*CreateCertificateFromCsrInput) CreateCertificateFromCsrRequest
}

// Send marshals and sends the CreateCertificateFromCsr API request.
func (r CreateCertificateFromCsrRequest) Send(ctx context.Context) (*CreateCertificateFromCsrResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCertificateFromCsrResponse{
		CreateCertificateFromCsrOutput: r.Request.Data.(*CreateCertificateFromCsrOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCertificateFromCsrResponse is the response type for the
// CreateCertificateFromCsr API operation.
type CreateCertificateFromCsrResponse struct {
	*CreateCertificateFromCsrOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCertificateFromCsr request.
func (r *CreateCertificateFromCsrResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
