// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios system managementtunnel
func LookupSystemManagementtunnel(ctx *pulumi.Context, args *LookupSystemManagementtunnelArgs, opts ...pulumi.InvokeOption) (*LookupSystemManagementtunnelResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSystemManagementtunnelResult
	err := ctx.Invoke("fortios:index/getSystemManagementtunnel:getSystemManagementtunnel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSystemManagementtunnel.
type LookupSystemManagementtunnelArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getSystemManagementtunnel.
type LookupSystemManagementtunnelResult struct {
	// Enable/disable collection of run time statistics.
	AllowCollectStatistics string `pulumi:"allowCollectStatistics"`
	// Enable/disable allow config restore.
	AllowConfigRestore string `pulumi:"allowConfigRestore"`
	// Enable/disable push configuration.
	AllowPushConfiguration string `pulumi:"allowPushConfiguration"`
	// Enable/disable push firmware.
	AllowPushFirmware string `pulumi:"allowPushFirmware"`
	// Enable/disable restriction of authorized manager only.
	AuthorizedManagerOnly string `pulumi:"authorizedManagerOnly"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Serial number.
	SerialNumber string `pulumi:"serialNumber"`
	// Enable/disable FGFM tunnel.
	Status    string  `pulumi:"status"`
	Vdomparam *string `pulumi:"vdomparam"`
}

func LookupSystemManagementtunnelOutput(ctx *pulumi.Context, args LookupSystemManagementtunnelOutputArgs, opts ...pulumi.InvokeOption) LookupSystemManagementtunnelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSystemManagementtunnelResult, error) {
			args := v.(LookupSystemManagementtunnelArgs)
			r, err := LookupSystemManagementtunnel(ctx, &args, opts...)
			var s LookupSystemManagementtunnelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSystemManagementtunnelResultOutput)
}

// A collection of arguments for invoking getSystemManagementtunnel.
type LookupSystemManagementtunnelOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupSystemManagementtunnelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemManagementtunnelArgs)(nil)).Elem()
}

// A collection of values returned by getSystemManagementtunnel.
type LookupSystemManagementtunnelResultOutput struct{ *pulumi.OutputState }

func (LookupSystemManagementtunnelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSystemManagementtunnelResult)(nil)).Elem()
}

func (o LookupSystemManagementtunnelResultOutput) ToLookupSystemManagementtunnelResultOutput() LookupSystemManagementtunnelResultOutput {
	return o
}

func (o LookupSystemManagementtunnelResultOutput) ToLookupSystemManagementtunnelResultOutputWithContext(ctx context.Context) LookupSystemManagementtunnelResultOutput {
	return o
}

// Enable/disable collection of run time statistics.
func (o LookupSystemManagementtunnelResultOutput) AllowCollectStatistics() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.AllowCollectStatistics }).(pulumi.StringOutput)
}

// Enable/disable allow config restore.
func (o LookupSystemManagementtunnelResultOutput) AllowConfigRestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.AllowConfigRestore }).(pulumi.StringOutput)
}

// Enable/disable push configuration.
func (o LookupSystemManagementtunnelResultOutput) AllowPushConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.AllowPushConfiguration }).(pulumi.StringOutput)
}

// Enable/disable push firmware.
func (o LookupSystemManagementtunnelResultOutput) AllowPushFirmware() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.AllowPushFirmware }).(pulumi.StringOutput)
}

// Enable/disable restriction of authorized manager only.
func (o LookupSystemManagementtunnelResultOutput) AuthorizedManagerOnly() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.AuthorizedManagerOnly }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSystemManagementtunnelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.Id }).(pulumi.StringOutput)
}

// Serial number.
func (o LookupSystemManagementtunnelResultOutput) SerialNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.SerialNumber }).(pulumi.StringOutput)
}

// Enable/disable FGFM tunnel.
func (o LookupSystemManagementtunnelResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSystemManagementtunnelResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSystemManagementtunnelResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSystemManagementtunnelResultOutput{})
}
