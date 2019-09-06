// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the DescribeIdentityPool action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/DescribeIdentityPoolInput
type DescribeIdentityPoolInput struct {
	_ struct{} `type:"structure"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIdentityPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIdentityPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIdentityPoolInput"}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An object representing an Amazon Cognito identity pool.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/IdentityPool
type DescribeIdentityPoolOutput struct {
	_ struct{} `type:"structure"`

	// TRUE if the identity pool supports unauthenticated logins.
	//
	// AllowUnauthenticatedIdentities is a required field
	AllowUnauthenticatedIdentities *bool `type:"boolean" required:"true"`

	// A list representing an Amazon Cognito user pool and its client ID.
	CognitoIdentityProviders []CognitoIdentityProvider `type:"list"`

	// The "domain" by which Cognito will refer to your users.
	DeveloperProviderName *string `min:"1" type:"string"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`

	// A string that you provide.
	//
	// IdentityPoolName is a required field
	IdentityPoolName *string `min:"1" type:"string" required:"true"`

	// The tags that are assigned to the identity pool. A tag is a label that you
	// can apply to identity pools to categorize and manage them in different ways,
	// such as by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]string `type:"map"`

	// A list of OpendID Connect provider ARNs.
	OpenIdConnectProviderARNs []string `type:"list"`

	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []string `type:"list"`

	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string `type:"map"`
}

// String returns the string representation
func (s DescribeIdentityPoolOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeIdentityPool = "DescribeIdentityPool"

// DescribeIdentityPoolRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Gets details about a particular identity pool, including the pool name, ID
// description, creation date, and current number of users.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using DescribeIdentityPoolRequest.
//    req := client.DescribeIdentityPoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/DescribeIdentityPool
func (c *Client) DescribeIdentityPoolRequest(input *DescribeIdentityPoolInput) DescribeIdentityPoolRequest {
	op := &aws.Operation{
		Name:       opDescribeIdentityPool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIdentityPoolInput{}
	}

	req := c.newRequest(op, input, &DescribeIdentityPoolOutput{})
	return DescribeIdentityPoolRequest{Request: req, Input: input, Copy: c.DescribeIdentityPoolRequest}
}

// DescribeIdentityPoolRequest is the request type for the
// DescribeIdentityPool API operation.
type DescribeIdentityPoolRequest struct {
	*aws.Request
	Input *DescribeIdentityPoolInput
	Copy  func(*DescribeIdentityPoolInput) DescribeIdentityPoolRequest
}

// Send marshals and sends the DescribeIdentityPool API request.
func (r DescribeIdentityPoolRequest) Send(ctx context.Context) (*DescribeIdentityPoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIdentityPoolResponse{
		DescribeIdentityPoolOutput: r.Request.Data.(*DescribeIdentityPoolOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIdentityPoolResponse is the response type for the
// DescribeIdentityPool API operation.
type DescribeIdentityPoolResponse struct {
	*DescribeIdentityPoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIdentityPool request.
func (r *DescribeIdentityPoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}