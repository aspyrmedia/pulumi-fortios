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
    public sealed class Wirelesscontrollerhotspot20AnqproamingconsortiumOiList
    {
        /// <summary>
        /// Comment.
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// OI index.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// Organization identifier.
        /// </summary>
        public readonly string? Oi;

        [OutputConstructor]
        private Wirelesscontrollerhotspot20AnqproamingconsortiumOiList(
            string? comment,

            int? index,

            string? oi)
        {
            Comment = comment;
            Index = index;
            Oi = oi;
        }
    }
}
