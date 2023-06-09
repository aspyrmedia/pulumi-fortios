// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fortios

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on fortios router ospf
func LookupRouterOspf(ctx *pulumi.Context, args *LookupRouterOspfArgs, opts ...pulumi.InvokeOption) (*LookupRouterOspfResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouterOspfResult
	err := ctx.Invoke("fortios:index/getRouterOspf:getRouterOspf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouterOspf.
type LookupRouterOspfArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam *string `pulumi:"vdomparam"`
}

// A collection of values returned by getRouterOspf.
type LookupRouterOspfResult struct {
	// Area border router type.
	AbrType string `pulumi:"abrType"`
	// Attach the network to area.
	Areas []GetRouterOspfArea `pulumi:"areas"`
	// Reference bandwidth in terms of megabits per second.
	AutoCostRefBandwidth int `pulumi:"autoCostRefBandwidth"`
	// Bidirectional Forwarding Detection (BFD).
	Bfd string `pulumi:"bfd"`
	// Enable/disable database overflow.
	DatabaseOverflow string `pulumi:"databaseOverflow"`
	// Database overflow maximum LSAs.
	DatabaseOverflowMaxLsas int `pulumi:"databaseOverflowMaxLsas"`
	// Database overflow time to recover (sec).
	DatabaseOverflowTimeToRecover int `pulumi:"databaseOverflowTimeToRecover"`
	// Default information metric.
	DefaultInformationMetric int `pulumi:"defaultInformationMetric"`
	// Default information metric type.
	DefaultInformationMetricType string `pulumi:"defaultInformationMetricType"`
	// Enable/disable generation of default route.
	DefaultInformationOriginate string `pulumi:"defaultInformationOriginate"`
	// Default information route map.
	DefaultInformationRouteMap string `pulumi:"defaultInformationRouteMap"`
	// Default metric of redistribute routes.
	DefaultMetric int `pulumi:"defaultMetric"`
	// Distance of the route.
	Distance int `pulumi:"distance"`
	// Administrative external distance.
	DistanceExternal int `pulumi:"distanceExternal"`
	// Administrative inter-area distance.
	DistanceInterArea int `pulumi:"distanceInterArea"`
	// Administrative intra-area distance.
	DistanceIntraArea int `pulumi:"distanceIntraArea"`
	// Filter incoming routes.
	DistributeListIn string `pulumi:"distributeListIn"`
	// Distribute list configuration. The structure of `distributeList` block is documented below.
	DistributeLists []GetRouterOspfDistributeList `pulumi:"distributeLists"`
	// Filter incoming external routes by route-map.
	DistributeRouteMapIn string `pulumi:"distributeRouteMapIn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Enable logging of OSPF neighbour's changes
	LogNeighbourChanges string `pulumi:"logNeighbourChanges"`
	// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
	Neighbors []GetRouterOspfNeighbor `pulumi:"neighbors"`
	// OSPF network configuration. The structure of `network` block is documented below.
	Networks []GetRouterOspfNetwork `pulumi:"networks"`
	// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
	OspfInterfaces []GetRouterOspfOspfInterface `pulumi:"ospfInterfaces"`
	// Passive interface configuration. The structure of `passiveInterface` block is documented below.
	PassiveInterfaces []GetRouterOspfPassiveInterface `pulumi:"passiveInterfaces"`
	// Redistribute configuration. The structure of `redistribute` block is documented below.
	Redistributes []GetRouterOspfRedistribute `pulumi:"redistributes"`
	// OSPF restart mode (graceful or LLS).
	RestartMode string `pulumi:"restartMode"`
	// Enable/disable continuing graceful restart upon topology change.
	RestartOnTopologyChange string `pulumi:"restartOnTopologyChange"`
	// Graceful restart period.
	RestartPeriod int `pulumi:"restartPeriod"`
	// Enable/disable RFC1583 compatibility.
	Rfc1583Compatible string `pulumi:"rfc1583Compatible"`
	// Router ID.
	RouterId string `pulumi:"routerId"`
	// SPF calculation frequency.
	SpfTimers string `pulumi:"spfTimers"`
	// IP address summary configuration. The structure of `summaryAddress` block is documented below.
	SummaryAddresses []GetRouterOspfSummaryAddress `pulumi:"summaryAddresses"`
	Vdomparam        *string                       `pulumi:"vdomparam"`
}

func LookupRouterOspfOutput(ctx *pulumi.Context, args LookupRouterOspfOutputArgs, opts ...pulumi.InvokeOption) LookupRouterOspfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouterOspfResult, error) {
			args := v.(LookupRouterOspfArgs)
			r, err := LookupRouterOspf(ctx, &args, opts...)
			var s LookupRouterOspfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouterOspfResultOutput)
}

// A collection of arguments for invoking getRouterOspf.
type LookupRouterOspfOutputArgs struct {
	// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
	Vdomparam pulumi.StringPtrInput `pulumi:"vdomparam"`
}

func (LookupRouterOspfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterOspfArgs)(nil)).Elem()
}

// A collection of values returned by getRouterOspf.
type LookupRouterOspfResultOutput struct{ *pulumi.OutputState }

func (LookupRouterOspfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouterOspfResult)(nil)).Elem()
}

func (o LookupRouterOspfResultOutput) ToLookupRouterOspfResultOutput() LookupRouterOspfResultOutput {
	return o
}

func (o LookupRouterOspfResultOutput) ToLookupRouterOspfResultOutputWithContext(ctx context.Context) LookupRouterOspfResultOutput {
	return o
}

// Area border router type.
func (o LookupRouterOspfResultOutput) AbrType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.AbrType }).(pulumi.StringOutput)
}

// Attach the network to area.
func (o LookupRouterOspfResultOutput) Areas() GetRouterOspfAreaArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfArea { return v.Areas }).(GetRouterOspfAreaArrayOutput)
}

// Reference bandwidth in terms of megabits per second.
func (o LookupRouterOspfResultOutput) AutoCostRefBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.AutoCostRefBandwidth }).(pulumi.IntOutput)
}

// Bidirectional Forwarding Detection (BFD).
func (o LookupRouterOspfResultOutput) Bfd() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.Bfd }).(pulumi.StringOutput)
}

// Enable/disable database overflow.
func (o LookupRouterOspfResultOutput) DatabaseOverflow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.DatabaseOverflow }).(pulumi.StringOutput)
}

// Database overflow maximum LSAs.
func (o LookupRouterOspfResultOutput) DatabaseOverflowMaxLsas() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DatabaseOverflowMaxLsas }).(pulumi.IntOutput)
}

// Database overflow time to recover (sec).
func (o LookupRouterOspfResultOutput) DatabaseOverflowTimeToRecover() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DatabaseOverflowTimeToRecover }).(pulumi.IntOutput)
}

// Default information metric.
func (o LookupRouterOspfResultOutput) DefaultInformationMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DefaultInformationMetric }).(pulumi.IntOutput)
}

// Default information metric type.
func (o LookupRouterOspfResultOutput) DefaultInformationMetricType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.DefaultInformationMetricType }).(pulumi.StringOutput)
}

// Enable/disable generation of default route.
func (o LookupRouterOspfResultOutput) DefaultInformationOriginate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.DefaultInformationOriginate }).(pulumi.StringOutput)
}

// Default information route map.
func (o LookupRouterOspfResultOutput) DefaultInformationRouteMap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.DefaultInformationRouteMap }).(pulumi.StringOutput)
}

// Default metric of redistribute routes.
func (o LookupRouterOspfResultOutput) DefaultMetric() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DefaultMetric }).(pulumi.IntOutput)
}

// Distance of the route.
func (o LookupRouterOspfResultOutput) Distance() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.Distance }).(pulumi.IntOutput)
}

// Administrative external distance.
func (o LookupRouterOspfResultOutput) DistanceExternal() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DistanceExternal }).(pulumi.IntOutput)
}

// Administrative inter-area distance.
func (o LookupRouterOspfResultOutput) DistanceInterArea() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DistanceInterArea }).(pulumi.IntOutput)
}

// Administrative intra-area distance.
func (o LookupRouterOspfResultOutput) DistanceIntraArea() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.DistanceIntraArea }).(pulumi.IntOutput)
}

// Filter incoming routes.
func (o LookupRouterOspfResultOutput) DistributeListIn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.DistributeListIn }).(pulumi.StringOutput)
}

// Distribute list configuration. The structure of `distributeList` block is documented below.
func (o LookupRouterOspfResultOutput) DistributeLists() GetRouterOspfDistributeListArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfDistributeList { return v.DistributeLists }).(GetRouterOspfDistributeListArrayOutput)
}

// Filter incoming external routes by route-map.
func (o LookupRouterOspfResultOutput) DistributeRouteMapIn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.DistributeRouteMapIn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouterOspfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.Id }).(pulumi.StringOutput)
}

// Enable logging of OSPF neighbour's changes
func (o LookupRouterOspfResultOutput) LogNeighbourChanges() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.LogNeighbourChanges }).(pulumi.StringOutput)
}

// OSPF neighbor configuration are used when OSPF runs on non-broadcast media The structure of `neighbor` block is documented below.
func (o LookupRouterOspfResultOutput) Neighbors() GetRouterOspfNeighborArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfNeighbor { return v.Neighbors }).(GetRouterOspfNeighborArrayOutput)
}

// OSPF network configuration. The structure of `network` block is documented below.
func (o LookupRouterOspfResultOutput) Networks() GetRouterOspfNetworkArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfNetwork { return v.Networks }).(GetRouterOspfNetworkArrayOutput)
}

// OSPF interface configuration. The structure of `ospfInterface` block is documented below.
func (o LookupRouterOspfResultOutput) OspfInterfaces() GetRouterOspfOspfInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfOspfInterface { return v.OspfInterfaces }).(GetRouterOspfOspfInterfaceArrayOutput)
}

// Passive interface configuration. The structure of `passiveInterface` block is documented below.
func (o LookupRouterOspfResultOutput) PassiveInterfaces() GetRouterOspfPassiveInterfaceArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfPassiveInterface { return v.PassiveInterfaces }).(GetRouterOspfPassiveInterfaceArrayOutput)
}

// Redistribute configuration. The structure of `redistribute` block is documented below.
func (o LookupRouterOspfResultOutput) Redistributes() GetRouterOspfRedistributeArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfRedistribute { return v.Redistributes }).(GetRouterOspfRedistributeArrayOutput)
}

// OSPF restart mode (graceful or LLS).
func (o LookupRouterOspfResultOutput) RestartMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.RestartMode }).(pulumi.StringOutput)
}

// Enable/disable continuing graceful restart upon topology change.
func (o LookupRouterOspfResultOutput) RestartOnTopologyChange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.RestartOnTopologyChange }).(pulumi.StringOutput)
}

// Graceful restart period.
func (o LookupRouterOspfResultOutput) RestartPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) int { return v.RestartPeriod }).(pulumi.IntOutput)
}

// Enable/disable RFC1583 compatibility.
func (o LookupRouterOspfResultOutput) Rfc1583Compatible() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.Rfc1583Compatible }).(pulumi.StringOutput)
}

// Router ID.
func (o LookupRouterOspfResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.RouterId }).(pulumi.StringOutput)
}

// SPF calculation frequency.
func (o LookupRouterOspfResultOutput) SpfTimers() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) string { return v.SpfTimers }).(pulumi.StringOutput)
}

// IP address summary configuration. The structure of `summaryAddress` block is documented below.
func (o LookupRouterOspfResultOutput) SummaryAddresses() GetRouterOspfSummaryAddressArrayOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) []GetRouterOspfSummaryAddress { return v.SummaryAddresses }).(GetRouterOspfSummaryAddressArrayOutput)
}

func (o LookupRouterOspfResultOutput) Vdomparam() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouterOspfResult) *string { return v.Vdomparam }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouterOspfResultOutput{})
}
