// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Pts.Outputs
{

    [OutputType]
    public sealed class GetScenarioWithJobsScenarioWithJobsSetScenarioLoadResult
    {
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioLoadGeoRegionsLoadDistributionResult> GeoRegionsLoadDistributions;
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioLoadLoadSpecResult> LoadSpecs;
        public readonly ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioLoadVpcLoadDistributionResult> VpcLoadDistributions;

        [OutputConstructor]
        private GetScenarioWithJobsScenarioWithJobsSetScenarioLoadResult(
            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioLoadGeoRegionsLoadDistributionResult> geoRegionsLoadDistributions,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioLoadLoadSpecResult> loadSpecs,

            ImmutableArray<Outputs.GetScenarioWithJobsScenarioWithJobsSetScenarioLoadVpcLoadDistributionResult> vpcLoadDistributions)
        {
            GeoRegionsLoadDistributions = geoRegionsLoadDistributions;
            LoadSpecs = loadSpecs;
            VpcLoadDistributions = vpcLoadDistributions;
        }
    }
}
