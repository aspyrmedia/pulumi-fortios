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
    public sealed class ApplicationListEntryParameter
    {
        /// <summary>
        /// Parameter ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Parameter tuple members. The structure of `members` block is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationListEntryParameterMember> Members;
        /// <summary>
        /// Parameter value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ApplicationListEntryParameter(
            int? id,

            ImmutableArray<Outputs.ApplicationListEntryParameterMember> members,

            string? value)
        {
            Id = id;
            Members = members;
            Value = value;
        }
    }
}
