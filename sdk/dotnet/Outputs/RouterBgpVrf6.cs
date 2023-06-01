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
    public sealed class RouterBgpVrf6
    {
        /// <summary>
        /// Target VRF table. The structure of `leak_target` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.RouterBgpVrf6LeakTarget> LeakTargets;
        /// <summary>
        /// BGP VRF leaking table. The structure of `vrf` block is documented below.
        /// </summary>
        public readonly string? Vrf;

        [OutputConstructor]
        private RouterBgpVrf6(
            ImmutableArray<Outputs.RouterBgpVrf6LeakTarget> leakTargets,

            string? vrf)
        {
            LeakTargets = leakTargets;
            Vrf = vrf;
        }
    }
}