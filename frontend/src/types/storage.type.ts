export interface IPersistentVolumeClaim {
	name: string
	namespace: string
	storageClass: string
	size: string
	age: string
	status: string
	label: string
}

export interface IPersistentVolume {
	name: string
	namespace: string
	storageClass: string
	capacity: string
	claim: string
	age: string
	status: string
	label: string
}
