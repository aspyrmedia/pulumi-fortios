// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configure FortiSwitch flow tracking and export via ipfix/netflow. Applies to FortiOS Version `>= 6.2.4`.
//
// ## Import
//
// # SwitchController FlowTracking can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerFlowtracking:SwitchcontrollerFlowtracking labelname SwitchControllerFlowTracking
//
// ```
//
//	If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
//
// ```sh
//
//	$ pulumi import fortios:index/switchcontrollerFlowtracking:SwitchcontrollerFlowtracking labelname SwitchControllerFlowTracking
//
// ```
//
//	$ unset "FORTIOS_IMPORT_TABLE"
type SwitchcontrollerFlowtracking struct {
	pulumi.CustomResourceState

	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates SwitchcontrollerFlowtrackingAggregateArrayOutput `pulumi:"aggregates"`
	// Configure collector ip address.
	CollectorIp pulumi.StringOutput `pulumi:"collectorIp"`
	// Configure collector port number(0-65535, default=0).
	CollectorPort pulumi.IntOutput `pulumi:"collectorPort"`
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors SwitchcontrollerFlowtrackingCollectorArrayOutput `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrOutput `pulumi:"dynamicSortSubtable"`
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format pulumi.StringOutput `pulumi:"format"`
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level pulumi.StringOutput `pulumi:"level"`
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize pulumi.IntOutput `pulumi:"maxExportPktSize"`
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode pulumi.StringOutput `pulumi:"sampleMode"`
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate pulumi.IntOutput `pulumi:"sampleRate"`
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod pulumi.IntOutput `pulumi:"templateExportPeriod"`
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral pulumi.IntOutput `pulumi:"timeoutGeneral"`
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp pulumi.IntOutput `pulumi:"timeoutIcmp"`
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax pulumi.IntOutput `pulumi:"timeoutMax"`
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp pulumi.IntOutput `pulumi:"timeoutTcp"`
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin pulumi.IntOutput `pulumi:"timeoutTcpFin"`
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst pulumi.IntOutput `pulumi:"timeoutTcpRst"`
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp pulumi.IntOutput `pulumi:"timeoutUdp"`
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport pulumi.StringOutput `pulumi:"transport"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrOutput `pulumi:"vdomparam"`
}

