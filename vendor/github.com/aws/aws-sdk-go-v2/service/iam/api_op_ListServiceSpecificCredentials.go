// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the service-specific credentials associated with the
// specified IAM user. If none exists, the operation returns an empty list. The
// service-specific credentials returned by this operation are used only for
// authenticating the IAM user to a specific service. For more information about
// using service-specific credentials to authenticate to an Amazon Web Services
// service, see Set up service-specific credentials (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-gc.html)
// in the CodeCommit User Guide.
func (c *Client) ListServiceSpecificCredentials(ctx context.Context, params *ListServiceSpecificCredentialsInput, optFns ...func(*Options)) (*ListServiceSpecificCredentialsOutput, error) {
	if params == nil {
		params = &ListServiceSpecificCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceSpecificCredentials", params, optFns, c.addOperationListServiceSpecificCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceSpecificCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceSpecificCredentialsInput struct {

	// Filters the returned results to only those for the specified Amazon Web
	// Services service. If not specified, then Amazon Web Services returns
	// service-specific credentials for all services.
	ServiceName *string

	// The name of the user whose service-specific credentials you want information
	// about. If this value is not specified, then the operation assumes the user whose
	// credentials are used to call the operation. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex) ) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	UserName *string

	noSmithyDocumentSerde
}

type ListServiceSpecificCredentialsOutput struct {

	// A list of structures that each contain details about a service-specific
	// credential.
	ServiceSpecificCredentials []types.ServiceSpecificCredentialMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceSpecificCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListServiceSpecificCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListServiceSpecificCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServiceSpecificCredentials"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceSpecificCredentials(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListServiceSpecificCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServiceSpecificCredentials",
	}
}
