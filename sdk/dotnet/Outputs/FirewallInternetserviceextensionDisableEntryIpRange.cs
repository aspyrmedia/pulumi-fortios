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
    public sealed class FirewallInternetserviceextensionDisableEntryIpRange
    {
        /// <summary>
        /// End IP address.
        /// 
        /// The `ip6_range` block supports:
        /// </summary>
        public readonly string? EndIp;
        /// <summary>
        /// Disable entry range ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Start IP address.
        /// </summary>
        public readonly string? StartIp;

        [OutputConstructor]
        private FirewallInternetserviceextensionDisableEntryIpRange(
            string? endIp,

            int? id,

            string? startIp)
        {
            EndIp = endIp;
            Id = id;
            StartIp = startIp;
        }
    }
}
