// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kubernetes.Inputs
{

    public sealed class ClusterAttachmentWorkerConfigOverridesGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataDisks")]
        private InputList<Inputs.ClusterAttachmentWorkerConfigOverridesDataDiskGetArgs>? _dataDisks;
        public InputList<Inputs.ClusterAttachmentWorkerConfigOverridesDataDiskGetArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.ClusterAttachmentWorkerConfigOverridesDataDiskGetArgs>());
            set => _dataDisks = value;
        }

        [Input("desiredPodNum")]
        public Input<int>? DesiredPodNum { get; set; }

        [Input("dockerGraphPath")]
        public Input<string>? DockerGraphPath { get; set; }

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;
        [Obsolete(@"This argument was no longer supported by TencentCloud TKE.")]
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        [Input("gpuArgs")]
        public Input<Inputs.ClusterAttachmentWorkerConfigOverridesGpuArgsGetArgs>? GpuArgs { get; set; }

        [Input("isSchedule")]
        public Input<bool>? IsSchedule { get; set; }

        [Input("mountTarget")]
        public Input<string>? MountTarget { get; set; }

        [Input("preStartUserScript")]
        public Input<string>? PreStartUserScript { get; set; }

        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public ClusterAttachmentWorkerConfigOverridesGetArgs()
        {
        }
        public static new ClusterAttachmentWorkerConfigOverridesGetArgs Empty => new ClusterAttachmentWorkerConfigOverridesGetArgs();
    }
}
