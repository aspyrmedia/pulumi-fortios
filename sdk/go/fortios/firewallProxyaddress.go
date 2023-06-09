// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Web proxy address configuration.
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
//			_, err := fortios.NewFirewallProxyaddress(ctx, "trname", &fortios.FirewallProxyaddressArgs{
//				CaseSensitivity: pulumi.String("disable"),
//				Color:           pulumi.Int(2),
//				Referrer:        pulumi.String("enable"),
//				Type:            pulumi.String("host-regex"),
//				Visibility:      pulumi.String("enable"),
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
// # Firewall ProxyAddress can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallProxyaddress:FirewallProxyaddress labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallProxyaddress:FirewallProxyaddress labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallProxyaddress struct {
	pulumi.CustomResourceState

	// SaaS application. The structure of `application` block is documented below.
	Applications FirewallProxyaddressApplicationArrayOutput `pulumi:"applications"`
	// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringOutput `pulumi:"caseSensitivity"`
	// FortiGuard category ID. The structure of `category` block is documented below.
	Categories FirewallProxyaddressCategoryArrayOutput `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntOutput `pulumi:"color"`
	// Optional comments.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// HTTP header name as a regular expression.
	Header pulumi.StringOutput `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups FirewallProxyaddressHeaderGroupArrayOutput `pulumi:"headerGroups"`
	// Name of HTTP header.
	HeaderName pulumi.StringOutput `pulumi:"headerName"`
	// Address object for the host.
	Host pulumi.StringOutput `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex pulumi.StringOutput `pulumi:"hostRegex"`
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method pulumi.StringOutput `pulumi:"method"`
	// Address name.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL path as a regular expression.
	Path pulumi.StringOutput `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query pulumi.StringOutput `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer pulumi.StringOutput `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyaddressTaggingArrayOutput `pulumi:"taggings"`
	// Proxy address type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua pulumi.StringOutput `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewFirewallProxyaddress registers a new resource with the given unique name, arguments, and options.
