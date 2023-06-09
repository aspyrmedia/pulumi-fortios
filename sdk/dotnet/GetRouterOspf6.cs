// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios
{
    public static class GetRouterOspf6
    {
        /// <summary>
        /// Use this data source to get information on fortios router ospf6
        /// </summary>
        public static Task<GetRouterOspf6Result> InvokeAsync(GetRouterOspf6Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouterOspf6Result>("fortios:index/getRouterOspf6:getRouterOspf6", args ?? new GetRouterOspf6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on fortios router ospf6
        /// </summary>
        public static Output<GetRouterOspf6Result> Invoke(GetRouterOspf6InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouterOspf6Result>("fortios:index/getRouterOspf6:getRouterOspf6", args ?? new GetRouterOspf6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouterOspf6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetRouterOspf6Args()
        {
        }
        public static new GetRouterOspf6Args Empty => new GetRouterOspf6Args();
    }

    public sealed class GetRouterOspf6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetRouterOspf6InvokeArgs()
        {
        }
        public static new GetRouterOspf6InvokeArgs Empty => new GetRouterOspf6InvokeArgs();
    }


    [OutputType]
    public sealed class GetRouterOspf6Result
    {
        /// <summary>
        /// Area border router type.
        /// </summary>
        public readonly string AbrType;
        /// <summary>
        /// OSPF6 area configuration. The structure of `area` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspf6AreaResult> Areas;
        /// <summary>
        /// Reference bandwidth in terms of megabits per second.
        /// </summary>
        public readonly int AutoCostRefBandwidth;
        /// <summary>
        /// Enable/disable Bidirectional Forwarding Detection (BFD).
        /// </summary>
        public readonly string Bfd;
        /// <summary>
        /// Default information metric.
        /// </summary>
        public readonly int DefaultInformationMetric;
        /// <summary>
        /// Default information metric type.
        /// </summary>
        public readonly string DefaultInformationMetricType;
        /// <summary>
        /// Enable/disable generation of default route.
        /// </summary>
        public readonly string DefaultInformationOriginate;
        /// <summary>
        /// Default information route map.
        /// </summary>
        public readonly string DefaultInformationRouteMap;
        /// <summary>
        /// Default metric of redistribute routes.
        /// </summary>
        public readonly int DefaultMetric;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Enable logging of OSPFv3 neighbour's changes
        /// </summary>
        public readonly string LogNeighbourChanges;
        /// <summary>
        /// OSPF6 interface configuration. The structure of `ospf6_interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspf6Ospf6InterfaceResult> Ospf6Interfaces;
        /// <summary>
        /// Passive interface configuration. The structure of `passive_interface` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspf6PassiveInterfaceResult> PassiveInterfaces;
        /// <summary>
        /// Redistribute configuration. The structure of `redistribute` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspf6RedistributeResult> Redistributes;
        /// <summary>
        /// OSPFv3 restart mode (graceful or none).
        /// </summary>
        public readonly string RestartMode;
        /// <summary>
        /// Enable/disable continuing graceful restart upon topology change.
        /// </summary>
        public readonly string RestartOnTopologyChange;
        /// <summary>
        /// Graceful restart period in seconds.
        /// </summary>
        public readonly int RestartPeriod;
        /// <summary>
        /// A.B.C.D, in IPv4 address format.
        /// </summary>
        public readonly string RouterId;
        /// <summary>
        /// SPF calculation frequency.
        /// </summary>
        public readonly string SpfTimers;
        /// <summary>
        /// IPv6 address summary configuration. The structure of `summary_address` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterOspf6SummaryAddressResult> SummaryAddresses;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetRouterOspf6Result(
            string abrType,

            ImmutableArray<Outputs.GetRouterOspf6AreaResult> areas,

            int autoCostRefBandwidth,

            string bfd,

            int defaultInformationMetric,

            string defaultInformationMetricType,

            string defaultInformationOriginate,

            string defaultInformationRouteMap,

            int defaultMetric,

            string id,

            string logNeighbourChanges,

            ImmutableArray<Outputs.GetRouterOspf6Ospf6InterfaceResult> ospf6Interfaces,

            ImmutableArray<Outputs.GetRouterOspf6PassiveInterfaceResult> passiveInterfaces,

            ImmutableArray<Outputs.GetRouterOspf6RedistributeResult> redistributes,

            string restartMode,

            string restartOnTopologyChange,

            int restartPeriod,

            string routerId,

            string spfTimers,

            ImmutableArray<Outputs.GetRouterOspf6SummaryAddressResult> summaryAddresses,

            string? vdomparam)
        {
            AbrType = abrType;
            Areas = areas;
            AutoCostRefBandwidth = autoCostRefBandwidth;
            Bfd = bfd;
            DefaultInformationMetric = defaultInformationMetric;
            DefaultInformationMetricType = defaultInformationMetricType;
            DefaultInformationOriginate = defaultInformationOriginate;
            DefaultInformationRouteMap = defaultInformationRouteMap;
            DefaultMetric = defaultMetric;
            Id = id;
            LogNeighbourChanges = logNeighbourChanges;
            Ospf6Interfaces = ospf6Interfaces;
            PassiveInterfaces = passiveInterfaces;
            Redistributes = redistributes;
            RestartMode = restartMode;
            RestartOnTopologyChange = restartOnTopologyChange;
            RestartPeriod = restartPeriod;
            RouterId = routerId;
            SpfTimers = spfTimers;
            SummaryAddresses = summaryAddresses;
            Vdomparam = vdomparam;
        }
    }
}
