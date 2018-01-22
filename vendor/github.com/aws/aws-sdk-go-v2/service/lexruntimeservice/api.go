// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

const opPostContent = "PostContent"

// PostContentRequest is a API request type for the PostContent API operation.
type PostContentRequest struct {
	*aws.Request
	Input *PostContentInput
}

// Send marshals and sends the PostContent API request.
func (r PostContentRequest) Send() (*PostContentOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*PostContentOutput), nil
}

// PostContentRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Sends user input (text or speech) to Amazon Lex. Clients use this API to
// send text and audio requests to Amazon Lex at runtime. Amazon Lex interprets
// the user input using the machine learning model that it built for the bot.
//
// The PostContent operation supports audio input at 8kHz and 16kHz. You can
// use 8kHz audio to achieve higher speech recognition accuracy in telephone
// audio applications.
//
// In response, Amazon Lex returns the next message to convey to the user. Consider
// the following example messages:
//
//    *  For a user input "I would like a pizza," Amazon Lex might return a
//    response with a message eliciting slot data (for example, PizzaSize):
//    "What size pizza would you like?".
//
//    *  After the user provides all of the pizza order information, Amazon
//    Lex might return a response with a message to get user confirmation: "Order
//    the pizza?".
//
//    *  After the user replies "Yes" to the confirmation prompt, Amazon Lex
//    might return a conclusion statement: "Thank you, your cheese pizza has
//    been ordered.".
//
// Not all Amazon Lex messages require a response from the user. For example,
// conclusion statements do not require a response. Some messages require only
// a yes or no response. In addition to the message, Amazon Lex provides additional
// context about the message in the response that you can use to enhance client
// behavior, such as displaying the appropriate client user interface. Consider
// the following examples:
//
//    *  If the message is to elicit slot data, Amazon Lex returns the following
//    context information:
//
// x-amz-lex-dialog-state header set to ElicitSlot
//
// x-amz-lex-intent-name header set to the intent name in the current context
//
//
// x-amz-lex-slot-to-elicit header set to the slot name for which the message
//    is eliciting information
//
// x-amz-lex-slots header set to a map of slots configured for the intent with
//    their current values
//
//    *  If the message is a confirmation prompt, the x-amz-lex-dialog-state
//    header is set to Confirmation and the x-amz-lex-slot-to-elicit header
//    is omitted.
//
//    *  If the message is a clarification prompt configured for the intent,
//    indicating that the user intent is not understood, the x-amz-dialog-state
//    header is set to ElicitIntent and the x-amz-slot-to-elicit header is omitted.
//
//
// In addition, Amazon Lex also returns your application-specific sessionAttributes.
// For more information, see Managing Conversation Context (http://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html).
//
//    // Example sending a request using the PostContentRequest method.
//    req := client.PostContentRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostContent
func (c *LexRuntimeService) PostContentRequest(input *PostContentInput) PostContentRequest {
	op := &aws.Operation{
		Name:       opPostContent,
		HTTPMethod: "POST",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/content",
	}

	if input == nil {
		input = &PostContentInput{}
	}

	output := &PostContentOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	req.Handlers.Sign.Remove(v4.SignRequestHandler)
	handler := v4.BuildNamedHandler("v4.CustomSignerHandler", v4.WithUnsignedPayload)
	req.Handlers.Sign.PushFrontNamed(handler)
	return PostContentRequest{Request: req, Input: input}
}

const opPostText = "PostText"

// PostTextRequest is a API request type for the PostText API operation.
type PostTextRequest struct {
	*aws.Request
	Input *PostTextInput
}

// Send marshals and sends the PostText API request.
func (r PostTextRequest) Send() (*PostTextOutput, error) {
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	return r.Request.Data.(*PostTextOutput), nil
}

// PostTextRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Sends user input (text-only) to Amazon Lex. Client applications can use this
// API to send requests to Amazon Lex at runtime. Amazon Lex then interprets
// the user input using the machine learning model it built for the bot.
//
// In response, Amazon Lex returns the next message to convey to the user an
// optional responseCard to display. Consider the following example messages:
//
//    *  For a user input "I would like a pizza", Amazon Lex might return a
//    response with a message eliciting slot data (for example, PizzaSize):
//    "What size pizza would you like?"
//
//    *  After the user provides all of the pizza order information, Amazon
//    Lex might return a response with a message to obtain user confirmation
//    "Proceed with the pizza order?".
//
//    *  After the user replies to a confirmation prompt with a "yes", Amazon
//    Lex might return a conclusion statement: "Thank you, your cheese pizza
//    has been ordered.".
//
// Not all Amazon Lex messages require a user response. For example, a conclusion
// statement does not require a response. Some messages require only a "yes"
// or "no" user response. In addition to the message, Amazon Lex provides additional
// context about the message in the response that you might use to enhance client
// behavior, for example, to display the appropriate client user interface.
// These are the slotToElicit, dialogState, intentName, and slots fields in
// the response. Consider the following examples:
//
//    * If the message is to elicit slot data, Amazon Lex returns the following
//    context information:
//
// dialogState set to ElicitSlot
//
// intentName set to the intent name in the current context
//
// slotToElicit set to the slot name for which the message is eliciting information
//
//
// slots set to a map of slots, configured for the intent, with currently known
//    values
//
//    *  If the message is a confirmation prompt, the dialogState is set to
//    ConfirmIntent and SlotToElicit is set to null.
//
//    * If the message is a clarification prompt (configured for the intent)
//    that indicates that user intent is not understood, the dialogState is
//    set to ElicitIntent and slotToElicit is set to null.
//
// In addition, Amazon Lex also returns your application-specific sessionAttributes.
// For more information, see Managing Conversation Context (http://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html).
//
//    // Example sending a request using the PostTextRequest method.
//    req := client.PostTextRequest(params)
//    resp, err := req.Send()
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostText
func (c *LexRuntimeService) PostTextRequest(input *PostTextInput) PostTextRequest {
	op := &aws.Operation{
		Name:       opPostText,
		HTTPMethod: "POST",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/text",
	}

	if input == nil {
		input = &PostTextInput{}
	}

	output := &PostTextOutput{}
	req := c.newRequest(op, input, output)
	output.responseMetadata = aws.Response{Request: req}

	return PostTextRequest{Request: req, Input: input}
}

