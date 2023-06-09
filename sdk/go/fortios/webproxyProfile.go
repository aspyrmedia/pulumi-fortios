// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure web proxy profiles.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-fortios/sdk/go/fortios"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := fortios.NewWebproxyProfile(ctx, "trname", &fortios.WebproxyProfileArgs{
//				HeaderClientIp:             pulumi.String("pass"),
//				HeaderFrontEndHttps:        pulumi.String("pass"),
//				HeaderViaRequest:           pulumi.String("add"),
//				HeaderViaResponse:          pulumi.String("pass"),
//				HeaderXAuthenticatedGroups: pulumi.String("pass"),
//				HeaderXAuthenticatedUser:   pulumi.String("pass"),
//				HeaderXForwardedFor:        pulumi.String("pass"),
//				LogHeaderChange:            pulumi.String("disable"),
//				StripEncoding:              pulumi.String("disable"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # WebProxy Profile can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/webproxyProfile:WebproxyProfile labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/webproxyProfile:WebproxyProfile labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type WebproxyProfile struct {
	pulumi.CustomResourceState

	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp pulumi.StringOutput `pulumi:"headerClientIp"`
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps pulumi.StringOutput `pulumi:"headerFrontEndHttps"`
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest pulumi.StringOutput `pulumi:"headerViaRequest"`
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse pulumi.StringOutput `pulumi:"headerViaResponse"`
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups pulumi.StringOutput `pulumi:"headerXAuthenticatedGroups"`
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser pulumi.StringOutput `pulumi:"headerXAuthenticatedUser"`
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert pulumi.StringOutput `pulumi:"headerXForwardedClientCert"`
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor pulumi.StringOutput `pulumi:"headerXForwardedFor"`
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers WebproxyProfileHeaderArrayOutput `pulumi:"headers"`
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange pulumi.StringOutput `pulumi:"logHeaderChange"`
	// Profile name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding pulumi.StringOutput `pulumi:"stripEncoding"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewWebproxyProfile registers a new resource with the given unique name, arguments, and options.
