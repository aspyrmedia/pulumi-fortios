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
    public sealed class GetRouterBgpVrfLeak6Result
    {
        /// <summary>
        /// Target VRF table. The structure of `target` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRouterBgpVrfLeak6TargetResult> Targets;
        /// <summary>
        /// Target VRF ID &lt;0 - 31&gt;.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetRouterBgpVrfLeak6Result(
            ImmutableArray<Outputs.GetRouterBgpVrfLeak6TargetResult> targets,

            string vrf)
        {
            Targets = targets;
            Vrf = vrf;
        }
    }
}
