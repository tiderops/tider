package model

type PersistentVolumeDto struct {
	Name         string
	Namespace    string
	StorageClass string
	Capacity     string
	Claim        string
	Labels       map[string]string
	VolumeSpec   VolumeSpec
	Age          string
	Status       string
}

type VolumeSpec struct {
	Local                         Local
	VolumeMode                    string
	AccessModes                   []string
	StorageClass                  string
	VolumeAttributesClassName     string
	PersistentVolumeReclaimPolicy string
	MountOptions                  []string
	Capacity                      Resource
	Host                          Host
	NFS                           NFS
}
type Local struct {
	Path   string
	FSType string
}
type Host struct {
	Path string
	Type string
}
type NFS struct {
	Path   string
	Server string
}

type PersistentVolumeClaimDto struct {
	Name            string
	Namespace       string
	storageClass    string
	size            string
	Labels          map[string]string
	VolumeClaimSpec VolumeClaimSpec
	Status          string
	Capacity        Resource
	Age             string
}

type VolumeClaimSpec struct {
	VolumeName                string
	VolumeMode                string
	AccessModes               []string
	DataSourceName            string
	DataSourceRef             string
	StorageClass              string
	VolumeAttributesClassName string
	Limit                     Resource
	Request                   Resource
}
