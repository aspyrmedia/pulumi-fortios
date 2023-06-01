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
    public sealed class FirewallInternetserviceextensionDisableEntryIp6Range
    {
        /// <summary>
        /// End IPv6 address.
        /// </summary>
        public readonly string? EndIp6;
        /// <summary>
        /// Entry ID(1-255).
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Start IPv6 address.
        /// </summary>
        public readonly string? StartIp6;

        [OutputConstructor]
        private FirewallInternetserviceextensionDisableEntryIp6Range(
            string? endIp6,

            int? id,

            string? startIp6)
        {
            EndIp6 = endIp6;
            Id = id;
            StartIp6 = startIp6;
        }
    }
}