// NewSwitchcontrollerFlowtracking registers a new resource with the given unique name, arguments, and options.
func NewSwitchcontrollerFlowtracking(ctx *pulumi.Context,
	name string, args *SwitchcontrollerFlowtrackingArgs, opts ...pulumi.ResourceOption) (*SwitchcontrollerFlowtracking, error) {
	if args == nil {
		args = &SwitchcontrollerFlowtrackingArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchcontrollerFlowtracking
	err := ctx.RegisterResource("fortios:index/switchcontrollerFlowtracking:SwitchcontrollerFlowtracking", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchcontrollerFlowtracking gets an existing SwitchcontrollerFlowtracking resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchcontrollerFlowtracking(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchcontrollerFlowtrackingState, opts ...pulumi.ResourceOption) (*SwitchcontrollerFlowtracking, error) {
	var resource SwitchcontrollerFlowtracking
	err := ctx.ReadResource("fortios:index/switchcontrollerFlowtracking:SwitchcontrollerFlowtracking", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchcontrollerFlowtracking resources.
type switchcontrollerFlowtrackingState struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates []SwitchcontrollerFlowtrackingAggregate `pulumi:"aggregates"`
	// Configure collector ip address.
	CollectorIp *string `pulumi:"collectorIp"`
	// Configure collector port number(0-65535, default=0).
	CollectorPort *int `pulumi:"collectorPort"`
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors []SwitchcontrollerFlowtrackingCollector `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format *string `pulumi:"format"`
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level *string `pulumi:"level"`
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize *int `pulumi:"maxExportPktSize"`
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode *string `pulumi:"sampleMode"`
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate *int `pulumi:"sampleRate"`
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod *int `pulumi:"templateExportPeriod"`
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral *int `pulumi:"timeoutGeneral"`
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp *int `pulumi:"timeoutIcmp"`
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax *int `pulumi:"timeoutMax"`
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp *int `pulumi:"timeoutTcp"`
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin *int `pulumi:"timeoutTcpFin"`
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst *int `pulumi:"timeoutTcpRst"`
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp *int `pulumi:"timeoutUdp"`
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport *string `pulumi:"transport"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

type SwitchcontrollerFlowtrackingState struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates SwitchcontrollerFlowtrackingAggregateArrayInput
	// Configure collector ip address.
	CollectorIp pulumi.StringPtrInput
	// Configure collector port number(0-65535, default=0).
	CollectorPort pulumi.IntPtrInput
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors SwitchcontrollerFlowtrackingCollectorArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format pulumi.StringPtrInput
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level pulumi.StringPtrInput
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize pulumi.IntPtrInput
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode pulumi.StringPtrInput
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate pulumi.IntPtrInput
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod pulumi.IntPtrInput
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral pulumi.IntPtrInput
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp pulumi.IntPtrInput
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax pulumi.IntPtrInput
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp pulumi.IntPtrInput
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin pulumi.IntPtrInput
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst pulumi.IntPtrInput
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp pulumi.IntPtrInput
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerFlowtrackingState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerFlowtrackingState)(nil)).Elem()
}

type switchcontrollerFlowtrackingArgs struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates []SwitchcontrollerFlowtrackingAggregate `pulumi:"aggregates"`
	// Configure collector ip address.
	CollectorIp *string `pulumi:"collectorIp"`
	// Configure collector port number(0-65535, default=0).
	CollectorPort *int `pulumi:"collectorPort"`
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors []SwitchcontrollerFlowtrackingCollector `pulumi:"collectors"`
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable *string `pulumi:"dynamicSortSubtable"`
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format *string `pulumi:"format"`
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level *string `pulumi:"level"`
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize *int `pulumi:"maxExportPktSize"`
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode *string `pulumi:"sampleMode"`
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate *int `pulumi:"sampleRate"`
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod *int `pulumi:"templateExportPeriod"`
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral *int `pulumi:"timeoutGeneral"`
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp *int `pulumi:"timeoutIcmp"`
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax *int `pulumi:"timeoutMax"`
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp *int `pulumi:"timeoutTcp"`
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin *int `pulumi:"timeoutTcpFin"`
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst *int `pulumi:"timeoutTcpRst"`
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp *int `pulumi:"timeoutUdp"`
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport *string `pulumi:"transport"`
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// The set of arguments for constructing a SwitchcontrollerFlowtracking resource.
type SwitchcontrollerFlowtrackingArgs struct {
	// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
	Aggregates SwitchcontrollerFlowtrackingAggregateArrayInput
	// Configure collector ip address.
	CollectorIp pulumi.StringPtrInput
	// Configure collector port number(0-65535, default=0).
	CollectorPort pulumi.IntPtrInput
	// Configure collectors for the flow. The structure of `collectors` block is documented below.
	Collectors SwitchcontrollerFlowtrackingCollectorArrayInput
	// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
	DynamicSortSubtable pulumi.StringPtrInput
	// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
	Format pulumi.StringPtrInput
	// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
	Level pulumi.StringPtrInput
	// Configure flow max export packet size (512-9216, default=512 bytes).
	MaxExportPktSize pulumi.IntPtrInput
	// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
	SampleMode pulumi.StringPtrInput
	// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
	SampleRate pulumi.IntPtrInput
	// Configure template export period (1-60, default=5 minutes).
	TemplateExportPeriod pulumi.IntPtrInput
	// Configure flow session general timeout (60-604800, default=3600 seconds).
	TimeoutGeneral pulumi.IntPtrInput
	// Configure flow session ICMP timeout (60-604800, default=300 seconds).
	TimeoutIcmp pulumi.IntPtrInput
	// Configure flow session max timeout (60-604800, default=604800 seconds).
	TimeoutMax pulumi.IntPtrInput
	// Configure flow session TCP timeout (60-604800, default=3600 seconds).
	TimeoutTcp pulumi.IntPtrInput
	// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
	TimeoutTcpFin pulumi.IntPtrInput
	// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
	TimeoutTcpRst pulumi.IntPtrInput
	// Configure flow session UDP timeout (60-604800, default=300 seconds).
	TimeoutUdp pulumi.IntPtrInput
	// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
	Transport pulumi.StringPtrInput
	// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput
}

func (SwitchcontrollerFlowtrackingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchcontrollerFlowtrackingArgs)(nil)).Elem()
}

type SwitchcontrollerFlowtrackingInput interface {
	pulumi.Input

	ToSwitchcontrollerFlowtrackingOutput() SwitchcontrollerFlowtrackingOutput
	ToSwitchcontrollerFlowtrackingOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingOutput
}

func (*SwitchcontrollerFlowtracking) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerFlowtracking)(nil)).Elem()
}

func (i *SwitchcontrollerFlowtracking) ToSwitchcontrollerFlowtrackingOutput() SwitchcontrollerFlowtrackingOutput {
	return i.ToSwitchcontrollerFlowtrackingOutputWithContext(context.Background())
}

func (i *SwitchcontrollerFlowtracking) ToSwitchcontrollerFlowtrackingOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerFlowtrackingOutput)
}

