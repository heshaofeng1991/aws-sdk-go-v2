// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the settings for a slot.
func (c *Client) UpdateSlot(ctx context.Context, params *UpdateSlotInput, optFns ...func(*Options)) (*UpdateSlotOutput, error) {
	if params == nil {
		params = &UpdateSlotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSlot", params, optFns, c.addOperationUpdateSlotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSlotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSlotInput struct {

	// The unique identifier of the bot that contains the slot.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the slot. Must always be DRAFT.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the intent that contains the slot.
	//
	// This member is required.
	IntentId *string

	// The identifier of the language and locale that contains the slot. The string
	// must match one of the supported locales. For more information, see Supported
	// languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// The unique identifier for the slot to update.
	//
	// This member is required.
	SlotId *string

	// The new name for the slot.
	//
	// This member is required.
	SlotName *string

	// The unique identifier of the new slot type to associate with this slot.
	//
	// This member is required.
	SlotTypeId *string

	// A new set of prompts that Amazon Lex sends to the user to elicit a response the
	// provides a value for the slot.
	//
	// This member is required.
	ValueElicitationSetting *types.SlotValueElicitationSetting

	// The new description for the slot.
	Description *string

	// Determines whether the slot accepts multiple values in one response. Multiple
	// value slots are only available in the en-US locale. If you set this value to
	// true in any other locale, Amazon Lex throws a ValidationException. If the
	// multipleValuesSetting is not set, the default value is false.
	MultipleValuesSetting *types.MultipleValuesSetting

	// New settings that determine how slot values are formatted in Amazon CloudWatch
	// logs.
	ObfuscationSetting *types.ObfuscationSetting

	noSmithyDocumentSerde
}

type UpdateSlotOutput struct {

	// The identifier of the bot that contains the slot.
	BotId *string

	// The identifier of the slot version that contains the slot. Will always be DRAFT.
	BotVersion *string

	// The timestamp of the date and time that the slot was created.
	CreationDateTime *time.Time

	// The updated description of the bot.
	Description *string

	// The intent that contains the slot.
	IntentId *string

	// The timestamp of the date and time that the slot was last updated.
	LastUpdatedDateTime *time.Time

	// The locale that contains the slot.
	LocaleId *string

	// Indicates whether the slot accepts multiple values in one response.
	MultipleValuesSetting *types.MultipleValuesSetting

	// The updated setting that determines whether the slot value is obfuscated in the
	// Amazon CloudWatch logs.
	ObfuscationSetting *types.ObfuscationSetting

	// The unique identifier of the slot that was updated.
	SlotId *string

	// The updated name of the slot.
	SlotName *string

	// The updated identifier of the slot type that provides values for the slot.
	SlotTypeId *string

	// The updated prompts that Amazon Lex sends to the user to elicit a response that
	// provides a value for the slot.
	ValueElicitationSetting *types.SlotValueElicitationSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSlotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSlot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSlot{}, middleware.After)
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
	if err = addOpUpdateSlotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSlot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSlot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "UpdateSlot",
	}
}