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
    /// <summary>
    /// Configure FortiSwitch QoS policy.
    /// 
    /// ## Import
    /// 
    /// SwitchControllerQos QosPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerqosQospolicy:SwitchcontrollerqosQospolicy labelname {{name}}
    /// ```
    /// 
    ///  If you do not want to import arguments of block$ export "FORTIOS_IMPORT_TABLE"="false"
    /// 
    /// ```sh
    ///  $ pulumi import fortios:index/switchcontrollerqosQospolicy:SwitchcontrollerqosQospolicy labelname {{name}}
    /// ```
    /// 
    ///  $ unset "FORTIOS_IMPORT_TABLE"
    /// </summary>
    [FortiosResourceType("fortios:index/switchcontrollerqosQospolicy:SwitchcontrollerqosQospolicy")]
    public partial class SwitchcontrollerqosQospolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Default cos queue for untagged packets.
        /// </summary>
        [Output("defaultCos")]
        public Output<int> DefaultCos { get; private set; } = null!;

        /// <summary>
        /// QoS policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// QoS egress queue policy.
        /// </summary>
        [Output("queuePolicy")]
        public Output<string> QueuePolicy { get; private set; } = null!;

        /// <summary>
        /// QoS trust 802.1p map.
        /// </summary>
        [Output("trustDot1pMap")]
        public Output<string> TrustDot1pMap { get; private set; } = null!;

        /// <summary>
        /// QoS trust ip dscp map.
        /// </summary>
        [Output("trustIpDscpMap")]
        public Output<string> TrustIpDscpMap { get; private set; } = null!;

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a SwitchcontrollerqosQospolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SwitchcontrollerqosQospolicy(string name, SwitchcontrollerqosQospolicyArgs args, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerqosQospolicy:SwitchcontrollerqosQospolicy", name, args ?? new SwitchcontrollerqosQospolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SwitchcontrollerqosQospolicy(string name, Input<string> id, SwitchcontrollerqosQospolicyState? state = null, CustomResourceOptions? options = null)
            : base("fortios:index/switchcontrollerqosQospolicy:SwitchcontrollerqosQospolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-fortios",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SwitchcontrollerqosQospolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SwitchcontrollerqosQospolicy Get(string name, Input<string> id, SwitchcontrollerqosQospolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new SwitchcontrollerqosQospolicy(name, id, state, options);
        }
    }

    public sealed class SwitchcontrollerqosQospolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default cos queue for untagged packets.
        /// </summary>
        [Input("defaultCos", required: true)]
        public Input<int> DefaultCos { get; set; } = null!;

        /// <summary>
        /// QoS policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// QoS egress queue policy.
        /// </summary>
        [Input("queuePolicy")]
        public Input<string>? QueuePolicy { get; set; }

        /// <summary>
        /// QoS trust 802.1p map.
        /// </summary>
        [Input("trustDot1pMap")]
        public Input<string>? TrustDot1pMap { get; set; }

        /// <summary>
        /// QoS trust ip dscp map.
        /// </summary>
        [Input("trustIpDscpMap")]
        public Input<string>? TrustIpDscpMap { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerqosQospolicyArgs()
        {
        }
        public static new SwitchcontrollerqosQospolicyArgs Empty => new SwitchcontrollerqosQospolicyArgs();
    }

    public sealed class SwitchcontrollerqosQospolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Default cos queue for untagged packets.
        /// </summary>
        [Input("defaultCos")]
        public Input<int>? DefaultCos { get; set; }

        /// <summary>
        /// QoS policy name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// QoS egress queue policy.
        /// </summary>
        [Input("queuePolicy")]
        public Input<string>? QueuePolicy { get; set; }

        /// <summary>
        /// QoS trust 802.1p map.
        /// </summary>
        [Input("trustDot1pMap")]
        public Input<string>? TrustDot1pMap { get; set; }

        /// <summary>
        /// QoS trust ip dscp map.
        /// </summary>
        [Input("trustIpDscpMap")]
        public Input<string>? TrustIpDscpMap { get; set; }

        /// <summary>
        /// Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
        /// </summary>
        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public SwitchcontrollerqosQospolicyState()
        {
        }
        public static new SwitchcontrollerqosQospolicyState Empty => new SwitchcontrollerqosQospolicyState();
    }
}