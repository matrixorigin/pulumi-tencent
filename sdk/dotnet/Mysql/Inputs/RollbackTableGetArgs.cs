// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Mysql.Inputs
{

    public sealed class RollbackTableGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("tables", required: true)]
        private InputList<Inputs.RollbackTableTableGetArgs>? _tables;
        public InputList<Inputs.RollbackTableTableGetArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.RollbackTableTableGetArgs>());
            set => _tables = value;
        }

        public RollbackTableGetArgs()
        {
        }
        public static new RollbackTableGetArgs Empty => new RollbackTableGetArgs();
    }
}