// SwitchcontrollerFlowtrackingArrayInput is an input type that accepts SwitchcontrollerFlowtrackingArray and SwitchcontrollerFlowtrackingArrayOutput values.
// You can construct a concrete instance of `SwitchcontrollerFlowtrackingArrayInput` via:
//
//	SwitchcontrollerFlowtrackingArray{ SwitchcontrollerFlowtrackingArgs{...} }
type SwitchcontrollerFlowtrackingArrayInput interface {
	pulumi.Input

	ToSwitchcontrollerFlowtrackingArrayOutput() SwitchcontrollerFlowtrackingArrayOutput
	ToSwitchcontrollerFlowtrackingArrayOutputWithContext(context.Context) SwitchcontrollerFlowtrackingArrayOutput
}

type SwitchcontrollerFlowtrackingArray []SwitchcontrollerFlowtrackingInput

func (SwitchcontrollerFlowtrackingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerFlowtracking)(nil)).Elem()
}

func (i SwitchcontrollerFlowtrackingArray) ToSwitchcontrollerFlowtrackingArrayOutput() SwitchcontrollerFlowtrackingArrayOutput {
	return i.ToSwitchcontrollerFlowtrackingArrayOutputWithContext(context.Background())
}

func (i SwitchcontrollerFlowtrackingArray) ToSwitchcontrollerFlowtrackingArrayOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerFlowtrackingArrayOutput)
}

// SwitchcontrollerFlowtrackingMapInput is an input type that accepts SwitchcontrollerFlowtrackingMap and SwitchcontrollerFlowtrackingMapOutput values.
// You can construct a concrete instance of `SwitchcontrollerFlowtrackingMapInput` via:
//
//	SwitchcontrollerFlowtrackingMap{ "key": SwitchcontrollerFlowtrackingArgs{...} }
type SwitchcontrollerFlowtrackingMapInput interface {
	pulumi.Input

	ToSwitchcontrollerFlowtrackingMapOutput() SwitchcontrollerFlowtrackingMapOutput
	ToSwitchcontrollerFlowtrackingMapOutputWithContext(context.Context) SwitchcontrollerFlowtrackingMapOutput
}

type SwitchcontrollerFlowtrackingMap map[string]SwitchcontrollerFlowtrackingInput

func (SwitchcontrollerFlowtrackingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerFlowtracking)(nil)).Elem()
}

func (i SwitchcontrollerFlowtrackingMap) ToSwitchcontrollerFlowtrackingMapOutput() SwitchcontrollerFlowtrackingMapOutput {
	return i.ToSwitchcontrollerFlowtrackingMapOutputWithContext(context.Background())
}

func (i SwitchcontrollerFlowtrackingMap) ToSwitchcontrollerFlowtrackingMapOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchcontrollerFlowtrackingMapOutput)
}

type SwitchcontrollerFlowtrackingOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerFlowtrackingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchcontrollerFlowtracking)(nil)).Elem()
}

func (o SwitchcontrollerFlowtrackingOutput) ToSwitchcontrollerFlowtrackingOutput() SwitchcontrollerFlowtrackingOutput {
	return o
}

func (o SwitchcontrollerFlowtrackingOutput) ToSwitchcontrollerFlowtrackingOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingOutput {
	return o
}

// Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow. The structure of `aggregates` block is documented below.
func (o SwitchcontrollerFlowtrackingOutput) Aggregates() SwitchcontrollerFlowtrackingAggregateArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) SwitchcontrollerFlowtrackingAggregateArrayOutput {
		return v.Aggregates
	}).(SwitchcontrollerFlowtrackingAggregateArrayOutput)
}

// Configure collector ip address.
func (o SwitchcontrollerFlowtrackingOutput) CollectorIp() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringOutput { return v.CollectorIp }).(pulumi.StringOutput)
}

// Configure collector port number(0-65535, default=0).
func (o SwitchcontrollerFlowtrackingOutput) CollectorPort() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.CollectorPort }).(pulumi.IntOutput)
}

// Configure collectors for the flow. The structure of `collectors` block is documented below.
func (o SwitchcontrollerFlowtrackingOutput) Collectors() SwitchcontrollerFlowtrackingCollectorArrayOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) SwitchcontrollerFlowtrackingCollectorArrayOutput {
		return v.Collectors
	}).(SwitchcontrollerFlowtrackingCollectorArrayOutput)
}

// Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] -> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] -> [ a10, a2 ].
func (o SwitchcontrollerFlowtrackingOutput) DynamicSortSubtable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringPtrOutput { return v.DynamicSortSubtable }).(pulumi.StringPtrOutput)
}

// Configure flow tracking protocol. Valid values: `netflow1`, `netflow5`, `netflow9`, `ipfix`.
func (o SwitchcontrollerFlowtrackingOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Configure flow tracking level. Valid values: `vlan`, `ip`, `port`, `proto`, `mac`.
func (o SwitchcontrollerFlowtrackingOutput) Level() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringOutput { return v.Level }).(pulumi.StringOutput)
}

