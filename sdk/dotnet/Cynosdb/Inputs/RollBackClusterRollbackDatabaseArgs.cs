// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cynosdb.Inputs
{

    public sealed class RollBackClusterRollbackDatabaseArgs : global::Pulumi.ResourceArgs
    {
        [Input("newDatabase", required: true)]
        public Input<string> NewDatabase { get; set; } = null!;

        [Input("oldDatabase", required: true)]
        public Input<string> OldDatabase { get; set; } = null!;

        public RollBackClusterRollbackDatabaseArgs()
        {
        }
        public static new RollBackClusterRollbackDatabaseArgs Empty => new RollBackClusterRollbackDatabaseArgs();
    }
}