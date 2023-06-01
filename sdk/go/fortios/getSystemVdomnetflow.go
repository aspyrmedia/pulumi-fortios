// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system vdomnetflow
func LookupSystemVdomnetflow(ctx *pulumi.Context, args *LookupSystemVdomnetflowArgs, opts ...pulumi.InvokeOption) (*LookupSystemVdomnetflowResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemVdomnetflowResult
	err := ctx.Invoke("fortios:index/getSystemVdomnetflow:getSystemVdomnetflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemVdomnetflow.
type LookupSystemVdomnetflowArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemVdomnetflow.
type LookupSystemVdomnetflowResult struct {
	// NetFlow collector IP address.
	CollectorIp string `pulumi:"collectorIp"`
	// NetFlow collector port number.
	CollectorPort int `pulumi:"collectorPort"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify outgoing interface to reach server.
	Interface string `pulumi:"interface"`
	// Specify how to select outgoing interface to reach server.
	InterfaceSelectMethod string `pulumi:"interfaceSelectMethod"`
	// Source IP address for communication with the NetFlow agent.
	SourceIp string `pulumi:"sourceIp"`
	// Enable/disable NetFlow per VDOM.
	VdomNetflow string  `pulumi:"vdomNetflow"`
	Vdomparam   *string `pulumi:"vdomparam"`
}

func LookupSystemVdomnetflowOutput(ctx *pulumi.Context, args LookupSystemVdomnetflowOutputArgs, opts ...pulumi.InvokeOption) LookupSystemVdomnetflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemVdomnetflowResult, error) {
			args := v.(LookupSystemVdomnetflowArgs)
			r, err := LookupSystemVdomnetflow(ctx, &args, opts...)
			var s LookupSystemVdomnetflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemVdomnetflowResultOutput)
}

// A collection of arguments for invoking getSystemVdomnetflow.
type LookupSystemVdomnetflowOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemVdomnetflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVdomnetflowArgs)(nil)).Elem()
}

// A collection of values returned by getSystemVdomnetflow.
type LookupSystemVdomnetflowResultOutput struct{ *pulumi.OutputState }

func (LookupSystemVdomnetflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemVdomnetflowResult)(nil)).Elem()
}

func (o LookupSystemVdomnetflowResultOutput) ToLookupSystemVdomnetflowResultOutput() LookupSystemVdomnetflowResultOutput {
	return o
}

func (o LookupSystemVdomnetflowResultOutput) ToLookupSystemVdomnetflowResultOutputWithContext(ctx context.Context) LookupSystemVdomnetflowResultOutput {
	return o
}

// NetFlow collector IP address.
func (o LookupSystemVdomnetflowResultOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) string { return v.CollectorIp }).(pulumi.StringOutput)
}

// NetFlow collector port number.
func (o LookupSystemVdomnetflowResultOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) int { return v.CollectorPort }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemVdomnetflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify outgoing interface to reach server.
func (o LookupSystemVdomnetflowResultOutput) Interface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) string { return v.Interface }).(pulumi.StringOutput)
}

// Specify how to select outgoing interface to reach server.
func (o LookupSystemVdomnetflowResultOutput) InterfaceSelectMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) string { return v.InterfaceSelectMethod }).(pulumi.StringOutput)
}

// Source IP address for communication with the NetFlow agent.
func (o LookupSystemVdomnetflowResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Enable/disable NetFlow per VDOM.
func (o LookupSystemVdomnetflowResultOutput) VdomNetflow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) string { return v.VdomNetflow }).(pulumi.StringOutput)
}

func (o LookupSystemVdomnetflowResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemVdomnetflowResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemVdomnetflowResultOutput{})
}