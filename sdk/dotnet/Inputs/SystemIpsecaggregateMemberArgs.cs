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

    public sealed class SystemIpsecaggregateMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Tunnel name.
        /// </summary>
        [Input("tunnelName")]
        public Input<string>? TunnelName { get; set; }

        public SystemIpsecaggregateMemberArgs()
        {
        }
        public static new SystemIpsecaggregateMemberArgs Empty => new SystemIpsecaggregateMemberArgs();
    }
}