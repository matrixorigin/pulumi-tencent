// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetFunctionsFunctionImageConfigResult
    {
        public readonly string Args;
        public readonly string Command;
        public readonly bool ContainerImageAccelerate;
        public readonly string EntryPoint;
        public readonly int ImagePort;
        public readonly string ImageType;
        public readonly string ImageUri;
        public readonly string RegistryId;

        [OutputConstructor]
        private GetFunctionsFunctionImageConfigResult(
            string args,

            string command,

            bool containerImageAccelerate,

            string entryPoint,

            int imagePort,

            string imageType,

            string imageUri,

            string registryId)
        {
            Args = args;
            Command = command;
            ContainerImageAccelerate = containerImageAccelerate;
            EntryPoint = entryPoint;
            ImagePort = imagePort;
            ImageType = imageType;
            ImageUri = imageUri;
            RegistryId = registryId;
        }
    }
}