// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// OCSP server configuration.
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
//			_, err := fortios.NewVpncertificateOcspserver(ctx, "trname", &fortios.VpncertificateOcspserverArgs{
//				Cert:          pulumi.String("ACCVRAIZ1"),
//				SourceIp:      pulumi.String("0.0.0.0"),
//				UnavailAction: pulumi.String("revoke"),
//				Url:           pulumi.String("www.tetserv.com"),
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
// # VpnCertificate OcspServer can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/vpncertificateOcspserver:VpncertificateOcspserver labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/vpncertificateOcspserver:VpncertificateOcspserver labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type VpncertificateOcspserver struct {
	pulumi.CustomResourceState

	// OCSP server certificate.
	Cert pulumi.StringOutput `pulumi:"cert"`
	// OCSP server entry name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringOutput `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringOutput `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringOutput `pulumi:"unavailAction"`
	// OCSP server URL.
	Url pulumi.StringOutput `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewVpncertificateOcspserver registers a new resource with the given unique name, arguments, and options.
func NewVpncertificateOcspserver(ctx *pulumi.Context,
	name string, args *VpncertificateOcspserverArgs, opts ...pulumi.ResourceOption) (*VpncertificateOcspserver, error) {
	if args == nil {
		args = &VpncertificateOcspserverArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpncertificateOcspserver
	err := ctx.RegisterResource("fortios:index/vpncertificateOcspserver:VpncertificateOcspserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpncertificateOcspserver gets an existing VpncertificateOcspserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpncertificateOcspserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpncertificateOcspserverState, opts ...pulumi.ResourceOption) (*VpncertificateOcspserver, error) {
	var resource VpncertificateOcspserver
	err := ctx.ReadResource("fortios:index/vpncertificateOcspserver:VpncertificateOcspserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpncertificateOcspserver resources.
type vpncertificateOcspserverState struct {
	// OCSP server certificate.
	Cert *string `pulumi:"cert"`
	// OCSP server entry name.
	Name *string `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert *string `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl *string `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction *string `pulumi:"unavailAction"`
	// OCSP server URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type VpncertificateOcspserverState struct {
	// OCSP server certificate.
	Cert pulumi.StringPtrInput
	// OCSP server entry name.
	Name pulumi.StringPtrInput
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringPtrInput
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringPtrInput
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringPtrInput
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringPtrInput
	// OCSP server URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpncertificateOcspserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpncertificateOcspserverState)(nil)).Elem()
}