func NewWebproxyProfile(ctx *pulumi.Context,
	name string, args *WebproxyProfileArgs, opts ...pulumi.ResourceOption) (*WebproxyProfile, error) {
	if args == nil {
		args = &WebproxyProfileArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource WebproxyProfile
	err := ctx.RegisterResource("fortios:index/webproxyProfile:WebproxyProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebproxyProfile gets an existing WebproxyProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebproxyProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebproxyProfileState, opts ...pulumi.ResourceOption) (*WebproxyProfile, error) {
	var resource WebproxyProfile
	err := ctx.ReadResource("fortios:index/webproxyProfile:WebproxyProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebproxyProfile resources.
type webproxyProfileState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp *string `pulumi:"headerClientIp"`
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps *string `pulumi:"headerFrontEndHttps"`
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest *string `pulumi:"headerViaRequest"`
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse *string `pulumi:"headerViaResponse"`
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups *string `pulumi:"headerXAuthenticatedGroups"`
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser *string `pulumi:"headerXAuthenticatedUser"`
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert *string `pulumi:"headerXForwardedClientCert"`
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor *string `pulumi:"headerXForwardedFor"`
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers []WebproxyProfileHeader `pulumi:"headers"`
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange *string `pulumi:"logHeaderChange"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding *string `pulumi:"stripEncoding"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type WebproxyProfileState struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp pulumi.StringPtrInput
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor pulumi.StringPtrInput
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers WebproxyProfileHeaderArrayInput
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebproxyProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*webproxyProfileState)(nil)).Elem()
}

type webproxyProfileArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp *string `pulumi:"headerClientIp"`
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps *string `pulumi:"headerFrontEndHttps"`
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest *string `pulumi:"headerViaRequest"`
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse *string `pulumi:"headerViaResponse"`
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups *string `pulumi:"headerXAuthenticatedGroups"`
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser *string `pulumi:"headerXAuthenticatedUser"`
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert *string `pulumi:"headerXForwardedClientCert"`
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor *string `pulumi:"headerXForwardedFor"`
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers []WebproxyProfileHeader `pulumi:"headers"`
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange *string `pulumi:"logHeaderChange"`
	// Profile name.
	Name *string `pulumi:"name"`
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding *string `pulumi:"stripEncoding"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a WebproxyProfile resource.
type WebproxyProfileArgs struct {
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderClientIp pulumi.StringPtrInput
	// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderFrontEndHttps pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaRequest pulumi.StringPtrInput
	// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderViaResponse pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedGroups pulumi.StringPtrInput
	// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXAuthenticatedUser pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedClientCert pulumi.StringPtrInput
	// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
	HeaderXForwardedFor pulumi.StringPtrInput
	// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
	Headers WebproxyProfileHeaderArrayInput
	// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
	LogHeaderChange pulumi.StringPtrInput
	// Profile name.
	Name pulumi.StringPtrInput
	// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
	StripEncoding pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (WebproxyProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webproxyProfileArgs)(nil)).Elem()
}

type WebproxyProfileInput interface {
	pulumi.Input

	ToWebproxyProfileOutput() WebproxyProfileOutput
	ToWebproxyProfileOutputWithContext(ctx context.Context) WebproxyProfileOutput
}

func (*WebproxyProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**WebproxyProfile)(nil)).Elem()
}

func (i *WebproxyProfile) ToWebproxyProfileOutput() WebproxyProfileOutput {
	return i.ToWebproxyProfileOutputWithContext(context.Background())
}

func (i *WebproxyProfile) ToWebproxyProfileOutputWithContext(ctx context.Context) WebproxyProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyProfileOutput)
}

// WebproxyProfileArrayInput is an input type that accepts WebproxyProfileArray and WebproxyProfileArrayOutput values.
// You can construct a concrete instance of `WebproxyProfileArrayInput` via:
//
//	WebproxyProfileArray{ WebproxyProfileArgs{...} }
type WebproxyProfileArrayInput interface {
	pulumi.Input

	ToWebproxyProfileArrayOutput() WebproxyProfileArrayOutput
	ToWebproxyProfileArrayOutputWithContext(context.Context) WebproxyProfileArrayOutput
}

type WebproxyProfileArray []WebproxyProfileInput

func (WebproxyProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebproxyProfile)(nil)).Elem()
}

func (i WebproxyProfileArray) ToWebproxyProfileArrayOutput() WebproxyProfileArrayOutput {
	return i.ToWebproxyProfileArrayOutputWithContext(context.Background())
}

func (i WebproxyProfileArray) ToWebproxyProfileArrayOutputWithContext(ctx context.Context) WebproxyProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyProfileArrayOutput)
}

// WebproxyProfileMapInput is an input type that accepts WebproxyProfileMap and WebproxyProfileMapOutput values.
// You can construct a concrete instance of `WebproxyProfileMapInput` via:
//
//	WebproxyProfileMap{ "key": WebproxyProfileArgs{...} }
type WebproxyProfileMapInput interface {
	pulumi.Input

	ToWebproxyProfileMapOutput() WebproxyProfileMapOutput
	ToWebproxyProfileMapOutputWithContext(context.Context) WebproxyProfileMapOutput
}

type WebproxyProfileMap map[string]WebproxyProfileInput

func (WebproxyProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebproxyProfile)(nil)).Elem()
}

func (i WebproxyProfileMap) ToWebproxyProfileMapOutput() WebproxyProfileMapOutput {
	return i.ToWebproxyProfileMapOutputWithContext(context.Background())
}

func (i WebproxyProfileMap) ToWebproxyProfileMapOutputWithContext(ctx context.Context) WebproxyProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebproxyProfileMapOutput)
}

type WebproxyProfileOutput struct{ *pulumi.OutputState }

func (WebproxyProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebproxyProfile)(nil)).Elem()
}

func (o WebproxyProfileOutput) ToWebproxyProfileOutput() WebproxyProfileOutput {
	return o
}

func (o WebproxyProfileOutput) ToWebproxyProfileOutputWithContext(ctx context.Context) WebproxyProfileOutput {
	return o
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o WebproxyProfileOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderClientIp() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderClientIp }).(pulumi.StringOutput)
}

// Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderFrontEndHttps() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderFrontEndHttps }).(pulumi.StringOutput)
}

// Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderViaRequest() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderViaRequest }).(pulumi.StringOutput)
}

// Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderViaResponse() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderViaResponse }).(pulumi.StringOutput)
}

// Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderXAuthenticatedGroups() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderXAuthenticatedGroups }).(pulumi.StringOutput)
}

// Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderXAuthenticatedUser() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderXAuthenticatedUser }).(pulumi.StringOutput)
}

// Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderXForwardedClientCert() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderXForwardedClientCert }).(pulumi.StringOutput)
}

// Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.
func (o WebproxyProfileOutput) HeaderXForwardedFor() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.HeaderXForwardedFor }).(pulumi.StringOutput)
}

// Configure HTTP forwarded requests headers. The structure of `headers` block is documented below.
func (o WebproxyProfileOutput) Headers() WebproxyProfileHeaderArrayOutput {
	return o.ApplyT(func(v *WebproxyProfile) WebproxyProfileHeaderArrayOutput { return v.Headers }).(WebproxyProfileHeaderArrayOutput)
}

// Enable/disable logging HTTP header changes. Valid values: `enable`, `disable`.
func (o WebproxyProfileOutput) LogHeaderChange() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.LogHeaderChange }).(pulumi.StringOutput)
}

// Profile name.
func (o WebproxyProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable/disable stripping unsupported encoding from the request header. Valid values: `enable`, `disable`.
func (o WebproxyProfileOutput) StripEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringOutput { return v.StripEncoding }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o WebproxyProfileOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebproxyProfile) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type WebproxyProfileArrayOutput struct{ *pulumi.OutputState }

func (WebproxyProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebproxyProfile)(nil)).Elem()
}

func (o WebproxyProfileArrayOutput) ToWebproxyProfileArrayOutput() WebproxyProfileArrayOutput {
	return o
}

func (o WebproxyProfileArrayOutput) ToWebproxyProfileArrayOutputWithContext(ctx context.Context) WebproxyProfileArrayOutput {
	return o
}

func (o WebproxyProfileArrayOutput) Index(i pulumi.IntInput) WebproxyProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebproxyProfile {
		return vs[0].([]*WebproxyProfile)[vs[1].(int)]
	}).(WebproxyProfileOutput)
}

type WebproxyProfileMapOutput struct{ *pulumi.OutputState }

func (WebproxyProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebproxyProfile)(nil)).Elem()
}

func (o WebproxyProfileMapOutput) ToWebproxyProfileMapOutput() WebproxyProfileMapOutput {
	return o
}

func (o WebproxyProfileMapOutput) ToWebproxyProfileMapOutputWithContext(ctx context.Context) WebproxyProfileMapOutput {
	return o
}

func (o WebproxyProfileMapOutput) MapIndex(k pulumi.StringInput) WebproxyProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebproxyProfile {
		return vs[0].(map[string]*WebproxyProfile)[vs[1].(string)]
	}).(WebproxyProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyProfileInput)(nil)).Elem(), &WebproxyProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyProfileArrayInput)(nil)).Elem(), WebproxyProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebproxyProfileMapInput)(nil)).Elem(), WebproxyProfileMap{})
	pulumi.RegisterOutputType(WebproxyProfileOutput{})
	pulumi.RegisterOutputType(WebproxyProfileArrayOutput{})
	pulumi.RegisterOutputType(WebproxyProfileMapOutput{})
}
