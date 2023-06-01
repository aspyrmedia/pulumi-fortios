// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemInterfaceSecondaryipGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Management access settings for the secondary IP address.
        /// </summary>
        [Input("allowaccess")]
        public Input<string>? Allowaccess { get; set; }

        /// <summary>
        /// Protocols used to detect the server. Valid values: `ping`, `tcp-echo`, `udp-echo`.
        /// </summary>
        [Input("detectprotocol")]
        public Input<string>? Detectprotocol { get; set; }

        /// <summary>
        /// Gateway's ping server for this IP.
        /// </summary>
        [Input("detectserver")]
        public Input<string>? Detectserver { get; set; }

        /// <summary>
        /// Enable/disable detect gateway alive for first. Valid values: `enable`, `disable`.
        /// </summary>
        [Input("gwdetect")]
        public Input<string>? Gwdetect { get; set; }

        /// <summary>
        /// HA election priority for the PING server.
        /// </summary>
        [Input("haPriority")]
        public Input<int>? HaPriority { get; set; }

        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Secondary IP address of the interface.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// PING server status.
        /// </summary>
        [Input("pingServStatus")]
        public Input<int>? PingServStatus { get; set; }

        public SystemInterfaceSecondaryipGetArgs()
        {
        }
        public static new SystemInterfaceSecondaryipGetArgs Empty => new SystemInterfaceSecondaryipGetArgs();
    }
}