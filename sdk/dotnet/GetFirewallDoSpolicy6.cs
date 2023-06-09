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
    public static class GetFirewallDoSpolicy6
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall DoSpolicy6
        /// </summary>
        public static Task<GetFirewallDoSpolicy6Result> InvokeAsync(GetFirewallDoSpolicy6Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallDoSpolicy6Result>("fortios:index/getFirewallDoSpolicy6:getFirewallDoSpolicy6", args ?? new GetFirewallDoSpolicy6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall DoSpolicy6
        /// </summary>
        public static Output<GetFirewallDoSpolicy6Result> Invoke(GetFirewallDoSpolicy6InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallDoSpolicy6Result>("fortios:index/getFirewallDoSpolicy6:getFirewallDoSpolicy6", args ?? new GetFirewallDoSpolicy6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallDoSpolicy6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall DoSpolicy6.
        /// </summary>
        [Input("policyid", required: true)]
        public int Policyid { get; set; }

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallDoSpolicy6Args()
        {
        }
        public static new GetFirewallDoSpolicy6Args Empty => new GetFirewallDoSpolicy6Args();
    }

    public sealed class GetFirewallDoSpolicy6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the policyid of the desired firewall DoSpolicy6.
        /// </summary>
        [Input("policyid", required: true)]
        public Input<int> Policyid { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallDoSpolicy6InvokeArgs()
        {
        }
        public static new GetFirewallDoSpolicy6InvokeArgs Empty => new GetFirewallDoSpolicy6InvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallDoSpolicy6Result
    {
        /// <summary>
        /// Anomaly name. The structure of `anomaly` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallDoSpolicy6AnomalyResult> Anomalies;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comments;
        /// <summary>
        /// Destination address name from available addresses. The structure of `dstaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallDoSpolicy6DstaddrResult> Dstaddrs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Incoming interface name from available interfaces.
        /// </summary>
        public readonly string Interface;
        /// <summary>
        /// Anomaly name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Policy ID.
        /// </summary>
        public readonly int Policyid;
        /// <summary>
        /// Service object from available options. The structure of `service` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallDoSpolicy6ServiceResult> Services;
        /// <summary>
        /// Source address name from available addresses. The structure of `srcaddr` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallDoSpolicy6SrcaddrResult> Srcaddrs;
        /// <summary>
        /// Enable/disable this anomaly.
        /// </summary>
        public readonly string Status;
        public readonly string? Vdomparam;

        [OutputConstructor]
        private GetFirewallDoSpolicy6Result(
            ImmutableArray<Outputs.GetFirewallDoSpolicy6AnomalyResult> anomalies,

            string comments,

            ImmutableArray<Outputs.GetFirewallDoSpolicy6DstaddrResult> dstaddrs,

            string id,

            string @interface,

            string name,

            int policyid,

            ImmutableArray<Outputs.GetFirewallDoSpolicy6ServiceResult> services,

            ImmutableArray<Outputs.GetFirewallDoSpolicy6SrcaddrResult> srcaddrs,

            string status,

            string? vdomparam)
        {
            Anomalies = anomalies;
            Comments = comments;
            Dstaddrs = dstaddrs;
            Id = id;
            Interface = @interface;
            Name = name;
            Policyid = policyid;
            Services = services;
            Srcaddrs = srcaddrs;
            Status = status;
            Vdomparam = vdomparam;
        }
    }
}
