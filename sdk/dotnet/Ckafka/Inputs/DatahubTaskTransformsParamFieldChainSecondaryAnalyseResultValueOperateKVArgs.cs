// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateKVArgs : global::Pulumi.ResourceArgs
    {
        [Input("delimiter", required: true)]
        public Input<string> Delimiter { get; set; } = null!;

        [Input("keepOriginalKey")]
        public Input<string>? KeepOriginalKey { get; set; }

        [Input("regex", required: true)]
        public Input<string> Regex { get; set; } = null!;

        public DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateKVArgs()
        {
        }
        public static new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateKVArgs Empty => new DatahubTaskTransformsParamFieldChainSecondaryAnalyseResultValueOperateKVArgs();
    }
}