// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure Authentication Rules.
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
//			_, err := fortios.NewAuthenticationRule(ctx, "trname", &fortios.AuthenticationRuleArgs{
//				IpBased:          pulumi.String("enable"),
//				Protocol:         pulumi.String("ftp"),
//				Status:           pulumi.String("enable"),
//				TransactionBased: pulumi.String("disable"),
//				WebAuthCookie:    pulumi.String("disable"),
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
// # Authentication Rule can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/authenticationRule:AuthenticationRule labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/authenticationRule:AuthenticationRule labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type AuthenticationRule struct {
	pulumi.CustomResourceState

	// Select an active authentication method.
	ActiveAuthMethod pulumi.StringOutput `pulumi:"activeAuthMethod"`
	// Comment.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
	Dstaddr6s AuthenticationRuleDstaddr6ArrayOutput `pulumi:"dstaddr6s"`
	// Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
	Dstaddrs AuthenticationRuleDstaddrArrayOutput `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
	IpBased pulumi.StringOutput `pulumi:"ipBased"`
	// Authentication rule name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
	Srcaddr6s AuthenticationRuleSrcaddr6ArrayOutput `pulumi:"srcaddr6s"`
	// Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
	Srcaddrs AuthenticationRuleSrcaddrArrayOutput `pulumi:"srcaddrs"`
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs AuthenticationRuleSrcintfArrayOutput `pulumi:"srcintfs"`
	// Select a single-sign on (SSO) authentication method.
	SsoAuthMethod pulumi.StringOutput `pulumi:"ssoAuthMethod"`
	// Enable/disable this authentication rule. Valid values: `enable`, `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
	TransactionBased pulumi.StringOutput `pulumi:"transactionBased"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
	WebAuthCookie pulumi.StringOutput `pulumi:"webAuthCookie"`
	// Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
	WebPortal pulumi.StringOutput `pulumi:"webPortal"`
}

// NewAuthenticationRule registers a new resource with the given unique name, arguments, and options.
func NewAuthenticationRule(ctx *pulumi.Context,
	name string, args *AuthenticationRuleArgs, opts ...pulumi.ResourceOption) (*AuthenticationRule, error) {
	if args == nil {
		args = &AuthenticationRuleArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AuthenticationRule
	err := ctx.RegisterResource("fortios:index/authenticationRule:AuthenticationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthenticationRule gets an existing AuthenticationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthenticationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthenticationRuleState, opts ...pulumi.ResourceOption) (*AuthenticationRule, error) {
	var resource AuthenticationRule
	err := ctx.ReadResource("fortios:index/authenticationRule:AuthenticationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthenticationRule resources.
type authenticationRuleState struct {
	// Select an active authentication method.
	ActiveAuthMethod *string `pulumi:"activeAuthMethod"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
	Dstaddr6s []AuthenticationRuleDstaddr6 `pulumi:"dstaddr6s"`
	// Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
	Dstaddrs []AuthenticationRuleDstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
	IpBased *string `pulumi:"ipBased"`
	// Authentication rule name.
	Name *string `pulumi:"name"`
	// Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.
	Protocol *string `pulumi:"protocol"`
	// Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
	Srcaddr6s []AuthenticationRuleSrcaddr6 `pulumi:"srcaddr6s"`
	// Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
	Srcaddrs []AuthenticationRuleSrcaddr `pulumi:"srcaddrs"`
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs []AuthenticationRuleSrcintf `pulumi:"srcintfs"`
	// Select a single-sign on (SSO) authentication method.
	SsoAuthMethod *string `pulumi:"ssoAuthMethod"`
	// Enable/disable this authentication rule. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
	TransactionBased *string `pulumi:"transactionBased"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
	WebAuthCookie *string `pulumi:"webAuthCookie"`
	// Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
	WebPortal *string `pulumi:"webPortal"`
}

type AuthenticationRuleState struct {
	// Select an active authentication method.
	ActiveAuthMethod pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
	Dstaddr6s AuthenticationRuleDstaddr6ArrayInput
	// Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
	Dstaddrs AuthenticationRuleDstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
	IpBased pulumi.StringPtrInput
	// Authentication rule name.
	Name pulumi.StringPtrInput
	// Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.
	Protocol pulumi.StringPtrInput
	// Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
	Srcaddr6s AuthenticationRuleSrcaddr6ArrayInput
	// Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
	Srcaddrs AuthenticationRuleSrcaddrArrayInput
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs AuthenticationRuleSrcintfArrayInput
	// Select a single-sign on (SSO) authentication method.
	SsoAuthMethod pulumi.StringPtrInput
	// Enable/disable this authentication rule. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
	TransactionBased pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
	WebAuthCookie pulumi.StringPtrInput
	// Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
	WebPortal pulumi.StringPtrInput
}

func (AuthenticationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationRuleState)(nil)).Elem()
}

type authenticationRuleArgs struct {
	// Select an active authentication method.
	ActiveAuthMethod *string `pulumi:"activeAuthMethod"`
	// Comment.
	Comments *string `pulumi:"comments"`
	// Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
	Dstaddr6s []AuthenticationRuleDstaddr6 `pulumi:"dstaddr6s"`
	// Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
	Dstaddrs []AuthenticationRuleDstaddr `pulumi:"dstaddrs"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
	IpBased *string `pulumi:"ipBased"`
	// Authentication rule name.
	Name *string `pulumi:"name"`
	// Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.
	Protocol *string `pulumi:"protocol"`
	// Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
	Srcaddr6s []AuthenticationRuleSrcaddr6 `pulumi:"srcaddr6s"`
	// Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
	Srcaddrs []AuthenticationRuleSrcaddr `pulumi:"srcaddrs"`
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs []AuthenticationRuleSrcintf `pulumi:"srcintfs"`
	// Select a single-sign on (SSO) authentication method.
	SsoAuthMethod *string `pulumi:"ssoAuthMethod"`
	// Enable/disable this authentication rule. Valid values: `enable`, `disable`.
	Status *string `pulumi:"status"`
	// Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
	TransactionBased *string `pulumi:"transactionBased"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
	WebAuthCookie *string `pulumi:"webAuthCookie"`
	// Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
	WebPortal *string `pulumi:"webPortal"`
}

// The set of arguments for constructing a AuthenticationRule resource.
type AuthenticationRuleArgs struct {
	// Select an active authentication method.
	ActiveAuthMethod pulumi.StringPtrInput
	// Comment.
	Comments pulumi.StringPtrInput
	// Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
	Dstaddr6s AuthenticationRuleDstaddr6ArrayInput
	// Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
	Dstaddrs AuthenticationRuleDstaddrArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
	IpBased pulumi.StringPtrInput
	// Authentication rule name.
	Name pulumi.StringPtrInput
	// Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.
	Protocol pulumi.StringPtrInput
	// Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
	Srcaddr6s AuthenticationRuleSrcaddr6ArrayInput
	// Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
	Srcaddrs AuthenticationRuleSrcaddrArrayInput
	// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
	Srcintfs AuthenticationRuleSrcintfArrayInput
	// Select a single-sign on (SSO) authentication method.
	SsoAuthMethod pulumi.StringPtrInput
	// Enable/disable this authentication rule. Valid values: `enable`, `disable`.
	Status pulumi.StringPtrInput
	// Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
	TransactionBased pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
	WebAuthCookie pulumi.StringPtrInput
	// Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
	WebPortal pulumi.StringPtrInput
}

func (AuthenticationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationRuleArgs)(nil)).Elem()
}

type AuthenticationRuleInput interface {
	pulumi.Input

	ToAuthenticationRuleOutput() AuthenticationRuleOutput
	ToAuthenticationRuleOutputWithContext(ctx context.Context) AuthenticationRuleOutput
}

func (*AuthenticationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationRule)(nil)).Elem()
}

func (i *AuthenticationRule) ToAuthenticationRuleOutput() AuthenticationRuleOutput {
	return i.ToAuthenticationRuleOutputWithContext(context.Background())
}

func (i *AuthenticationRule) ToAuthenticationRuleOutputWithContext(ctx context.Context) AuthenticationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationRuleOutput)
}

// AuthenticationRuleArrayInput is an input type that accepts AuthenticationRuleArray and AuthenticationRuleArrayOutput values.
// You can construct a concrete instance of `AuthenticationRuleArrayInput` via:
//
//	AuthenticationRuleArray{ AuthenticationRuleArgs{...} }
type AuthenticationRuleArrayInput interface {
	pulumi.Input

	ToAuthenticationRuleArrayOutput() AuthenticationRuleArrayOutput
	ToAuthenticationRuleArrayOutputWithContext(context.Context) AuthenticationRuleArrayOutput
}

type AuthenticationRuleArray []AuthenticationRuleInput

func (AuthenticationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationRule)(nil)).Elem()
}

func (i AuthenticationRuleArray) ToAuthenticationRuleArrayOutput() AuthenticationRuleArrayOutput {
	return i.ToAuthenticationRuleArrayOutputWithContext(context.Background())
}

func (i AuthenticationRuleArray) ToAuthenticationRuleArrayOutputWithContext(ctx context.Context) AuthenticationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationRuleArrayOutput)
}

// AuthenticationRuleMapInput is an input type that accepts AuthenticationRuleMap and AuthenticationRuleMapOutput values.
// You can construct a concrete instance of `AuthenticationRuleMapInput` via:
//
//	AuthenticationRuleMap{ "key": AuthenticationRuleArgs{...} }
type AuthenticationRuleMapInput interface {
	pulumi.Input

	ToAuthenticationRuleMapOutput() AuthenticationRuleMapOutput
	ToAuthenticationRuleMapOutputWithContext(context.Context) AuthenticationRuleMapOutput
}

type AuthenticationRuleMap map[string]AuthenticationRuleInput

func (AuthenticationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationRule)(nil)).Elem()
}

func (i AuthenticationRuleMap) ToAuthenticationRuleMapOutput() AuthenticationRuleMapOutput {
	return i.ToAuthenticationRuleMapOutputWithContext(context.Background())
}

func (i AuthenticationRuleMap) ToAuthenticationRuleMapOutputWithContext(ctx context.Context) AuthenticationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationRuleMapOutput)
}

type AuthenticationRuleOutput struct{ *pulumi.OutputState }

func (AuthenticationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationRule)(nil)).Elem()
}

func (o AuthenticationRuleOutput) ToAuthenticationRuleOutput() AuthenticationRuleOutput {
	return o
}

func (o AuthenticationRuleOutput) ToAuthenticationRuleOutputWithContext(ctx context.Context) AuthenticationRuleOutput {
	return o
}

// Select an active authentication method.
func (o AuthenticationRuleOutput) ActiveAuthMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.ActiveAuthMethod }).(pulumi.StringOutput)
}

// Comment.
func (o AuthenticationRuleOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
func (o AuthenticationRuleOutput) Dstaddr6s() AuthenticationRuleDstaddr6ArrayOutput {
	return o.ApplyT(func(v *AuthenticationRule) AuthenticationRuleDstaddr6ArrayOutput { return v.Dstaddr6s }).(AuthenticationRuleDstaddr6ArrayOutput)
}

// Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
func (o AuthenticationRuleOutput) Dstaddrs() AuthenticationRuleDstaddrArrayOutput {
	return o.ApplyT(func(v *AuthenticationRule) AuthenticationRuleDstaddrArrayOutput { return v.Dstaddrs }).(AuthenticationRuleDstaddrArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o AuthenticationRuleOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
func (o AuthenticationRuleOutput) IpBased() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.IpBased }).(pulumi.StringOutput)
}

// Authentication rule name.
func (o AuthenticationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate. Valid values: `http`, `ftp`, `socks`, `ssh`.
func (o AuthenticationRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
func (o AuthenticationRuleOutput) Srcaddr6s() AuthenticationRuleSrcaddr6ArrayOutput {
	return o.ApplyT(func(v *AuthenticationRule) AuthenticationRuleSrcaddr6ArrayOutput { return v.Srcaddr6s }).(AuthenticationRuleSrcaddr6ArrayOutput)
}

// Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
func (o AuthenticationRuleOutput) Srcaddrs() AuthenticationRuleSrcaddrArrayOutput {
	return o.ApplyT(func(v *AuthenticationRule) AuthenticationRuleSrcaddrArrayOutput { return v.Srcaddrs }).(AuthenticationRuleSrcaddrArrayOutput)
}

// Incoming (ingress) interface. The structure of `srcintf` block is documented below.
func (o AuthenticationRuleOutput) Srcintfs() AuthenticationRuleSrcintfArrayOutput {
	return o.ApplyT(func(v *AuthenticationRule) AuthenticationRuleSrcintfArrayOutput { return v.Srcintfs }).(AuthenticationRuleSrcintfArrayOutput)
}

// Select a single-sign on (SSO) authentication method.
func (o AuthenticationRuleOutput) SsoAuthMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.SsoAuthMethod }).(pulumi.StringOutput)
}

// Enable/disable this authentication rule. Valid values: `enable`, `disable`.
func (o AuthenticationRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
func (o AuthenticationRuleOutput) TransactionBased() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.TransactionBased }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o AuthenticationRuleOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
func (o AuthenticationRuleOutput) WebAuthCookie() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.WebAuthCookie }).(pulumi.StringOutput)
}

// Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
func (o AuthenticationRuleOutput) WebPortal() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.WebPortal }).(pulumi.StringOutput)
}

type AuthenticationRuleArrayOutput struct{ *pulumi.OutputState }

func (AuthenticationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationRule)(nil)).Elem()
}

func (o AuthenticationRuleArrayOutput) ToAuthenticationRuleArrayOutput() AuthenticationRuleArrayOutput {
	return o
}

func (o AuthenticationRuleArrayOutput) ToAuthenticationRuleArrayOutputWithContext(ctx context.Context) AuthenticationRuleArrayOutput {
	return o
}

func (o AuthenticationRuleArrayOutput) Index(i pulumi.IntInput) AuthenticationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthenticationRule {
		return vs[0].([]*AuthenticationRule)[vs[1].(int)]
	}).(AuthenticationRuleOutput)
}

type AuthenticationRuleMapOutput struct{ *pulumi.OutputState }

func (AuthenticationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationRule)(nil)).Elem()
}

func (o AuthenticationRuleMapOutput) ToAuthenticationRuleMapOutput() AuthenticationRuleMapOutput {
	return o
}

func (o AuthenticationRuleMapOutput) ToAuthenticationRuleMapOutputWithContext(ctx context.Context) AuthenticationRuleMapOutput {
	return o
}

func (o AuthenticationRuleMapOutput) MapIndex(k pulumi.StringInput) AuthenticationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthenticationRule {
		return vs[0].(map[string]*AuthenticationRule)[vs[1].(string)]
	}).(AuthenticationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationRuleInput)(nil)).Elem(), &AuthenticationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationRuleArrayInput)(nil)).Elem(), AuthenticationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationRuleMapInput)(nil)).Elem(), AuthenticationRuleMap{})
	pulumi.RegisterOutputType(AuthenticationRuleOutput{})
	pulumi.RegisterOutputType(AuthenticationRuleArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationRuleMapOutput{})
}
