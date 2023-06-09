// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure profile groups.
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
//			_, err := fortios.NewFirewallProfilegroup(ctx, "trname", &fortios.FirewallProfilegroupArgs{
//				ProfileProtocolOptions: pulumi.String("default"),
//				SslSshProfile:          pulumi.String("deep-inspection"),
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
// # Firewall ProfileGroup can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/firewallProfilegroup:FirewallProfilegroup labelname {{name}}
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/firewallProfilegroup:FirewallProfilegroup labelname {{name}}
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type FirewallProfilegroup struct {
	pulumi.CustomResourceState

	// Name of an existing Application list.
	ApplicationList pulumi.StringOutput `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile pulumi.StringOutput `pulumi:"avProfile"`
	// Name of an existing CIFS profile.
	CifsProfile pulumi.StringOutput `pulumi:"cifsProfile"`
	// Name of an existing DLP profile.
	DlpProfile pulumi.StringOutput `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringOutput `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile pulumi.StringOutput `pulumi:"dnsfilterProfile"`
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringOutput `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringOutput `pulumi:"fileFilterProfile"`
	// Name of an existing ICAP profile.
	IcapProfile pulumi.StringOutput `pulumi:"icapProfile"`
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringOutput `pulumi:"ipsSensor"`
	// Profile group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions pulumi.StringOutput `pulumi:"profileProtocolOptions"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile pulumi.StringOutput `pulumi:"sctpFilterProfile"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile pulumi.StringOutput `pulumi:"spamfilterProfile"`
	// Name of an existing SSH filter profile.
	SshFilterProfile pulumi.StringOutput `pulumi:"sshFilterProfile"`
	// Name of an existing SSL SSH profile.
	SslSshProfile pulumi.StringOutput `pulumi:"sslSshProfile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile pulumi.StringOutput `pulumi:"videofilterProfile"`
	// Name of an existing VoIP profile.
	VoipProfile pulumi.StringOutput `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile pulumi.StringOutput `pulumi:"wafProfile"`
	// Name of an existing Web filter profile.
	WebfilterProfile pulumi.StringOutput `pulumi:"webfilterProfile"`
}

