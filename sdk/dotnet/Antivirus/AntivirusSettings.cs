// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Aspyrmedia.Fortios.Antivirus
{
    [FortiosResourceType("fortios:antivirus/antivirusSettings:antivirusSettings")]
    public partial class AntivirusSettings : global::Pulumi.CustomResource
    {
        [Output("cacheCleanResult")]
        public Output<string> CacheCleanResult { get; private set; } = null!;

        [Output("cacheInfectedResult")]
        public Output<string> CacheInfectedResult { get; private set; } = null!;

        [Output("defaultDb")]
        public Output<string> DefaultDb { get; private set; } = null!;

        [Output("grayware")]
        public Output<string> Grayware { get; private set; } = null!;

        [Output("machineLearningDetection")]
        public Output<string> MachineLearningDetection { get; private set; } = null!;

        [Output("overrideTimeout")]
        public Output<int> OverrideTimeout { get; private set; } = null!;

        [Output("useExtremeDb")]
        public Output<string> UseExtremeDb { get; private set; } = null!;

        [Output("vdomparam")]
        public Output<string?> Vdomparam { get; private set; } = null!;


        /// <summary>
        /// Create a AntivirusSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AntivirusSettings(string name, AntivirusSettingsArgs? args = null, CustomResourceOptions? options = null)
            : base("fortios:antivirus/antivirusSettings:antivirusSettings", name, args ?? new AntivirusSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AntivirusSettings(string name, Input<string> id, AntivirusSettingsState? state = null, CustomResourceOptions? options = null)
            : base("fortios:antivirus/antivirusSettings:antivirusSettings", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/aspyrmedia/pulumi-fortios/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AntivirusSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AntivirusSettings Get(string name, Input<string> id, AntivirusSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new AntivirusSettings(name, id, state, options);
        }
    }

    public sealed class AntivirusSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cacheCleanResult")]
        public Input<string>? CacheCleanResult { get; set; }

        [Input("cacheInfectedResult")]
        public Input<string>? CacheInfectedResult { get; set; }

        [Input("defaultDb")]
        public Input<string>? DefaultDb { get; set; }

        [Input("grayware")]
        public Input<string>? Grayware { get; set; }

        [Input("machineLearningDetection")]
        public Input<string>? MachineLearningDetection { get; set; }

        [Input("overrideTimeout")]
        public Input<int>? OverrideTimeout { get; set; }

        [Input("useExtremeDb")]
        public Input<string>? UseExtremeDb { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AntivirusSettingsArgs()
        {
        }
        public static new AntivirusSettingsArgs Empty => new AntivirusSettingsArgs();
    }

    public sealed class AntivirusSettingsState : global::Pulumi.ResourceArgs
    {
        [Input("cacheCleanResult")]
        public Input<string>? CacheCleanResult { get; set; }

        [Input("cacheInfectedResult")]
        public Input<string>? CacheInfectedResult { get; set; }

        [Input("defaultDb")]
        public Input<string>? DefaultDb { get; set; }

        [Input("grayware")]
        public Input<string>? Grayware { get; set; }

        [Input("machineLearningDetection")]
        public Input<string>? MachineLearningDetection { get; set; }

        [Input("overrideTimeout")]
        public Input<int>? OverrideTimeout { get; set; }

        [Input("useExtremeDb")]
        public Input<string>? UseExtremeDb { get; set; }

        [Input("vdomparam")]
        public Input<string>? Vdomparam { get; set; }

        public AntivirusSettingsState()
        {
        }
        public static new AntivirusSettingsState Empty => new AntivirusSettingsState();
    }
}
