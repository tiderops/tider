export interface PersistentVolumeClaimRow {
	name: string
	namespace: string
	storageClass: string
	size: string
	age: string
	status: string
}

export interface PersistentVolumeRow {
	name: string
	namespace: string
	storageClass: string
	capacity: string
	claim: string
	age: string
	status: string
}
