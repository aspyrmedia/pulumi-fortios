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
    public static class GetFirewallMulticastaddress6
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall multicastaddress6
        /// </summary>
        public static Task<GetFirewallMulticastaddress6Result> InvokeAsync(GetFirewallMulticastaddress6Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallMulticastaddress6Result>("fortios:index/getFirewallMulticastaddress6:getFirewallMulticastaddress6", args ?? new GetFirewallMulticastaddress6Args(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall multicastaddress6
        /// </summary>
        public static Output<GetFirewallMulticastaddress6Result> Invoke(GetFirewallMulticastaddress6InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallMulticastaddress6Result>("fortios:index/getFirewallMulticastaddress6:getFirewallMulticastaddress6", args ?? new GetFirewallMulticastaddress6InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallMulticastaddress6Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall multicastaddress6.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallMulticastaddress6Args()
        {
        }
        public static new GetFirewallMulticastaddress6Args Empty => new GetFirewallMulticastaddress6Args();
    }

    public sealed class GetFirewallMulticastaddress6InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall multicastaddress6.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallMulticastaddress6InvokeArgs()
        {
        }
        public static new GetFirewallMulticastaddress6InvokeArgs Empty => new GetFirewallMulticastaddress6InvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallMulticastaddress6Result
    {
        /// <summary>
        /// Color of icon on the GUI.
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
        /// </summary>
        public readonly string Ip6;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallMulticastaddress6TaggingResult> Taggings;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable visibility of the IPv6 multicast address on the GUI.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetFirewallMulticastaddress6Result(
            int color,

            string comment,

            string id,

            string ip6,

            string name,

            ImmutableArray<Outputs.GetFirewallMulticastaddress6TaggingResult> taggings,

            string? vdomparam,

            string visibility)
        {
            Color = color;
            Comment = comment;
            Id = id;
            Ip6 = ip6;
            Name = name;
            Taggings = taggings;
            Vdomparam = vdomparam;
            Visibility = visibility;
        }
    }
}
