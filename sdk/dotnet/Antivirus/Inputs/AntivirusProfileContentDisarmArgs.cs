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

    public sealed class AntivirusProfileContentDisarmArgs : global::Pulumi.ResourceArgs
    {
        [Input("coverPage")]
        public Input<string>? CoverPage { get; set; }

        [Input("detectOnly")]
        public Input<string>? DetectOnly { get; set; }

        [Input("errorAction")]
        public Input<string>? ErrorAction { get; set; }

        [Input("officeAction")]
        public Input<string>? OfficeAction { get; set; }

        [Input("officeDde")]
        public Input<string>? OfficeDde { get; set; }

        [Input("officeEmbed")]
        public Input<string>? OfficeEmbed { get; set; }

        [Input("officeHylink")]
        public Input<string>? OfficeHylink { get; set; }

        [Input("officeLinked")]
        public Input<string>? OfficeLinked { get; set; }

        [Input("officeMacro")]
        public Input<string>? OfficeMacro { get; set; }

        [Input("originalFileDestination")]
        public Input<string>? OriginalFileDestination { get; set; }

        [Input("pdfActForm")]
        public Input<string>? PdfActForm { get; set; }

        [Input("pdfActGotor")]
        public Input<string>? PdfActGotor { get; set; }

        [Input("pdfActJava")]
        public Input<string>? PdfActJava { get; set; }

        [Input("pdfActLaunch")]
        public Input<string>? PdfActLaunch { get; set; }

        [Input("pdfActMovie")]
        public Input<string>? PdfActMovie { get; set; }

        [Input("pdfActSound")]
        public Input<string>? PdfActSound { get; set; }

        [Input("pdfEmbedfile")]
        public Input<string>? PdfEmbedfile { get; set; }

        [Input("pdfHyperlink")]
        public Input<string>? PdfHyperlink { get; set; }

        [Input("pdfJavacode")]
        public Input<string>? PdfJavacode { get; set; }

        public AntivirusProfileContentDisarmArgs()
        {
        }
        public static new AntivirusProfileContentDisarmArgs Empty => new AntivirusProfileContentDisarmArgs();
    }
}