// NewFirewallProfilegroup registers a new resource with the given unique name, arguments, and options.
func NewFirewallProfilegroup(ctx *pulumi.Context,
	name string, args *FirewallProfilegroupArgs, opts ...pulumi.ResourceOption) (*FirewallProfilegroup, error) {
	if args == nil {
		args = &FirewallProfilegroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource FirewallProfilegroup
	err := ctx.RegisterResource("fortios:index/firewallProfilegroup:FirewallProfilegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallProfilegroup gets an existing FirewallProfilegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallProfilegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallProfilegroupState, opts ...pulumi.ResourceOption) (*FirewallProfilegroup, error) {
	var resource FirewallProfilegroup
	err := ctx.ReadResource("fortios:index/firewallProfilegroup:FirewallProfilegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallProfilegroup resources.
type firewallProfilegroupState struct {
	// Name of an existing Application list.
	ApplicationList *string `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile *string `pulumi:"avProfile"`
	// Name of an existing CIFS profile.
	CifsProfile *string `pulumi:"cifsProfile"`
	// Name of an existing DLP profile.
	DlpProfile *string `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor *string `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile *string `pulumi:"dnsfilterProfile"`
	// Name of an existing email filter profile.
	EmailfilterProfile *string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile *string `pulumi:"fileFilterProfile"`
	// Name of an existing ICAP profile.
	IcapProfile *string `pulumi:"icapProfile"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Profile group name.
	Name *string `pulumi:"name"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions *string `pulumi:"profileProtocolOptions"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile *string `pulumi:"sctpFilterProfile"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile *string `pulumi:"spamfilterProfile"`
	// Name of an existing SSH filter profile.
	SshFilterProfile *string `pulumi:"sshFilterProfile"`
	// Name of an existing SSL SSH profile.
	SslSshProfile *string `pulumi:"sslSshProfile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile *string `pulumi:"videofilterProfile"`
	// Name of an existing VoIP profile.
	VoipProfile *string `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile *string `pulumi:"wafProfile"`
	// Name of an existing Web filter profile.
	WebfilterProfile *string `pulumi:"webfilterProfile"`
}

type FirewallProfilegroupState struct {
	// Name of an existing Application list.
	ApplicationList pulumi.StringPtrInput
	// Name of an existing Antivirus profile.
	AvProfile pulumi.StringPtrInput
	// Name of an existing CIFS profile.
	CifsProfile pulumi.StringPtrInput
	// Name of an existing DLP profile.
	DlpProfile pulumi.StringPtrInput
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringPtrInput
	// Name of an existing DNS filter profile.
	DnsfilterProfile pulumi.StringPtrInput
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringPtrInput
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringPtrInput
	// Name of an existing ICAP profile.
	IcapProfile pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Profile group name.
	Name pulumi.StringPtrInput
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions pulumi.StringPtrInput
	// Name of an existing SCTP filter profile.
	SctpFilterProfile pulumi.StringPtrInput
	// Name of an existing Spam filter profile.
	SpamfilterProfile pulumi.StringPtrInput
	// Name of an existing SSH filter profile.
	SshFilterProfile pulumi.StringPtrInput
	// Name of an existing SSL SSH profile.
	SslSshProfile pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Name of an existing VideoFilter profile.
	VideofilterProfile pulumi.StringPtrInput
	// Name of an existing VoIP profile.
	VoipProfile pulumi.StringPtrInput
	// Name of an existing Web application firewall profile.
	WafProfile pulumi.StringPtrInput
	// Name of an existing Web filter profile.
	WebfilterProfile pulumi.StringPtrInput
}

func (FirewallProfilegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProfilegroupState)(nil)).Elem()
}

type firewallProfilegroupArgs struct {
	// Name of an existing Application list.
	ApplicationList *string `pulumi:"applicationList"`
	// Name of an existing Antivirus profile.
	AvProfile *string `pulumi:"avProfile"`
	// Name of an existing CIFS profile.
	CifsProfile *string `pulumi:"cifsProfile"`
	// Name of an existing DLP profile.
	DlpProfile *string `pulumi:"dlpProfile"`
	// Name of an existing DLP sensor.
	DlpSensor *string `pulumi:"dlpSensor"`
	// Name of an existing DNS filter profile.
	DnsfilterProfile *string `pulumi:"dnsfilterProfile"`
	// Name of an existing email filter profile.
	EmailfilterProfile *string `pulumi:"emailfilterProfile"`
	// Name of an existing file-filter profile.
	FileFilterProfile *string `pulumi:"fileFilterProfile"`
	// Name of an existing ICAP profile.
	IcapProfile *string `pulumi:"icapProfile"`
	// Name of an existing IPS sensor.
	IpsSensor *string `pulumi:"ipsSensor"`
	// Profile group name.
	Name *string `pulumi:"name"`
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions *string `pulumi:"profileProtocolOptions"`
	// Name of an existing SCTP filter profile.
	SctpFilterProfile *string `pulumi:"sctpFilterProfile"`
	// Name of an existing Spam filter profile.
	SpamfilterProfile *string `pulumi:"spamfilterProfile"`
	// Name of an existing SSH filter profile.
	SshFilterProfile *string `pulumi:"sshFilterProfile"`
	// Name of an existing SSL SSH profile.
	SslSshProfile *string `pulumi:"sslSshProfile"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
	// Name of an existing VideoFilter profile.
	VideofilterProfile *string `pulumi:"videofilterProfile"`
	// Name of an existing VoIP profile.
	VoipProfile *string `pulumi:"voipProfile"`
	// Name of an existing Web application firewall profile.
	WafProfile *string `pulumi:"wafProfile"`
	// Name of an existing Web filter profile.
	WebfilterProfile *string `pulumi:"webfilterProfile"`
}

// The set of arguments for constructing a FirewallProfilegroup resource.
type FirewallProfilegroupArgs struct {
	// Name of an existing Application list.
	ApplicationList pulumi.StringPtrInput
	// Name of an existing Antivirus profile.
	AvProfile pulumi.StringPtrInput
	// Name of an existing CIFS profile.
	CifsProfile pulumi.StringPtrInput
	// Name of an existing DLP profile.
	DlpProfile pulumi.StringPtrInput
	// Name of an existing DLP sensor.
	DlpSensor pulumi.StringPtrInput
	// Name of an existing DNS filter profile.
	DnsfilterProfile pulumi.StringPtrInput
	// Name of an existing email filter profile.
	EmailfilterProfile pulumi.StringPtrInput
	// Name of an existing file-filter profile.
	FileFilterProfile pulumi.StringPtrInput
	// Name of an existing ICAP profile.
	IcapProfile pulumi.StringPtrInput
	// Name of an existing IPS sensor.
	IpsSensor pulumi.StringPtrInput
	// Profile group name.
	Name pulumi.StringPtrInput
	// Name of an existing Protocol options profile.
	ProfileProtocolOptions pulumi.StringPtrInput
	// Name of an existing SCTP filter profile.
	SctpFilterProfile pulumi.StringPtrInput
	// Name of an existing Spam filter profile.
	SpamfilterProfile pulumi.StringPtrInput
	// Name of an existing SSH filter profile.
	SshFilterProfile pulumi.StringPtrInput
	// Name of an existing SSL SSH profile.
	SslSshProfile pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
	// Name of an existing VideoFilter profile.
	VideofilterProfile pulumi.StringPtrInput
	// Name of an existing VoIP profile.
	VoipProfile pulumi.StringPtrInput
	// Name of an existing Web application firewall profile.
	WafProfile pulumi.StringPtrInput
	// Name of an existing Web filter profile.
	WebfilterProfile pulumi.StringPtrInput
}

func (FirewallProfilegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallProfilegroupArgs)(nil)).Elem()
}

type FirewallProfilegroupInput interface {
	pulumi.Input

	ToFirewallProfilegroupOutput() FirewallProfilegroupOutput
	ToFirewallProfilegroupOutputWithContext(ctx context.Context) FirewallProfilegroupOutput
}

func (*FirewallProfilegroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProfilegroup)(nil)).Elem()
}

