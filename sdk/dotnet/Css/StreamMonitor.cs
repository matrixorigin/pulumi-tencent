// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Css
{
    [TencentcloudResourceType("tencentcloud:Css/streamMonitor:StreamMonitor")]
    public partial class StreamMonitor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AI asr input index list.(first input index is 1.).
        /// </summary>
        [Output("aiAsrInputIndexLists")]
        public Output<ImmutableArray<int>> AiAsrInputIndexLists { get; private set; } = null!;

        /// <summary>
        /// If enable format diagnose.
        /// </summary>
        [Output("aiFormatDiagnose")]
        public Output<int?> AiFormatDiagnose { get; private set; } = null!;

        /// <summary>
        /// Ai ocr input index list(first input index is 1.).
        /// </summary>
        [Output("aiOcrInputIndexLists")]
        public Output<ImmutableArray<int>> AiOcrInputIndexLists { get; private set; } = null!;

        /// <summary>
        /// If store monitor event.
        /// </summary>
        [Output("allowMonitorReport")]
        public Output<int?> AllowMonitorReport { get; private set; } = null!;

        /// <summary>
        /// Asr language.0: close.1: Chinese2: English3: Japanese4: Korean.
        /// </summary>
        [Output("asrLanguage")]
        public Output<int?> AsrLanguage { get; private set; } = null!;

        /// <summary>
        /// If enable stream broken check.
        /// </summary>
        [Output("checkStreamBroken")]
        public Output<int?> CheckStreamBroken { get; private set; } = null!;

        /// <summary>
        /// If enable low frame rate check.
        /// </summary>
        [Output("checkStreamLowFrameRate")]
        public Output<int?> CheckStreamLowFrameRate { get; private set; } = null!;

        /// <summary>
        /// Wait monitor input info list.
        /// </summary>
        [Output("inputLists")]
        public Output<ImmutableArray<Outputs.StreamMonitorInputList>> InputLists { get; private set; } = null!;

        /// <summary>
        /// Monitor task name.
        /// </summary>
        [Output("monitorName")]
        public Output<string?> MonitorName { get; private set; } = null!;

        /// <summary>
        /// Monitor event notify policy.
        /// </summary>
        [Output("notifyPolicy")]
        public Output<Outputs.StreamMonitorNotifyPolicy?> NotifyPolicy { get; private set; } = null!;

        /// <summary>
        /// Intelligent text recognition language settings: ocr language.0: close.1. Chinese,English.
        /// </summary>
        [Output("ocrLanguage")]
        public Output<int?> OcrLanguage { get; private set; } = null!;

        /// <summary>
        /// Monitor task output info.
        /// </summary>
        [Output("outputInfo")]
        public Output<Outputs.StreamMonitorOutputInfo> OutputInfo { get; private set; } = null!;


        /// <summary>
        /// Create a StreamMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamMonitor(string name, StreamMonitorArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/streamMonitor:StreamMonitor", name, args ?? new StreamMonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamMonitor(string name, Input<string> id, StreamMonitorState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/streamMonitor:StreamMonitor", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/matrixorigin",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamMonitor Get(string name, Input<string> id, StreamMonitorState? state = null, CustomResourceOptions? options = null)
        {
            return new StreamMonitor(name, id, state, options);
        }
    }

    public sealed class StreamMonitorArgs : global::Pulumi.ResourceArgs
    {
        [Input("aiAsrInputIndexLists")]
        private InputList<int>? _aiAsrInputIndexLists;

        /// <summary>
        /// AI asr input index list.(first input index is 1.).
        /// </summary>
        public InputList<int> AiAsrInputIndexLists
        {
            get => _aiAsrInputIndexLists ?? (_aiAsrInputIndexLists = new InputList<int>());
            set => _aiAsrInputIndexLists = value;
        }

        /// <summary>
        /// If enable format diagnose.
        /// </summary>
        [Input("aiFormatDiagnose")]
        public Input<int>? AiFormatDiagnose { get; set; }

        [Input("aiOcrInputIndexLists")]
        private InputList<int>? _aiOcrInputIndexLists;

        /// <summary>
        /// Ai ocr input index list(first input index is 1.).
        /// </summary>
        public InputList<int> AiOcrInputIndexLists
        {
            get => _aiOcrInputIndexLists ?? (_aiOcrInputIndexLists = new InputList<int>());
            set => _aiOcrInputIndexLists = value;
        }

        /// <summary>
        /// If store monitor event.
        /// </summary>
        [Input("allowMonitorReport")]
        public Input<int>? AllowMonitorReport { get; set; }

        /// <summary>
        /// Asr language.0: close.1: Chinese2: English3: Japanese4: Korean.
        /// </summary>
        [Input("asrLanguage")]
        public Input<int>? AsrLanguage { get; set; }

        /// <summary>
        /// If enable stream broken check.
        /// </summary>
        [Input("checkStreamBroken")]
        public Input<int>? CheckStreamBroken { get; set; }

        /// <summary>
        /// If enable low frame rate check.
        /// </summary>
        [Input("checkStreamLowFrameRate")]
        public Input<int>? CheckStreamLowFrameRate { get; set; }

        [Input("inputLists", required: true)]
        private InputList<Inputs.StreamMonitorInputListArgs>? _inputLists;

        /// <summary>
        /// Wait monitor input info list.
        /// </summary>
        public InputList<Inputs.StreamMonitorInputListArgs> InputLists
        {
            get => _inputLists ?? (_inputLists = new InputList<Inputs.StreamMonitorInputListArgs>());
            set => _inputLists = value;
        }

        /// <summary>
        /// Monitor task name.
        /// </summary>
        [Input("monitorName")]
        public Input<string>? MonitorName { get; set; }

        /// <summary>
        /// Monitor event notify policy.
        /// </summary>
        [Input("notifyPolicy")]
        public Input<Inputs.StreamMonitorNotifyPolicyArgs>? NotifyPolicy { get; set; }

        /// <summary>
        /// Intelligent text recognition language settings: ocr language.0: close.1. Chinese,English.
        /// </summary>
        [Input("ocrLanguage")]
        public Input<int>? OcrLanguage { get; set; }

        /// <summary>
        /// Monitor task output info.
        /// </summary>
        [Input("outputInfo", required: true)]
        public Input<Inputs.StreamMonitorOutputInfoArgs> OutputInfo { get; set; } = null!;

        public StreamMonitorArgs()
        {
        }
        public static new StreamMonitorArgs Empty => new StreamMonitorArgs();
    }

    public sealed class StreamMonitorState : global::Pulumi.ResourceArgs
    {
        [Input("aiAsrInputIndexLists")]
        private InputList<int>? _aiAsrInputIndexLists;

        /// <summary>
        /// AI asr input index list.(first input index is 1.).
        /// </summary>
        public InputList<int> AiAsrInputIndexLists
        {
            get => _aiAsrInputIndexLists ?? (_aiAsrInputIndexLists = new InputList<int>());
            set => _aiAsrInputIndexLists = value;
        }

        /// <summary>
        /// If enable format diagnose.
        /// </summary>
        [Input("aiFormatDiagnose")]
        public Input<int>? AiFormatDiagnose { get; set; }

        [Input("aiOcrInputIndexLists")]
        private InputList<int>? _aiOcrInputIndexLists;

        /// <summary>
        /// Ai ocr input index list(first input index is 1.).
        /// </summary>
        public InputList<int> AiOcrInputIndexLists
        {
            get => _aiOcrInputIndexLists ?? (_aiOcrInputIndexLists = new InputList<int>());
            set => _aiOcrInputIndexLists = value;
        }

        /// <summary>
        /// If store monitor event.
        /// </summary>
        [Input("allowMonitorReport")]
        public Input<int>? AllowMonitorReport { get; set; }

        /// <summary>
        /// Asr language.0: close.1: Chinese2: English3: Japanese4: Korean.
        /// </summary>
        [Input("asrLanguage")]
        public Input<int>? AsrLanguage { get; set; }

        /// <summary>
        /// If enable stream broken check.
        /// </summary>
        [Input("checkStreamBroken")]
        public Input<int>? CheckStreamBroken { get; set; }

        /// <summary>
        /// If enable low frame rate check.
        /// </summary>
        [Input("checkStreamLowFrameRate")]
        public Input<int>? CheckStreamLowFrameRate { get; set; }

        [Input("inputLists")]
        private InputList<Inputs.StreamMonitorInputListGetArgs>? _inputLists;

        /// <summary>
        /// Wait monitor input info list.
        /// </summary>
        public InputList<Inputs.StreamMonitorInputListGetArgs> InputLists
        {
            get => _inputLists ?? (_inputLists = new InputList<Inputs.StreamMonitorInputListGetArgs>());
            set => _inputLists = value;
        }

        /// <summary>
        /// Monitor task name.
        /// </summary>
        [Input("monitorName")]
        public Input<string>? MonitorName { get; set; }

        /// <summary>
        /// Monitor event notify policy.
        /// </summary>
        [Input("notifyPolicy")]
        public Input<Inputs.StreamMonitorNotifyPolicyGetArgs>? NotifyPolicy { get; set; }

        /// <summary>
        /// Intelligent text recognition language settings: ocr language.0: close.1. Chinese,English.
        /// </summary>
        [Input("ocrLanguage")]
        public Input<int>? OcrLanguage { get; set; }

        /// <summary>
        /// Monitor task output info.
        /// </summary>
        [Input("outputInfo")]
        public Input<Inputs.StreamMonitorOutputInfoGetArgs>? OutputInfo { get; set; }

        public StreamMonitorState()
        {
        }
        public static new StreamMonitorState Empty => new StreamMonitorState();
    }
}