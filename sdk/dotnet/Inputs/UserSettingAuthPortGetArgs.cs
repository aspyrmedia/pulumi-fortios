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

    public sealed class UserSettingAuthPortGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Non-standard port for firewall user authentication.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Service type. Valid values: `http`, `https`, `ftp`, `telnet`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public UserSettingAuthPortGetArgs()
        {
        }
        public static new UserSettingAuthPortGetArgs Empty => new UserSettingAuthPortGetArgs();
    }
}