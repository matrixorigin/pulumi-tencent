// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse.Inputs
{

    public sealed class InstanceContainerGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("command")]
        public Input<string>? Command { get; set; }

        [Input("containerImage")]
        public Input<string>? ContainerImage { get; set; }

        [Input("containerName")]
        public Input<string>? ContainerName { get; set; }

        [Input("envs")]
        private InputList<Inputs.InstanceContainerEnvGetArgs>? _envs;
        public InputList<Inputs.InstanceContainerEnvGetArgs> Envs
        {
            get => _envs ?? (_envs = new InputList<Inputs.InstanceContainerEnvGetArgs>());
            set => _envs = value;
        }

        [Input("publishPorts")]
        private InputList<Inputs.InstanceContainerPublishPortGetArgs>? _publishPorts;
        public InputList<Inputs.InstanceContainerPublishPortGetArgs> PublishPorts
        {
            get => _publishPorts ?? (_publishPorts = new InputList<Inputs.InstanceContainerPublishPortGetArgs>());
            set => _publishPorts = value;
        }

        [Input("volumes")]
        private InputList<Inputs.InstanceContainerVolumeGetArgs>? _volumes;
        public InputList<Inputs.InstanceContainerVolumeGetArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.InstanceContainerVolumeGetArgs>());
            set => _volumes = value;
        }

        public InstanceContainerGetArgs()
        {
        }
        public static new InstanceContainerGetArgs Empty => new InstanceContainerGetArgs();
    }
}