func (i *FirewallProfilegroup) ToFirewallProfilegroupOutput() FirewallProfilegroupOutput {
	return i.ToFirewallProfilegroupOutputWithContext(context.Background())
}

func (i *FirewallProfilegroup) ToFirewallProfilegroupOutputWithContext(ctx context.Context) FirewallProfilegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProfilegroupOutput)
}

// FirewallProfilegroupArrayInput is an input type that accepts FirewallProfilegroupArray and FirewallProfilegroupArrayOutput values.
// You can construct a concrete instance of `FirewallProfilegroupArrayInput` via:
//
//	FirewallProfilegroupArray{ FirewallProfilegroupArgs{...} }
type FirewallProfilegroupArrayInput interface {
	pulumi.Input

	ToFirewallProfilegroupArrayOutput() FirewallProfilegroupArrayOutput
	ToFirewallProfilegroupArrayOutputWithContext(context.Context) FirewallProfilegroupArrayOutput
}

type FirewallProfilegroupArray []FirewallProfilegroupInput

func (FirewallProfilegroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProfilegroup)(nil)).Elem()
}

func (i FirewallProfilegroupArray) ToFirewallProfilegroupArrayOutput() FirewallProfilegroupArrayOutput {
	return i.ToFirewallProfilegroupArrayOutputWithContext(context.Background())
}

func (i FirewallProfilegroupArray) ToFirewallProfilegroupArrayOutputWithContext(ctx context.Context) FirewallProfilegroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProfilegroupArrayOutput)
}

// FirewallProfilegroupMapInput is an input type that accepts FirewallProfilegroupMap and FirewallProfilegroupMapOutput values.
// You can construct a concrete instance of `FirewallProfilegroupMapInput` via:
//
//	FirewallProfilegroupMap{ "key": FirewallProfilegroupArgs{...} }
type FirewallProfilegroupMapInput interface {
	pulumi.Input

	ToFirewallProfilegroupMapOutput() FirewallProfilegroupMapOutput
	ToFirewallProfilegroupMapOutputWithContext(context.Context) FirewallProfilegroupMapOutput
}

type FirewallProfilegroupMap map[string]FirewallProfilegroupInput

func (FirewallProfilegroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProfilegroup)(nil)).Elem()
}

func (i FirewallProfilegroupMap) ToFirewallProfilegroupMapOutput() FirewallProfilegroupMapOutput {
	return i.ToFirewallProfilegroupMapOutputWithContext(context.Background())
}

func (i FirewallProfilegroupMap) ToFirewallProfilegroupMapOutputWithContext(ctx context.Context) FirewallProfilegroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallProfilegroupMapOutput)
}

type FirewallProfilegroupOutput struct{ *pulumi.OutputState }

func (FirewallProfilegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallProfilegroup)(nil)).Elem()
}

func (o FirewallProfilegroupOutput) ToFirewallProfilegroupOutput() FirewallProfilegroupOutput {
	return o
}

func (o FirewallProfilegroupOutput) ToFirewallProfilegroupOutputWithContext(ctx context.Context) FirewallProfilegroupOutput {
	return o
}

// Name of an existing Application list.
func (o FirewallProfilegroupOutput) ApplicationList() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.ApplicationList }).(pulumi.StringOutput)
}

// Name of an existing Antivirus profile.
func (o FirewallProfilegroupOutput) AvProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.AvProfile }).(pulumi.StringOutput)
}

