// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Fortios.Inputs
{

    public sealed class SystemClustersyncSyncvdGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// VDOM name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SystemClustersyncSyncvdGetArgs()
        {
        }
        public static new SystemClustersyncSyncvdGetArgs Empty => new SystemClustersyncSyncvdGetArgs();
    }
}
