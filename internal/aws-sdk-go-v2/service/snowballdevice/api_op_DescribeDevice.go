// Code generated by smithy-go-codegen DO NOT EDIT.

package snowballdevice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/eks-anywhere/internal/aws-sdk-go-v2/service/snowballdevice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) DescribeDevice(ctx context.Context, params *DescribeDeviceInput, optFns ...func(*Options)) (*DescribeDeviceOutput, error) {
	if params == nil {
		params = &DescribeDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDevice", params, optFns, c.addOperationDescribeDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDeviceInput struct {
	noSmithyDocumentSerde
}

type DescribeDeviceOutput struct {

	// This member is required.
	DeviceId *string

	// This member is required.
	UnlockStatus *types.UnlockStatus

	ActiveNetworkInterface *types.NetworkInterface

	ClusterAssociation *types.ClusterAssociation

	DeviceCapacities []types.Capacity

	DeviceType *string

	PhysicalNetworkInterfaces []types.PhysicalNetworkInterface

	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDevice(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDescribeDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowballdevice",
		OperationName: "DescribeDevice",
	}
}