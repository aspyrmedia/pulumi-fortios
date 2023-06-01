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
    public sealed class SystemGeoipoverrideIpRange
    {
        /// <summary>
        /// Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        public readonly string? EndIp;
        /// <summary>
        /// ID of individual entry in the IPv6 range table.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
        /// </summary>
        public readonly string? StartIp;

        [OutputConstructor]
        private SystemGeoipoverrideIpRange(
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