// Name of an existing CIFS profile.
func (o FirewallProfilegroupOutput) CifsProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.CifsProfile }).(pulumi.StringOutput)
}

// Name of an existing DLP profile.
func (o FirewallProfilegroupOutput) DlpProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.DlpProfile }).(pulumi.StringOutput)
}

// Name of an existing DLP sensor.
func (o FirewallProfilegroupOutput) DlpSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.DlpSensor }).(pulumi.StringOutput)
}

// Name of an existing DNS filter profile.
func (o FirewallProfilegroupOutput) DnsfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.DnsfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing email filter profile.
func (o FirewallProfilegroupOutput) EmailfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.EmailfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing file-filter profile.
func (o FirewallProfilegroupOutput) FileFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.FileFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing ICAP profile.
func (o FirewallProfilegroupOutput) IcapProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.IcapProfile }).(pulumi.StringOutput)
}

// Name of an existing IPS sensor.
func (o FirewallProfilegroupOutput) IpsSensor() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.IpsSensor }).(pulumi.StringOutput)
}

// Profile group name.
func (o FirewallProfilegroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name of an existing Protocol options profile.
func (o FirewallProfilegroupOutput) ProfileProtocolOptions() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.ProfileProtocolOptions }).(pulumi.StringOutput)
}

// Name of an existing SCTP filter profile.
func (o FirewallProfilegroupOutput) SctpFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.SctpFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing Spam filter profile.
func (o FirewallProfilegroupOutput) SpamfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.SpamfilterProfile }).(pulumi.StringOutput)
}

// Name of an existing SSH filter profile.
func (o FirewallProfilegroupOutput) SshFilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.SshFilterProfile }).(pulumi.StringOutput)
}

// Name of an existing SSL SSH profile.
func (o FirewallProfilegroupOutput) SslSshProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.SslSshProfile }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o FirewallProfilegroupOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

// Name of an existing VideoFilter profile.
func (o FirewallProfilegroupOutput) VideofilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.VideofilterProfile }).(pulumi.StringOutput)
}

// Name of an existing VoIP profile.
func (o FirewallProfilegroupOutput) VoipProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.VoipProfile }).(pulumi.StringOutput)
}

// Name of an existing Web application firewall profile.
func (o FirewallProfilegroupOutput) WafProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.WafProfile }).(pulumi.StringOutput)
}

// Name of an existing Web filter profile.
func (o FirewallProfilegroupOutput) WebfilterProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallProfilegroup) pulumi.StringOutput { return v.WebfilterProfile }).(pulumi.StringOutput)
}

type FirewallProfilegroupArrayOutput struct{ *pulumi.OutputState }

func (FirewallProfilegroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallProfilegroup)(nil)).Elem()
}

func (o FirewallProfilegroupArrayOutput) ToFirewallProfilegroupArrayOutput() FirewallProfilegroupArrayOutput {
	return o
}

func (o FirewallProfilegroupArrayOutput) ToFirewallProfilegroupArrayOutputWithContext(ctx context.Context) FirewallProfilegroupArrayOutput {
	return o
}

func (o FirewallProfilegroupArrayOutput) Index(i pulumi.IntInput) FirewallProfilegroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallProfilegroup {
		return vs[0].([]*FirewallProfilegroup)[vs[1].(int)]
	}).(FirewallProfilegroupOutput)
}

type FirewallProfilegroupMapOutput struct{ *pulumi.OutputState }

func (FirewallProfilegroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallProfilegroup)(nil)).Elem()
}

func (o FirewallProfilegroupMapOutput) ToFirewallProfilegroupMapOutput() FirewallProfilegroupMapOutput {
	return o
}

func (o FirewallProfilegroupMapOutput) ToFirewallProfilegroupMapOutputWithContext(ctx context.Context) FirewallProfilegroupMapOutput {
	return o
}

func (o FirewallProfilegroupMapOutput) MapIndex(k pulumi.StringInput) FirewallProfilegroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallProfilegroup {
		return vs[0].(map[string]*FirewallProfilegroup)[vs[1].(string)]
	}).(FirewallProfilegroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProfilegroupInput)(nil)).Elem(), &FirewallProfilegroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProfilegroupArrayInput)(nil)).Elem(), FirewallProfilegroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallProfilegroupMapInput)(nil)).Elem(), FirewallProfilegroupMap{})
	pulumi.RegisterOutputType(FirewallProfilegroupOutput{})
	pulumi.RegisterOutputType(FirewallProfilegroupArrayOutput{})
	pulumi.RegisterOutputType(FirewallProfilegroupMapOutput{})
}
