// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTransformsParamArgs : global::Pulumi.ResourceArgs
    {
        [Input("batchAnalyse")]
        public Input<Inputs.DatahubTaskTransformsParamBatchAnalyseArgs>? BatchAnalyse { get; set; }

        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("failureParam")]
        public Input<Inputs.DatahubTaskTransformsParamFailureParamArgs>? FailureParam { get; set; }

        [Input("fieldChains", required: true)]
        private InputList<Inputs.DatahubTaskTransformsParamFieldChainArgs>? _fieldChains;
        public InputList<Inputs.DatahubTaskTransformsParamFieldChainArgs> FieldChains
        {
            get => _fieldChains ?? (_fieldChains = new InputList<Inputs.DatahubTaskTransformsParamFieldChainArgs>());
            set => _fieldChains = value;
        }

        [Input("filterParams")]
        private InputList<Inputs.DatahubTaskTransformsParamFilterParamArgs>? _filterParams;
        public InputList<Inputs.DatahubTaskTransformsParamFilterParamArgs> FilterParams
        {
            get => _filterParams ?? (_filterParams = new InputList<Inputs.DatahubTaskTransformsParamFilterParamArgs>());
            set => _filterParams = value;
        }

        [Input("keepMetadata")]
        public Input<bool>? KeepMetadata { get; set; }

        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        [Input("result")]
        public Input<string>? Result { get; set; }

        [Input("rowParam")]
        public Input<Inputs.DatahubTaskTransformsParamRowParamArgs>? RowParam { get; set; }

        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        public DatahubTaskTransformsParamArgs()
        {
        }
        public static new DatahubTaskTransformsParamArgs Empty => new DatahubTaskTransformsParamArgs();
    }
}