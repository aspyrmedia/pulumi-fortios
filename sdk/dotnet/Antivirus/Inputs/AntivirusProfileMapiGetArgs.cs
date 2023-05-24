// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Aspyrmedia.Fortios.Antivirus.Inputs
{

    public sealed class AntivirusProfileMapiGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("archiveBlock")]
        public Input<string>? ArchiveBlock { get; set; }

        [Input("archiveLog")]
        public Input<string>? ArchiveLog { get; set; }

        [Input("avScan")]
        public Input<string>? AvScan { get; set; }

        [Input("emulator")]
        public Input<string>? Emulator { get; set; }

        [Input("executables")]
        public Input<string>? Executables { get; set; }

        [Input("externalBlocklist")]
        public Input<string>? ExternalBlocklist { get; set; }

        [Input("fortiai")]
        public Input<string>? Fortiai { get; set; }

        [Input("fortindr")]
        public Input<string>? Fortindr { get; set; }

        [Input("fortisandbox")]
        public Input<string>? Fortisandbox { get; set; }

        [Input("options")]
        public Input<string>? Options { get; set; }

        [Input("outbreakPrevention")]
        public Input<string>? OutbreakPrevention { get; set; }

        [Input("quarantine")]
        public Input<string>? Quarantine { get; set; }

        public AntivirusProfileMapiGetArgs()
        {
        }
        public static new AntivirusProfileMapiGetArgs Empty => new AntivirusProfileMapiGetArgs();
    }
}
