export var database;
(function (database) {
    class CommonParameterDto {
        Name;
        Link;
        Icon;
        static createFrom(source = {}) {
            return new CommonParameterDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Link = source["Link"];
            this.Icon = source["Icon"];
        }
    }
    database.CommonParameterDto = CommonParameterDto;
    class HeadParamsDto {
        Title;
        Key;
        Align;
        Sortable;
        static createFrom(source = {}) {
            return new HeadParamsDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Title = source["Title"];
            this.Key = source["Key"];
            this.Align = source["Align"];
            this.Sortable = source["Sortable"];
        }
    }
    database.HeadParamsDto = HeadParamsDto;
    class K8sObject {
        Name;
        Link;
        IsVisible;
        IsEditable;
        static createFrom(source = {}) {
            return new K8sObject(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Link = source["Link"];
            this.IsVisible = source["IsVisible"];
            this.IsEditable = source["IsEditable"];
        }
    }
    database.K8sObject = K8sObject;
    class ObjectType {
        Name;
        IsVisible;
        IsEditable;
        K8sObject;
        static createFrom(source = {}) {
            return new ObjectType(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.IsVisible = source["IsVisible"];
            this.IsEditable = source["IsEditable"];
            this.K8sObject = this.convertValues(source["K8sObject"], K8sObject);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    database.ObjectType = ObjectType;
})(database || (database = {}));
export var model;
(function (model) {
    class ClusterInfo {
        Name;
        Cluster;
        Server;
        User;
        Namespace;
        Status;
        Source;
        static createFrom(source = {}) {
            return new ClusterInfo(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Cluster = source["Cluster"];
            this.Server = source["Server"];
            this.User = source["User"];
            this.Namespace = source["Namespace"];
            this.Status = source["Status"];
            this.Source = source["Source"];
        }
    }
    model.ClusterInfo = ClusterInfo;
    class ConfigMap {
        Name;
        Namespace;
        UsedByPods;
        static createFrom(source = {}) {
            return new ConfigMap(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.UsedByPods = source["UsedByPods"];
        }
    }
    model.ConfigMap = ConfigMap;
    class Resource {
        Cpu;
        Memory;
        Storage;
        StorageEphemeral;
        static createFrom(source = {}) {
            return new Resource(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Cpu = source["Cpu"];
            this.Memory = source["Memory"];
            this.Storage = source["Storage"];
            this.StorageEphemeral = source["StorageEphemeral"];
        }
    }
    model.Resource = Resource;
    class Container {
        Limit;
        Request;
        static createFrom(source = {}) {
            return new Container(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Limit = this.convertValues(source["Limit"], Resource);
            this.Request = this.convertValues(source["Request"], Resource);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.Container = Container;
    class CronJob {
        Name;
        Namespace;
        JobNames;
        static createFrom(source = {}) {
            return new CronJob(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.JobNames = source["JobNames"];
        }
    }
    model.CronJob = CronJob;
    class Deployment {
        Name;
        Namespace;
        PodNames;
        static createFrom(source = {}) {
            return new Deployment(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.PodNames = source["PodNames"];
        }
    }
    model.Deployment = Deployment;
    class DeploymentDto {
        Name;
        Namespace;
        Status;
        Age;
        static createFrom(source = {}) {
            return new DeploymentDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Status = source["Status"];
            this.Age = source["Age"];
        }
    }
    model.DeploymentDto = DeploymentDto;
    class EnvironmentDto {
        Name;
        Description;
        Env;
        Status;
        static createFrom(source = {}) {
            return new EnvironmentDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Description = source["Description"];
            this.Env = source["Env"];
            this.Status = source["Status"];
        }
    }
    model.EnvironmentDto = EnvironmentDto;
    class Host {
        Path;
        Type;
        static createFrom(source = {}) {
            return new Host(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Path = source["Path"];
            this.Type = source["Type"];
        }
    }
    model.Host = Host;
    class Ingress {
        Name;
        Namespace;
        ServiceRefs;
        static createFrom(source = {}) {
            return new Ingress(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.ServiceRefs = source["ServiceRefs"];
        }
    }
    model.Ingress = Ingress;
    class RuleDto {
        Host;
        Path;
        static createFrom(source = {}) {
            return new RuleDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Host = source["Host"];
            this.Path = source["Path"];
        }
    }
    model.RuleDto = RuleDto;
    class IngressDto {
        Name;
        Namespace;
        Rules;
        Creation;
        Labels;
        static createFrom(source = {}) {
            return new IngressDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Rules = this.convertValues(source["Rules"], RuleDto);
            this.Creation = source["Creation"];
            this.Labels = source["Labels"];
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.IngressDto = IngressDto;
    class Job {
        Name;
        Namespace;
        PodNames;
        OwnerKind;
        OwnerName;
        static createFrom(source = {}) {
            return new Job(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.PodNames = source["PodNames"];
            this.OwnerKind = source["OwnerKind"];
            this.OwnerName = source["OwnerName"];
        }
    }
    model.Job = Job;
    class Local {
        Path;
        FSType;
        static createFrom(source = {}) {
            return new Local(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Path = source["Path"];
            this.FSType = source["FSType"];
        }
    }
    model.Local = Local;
    class NFS {
        Path;
        Server;
        static createFrom(source = {}) {
            return new NFS(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Path = source["Path"];
            this.Server = source["Server"];
        }
    }
    model.NFS = NFS;
    class Namespace {
        Name;
        PodNames;
        ServiceNames;
        static createFrom(source = {}) {
            return new Namespace(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.PodNames = source["PodNames"];
            this.ServiceNames = source["ServiceNames"];
        }
    }
    model.Namespace = Namespace;
    class NamespaceDto {
        Name;
        Version;
        CreationTime;
        Labels;
        Status;
        static createFrom(source = {}) {
            return new NamespaceDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Version = source["Version"];
            this.CreationTime = source["CreationTime"];
            this.Labels = source["Labels"];
            this.Status = source["Status"];
        }
    }
    model.NamespaceDto = NamespaceDto;
    class Node {
        Name;
        PodNames;
        static createFrom(source = {}) {
            return new Node(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.PodNames = source["PodNames"];
        }
    }
    model.Node = Node;
    class NodeDtoV2 {
        Name;
        Namespace;
        Resource;
        Version;
        CreationTimestamp;
        Labels;
        static createFrom(source = {}) {
            return new NodeDtoV2(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Resource = this.convertValues(source["Resource"], Resource);
            this.Version = source["Version"];
            this.CreationTimestamp = source["CreationTimestamp"];
            this.Labels = source["Labels"];
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.NodeDtoV2 = NodeDtoV2;
    class Secret {
        Name;
        Namespace;
        UsedByPods;
        static createFrom(source = {}) {
            return new Secret(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.UsedByPods = source["UsedByPods"];
        }
    }
    model.Secret = Secret;
    class PersistentVolume {
        Name;
        ClaimName;
        Namespace;
        static createFrom(source = {}) {
            return new PersistentVolume(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.ClaimName = source["ClaimName"];
            this.Namespace = source["Namespace"];
        }
    }
    model.PersistentVolume = PersistentVolume;
    class PersistentVolumeClaim {
        Name;
        Namespace;
        VolumeName;
        UsedByPods;
        static createFrom(source = {}) {
            return new PersistentVolumeClaim(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.VolumeName = source["VolumeName"];
            this.UsedByPods = source["UsedByPods"];
        }
    }
    model.PersistentVolumeClaim = PersistentVolumeClaim;
    class Service {
        Name;
        Namespace;
        Selector;
        PodNames;
        static createFrom(source = {}) {
            return new Service(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Selector = source["Selector"];
            this.PodNames = source["PodNames"];
        }
    }
    model.Service = Service;
    class ReplicaSet {
        Name;
        Namespace;
        Deployment;
        PodNames;
        static createFrom(source = {}) {
            return new ReplicaSet(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Deployment = source["Deployment"];
            this.PodNames = source["PodNames"];
        }
    }
    model.ReplicaSet = ReplicaSet;
    class Pod {
        Name;
        Namespace;
        Labels;
        NodeName;
        OwnerKind;
        OwnerName;
        Deployment;
        PVCNames;
        ConfigMapRefs;
        SecretRefs;
        ServiceRefs;
        static createFrom(source = {}) {
            return new Pod(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Labels = source["Labels"];
            this.NodeName = source["NodeName"];
            this.OwnerKind = source["OwnerKind"];
            this.OwnerName = source["OwnerName"];
            this.Deployment = source["Deployment"];
            this.PVCNames = source["PVCNames"];
            this.ConfigMapRefs = source["ConfigMapRefs"];
            this.SecretRefs = source["SecretRefs"];
            this.ServiceRefs = source["ServiceRefs"];
        }
    }
    model.Pod = Pod;
    class ObjectMapDto {
        ClusterInfo;
        Nodes;
        Namespaces;
        Pods;
        Deployments;
        ReplicaSets;
        Services;
        Ingresses;
        PersistentVolumeClaims;
        PersistentVolumes;
        ConfigMaps;
        Secrets;
        Jobs;
        CronJobs;
        static createFrom(source = {}) {
            return new ObjectMapDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.ClusterInfo = this.convertValues(source["ClusterInfo"], ClusterInfo);
            this.Nodes = this.convertValues(source["Nodes"], Node);
            this.Namespaces = this.convertValues(source["Namespaces"], Namespace);
            this.Pods = this.convertValues(source["Pods"], Pod);
            this.Deployments = this.convertValues(source["Deployments"], Deployment);
            this.ReplicaSets = this.convertValues(source["ReplicaSets"], ReplicaSet);
            this.Services = this.convertValues(source["Services"], Service);
            this.Ingresses = this.convertValues(source["Ingresses"], Ingress);
            this.PersistentVolumeClaims = this.convertValues(source["PersistentVolumeClaims"], PersistentVolumeClaim);
            this.PersistentVolumes = this.convertValues(source["PersistentVolumes"], PersistentVolume);
            this.ConfigMaps = this.convertValues(source["ConfigMaps"], ConfigMap);
            this.Secrets = this.convertValues(source["Secrets"], Secret);
            this.Jobs = this.convertValues(source["Jobs"], Job);
            this.CronJobs = this.convertValues(source["CronJobs"], CronJob);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.ObjectMapDto = ObjectMapDto;
    class VolumeClaimSpec {
        VolumeName;
        VolumeMode;
        AccessModes;
        DataSourceName;
        StorageClass;
        VolumeAttributesClassName;
        Limit;
        Request;
        static createFrom(source = {}) {
            return new VolumeClaimSpec(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.VolumeName = source["VolumeName"];
            this.VolumeMode = source["VolumeMode"];
            this.AccessModes = source["AccessModes"];
            this.DataSourceName = source["DataSourceName"];
            this.StorageClass = source["StorageClass"];
            this.VolumeAttributesClassName = source["VolumeAttributesClassName"];
            this.Limit = this.convertValues(source["Limit"], Resource);
            this.Request = this.convertValues(source["Request"], Resource);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.VolumeClaimSpec = VolumeClaimSpec;
    class PersistentVolumeClaimDto {
        Name;
        Namespace;
        CreationTimestamp;
        Labels;
        VolumeClaimSpec;
        static createFrom(source = {}) {
            return new PersistentVolumeClaimDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.CreationTimestamp = source["CreationTimestamp"];
            this.Labels = source["Labels"];
            this.VolumeClaimSpec = this.convertValues(source["VolumeClaimSpec"], VolumeClaimSpec);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.PersistentVolumeClaimDto = PersistentVolumeClaimDto;
    class VolumeSpec {
        Local;
        VolumeMode;
        AccessModes;
        StorageClass;
        VolumeAttributesClassName;
        PersistentVolumeReclaimPolicy;
        MountOptions;
        Capacity;
        Host;
        NFS;
        static createFrom(source = {}) {
            return new VolumeSpec(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Local = this.convertValues(source["Local"], Local);
            this.VolumeMode = source["VolumeMode"];
            this.AccessModes = source["AccessModes"];
            this.StorageClass = source["StorageClass"];
            this.VolumeAttributesClassName = source["VolumeAttributesClassName"];
            this.PersistentVolumeReclaimPolicy = source["PersistentVolumeReclaimPolicy"];
            this.MountOptions = source["MountOptions"];
            this.Capacity = this.convertValues(source["Capacity"], Resource);
            this.Host = this.convertValues(source["Host"], Host);
            this.NFS = this.convertValues(source["NFS"], NFS);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.VolumeSpec = VolumeSpec;
    class PersistentVolumeDto {
        Name;
        Namespace;
        CreationTimestamp;
        Labels;
        VolumeSpec;
        static createFrom(source = {}) {
            return new PersistentVolumeDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.CreationTimestamp = source["CreationTimestamp"];
            this.Labels = source["Labels"];
            this.VolumeSpec = this.convertValues(source["VolumeSpec"], VolumeSpec);
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.PersistentVolumeDto = PersistentVolumeDto;
    class PodDto {
        Name;
        Namespace;
        Replicas;
        Container;
        Age;
        Status;
        static createFrom(source = {}) {
            return new PodDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Replicas = source["Replicas"];
            this.Container = this.convertValues(source["Container"], Container);
            this.Age = source["Age"];
            this.Status = source["Status"];
        }
        convertValues(a, classs, asMap = false) {
            if (!a) {
                return a;
            }
            if (a.slice && a.map) {
                return a.map(elem => this.convertValues(elem, classs));
            }
            else if ("object" === typeof a) {
                if (asMap) {
                    for (const key of Object.keys(a)) {
                        a[key] = new classs(a[key]);
                    }
                    return a;
                }
                return new classs(a);
            }
            return a;
        }
    }
    model.PodDto = PodDto;
    class ServiceDto {
        Name;
        Namespace;
        Labels;
        Status;
        CreationTimestamp;
        Spec;
        static createFrom(source = {}) {
            return new ServiceDto(source);
        }
        constructor(source = {}) {
            if ('string' === typeof source)
                source = JSON.parse(source);
            this.Name = source["Name"];
            this.Namespace = source["Namespace"];
            this.Labels = source["Labels"];
            this.Status = source["Status"];
            this.CreationTimestamp = source["CreationTimestamp"];
            this.Spec = source["Spec"];
        }
    }
    model.ServiceDto = ServiceDto;
})(model || (model = {}));