// Represents an option to be shown on the client platform (Facebook, Slack,
// etc.)
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/Button
type Button struct {
	_ struct{} `type:"structure"`

	// Text that is visible to the user on the button.
	//
	// Text is a required field
	Text *string `locationName:"text" min:"1" type:"string" required:"true"`

	// The value sent to Amazon Lex when a user chooses the button. For example,
	// consider button text "NYC." When the user chooses the button, the value sent
	// can be "New York City."
	//
	// Value is a required field
	Value *string `locationName:"value" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Button) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Button) GoString() string {
	return s.String()
}

// Represents an option rendered to the user when a prompt is shown. It could
// be an image, a button, a link, or text.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/GenericAttachment
type GenericAttachment struct {
	_ struct{} `type:"structure"`

	// The URL of an attachment to the response card.
	AttachmentLinkUrl *string `locationName:"attachmentLinkUrl" min:"1" type:"string"`

	// The list of options to show to the user.
	Buttons []Button `locationName:"buttons" type:"list"`

	// The URL of an image that is displayed to the user.
	ImageUrl *string `locationName:"imageUrl" min:"1" type:"string"`

	// The subtitle shown below the title.
	SubTitle *string `locationName:"subTitle" min:"1" type:"string"`

	// The title of the option.
	Title *string `locationName:"title" min:"1" type:"string"`
}

// String returns the string representation
func (s GenericAttachment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenericAttachment) GoString() string {
	return s.String()
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostContentRequest
type PostContentInput struct {
	_ struct{} `type:"structure" payload:"InputStream"`

	// You pass this value as the Accept HTTP header.
	//
	// The message Amazon Lex returns in the response can be either text or speech
	// based on the Accept HTTP header value in the request.
	//
	//    *  If the value is text/plain; charset=utf-8, Amazon Lex returns text
	//    in the response.
	//
	//    *  If the value begins with audio/, Amazon Lex returns speech in the response.
	//    Amazon Lex uses Amazon Polly to generate the speech (using the configuration
	//    you specified in the Accept header). For example, if you specify audio/mpeg
	//    as the value, Amazon Lex returns speech in the MPEG format.
	//
	// The following are the accepted values:
	//
	// audio/mpeg
	//
	// audio/ogg
	//
	// audio/pcm
	//
	// text/plain; charset=utf-8
	//
	// audio/* (defaults to mpeg)
	Accept *string `location:"header" locationName:"Accept" type:"string"`

	// Alias of the Amazon Lex bot.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// Name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// You pass this value as the Content-Type HTTP header.
	//
	// Indicates the audio format or text. The header value must start with one
	// of the following prefixes:
	//
	//    * PCM format, audio data must be in little-endian byte order.
	//
	// audio/l16; rate=16000; channels=1
	//
	// audio/x-l16; sample-rate=16000; channel-count=1
	//
	// audio/lpcm; sample-rate=8000; sample-size-bits=16; channel-count=1; is-big-endian=false
	//
	//
	//    * Opus format
	//
	// audio/x-cbr-opus-with-preamble; preamble-size=0; bit-rate=256000; frame-size-milliseconds=4
	//
	//    * Text format
	//
	// text/plain; charset=utf-8
	//
	// ContentType is a required field
	ContentType *string `location:"header" locationName:"Content-Type" type:"string" required:"true"`

	// User input in PCM or Opus audio format or text format as described in the
	// Content-Type HTTP header.
	//
	// You can stream audio data to Amazon Lex or you can create a local buffer
	// that captures all of the audio data before sending. In general, you get better
	// performance if you stream audio data rather than buffering the data locally.
	//
	// InputStream is a required field
	InputStream io.ReadSeeker `locationName:"inputStream" type:"blob" required:"true"`

	// You pass this value as the x-amz-lex-request-attributes HTTP header.
	//
	// Request-specific information passed between Amazon Lex and a client application.
	// The value must be a JSON serialized and base64 encoded map with string keys
	// and values. The total size of the requestAttributes and sessionAttributes
	// headers is limited to 12 KB.
	//
	// The namespace x-amz-lex: is reserved for special attributes. Don't create
	// any request attributes with the prefix x-amz-lex:.
	//
	// For more information, see Setting Request Attributes (http://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs).
	RequestAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-request-attributes" type:"jsonvalue"`

	// You pass this value as the x-amz-lex-session-attributes HTTP header.
	//
	// Application-specific information passed between Amazon Lex and a client application.
	// The value must be a JSON serialized and base64 encoded map with string keys
	// and values. The total size of the sessionAttributes and requestAttributes
	// headers is limited to 12 KB.
	//
	// For more information, see Setting Session Attributes (http://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs).
	SessionAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-session-attributes" type:"jsonvalue"`

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot. At runtime, each request must contain
	// the userID field.
	//
	// To decide the user ID to use for your application, consider the following
	// factors.
	//
	//    * The userID field must not contain any personally identifiable information
	//    of the user, for example, name, personal identification numbers, or other
	//    end user personal information.
	//
	//    * If you want a user to start a conversation on one device and continue
	//    on another device, use a user-specific identifier.
	//
	//    * If you want the same user to be able to have two independent conversations
	//    on two different devices, choose a device-specific identifier.
	//
	//    * A user can't have two independent conversations with two different versions
	//    of the same bot. For example, a user can't have a conversation with the
	//    PROD and BETA versions of the same bot. If you anticipate that a user
	//    will need to have conversation with two different versions, for example,
	//    while testing, include the bot alias in the user ID to separate the two
	//    conversations.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PostContentInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostContentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostContentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostContentInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}

	if s.InputStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputStream"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostContentResponse
type PostContentOutput struct {
	_ struct{} `type:"structure" payload:"AudioStream"`

	responseMetadata aws.Response

	// The prompt (or statement) to convey to the user. This is based on the bot
	// configuration and context. For example, if Amazon Lex did not understand
	// the user intent, it sends the clarificationPrompt configured for the bot.
	// If the intent requires confirmation before taking the fulfillment action,
	// it sends the confirmationPrompt. Another example: Suppose that the Lambda
	// function successfully fulfilled the intent, and sent a message to convey
	// to the user. Then Amazon Lex sends that message in the response.
	AudioStream io.ReadCloser `locationName:"audioStream" type:"blob"`

	// Content type as specified in the Accept HTTP header in the request.
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// Identifies the current state of the user interaction. Amazon Lex returns
	// one of the following values as dialogState. The client can optionally use
	// this information to customize the user interface.
	//
	//    * ElicitIntent - Amazon Lex wants to elicit the user's intent. Consider
	//    the following examples:
	//
	//  For example, a user might utter an intent ("I want to order a pizza"). If
	//    Amazon Lex cannot infer the user intent from this utterance, it will return
	//    this dialog state.
	//
	//    * ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response.
	//
	// For example, Amazon Lex wants user confirmation before fulfilling an intent.
	//    Instead of a simple "yes" or "no" response, a user might respond with
	//    additional information. For example, "yes, but make it a thick crust pizza"
	//    or "no, I want to order a drink." Amazon Lex can process such additional
	//    information (in these examples, update the crust type slot or change the
	//    intent from OrderPizza to OrderDrink).
	//
	//    * ElicitSlot - Amazon Lex is expecting the value of a slot for the current
	//    intent.
	//
	//  For example, suppose that in the response Amazon Lex sends this message:
	//    "What size pizza would you like?". A user might reply with the slot value
	//    (e.g., "medium"). The user might also provide additional information in
	//    the response (e.g., "medium thick crust pizza"). Amazon Lex can process
	//    such additional information appropriately.
	//
	//    * Fulfilled - Conveys that the Lambda function has successfully fulfilled
	//    the intent.
	//
	//    * ReadyForFulfillment - Conveys that the client has to fulfill the request.
	//
	//
	//    * Failed - Conveys that the conversation with the user failed.
	//
	//  This can happen for various reasons, including that the user does not provide
	//    an appropriate response to prompts from the service (you can configure
	//    how many times Amazon Lex can prompt a user for specific information),
	//    or if the Lambda function fails to fulfill the intent.
	DialogState DialogState `location:"header" locationName:"x-amz-lex-dialog-state" type:"string" enum:"true"`

	// The text used to process the request.
	//
	// If the input was an audio stream, the inputTranscript field contains the
	// text extracted from the audio stream. This is the text that is actually processed
	// to recognize intents and slot values. You can use this information to determine
	// if Amazon Lex is correctly processing the audio that you send.
	InputTranscript *string `location:"header" locationName:"x-amz-lex-input-transcript" type:"string"`

	// Current user intent that Amazon Lex is aware of.
	IntentName *string `location:"header" locationName:"x-amz-lex-intent-name" type:"string"`

	// Message to convey to the user. It can come from the bot's configuration or
	// a code hook (Lambda function). If the current intent is not configured with
	// a code hook or if the code hook returned Delegate as the dialogAction.type
	// in its response, then Amazon Lex decides the next course of action and selects
	// an appropriate message from the bot configuration based on the current user
	// interaction context. For example, if Amazon Lex is not able to understand
	// the user input, it uses a clarification prompt message (For more information,
	// see the Error Handling section in the Amazon Lex console). Another example:
	// if the intent requires confirmation before fulfillment, then Amazon Lex uses
	// the confirmation prompt message in the intent configuration. If the code
	// hook returns a message, Amazon Lex passes it as-is in its response to the
	// client.
	Message *string `location:"header" locationName:"x-amz-lex-message" min:"1" type:"string"`

	// Map of key/value pairs representing the session-specific context information.
	SessionAttributes aws.JSONValue `location:"header" locationName:"x-amz-lex-session-attributes" type:"jsonvalue"`

	// If the dialogState value is ElicitSlot, returns the name of the slot for
	// which Amazon Lex is eliciting a value.
	SlotToElicit *string `location:"header" locationName:"x-amz-lex-slot-to-elicit" type:"string"`

	// Map of zero or more intent slots (name/value pairs) Amazon Lex detected from
	// the user input during the conversation.
	//
	// Amazon Lex creates a resolution list containing likely values for a slot.
	// The value that it returns is determined by the valueSelectionStrategy selected
	// when the slot type was created or updated. If valueSelectionStrategy is set
	// to ORIGINAL_VALUE, the value provided by the user is returned, if the user
	// value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list
	// or, if there is no resolution list, null. If you don't specify a valueSelectionStrategy,
	// the default is ORIGINAL_VALUE.
	Slots aws.JSONValue `location:"header" locationName:"x-amz-lex-slots" type:"jsonvalue"`
}

// String returns the string representation
func (s PostContentOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostContentOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s PostContentOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostTextRequest
type PostTextInput struct {
	_ struct{} `type:"structure"`

	// The alias of the Amazon Lex bot.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// The name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// The text that the user entered (Amazon Lex interprets this text).
	//
	// InputText is a required field
	InputText *string `locationName:"inputText" min:"1" type:"string" required:"true"`

	// Request-specific information passed between Amazon Lex and a client application.
	//
	// The namespace x-amz-lex: is reserved for special attributes. Don't create
	// any request attributes with the prefix x-amz-lex:.
	//
	// For more information, see Setting Request Attributes (http://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs).
	RequestAttributes map[string]string `locationName:"requestAttributes" type:"map"`

	// Application-specific information passed between Amazon Lex and a client application.
	//
	// For more information, see Setting Session Attributes (http://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs).
	SessionAttributes map[string]string `locationName:"sessionAttributes" type:"map"`

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot. At runtime, each request must contain
	// the userID field.
	//
	// To decide the user ID to use for your application, consider the following
	// factors.
	//
	//    * The userID field must not contain any personally identifiable information
	//    of the user, for example, name, personal identification numbers, or other
	//    end user personal information.
	//
	//    * If you want a user to start a conversation on one device and continue
	//    on another device, use a user-specific identifier.
	//
	//    * If you want the same user to be able to have two independent conversations
	//    on two different devices, choose a device-specific identifier.
	//
	//    * A user can't have two independent conversations with two different versions
	//    of the same bot. For example, a user can't have a conversation with the
	//    PROD and BETA versions of the same bot. If you anticipate that a user
	//    will need to have conversation with two different versions, for example,
	//    while testing, include the bot alias in the user ID to separate the two
	//    conversations.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PostTextInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostTextInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostTextInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostTextInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}

	if s.InputText == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputText"))
	}
	if s.InputText != nil && len(*s.InputText) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputText", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostTextResponse
type PostTextOutput struct {
	_ struct{} `type:"structure"`

	responseMetadata aws.Response

	// Identifies the current state of the user interaction. Amazon Lex returns
	// one of the following values as dialogState. The client can optionally use
	// this information to customize the user interface.
	//
	//    * ElicitIntent - Amazon Lex wants to elicit user intent.
	//
	// For example, a user might utter an intent ("I want to order a pizza"). If
	//    Amazon Lex cannot infer the user intent from this utterance, it will return
	//    this dialogState.
	//
	//    * ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response.
	//
	//  For example, Amazon Lex wants user confirmation before fulfilling an intent.
	//
	//
	// Instead of a simple "yes" or "no," a user might respond with additional information.
	//    For example, "yes, but make it thick crust pizza" or "no, I want to order
	//    a drink". Amazon Lex can process such additional information (in these
	//    examples, update the crust type slot value, or change intent from OrderPizza
	//    to OrderDrink).
	//
	//    * ElicitSlot - Amazon Lex is expecting a slot value for the current intent.
	//
	//
	// For example, suppose that in the response Amazon Lex sends this message:
	//    "What size pizza would you like?". A user might reply with the slot value
	//    (e.g., "medium"). The user might also provide additional information in
	//    the response (e.g., "medium thick crust pizza"). Amazon Lex can process
	//    such additional information appropriately.
	//
	//    * Fulfilled - Conveys that the Lambda function configured for the intent
	//    has successfully fulfilled the intent.
	//
	//    * ReadyForFulfillment - Conveys that the client has to fulfill the intent.
	//
	//
	//    * Failed - Conveys that the conversation with the user failed.
	//
	//  This can happen for various reasons including that the user did not provide
	//    an appropriate response to prompts from the service (you can configure
	//    how many times Amazon Lex can prompt a user for specific information),
	//    or the Lambda function failed to fulfill the intent.
	DialogState DialogState `locationName:"dialogState" type:"string" enum:"true"`

	// The current user intent that Amazon Lex is aware of.
	IntentName *string `locationName:"intentName" type:"string"`

	// A message to convey to the user. It can come from the bot's configuration
	// or a code hook (Lambda function). If the current intent is not configured
	// with a code hook or the code hook returned Delegate as the dialogAction.type
	// in its response, then Amazon Lex decides the next course of action and selects
	// an appropriate message from the bot configuration based on the current user
	// interaction context. For example, if Amazon Lex is not able to understand
	// the user input, it uses a clarification prompt message (for more information,
	// see the Error Handling section in the Amazon Lex console). Another example:
	// if the intent requires confirmation before fulfillment, then Amazon Lex uses
	// the confirmation prompt message in the intent configuration. If the code
	// hook returns a message, Amazon Lex passes it as-is in its response to the
	// client.
	Message *string `locationName:"message" min:"1" type:"string"`

	// Represents the options that the user has to respond to the current prompt.
	// Response Card can come from the bot configuration (in the Amazon Lex console,
	// choose the settings button next to a slot) or from a code hook (Lambda function).
	ResponseCard *ResponseCard `locationName:"responseCard" type:"structure"`

	// A map of key-value pairs representing the session-specific context information.
	SessionAttributes map[string]string `locationName:"sessionAttributes" type:"map"`

	// If the dialogState value is ElicitSlot, returns the name of the slot for
	// which Amazon Lex is eliciting a value.
	SlotToElicit *string `locationName:"slotToElicit" type:"string"`

	// The intent slots that Amazon Lex detected from the user input in the conversation.
	//
	// Amazon Lex creates a resolution list containing likely values for a slot.
	// The value that it returns is determined by the valueSelectionStrategy selected
	// when the slot type was created or updated. If valueSelectionStrategy is set
	// to ORIGINAL_VALUE, the value provided by the user is returned, if the user
	// value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list
	// or, if there is no resolution list, null. If you don't specify a valueSelectionStrategy,
	// the default is ORIGINAL_VALUE.
	Slots map[string]string `locationName:"slots" type:"map"`
}

// String returns the string representation
func (s PostTextOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostTextOutput) GoString() string {
	return s.String()
}

// SDKResponseMetdata return sthe response metadata for the API.
func (s PostTextOutput) SDKResponseMetadata() aws.Response {
	return s.responseMetadata
}

// If you configure a response card when creating your bots, Amazon Lex substitutes
// the session attributes and slot values that are available, and then returns
// it. The response card can also come from a Lambda function ( dialogCodeHook
// and fulfillmentActivity on an intent).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/ResponseCard
type ResponseCard struct {
	_ struct{} `type:"structure"`

	// The content type of the response.
	ContentType ContentType `locationName:"contentType" type:"string" enum:"true"`

	// An array of attachment objects representing options.
	GenericAttachments []GenericAttachment `locationName:"genericAttachments" type:"list"`

	// The version of the response card format.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s ResponseCard) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponseCard) GoString() string {
	return s.String()
}

type ContentType string

// Enum values for ContentType
const (
	ContentTypeApplicationVndAmazonawsCardGeneric ContentType = "application/vnd.amazonaws.card.generic"
)

type DialogState string

// Enum values for DialogState
const (
	DialogStateElicitIntent        DialogState = "ElicitIntent"
	DialogStateConfirmIntent       DialogState = "ConfirmIntent"
	DialogStateElicitSlot          DialogState = "ElicitSlot"
	DialogStateFulfilled           DialogState = "Fulfilled"
	DialogStateReadyForFulfillment DialogState = "ReadyForFulfillment"
	DialogStateFailed              DialogState = "Failed"
)
