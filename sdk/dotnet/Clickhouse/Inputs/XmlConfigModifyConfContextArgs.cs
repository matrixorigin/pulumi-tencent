// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Clickhouse.Inputs
{

    public sealed class XmlConfigModifyConfContextArgs : global::Pulumi.ResourceArgs
    {
        [Input("fileName", required: true)]
        public Input<string> FileName { get; set; } = null!;

        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        [Input("newConfValue", required: true)]
        public Input<string> NewConfValue { get; set; } = null!;

        public XmlConfigModifyConfContextArgs()
        {
        }
        public static new XmlConfigModifyConfContextArgs Empty => new XmlConfigModifyConfContextArgs();
    }
}
