// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Outputs
{

    [OutputType]
    public sealed class SystemInterfaceIpv6Ip6ExtraAddr
    {
        /// <summary>
        /// IPv6 address prefix.
        /// 
        /// The `ip6_prefix_list` block supports:
        /// </summary>
        public readonly string? Prefix;

        [OutputConstructor]
        private SystemInterfaceIpv6Ip6ExtraAddr(string? prefix)
        {
            Prefix = prefix;
        }
    }
}
