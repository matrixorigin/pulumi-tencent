// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Wedata.Inputs
{

    public sealed class IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaPropertyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaPropertyGetArgs()
        {
        }
        public static new IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaPropertyGetArgs Empty => new IntegrationTaskNodeNodeInfoNodeMappingSourceSchemaPropertyGetArgs();
    }
}