// Configure flow max export packet size (512-9216, default=512 bytes).
func (o SwitchcontrollerFlowtrackingOutput) MaxExportPktSize() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.MaxExportPktSize }).(pulumi.IntOutput)
}

// Configure sample mode for the flow tracking. Valid values: `local`, `perimeter`, `device-ingress`.
func (o SwitchcontrollerFlowtrackingOutput) SampleMode() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringOutput { return v.SampleMode }).(pulumi.StringOutput)
}

// Configure sample rate for the perimeter and device-ingress sampling(0 - 99999).
func (o SwitchcontrollerFlowtrackingOutput) SampleRate() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.SampleRate }).(pulumi.IntOutput)
}

// Configure template export period (1-60, default=5 minutes).
func (o SwitchcontrollerFlowtrackingOutput) TemplateExportPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TemplateExportPeriod }).(pulumi.IntOutput)
}

// Configure flow session general timeout (60-604800, default=3600 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutGeneral() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutGeneral }).(pulumi.IntOutput)
}

// Configure flow session ICMP timeout (60-604800, default=300 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutIcmp() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutIcmp }).(pulumi.IntOutput)
}

// Configure flow session max timeout (60-604800, default=604800 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutMax() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutMax }).(pulumi.IntOutput)
}

// Configure flow session TCP timeout (60-604800, default=3600 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutTcp() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutTcp }).(pulumi.IntOutput)
}

// Configure flow session TCP FIN timeout (60-604800, default=300 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutTcpFin() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutTcpFin }).(pulumi.IntOutput)
}

// Configure flow session TCP RST timeout (60-604800, default=120 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutTcpRst() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutTcpRst }).(pulumi.IntOutput)
}

// Configure flow session UDP timeout (60-604800, default=300 seconds).
func (o SwitchcontrollerFlowtrackingOutput) TimeoutUdp() pulumi.IntOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.IntOutput { return v.TimeoutUdp }).(pulumi.IntOutput)
}

// Configure L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.
func (o SwitchcontrollerFlowtrackingOutput) Transport() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringOutput { return v.Transport }).(pulumi.StringOutput)
}

// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
func (o SwitchcontrollerFlowtrackingOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SwitchcontrollerFlowtracking) pulumi.StringPtrOutput { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

type SwitchcontrollerFlowtrackingArrayOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerFlowtrackingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchcontrollerFlowtracking)(nil)).Elem()
}

func (o SwitchcontrollerFlowtrackingArrayOutput) ToSwitchcontrollerFlowtrackingArrayOutput() SwitchcontrollerFlowtrackingArrayOutput {
	return o
}

func (o SwitchcontrollerFlowtrackingArrayOutput) ToSwitchcontrollerFlowtrackingArrayOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingArrayOutput {
	return o
}

func (o SwitchcontrollerFlowtrackingArrayOutput) Index(i pulumi.IntInput) SwitchcontrollerFlowtrackingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchcontrollerFlowtracking {
		return vs[0].([]*SwitchcontrollerFlowtracking)[vs[1].(int)]
	}).(SwitchcontrollerFlowtrackingOutput)
}

type SwitchcontrollerFlowtrackingMapOutput struct{ *pulumi.OutputState }

func (SwitchcontrollerFlowtrackingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchcontrollerFlowtracking)(nil)).Elem()
}

func (o SwitchcontrollerFlowtrackingMapOutput) ToSwitchcontrollerFlowtrackingMapOutput() SwitchcontrollerFlowtrackingMapOutput {
	return o
}

func (o SwitchcontrollerFlowtrackingMapOutput) ToSwitchcontrollerFlowtrackingMapOutputWithContext(ctx context.Context) SwitchcontrollerFlowtrackingMapOutput {
	return o
}

func (o SwitchcontrollerFlowtrackingMapOutput) MapIndex(k pulumi.StringInput) SwitchcontrollerFlowtrackingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchcontrollerFlowtracking {
		return vs[0].(map[string]*SwitchcontrollerFlowtracking)[vs[1].(string)]
	}).(SwitchcontrollerFlowtrackingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerFlowtrackingInput)(nil)).Elem(), &SwitchcontrollerFlowtracking{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerFlowtrackingArrayInput)(nil)).Elem(), SwitchcontrollerFlowtrackingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchcontrollerFlowtrackingMapInput)(nil)).Elem(), SwitchcontrollerFlowtrackingMap{})
	pulumi.RegisterOutputType(SwitchcontrollerFlowtrackingOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerFlowtrackingArrayOutput{})
	pulumi.RegisterOutputType(SwitchcontrollerFlowtrackingMapOutput{})
}