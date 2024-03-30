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

// Lists all IAM users, groups, and roles that the specified managed policy is
// attached to. You can use the optional EntityFilter parameter to limit the
// results to a particular type of entity (users, groups, or roles). For example,
// to list only the roles that are attached to the specified policy, set
// EntityFilter to Role . You can paginate the results using the MaxItems and
// Marker parameters.
func (c *Client) ListEntitiesForPolicy(ctx context.Context, params *ListEntitiesForPolicyInput, optFns ...func(*Options)) (*ListEntitiesForPolicyOutput, error) {
	if params == nil {
		params = &ListEntitiesForPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitiesForPolicy", params, optFns, c.addOperationListEntitiesForPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitiesForPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesForPolicyInput struct {

	// The Amazon Resource Name (ARN) of the IAM policy for which you want the
	// versions. For more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	//
	// This member is required.
	PolicyArn *string

	// The entity type to use for filtering the results. For example, when EntityFilter
	// is Role , only the roles that are attached to the specified policy are returned.
	// This parameter is optional. If it is not included, all attached entities (users,
	// groups, and roles) are returned. The argument for this parameter must be one of
	// the valid values listed below.
	EntityFilter types.EntityType

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The path prefix for filtering the results. This parameter is optional. If it is
	// not included, it defaults to a slash (/), listing all entities. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string
	// of characters consisting of either a forward slash (/) by itself or a string
	// that must begin and end with forward slashes. In addition, it can contain any
	// ASCII character from the ! ( \u0021 ) through the DEL character ( \u007F ),
	// including most punctuation characters, digits, and upper and lowercased letters.
	PathPrefix *string

	// The policy usage method to use for filtering the results. To list only
	// permissions policies, set PolicyUsageFilter to PermissionsPolicy . To list only
	// the policies used to set permissions boundaries, set the value to
	// PermissionsBoundary . This parameter is optional. If it is not included, all
	// policies are returned.
	PolicyUsageFilter types.PolicyUsageType

	noSmithyDocumentSerde
}

// Contains the response to a successful ListEntitiesForPolicy request.
type ListEntitiesForPolicyOutput struct {

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated bool

	// When IsTruncated is true , this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// A list of IAM groups that the policy is attached to.
	PolicyGroups []types.PolicyGroup

	// A list of IAM roles that the policy is attached to.
	PolicyRoles []types.PolicyRole

	// A list of IAM users that the policy is attached to.
	PolicyUsers []types.PolicyUser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntitiesForPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListEntitiesForPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListEntitiesForPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEntitiesForPolicy"); err != nil {
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
	if err = addOpListEntitiesForPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitiesForPolicy(options.Region), middleware.Before); err != nil {
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

// ListEntitiesForPolicyAPIClient is a client that implements the
// ListEntitiesForPolicy operation.
type ListEntitiesForPolicyAPIClient interface {
	ListEntitiesForPolicy(context.Context, *ListEntitiesForPolicyInput, ...func(*Options)) (*ListEntitiesForPolicyOutput, error)
}

var _ ListEntitiesForPolicyAPIClient = (*Client)(nil)

// ListEntitiesForPolicyPaginatorOptions is the paginator options for
// ListEntitiesForPolicy
type ListEntitiesForPolicyPaginatorOptions struct {
	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true . If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true , and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEntitiesForPolicyPaginator is a paginator for ListEntitiesForPolicy
type ListEntitiesForPolicyPaginator struct {
	options   ListEntitiesForPolicyPaginatorOptions
	client    ListEntitiesForPolicyAPIClient
	params    *ListEntitiesForPolicyInput
	nextToken *string
	firstPage bool
}

// NewListEntitiesForPolicyPaginator returns a new ListEntitiesForPolicyPaginator
func NewListEntitiesForPolicyPaginator(client ListEntitiesForPolicyAPIClient, params *ListEntitiesForPolicyInput, optFns ...func(*ListEntitiesForPolicyPaginatorOptions)) *ListEntitiesForPolicyPaginator {
	if params == nil {
		params = &ListEntitiesForPolicyInput{}
	}

	options := ListEntitiesForPolicyPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEntitiesForPolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEntitiesForPolicyPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEntitiesForPolicy page.
func (p *ListEntitiesForPolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEntitiesForPolicyOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListEntitiesForPolicy(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListEntitiesForPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEntitiesForPolicy",
	}
}