type vpncertificateOcspserverArgs struct {
	// OCSP server certificate.
	Cert *string `pulumi:"cert"`
	// OCSP server entry name.
	Name *string `pulumi:"name"`
	// Secondary OCSP server certificate.
	SecondaryCert *string `pulumi:"secondaryCert"`
	// Secondary OCSP server URL.
	SecondaryUrl *string `pulumi:"secondaryUrl"`
	// Source IP address for communications to the OCSP server.
	SourceIp *string `pulumi:"sourceIp"`
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction *string `pulumi:"unavailAction"`
	// OCSP server URL.
	Url *string `pulumi:"url"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a VpncertificateOcspserver resource.
type VpncertificateOcspserverArgs struct {
	// OCSP server certificate.
	Cert pulumi.StringPtrInput
	// OCSP server entry name.
	Name pulumi.StringPtrInput
	// Secondary OCSP server certificate.
	SecondaryCert pulumi.StringPtrInput
	// Secondary OCSP server URL.
	SecondaryUrl pulumi.StringPtrInput
	// Source IP address for communications to the OCSP server.
	SourceIp pulumi.StringPtrInput
	// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
	UnavailAction pulumi.StringPtrInput
	// OCSP server URL.
	Url pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (VpncertificateOcspserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpncertificateOcspserverArgs)(nil)).Elem()
}

type VpncertificateOcspserverInput interface {
	pulumi.Input

	ToVpncertificateOcspserverOutput() VpncertificateOcspserverOutput
	ToVpncertificateOcspserverOutputWithContext(ctx context.Context) VpncertificateOcspserverOutput
}

func (*VpncertificateOcspserver) ElementType() reflect.Type {
	return reflect.TypeOf((**VpncertificateOcspserver)(nil)).Elem()
}

func (i *VpncertificateOcspserver) ToVpncertificateOcspserverOutput() VpncertificateOcspserverOutput {
	return i.ToVpncertificateOcspserverOutputWithContext(context.Background())
}

func (i *VpncertificateOcspserver) ToVpncertificateOcspserverOutputWithContext(ctx context.Context) VpncertificateOcspserverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateOcspserverOutput)
}

// VpncertificateOcspserverArrayInput is an input type that accepts VpncertificateOcspserverArray and VpncertificateOcspserverArrayOutput values.
// You can construct a concrete instance of `VpncertificateOcspserverArrayInput` via:
//
//	VpncertificateOcspserverArray{ VpncertificateOcspserverArgs{...} }
type VpncertificateOcspserverArrayInput interface {
	pulumi.Input

	ToVpncertificateOcspserverArrayOutput() VpncertificateOcspserverArrayOutput
	ToVpncertificateOcspserverArrayOutputWithContext(context.Context) VpncertificateOcspserverArrayOutput
}

type VpncertificateOcspserverArray []VpncertificateOcspserverInput

func (VpncertificateOcspserverArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpncertificateOcspserver)(nil)).Elem()
}

func (i VpncertificateOcspserverArray) ToVpncertificateOcspserverArrayOutput() VpncertificateOcspserverArrayOutput {
	return i.ToVpncertificateOcspserverArrayOutputWithContext(context.Background())
}

func (i VpncertificateOcspserverArray) ToVpncertificateOcspserverArrayOutputWithContext(ctx context.Context) VpncertificateOcspserverArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateOcspserverArrayOutput)
}

// VpncertificateOcspserverMapInput is an input type that accepts VpncertificateOcspserverMap and VpncertificateOcspserverMapOutput values.
// You can construct a concrete instance of `VpncertificateOcspserverMapInput` via:
//
//	VpncertificateOcspserverMap{ "key": VpncertificateOcspserverArgs{...} }
type VpncertificateOcspserverMapInput interface {
	pulumi.Input

	ToVpncertificateOcspserverMapOutput() VpncertificateOcspserverMapOutput
	ToVpncertificateOcspserverMapOutputWithContext(context.Context) VpncertificateOcspserverMapOutput
}

type VpncertificateOcspserverMap map[string]VpncertificateOcspserverInput

func (VpncertificateOcspserverMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpncertificateOcspserver)(nil)).Elem()
}

func (i VpncertificateOcspserverMap) ToVpncertificateOcspserverMapOutput() VpncertificateOcspserverMapOutput {
	return i.ToVpncertificateOcspserverMapOutputWithContext(context.Background())
}

func (i VpncertificateOcspserverMap) ToVpncertificateOcspserverMapOutputWithContext(ctx context.Context) VpncertificateOcspserverMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpncertificateOcspserverMapOutput)
}

type VpncertificateOcspserverOutput struct{ *pulumi.OutputState }

func (VpncertificateOcspserverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpncertificateOcspserver)(nil)).Elem()
}

func (o VpncertificateOcspserverOutput) ToVpncertificateOcspserverOutput() VpncertificateOcspserverOutput {
	return o
}

func (o VpncertificateOcspserverOutput) ToVpncertificateOcspserverOutputWithContext(ctx context.Context) VpncertificateOcspserverOutput {
	return o
}

// OCSP server certificate.
func (o VpncertificateOcspserverOutput) Cert() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.Cert }).(pulumi.StringOutput)
}

// OCSP server entry name.
func (o VpncertificateOcspserverOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Secondary OCSP server certificate.
func (o VpncertificateOcspserverOutput) SecondaryCert() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.SecondaryCert }).(pulumi.StringOutput)
}

// Secondary OCSP server URL.
func (o VpncertificateOcspserverOutput) SecondaryUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.SecondaryUrl }).(pulumi.StringOutput)
}

// Source IP address for communications to the OCSP server.
func (o VpncertificateOcspserverOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// Action when server is unavailable (revoke the certificate or ignore the result of the check). Valid values: `revoke`, `ignore`.
func (o VpncertificateOcspserverOutput) UnavailAction() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.UnavailAction }).(pulumi.StringOutput)
}

// OCSP server URL.
func (o VpncertificateOcspserverOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o VpncertificateOcspserverOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpncertificateOcspserver) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type VpncertificateOcspserverArrayOutput struct{ *pulumi.OutputState }

func (VpncertificateOcspserverArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpncertificateOcspserver)(nil)).Elem()
}

func (o VpncertificateOcspserverArrayOutput) ToVpncertificateOcspserverArrayOutput() VpncertificateOcspserverArrayOutput {
	return o
}

func (o VpncertificateOcspserverArrayOutput) ToVpncertificateOcspserverArrayOutputWithContext(ctx context.Context) VpncertificateOcspserverArrayOutput {
	return o
}

func (o VpncertificateOcspserverArrayOutput) Index(i pulumi.IntInput) VpncertificateOcspserverOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpncertificateOcspserver {
		return vs[0].([]*VpncertificateOcspserver)[vs[1].(int)]
	}).(VpncertificateOcspserverOutput)
}

type VpncertificateOcspserverMapOutput struct{ *pulumi.OutputState }

func (VpncertificateOcspserverMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpncertificateOcspserver)(nil)).Elem()
}

func (o VpncertificateOcspserverMapOutput) ToVpncertificateOcspserverMapOutput() VpncertificateOcspserverMapOutput {
	return o
}

func (o VpncertificateOcspserverMapOutput) ToVpncertificateOcspserverMapOutputWithContext(ctx context.Context) VpncertificateOcspserverMapOutput {
	return o
}

func (o VpncertificateOcspserverMapOutput) MapIndex(k pulumi.StringInput) VpncertificateOcspserverOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpncertificateOcspserver {
		return vs[0].(map[string]*VpncertificateOcspserver)[vs[1].(string)]
	}).(VpncertificateOcspserverOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateOcspserverInput)(nil)).Elem(), &VpncertificateOcspserver{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateOcspserverArrayInput)(nil)).Elem(), VpncertificateOcspserverArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpncertificateOcspserverMapInput)(nil)).Elem(), VpncertificateOcspserverMap{})
	pulumi.RegisterOutputType(VpncertificateOcspserverOutput{})
	pulumi.RegisterOutputType(VpncertificateOcspserverArrayOutput{})
	pulumi.RegisterOutputType(VpncertificateOcspserverMapOutput{})
}
