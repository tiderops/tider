export declare namespace database {
    class CommonParameterDto {
        Name: string;
        Link: string;
        Icon: string;
        static createFrom(source?: any): CommonParameterDto;
        constructor(source?: any);
    }
    class HeadParamsDto {
        Title: string;
        Key: string;
        Align: string;
        Sortable: boolean;
        static createFrom(source?: any): HeadParamsDto;
        constructor(source?: any);
    }
    class K8sObject {
        Name: string;
        Link: string;
        IsVisible: boolean;
        IsEditable: boolean;
        static createFrom(source?: any): K8sObject;
        constructor(source?: any);
    }
    class ObjectType {
        Name: string;
        IsVisible: boolean;
        IsEditable: boolean;
        K8sObject: K8sObject[];
        static createFrom(source?: any): ObjectType;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
}
export declare namespace model {
    class ClusterInfo {
        Name: string;
        Cluster: string;
        Server: string;
        User: string;
        Namespace: string;
        Status: boolean;
        Source: string;
        static createFrom(source?: any): ClusterInfo;
        constructor(source?: any);
    }
    class ConfigMap {
        Name: string;
        Namespace: string;
        UsedByPods: string[];
        static createFrom(source?: any): ConfigMap;
        constructor(source?: any);
    }
    class Resource {
        Cpu: string;
        Memory: string;
        Storage: string;
        StorageEphemeral: string;
        static createFrom(source?: any): Resource;
        constructor(source?: any);
    }
    class Container {
        Limit: Resource;
        Request: Resource;
        static createFrom(source?: any): Container;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class CronJob {
        Name: string;
        Namespace: string;
        JobNames: string[];
        static createFrom(source?: any): CronJob;
        constructor(source?: any);
    }
    class Deployment {
        Name: string;
        Namespace: string;
        PodNames: string[];
        static createFrom(source?: any): Deployment;
        constructor(source?: any);
    }
    class DeploymentDto {
        Name: string;
        Namespace: string;
        Status: string;
        Age: string;
        static createFrom(source?: any): DeploymentDto;
        constructor(source?: any);
    }
    class EnvironmentDto {
        Name: string;
        Description: string;
        Env: string;
        Status: boolean;
        static createFrom(source?: any): EnvironmentDto;
        constructor(source?: any);
    }
    class Host {
        Path: string;
        Type: string;
        static createFrom(source?: any): Host;
        constructor(source?: any);
    }
    class Ingress {
        Name: string;
        Namespace: string;
        ServiceRefs: string[];
        static createFrom(source?: any): Ingress;
        constructor(source?: any);
    }
    class RuleDto {
        Host: string;
        Path: string;
        static createFrom(source?: any): RuleDto;
        constructor(source?: any);
    }
    class IngressDto {
        Name: string;
        Namespace: string;
        Rules: RuleDto[];
        Creation: string;
        Labels: Record<string, string>;
        static createFrom(source?: any): IngressDto;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class Job {
        Name: string;
        Namespace: string;
        PodNames: string[];
        OwnerKind: string;
        OwnerName: string;
        static createFrom(source?: any): Job;
        constructor(source?: any);
    }
    class Local {
        Path: string;
        FSType: string;
        static createFrom(source?: any): Local;
        constructor(source?: any);
    }
    class NFS {
        Path: string;
        Server: string;
        static createFrom(source?: any): NFS;
        constructor(source?: any);
    }
    class Namespace {
        Name: string;
        PodNames: string[];
        ServiceNames: string[];
        static createFrom(source?: any): Namespace;
        constructor(source?: any);
    }
    class NamespaceDto {
        Name: string;
        Version: string;
        CreationTime: string;
        Labels: Record<string, string>;
        Status: string;
        static createFrom(source?: any): NamespaceDto;
        constructor(source?: any);
    }
    class Node {
        Name: string;
        PodNames: string[];
        static createFrom(source?: any): Node;
        constructor(source?: any);
    }
    class NodeDtoV2 {
        Name: string;
        Namespace: string;
        Resource: Resource;
        Version: string;
        CreationTimestamp: string;
        Labels: Record<string, string>;
        static createFrom(source?: any): NodeDtoV2;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class Secret {
        Name: string;
        Namespace: string;
        UsedByPods: string[];
        static createFrom(source?: any): Secret;
        constructor(source?: any);
    }
    class PersistentVolume {
        Name: string;
        ClaimName: string;
        Namespace: string;
        static createFrom(source?: any): PersistentVolume;
        constructor(source?: any);
    }
    class PersistentVolumeClaim {
        Name: string;
        Namespace: string;
        VolumeName: string;
        UsedByPods: string[];
        static createFrom(source?: any): PersistentVolumeClaim;
        constructor(source?: any);
    }
    class Service {
        Name: string;
        Namespace: string;
        Selector: Record<string, string>;
        PodNames: string[];
        static createFrom(source?: any): Service;
        constructor(source?: any);
    }
    class ReplicaSet {
        Name: string;
        Namespace: string;
        Deployment: string;
        PodNames: string[];
        static createFrom(source?: any): ReplicaSet;
        constructor(source?: any);
    }
    class Pod {
        Name: string;
        Namespace: string;
        Labels: Record<string, string>;
        NodeName: string;
        OwnerKind: string;
        OwnerName: string;
        Deployment: string;
        PVCNames: string[];
        ConfigMapRefs: string[];
        SecretRefs: string[];
        ServiceRefs: string[];
        static createFrom(source?: any): Pod;
        constructor(source?: any);
    }
    class ObjectMapDto {
        ClusterInfo: ClusterInfo;
        Nodes: Node[];
        Namespaces: Namespace[];
        Pods: Pod[];
        Deployments: Deployment[];
        ReplicaSets: ReplicaSet[];
        Services: Service[];
        Ingresses: Ingress[];
        PersistentVolumeClaims: PersistentVolumeClaim[];
        PersistentVolumes: PersistentVolume[];
        ConfigMaps: ConfigMap[];
        Secrets: Secret[];
        Jobs: Job[];
        CronJobs: CronJob[];
        static createFrom(source?: any): ObjectMapDto;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class VolumeClaimSpec {
        VolumeName: string;
        VolumeMode: string;
        AccessModes: string[];
        DataSourceName: string;
        StorageClass: string;
        VolumeAttributesClassName: string;
        Limit: Resource;
        Request: Resource;
        static createFrom(source?: any): VolumeClaimSpec;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class PersistentVolumeClaimDto {
        Name: string;
        Namespace: string;
        CreationTimestamp: string;
        Labels: Record<string, string>;
        VolumeClaimSpec: VolumeClaimSpec;
        static createFrom(source?: any): PersistentVolumeClaimDto;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class VolumeSpec {
        Local: Local;
        VolumeMode: string;
        AccessModes: string[];
        StorageClass: string;
        VolumeAttributesClassName: string;
        PersistentVolumeReclaimPolicy: string;
        MountOptions: string[];
        Capacity: Resource;
        Host: Host;
        NFS: NFS;
        static createFrom(source?: any): VolumeSpec;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class PersistentVolumeDto {
        Name: string;
        Namespace: string;
        CreationTimestamp: string;
        Labels: Record<string, string>;
        VolumeSpec: VolumeSpec;
        static createFrom(source?: any): PersistentVolumeDto;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class PodDto {
        Name: string;
        Namespace: string;
        Replicas: number;
        Container: Container;
        Age: string;
        Status: string;
        static createFrom(source?: any): PodDto;
        constructor(source?: any);
        convertValues(a: any, classs: any, asMap?: boolean): any;
    }
    class ServiceDto {
        Name: string;
        Namespace: string;
        Labels: Record<string, string>;
        Status: string;
        CreationTimestamp: string;
        Spec: string;
        static createFrom(source?: any): ServiceDto;
        constructor(source?: any);
    }
}
