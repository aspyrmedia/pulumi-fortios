// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an fortios router prefixlist6
func LookupRouterPrefixlist6(ctx *pulumi.Context, args *LookupRouterPrefixlist6Args, opts ...pulumi.InvokeOption) (*LookupRouterPrefixlist6Result, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterPrefixlist6Result
	err := ctx.Invoke("fortios:index/getRouterPrefixlist6:getRouterPrefixlist6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterPrefixlist6.
type LookupRouterPrefixlist6Args struct {
	// Specify the name of the desired router prefixlist6.
	Name string `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterPrefixlist6.
type LookupRouterPrefixlist6Result struct {
	// Comment.
	Comments string `pulumi:"comments"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name.
	Name string `pulumi:"name"`
	// IPv6 prefix list rule. The structure of `rule` block is documented below.
	Rules     []GetRouterPrefixlist6Rule `pulumi:"rules"`
	Vdomparam *string                    `pulumi:"vdomparam"`
}

func LookupRouterPrefixlist6Output(ctx *pulumi.Context, args LookupRouterPrefixlist6OutputArgs, opts ...pulumi.InvokeOption) LookupRouterPrefixlist6ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterPrefixlist6Result, error) {
			args := v.(LookupRouterPrefixlist6Args)
			r, err := LookupRouterPrefixlist6(ctx, &args, opts...)
			var s LookupRouterPrefixlist6Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterPrefixlist6ResultOutput)
}

// A collection of arguments for invoking getRouterPrefixlist6.
type LookupRouterPrefixlist6OutputArgs struct {
	// Specify the name of the desired router prefixlist6.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterPrefixlist6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterPrefixlist6Args)(nil)).Elem()
}

// A collection of values returned by getRouterPrefixlist6.
type LookupRouterPrefixlist6ResultOutput struct{ *pulumi.OutputState }

func (LookupRouterPrefixlist6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterPrefixlist6Result)(nil)).Elem()
}

func (o LookupRouterPrefixlist6ResultOutput) ToLookupRouterPrefixlist6ResultOutput() LookupRouterPrefixlist6ResultOutput {
	return o
}

func (o LookupRouterPrefixlist6ResultOutput) ToLookupRouterPrefixlist6ResultOutputWithContext(ctx context.Context) LookupRouterPrefixlist6ResultOutput {
	return o
}

// Comment.
func (o LookupRouterPrefixlist6ResultOutput) Comments() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPrefixlist6Result) string { return v.Comments }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterPrefixlist6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPrefixlist6Result) string { return v.Id }).(pulumi.StringOutput)
}

// Name.
func (o LookupRouterPrefixlist6ResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterPrefixlist6Result) string { return v.Name }).(pulumi.StringOutput)
}

// IPv6 prefix list rule. The structure of `rule` block is documented below.
func (o LookupRouterPrefixlist6ResultOutput) Rules() GetRouterPrefixlist6RuleArrayOutput {
	return o.ApplyT(func(v LookupRouterPrefixlist6Result) []GetRouterPrefixlist6Rule { return v.Rules }).(GetRouterPrefixlist6RuleArrayOutput)
}

func (o LookupRouterPrefixlist6ResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterPrefixlist6Result) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterPrefixlist6ResultOutput{})
}