func NewFirewallProxyaddress(ctx *pulumi.Context,
	name string, args *FirewallProxyaddressArgs, opts ...pulumi.ResourceOption) (*FirewallProxyaddress, error) {
	if args == nil {
		args = &FirewallProxyaddressArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallProxyaddress
	err := ctx.RegisterResource("fortios:index/firewallProxyaddress:FirewallProxyaddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProxyaddress gets an existing FirewallProxyaddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProxyaddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProxyaddressState, opts ...pulumi.ResourceOption) (*FirewallProxyaddress, error) {
	var resource FirewallProxyaddress
	err := ctx.ReadResource("fortios:index/firewallProxyaddress:FirewallProxyaddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProxyaddress resources.
type firewallProxyaddressState struct {
	// SaaS application. The structure of `application` block is documented below.
	Applications []FirewallProxyaddressApplication `pulumi:"applications"`
	// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// FortiGuard category ID. The structure of `category` block is documented below.
	Categories []FirewallProxyaddressCategory `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// HTTP header name as a regular expression.
	Header *string `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []FirewallProxyaddressHeaderGroup `pulumi:"headerGroups"`
	// Name of HTTP header.
	HeaderName *string `pulumi:"headerName"`
	// Address object for the host.
	Host *string `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex *string `pulumi:"hostRegex"`
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method *string `pulumi:"method"`
	// Address name.
	Name *string `pulumi:"name"`
	// URL path as a regular expression.
	Path *string `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query *string `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer *string `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyaddressTagging `pulumi:"taggings"`
	// Proxy address type.
	Type *string `pulumi:"type"`
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua *string `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

type FirewallProxyaddressState struct {
	// SaaS application. The structure of `application` block is documented below.
	Applications FirewallProxyaddressApplicationArrayInput
	// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringPtrInput
	// FortiGuard category ID. The structure of `category` block is documented below.
	Categories FirewallProxyaddressCategoryArrayInput
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// HTTP header name as a regular expression.
	Header pulumi.StringPtrInput
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups FirewallProxyaddressHeaderGroupArrayInput
	// Name of HTTP header.
	HeaderName pulumi.StringPtrInput
	// Address object for the host.
	Host pulumi.StringPtrInput
	// Host name as a regular expression.
	HostRegex pulumi.StringPtrInput
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// URL path as a regular expression.
	Path pulumi.StringPtrInput
	// Match the query part of the URL as a regular expression.
	Query pulumi.StringPtrInput
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyaddressTaggingArrayInput
	// Proxy address type.
	Type pulumi.StringPtrInput
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyaddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyaddressState)(nil)).Elem()
}

type firewallProxyaddressArgs struct {
	// SaaS application. The structure of `application` block is documented below.
	Applications []FirewallProxyaddressApplication `pulumi:"applications"`
	// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
	CaseSensitivity *string `pulumi:"caseSensitivity"`
	// FortiGuard category ID. The structure of `category` block is documented below.
	Categories []FirewallProxyaddressCategory `pulumi:"categories"`
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color *int `pulumi:"color"`
	// Optional comments.
	Comment *string `pulumi:"comment"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// HTTP header name as a regular expression.
	Header *string `pulumi:"header"`
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups []FirewallProxyaddressHeaderGroup `pulumi:"headerGroups"`
	// Name of HTTP header.
	HeaderName *string `pulumi:"headerName"`
	// Address object for the host.
	Host *string `pulumi:"host"`
	// Host name as a regular expression.
	HostRegex *string `pulumi:"hostRegex"`
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method *string `pulumi:"method"`
	// Address name.
	Name *string `pulumi:"name"`
	// URL path as a regular expression.
	Path *string `pulumi:"path"`
	// Match the query part of the URL as a regular expression.
	Query *string `pulumi:"query"`
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer *string `pulumi:"referrer"`
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings []FirewallProxyaddressTagging `pulumi:"taggings"`
	// Proxy address type.
	Type *string `pulumi:"type"`
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua *string `pulumi:"ua"`
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid *string `pulumi:"uuid"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility *string `pulumi:"visibility"`
}

// The set of arguments for constructing a FirewallProxyaddress resource.
type FirewallProxyaddressArgs struct {
	// SaaS application. The structure of `application` block is documented below.
	Applications FirewallProxyaddressApplicationArrayInput
	// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
	CaseSensitivity pulumi.StringPtrInput
	// FortiGuard category ID. The structure of `category` block is documented below.
	Categories FirewallProxyaddressCategoryArrayInput
	// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
	Color pulumi.IntPtrInput
	// Optional comments.
	Comment pulumi.StringPtrInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// HTTP header name as a regular expression.
	Header pulumi.StringPtrInput
	// HTTP header group. The structure of `headerGroup` block is documented below.
	HeaderGroups FirewallProxyaddressHeaderGroupArrayInput
	// Name of HTTP header.
	HeaderName pulumi.StringPtrInput
	// Address object for the host.
	Host pulumi.StringPtrInput
	// Host name as a regular expression.
	HostRegex pulumi.StringPtrInput
	// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
	Method pulumi.StringPtrInput
	// Address name.
	Name pulumi.StringPtrInput
	// URL path as a regular expression.
	Path pulumi.StringPtrInput
	// Match the query part of the URL as a regular expression.
	Query pulumi.StringPtrInput
	// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
	Referrer pulumi.StringPtrInput
	// Config object tagging. The structure of `tagging` block is documented below.
	Taggings FirewallProxyaddressTaggingArrayInput
	// Proxy address type.
	Type pulumi.StringPtrInput
	// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
	Ua pulumi.StringPtrInput
	// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
	Uuid pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
	Visibility pulumi.StringPtrInput
}

func (FirewallProxyaddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProxyaddressArgs)(nil)).Elem()
}

type FirewallProxyaddressInput interface {
	pulumi.Input

	ToFirewallProxyaddressOutput() FirewallProxyaddressOutput
	ToFirewallProxyaddressOutputWithContext(ctx context.Context) FirewallProxyaddressOutput
}

func (*FirewallProxyaddress) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyaddress)(nil)).Elem()
}

func (i *FirewallProxyaddress) ToFirewallProxyaddressOutput() FirewallProxyaddressOutput {
	return i.ToFirewallProxyaddressOutputWithContext(context.Background())
}

func (i *FirewallProxyaddress) ToFirewallProxyaddressOutputWithContext(ctx context.Context) FirewallProxyaddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyaddressOutput)
}

// FirewallProxyaddressArrayInput is an input type that accepts FirewallProxyaddressArray and FirewallProxyaddressArrayOutput values.
// You can construct a concrete instance of `FirewallProxyaddressArrayInput` via:
//
//	FirewallProxyaddressArray{ FirewallProxyaddressArgs{...} }
type FirewallProxyaddressArrayInput interface {
	pulumi.Input

	ToFirewallProxyaddressArrayOutput() FirewallProxyaddressArrayOutput
	ToFirewallProxyaddressArrayOutputWithContext(context.Context) FirewallProxyaddressArrayOutput
}

type FirewallProxyaddressArray []FirewallProxyaddressInput

func (FirewallProxyaddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxyaddress)(nil)).Elem()
}

func (i FirewallProxyaddressArray) ToFirewallProxyaddressArrayOutput() FirewallProxyaddressArrayOutput {
	return i.ToFirewallProxyaddressArrayOutputWithContext(context.Background())
}

func (i FirewallProxyaddressArray) ToFirewallProxyaddressArrayOutputWithContext(ctx context.Context) FirewallProxyaddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyaddressArrayOutput)
}

// FirewallProxyaddressMapInput is an input type that accepts FirewallProxyaddressMap and FirewallProxyaddressMapOutput values.
// You can construct a concrete instance of `FirewallProxyaddressMapInput` via:
//
//	FirewallProxyaddressMap{ "key": FirewallProxyaddressArgs{...} }
type FirewallProxyaddressMapInput interface {
	pulumi.Input

	ToFirewallProxyaddressMapOutput() FirewallProxyaddressMapOutput
	ToFirewallProxyaddressMapOutputWithContext(context.Context) FirewallProxyaddressMapOutput
}

type FirewallProxyaddressMap map[string]FirewallProxyaddressInput

func (FirewallProxyaddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxyaddress)(nil)).Elem()
}

func (i FirewallProxyaddressMap) ToFirewallProxyaddressMapOutput() FirewallProxyaddressMapOutput {
	return i.ToFirewallProxyaddressMapOutputWithContext(context.Background())
}

func (i FirewallProxyaddressMap) ToFirewallProxyaddressMapOutputWithContext(ctx context.Context) FirewallProxyaddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProxyaddressMapOutput)
}

type FirewallProxyaddressOutput struct{ *pulumi.OutputState }

func (FirewallProxyaddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProxyaddress)(nil)).Elem()
}

func (o FirewallProxyaddressOutput) ToFirewallProxyaddressOutput() FirewallProxyaddressOutput {
	return o
}

func (o FirewallProxyaddressOutput) ToFirewallProxyaddressOutputWithContext(ctx context.Context) FirewallProxyaddressOutput {
	return o
}

// SaaS application. The structure of `application` block is documented below.
func (o FirewallProxyaddressOutput) Applications() FirewallProxyaddressApplicationArrayOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) FirewallProxyaddressApplicationArrayOutput { return v.Applications }).(FirewallProxyaddressApplicationArrayOutput)
}

// Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.
func (o FirewallProxyaddressOutput) CaseSensitivity() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.CaseSensitivity }).(pulumi.StringOutput)
}

// FortiGuard category ID. The structure of `category` block is documented below.
func (o FirewallProxyaddressOutput) Categories() FirewallProxyaddressCategoryArrayOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) FirewallProxyaddressCategoryArrayOutput { return v.Categories }).(FirewallProxyaddressCategoryArrayOutput)
}

// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
func (o FirewallProxyaddressOutput) Color() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.IntOutput { return v.Color }).(pulumi.IntOutput)
}

// Optional comments.
func (o FirewallProxyaddressOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringPtrOutput { return v.Comment }).(pulumi.StringPtrOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o FirewallProxyaddressOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// HTTP header name as a regular expression.
func (o FirewallProxyaddressOutput) Header() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Header }).(pulumi.StringOutput)
}

// HTTP header group. The structure of `headerGroup` block is documented below.
func (o FirewallProxyaddressOutput) HeaderGroups() FirewallProxyaddressHeaderGroupArrayOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) FirewallProxyaddressHeaderGroupArrayOutput { return v.HeaderGroups }).(FirewallProxyaddressHeaderGroupArrayOutput)
}

// Name of HTTP header.
func (o FirewallProxyaddressOutput) HeaderName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.HeaderName }).(pulumi.StringOutput)
}

// Address object for the host.
func (o FirewallProxyaddressOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Host name as a regular expression.
func (o FirewallProxyaddressOutput) HostRegex() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.HostRegex }).(pulumi.StringOutput)
}

// HTTP request methods to be used. Valid values: `get`, `post`, `put`, `head`, `connect`, `trace`, `options`, `delete`.
func (o FirewallProxyaddressOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Method }).(pulumi.StringOutput)
}

// Address name.
func (o FirewallProxyaddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// URL path as a regular expression.
func (o FirewallProxyaddressOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Match the query part of the URL as a regular expression.
func (o FirewallProxyaddressOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `enable`, `disable`.
func (o FirewallProxyaddressOutput) Referrer() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Referrer }).(pulumi.StringOutput)
}

// Config object tagging. The structure of `tagging` block is documented below.
func (o FirewallProxyaddressOutput) Taggings() FirewallProxyaddressTaggingArrayOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) FirewallProxyaddressTaggingArrayOutput { return v.Taggings }).(FirewallProxyaddressTaggingArrayOutput)
}

// Proxy address type.
func (o FirewallProxyaddressOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`.
func (o FirewallProxyaddressOutput) Ua() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Ua }).(pulumi.StringOutput)
}

// Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
func (o FirewallProxyaddressOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallProxyaddressOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable visibility of the object in the GUI. Valid values: `enable`, `disable`.
func (o FirewallProxyaddressOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProxyaddress) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type FirewallProxyaddressArrayOutput struct{ *pulumi.OutputState }

func (FirewallProxyaddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProxyaddress)(nil)).Elem()
}

func (o FirewallProxyaddressArrayOutput) ToFirewallProxyaddressArrayOutput() FirewallProxyaddressArrayOutput {
	return o
}

func (o FirewallProxyaddressArrayOutput) ToFirewallProxyaddressArrayOutputWithContext(ctx context.Context) FirewallProxyaddressArrayOutput {
	return o
}

func (o FirewallProxyaddressArrayOutput) Index(i pulumi.IntInput) FirewallProxyaddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallProxyaddress {
		return vs[0].([]*FirewallProxyaddress)[vs[1].(int)]
	}).(FirewallProxyaddressOutput)
}

type FirewallProxyaddressMapOutput struct{ *pulumi.OutputState }

func (FirewallProxyaddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProxyaddress)(nil)).Elem()
}

func (o FirewallProxyaddressMapOutput) ToFirewallProxyaddressMapOutput() FirewallProxyaddressMapOutput {
	return o
}

func (o FirewallProxyaddressMapOutput) ToFirewallProxyaddressMapOutputWithContext(ctx context.Context) FirewallProxyaddressMapOutput {
	return o
}

func (o FirewallProxyaddressMapOutput) MapIndex(k pulumi.StringInput) FirewallProxyaddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallProxyaddress {
		return vs[0].(map[string]*FirewallProxyaddress)[vs[1].(string)]
	}).(FirewallProxyaddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyaddressInput)(nil)).Elem(), &FirewallProxyaddress{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyaddressArrayInput)(nil)).Elem(), FirewallProxyaddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProxyaddressMapInput)(nil)).Elem(), FirewallProxyaddressMap{})
	pulumi.RegisterOutputType(FirewallProxyaddressOutput{})
	pulumi.RegisterOutputType(FirewallProxyaddressArrayOutput{})
	pulumi.RegisterOutputType(FirewallProxyaddressMapOutput{})
}
