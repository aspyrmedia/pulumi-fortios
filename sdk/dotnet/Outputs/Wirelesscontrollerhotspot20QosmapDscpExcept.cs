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
    public sealed class Wirelesscontrollerhotspot20QosmapDscpExcept
    {
        /// <summary>
        /// DSCP value.
        /// </summary>
        public readonly int? Dscp;
        /// <summary>
        /// DSCP exception index.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// User priority.
        /// </summary>
        public readonly int? Up;

        [OutputConstructor]
        private Wirelesscontrollerhotspot20QosmapDscpExcept(
            int? dscp,

            int? index,

            int? up)
        {
            Dscp = dscp;
            Index = index;
            Up = up;
        }
    }
}