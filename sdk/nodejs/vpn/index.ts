// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ConnectionArgs, ConnectionState } from "./connection";
export type Connection = import("./connection").Connection;
export const Connection: typeof import("./connection").Connection = null as any;
utilities.lazyLoad(exports, ["Connection"], () => require("./connection"));

export { ConnectionResetArgs, ConnectionResetState } from "./connectionReset";
export type ConnectionReset = import("./connectionReset").ConnectionReset;
export const ConnectionReset: typeof import("./connectionReset").ConnectionReset = null as any;
utilities.lazyLoad(exports, ["ConnectionReset"], () => require("./connectionReset"));

export { CustomerGatewayArgs, CustomerGatewayState } from "./customerGateway";
export type CustomerGateway = import("./customerGateway").CustomerGateway;
export const CustomerGateway: typeof import("./customerGateway").CustomerGateway = null as any;
utilities.lazyLoad(exports, ["CustomerGateway"], () => require("./customerGateway"));

export { CustomerGatewayConfigurationDownloadArgs, CustomerGatewayConfigurationDownloadState } from "./customerGatewayConfigurationDownload";
export type CustomerGatewayConfigurationDownload = import("./customerGatewayConfigurationDownload").CustomerGatewayConfigurationDownload;
export const CustomerGatewayConfigurationDownload: typeof import("./customerGatewayConfigurationDownload").CustomerGatewayConfigurationDownload = null as any;
utilities.lazyLoad(exports, ["CustomerGatewayConfigurationDownload"], () => require("./customerGatewayConfigurationDownload"));

export { GatewayArgs, GatewayState } from "./gateway";
export type Gateway = import("./gateway").Gateway;
export const Gateway: typeof import("./gateway").Gateway = null as any;
utilities.lazyLoad(exports, ["Gateway"], () => require("./gateway"));

export { GatewayCcnRoutesArgs, GatewayCcnRoutesState } from "./gatewayCcnRoutes";
export type GatewayCcnRoutes = import("./gatewayCcnRoutes").GatewayCcnRoutes;
export const GatewayCcnRoutes: typeof import("./gatewayCcnRoutes").GatewayCcnRoutes = null as any;
utilities.lazyLoad(exports, ["GatewayCcnRoutes"], () => require("./gatewayCcnRoutes"));

export { GatewayRouteArgs, GatewayRouteState } from "./gatewayRoute";
export type GatewayRoute = import("./gatewayRoute").GatewayRoute;
export const GatewayRoute: typeof import("./gatewayRoute").GatewayRoute = null as any;
utilities.lazyLoad(exports, ["GatewayRoute"], () => require("./gatewayRoute"));

export { GatewaySslClientCertArgs, GatewaySslClientCertState } from "./gatewaySslClientCert";
export type GatewaySslClientCert = import("./gatewaySslClientCert").GatewaySslClientCert;
export const GatewaySslClientCert: typeof import("./gatewaySslClientCert").GatewaySslClientCert = null as any;
utilities.lazyLoad(exports, ["GatewaySslClientCert"], () => require("./gatewaySslClientCert"));

export { GetConnectionsArgs, GetConnectionsResult, GetConnectionsOutputArgs } from "./getConnections";
export const getConnections: typeof import("./getConnections").getConnections = null as any;
export const getConnectionsOutput: typeof import("./getConnections").getConnectionsOutput = null as any;
utilities.lazyLoad(exports, ["getConnections","getConnectionsOutput"], () => require("./getConnections"));

export { GetCustomerGatewayVendorsArgs, GetCustomerGatewayVendorsResult, GetCustomerGatewayVendorsOutputArgs } from "./getCustomerGatewayVendors";
export const getCustomerGatewayVendors: typeof import("./getCustomerGatewayVendors").getCustomerGatewayVendors = null as any;
export const getCustomerGatewayVendorsOutput: typeof import("./getCustomerGatewayVendors").getCustomerGatewayVendorsOutput = null as any;
utilities.lazyLoad(exports, ["getCustomerGatewayVendors","getCustomerGatewayVendorsOutput"], () => require("./getCustomerGatewayVendors"));

