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
    public static class GetFirewallMulticastaddress
    {
        /// <summary>
        /// Use this data source to get information on an fortios firewall multicastaddress
        /// </summary>
        public static Task<GetFirewallMulticastaddressResult> InvokeAsync(GetFirewallMulticastaddressArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallMulticastaddressResult>("fortios:index/getFirewallMulticastaddress:getFirewallMulticastaddress", args ?? new GetFirewallMulticastaddressArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information on an fortios firewall multicastaddress
        /// </summary>
        public static Output<GetFirewallMulticastaddressResult> Invoke(GetFirewallMulticastaddressInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallMulticastaddressResult>("fortios:index/getFirewallMulticastaddress:getFirewallMulticastaddress", args ?? new GetFirewallMulticastaddressInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallMulticastaddressArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall multicastaddress.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public string? Vdomparam { get; set; }

        public GetFirewallMulticastaddressArgs()
        {
        }
        public static new GetFirewallMulticastaddressArgs Empty => new GetFirewallMulticastaddressArgs();
    }

    public sealed class GetFirewallMulticastaddressInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specify the name of the desired firewall multicastaddress.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public GetFirewallMulticastaddressInvokeArgs()
        {
        }
        public static new GetFirewallMulticastaddressInvokeArgs Empty => new GetFirewallMulticastaddressInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallMulticastaddressResult
    {
        /// <summary>
        /// Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
        /// </summary>
        public readonly string AssociatedInterface;
        /// <summary>
        /// Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
        /// </summary>
        public readonly int Color;
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string Comment;
        /// <summary>
        /// Final IPv4 address (inclusive) in the range for the address.
        /// </summary>
        public readonly string EndIp;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// First IPv4 address (inclusive) in the range for the address.
        /// </summary>
        public readonly string StartIp;
        /// <summary>
        /// Broadcast address and subnet.
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// Config object tagging. The structure of `tagging` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFirewallMulticastaddressTaggingResult> Taggings;
        /// <summary>
        /// Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address.
        /// </summary>
        public readonly string Type;
        public readonly string? Vdomparam;
        /// <summary>
        /// Enable/disable visibility of the multicast address on the GUI.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetFirewallMulticastaddressResult(
            string associatedInterface,

            int color,

            string comment,

            string endIp,

            string id,

            string name,

            string startIp,

            string subnet,

            ImmutableArray<Outputs.GetFirewallMulticastaddressTaggingResult> taggings,

            string type,

            string? vdomparam,

            string visibility)
        {
            AssociatedInterface = associatedInterface;
            Color = color;
            Comment = comment;
            EndIp = endIp;
            Id = id;
            Name = name;
            StartIp = startIp;
            Subnet = subnet;
            Taggings = taggings;
            Type = type;
            Vdomparam = vdomparam;
            Visibility = visibility;
        }
    }
}
