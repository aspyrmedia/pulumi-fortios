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
    public sealed class GetSystemAdminGuiDashboardWidgetFilterResult
    {
        /// <summary>
        /// Select menu ID.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// Filter key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Filter value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetSystemAdminGuiDashboardWidgetFilterResult(
            int id,

            string key,

            string value)
        {
            Id = id;
            Key = key;
            Value = value;
        }
    }
}