export { GetCustomerGatewaysArgs, GetCustomerGatewaysResult, GetCustomerGatewaysOutputArgs } from "./getCustomerGateways";
export const getCustomerGateways: typeof import("./getCustomerGateways").getCustomerGateways = null as any;
export const getCustomerGatewaysOutput: typeof import("./getCustomerGateways").getCustomerGatewaysOutput = null as any;
utilities.lazyLoad(exports, ["getCustomerGateways","getCustomerGatewaysOutput"], () => require("./getCustomerGateways"));

export { GetDefaultHealthCheckIpArgs, GetDefaultHealthCheckIpResult, GetDefaultHealthCheckIpOutputArgs } from "./getDefaultHealthCheckIp";
export const getDefaultHealthCheckIp: typeof import("./getDefaultHealthCheckIp").getDefaultHealthCheckIp = null as any;
export const getDefaultHealthCheckIpOutput: typeof import("./getDefaultHealthCheckIp").getDefaultHealthCheckIpOutput = null as any;
utilities.lazyLoad(exports, ["getDefaultHealthCheckIp","getDefaultHealthCheckIpOutput"], () => require("./getDefaultHealthCheckIp"));

export { GetGatewayRoutesArgs, GetGatewayRoutesResult, GetGatewayRoutesOutputArgs } from "./getGatewayRoutes";
export const getGatewayRoutes: typeof import("./getGatewayRoutes").getGatewayRoutes = null as any;
export const getGatewayRoutesOutput: typeof import("./getGatewayRoutes").getGatewayRoutesOutput = null as any;
utilities.lazyLoad(exports, ["getGatewayRoutes","getGatewayRoutesOutput"], () => require("./getGatewayRoutes"));

export { GetGatewaysArgs, GetGatewaysResult, GetGatewaysOutputArgs } from "./getGateways";
export const getGateways: typeof import("./getGateways").getGateways = null as any;
export const getGatewaysOutput: typeof import("./getGateways").getGatewaysOutput = null as any;
utilities.lazyLoad(exports, ["getGateways","getGatewaysOutput"], () => require("./getGateways"));

export { SslClientArgs, SslClientState } from "./sslClient";
export type SslClient = import("./sslClient").SslClient;
export const SslClient: typeof import("./sslClient").SslClient = null as any;
utilities.lazyLoad(exports, ["SslClient"], () => require("./sslClient"));

export { SslServerArgs, SslServerState } from "./sslServer";
export type SslServer = import("./sslServer").SslServer;
export const SslServer: typeof import("./sslServer").SslServer = null as any;
utilities.lazyLoad(exports, ["SslServer"], () => require("./sslServer"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Vpn/connection:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/connectionReset:ConnectionReset":
                return new ConnectionReset(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/customerGateway:CustomerGateway":
                return new CustomerGateway(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/customerGatewayConfigurationDownload:CustomerGatewayConfigurationDownload":
                return new CustomerGatewayConfigurationDownload(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/gateway:Gateway":
                return new Gateway(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/gatewayCcnRoutes:GatewayCcnRoutes":
                return new GatewayCcnRoutes(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/gatewayRoute:GatewayRoute":
                return new GatewayRoute(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/gatewaySslClientCert:GatewaySslClientCert":
                return new GatewaySslClientCert(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/sslClient:SslClient":
                return new SslClient(name, <any>undefined, { urn })
            case "tencentcloud:Vpn/sslServer:SslServer":
                return new SslServer(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/connection", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/connectionReset", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/customerGateway", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/customerGatewayConfigurationDownload", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/gateway", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/gatewayCcnRoutes", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/gatewayRoute", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/gatewaySslClientCert", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/sslClient", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Vpn/sslServer", _module)