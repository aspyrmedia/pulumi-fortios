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
    public sealed class GetFirewallAddress6MacaddrResult
    {
        /// <summary>
        /// MAC address ranges &lt;start&gt;[-&lt;end&gt;] separated by space.
        /// </summary>
        public readonly string Macaddr;

        [OutputConstructor]
        private GetFirewallAddress6MacaddrResult(string macaddr)
        {
            Macaddr = macaddr;
        }
    }
}
