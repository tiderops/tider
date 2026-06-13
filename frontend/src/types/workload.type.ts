export interface PodRow {
	name: string
	namespace: string
	cpu: string
	memory: string
	storage: string
	node: string
	age: string
	status: string
}

export interface DeploymentRow {
	name: string
	namespace: string
	replicas: number
	age: string
	status: string
}

export interface PodDetail {
	name: string
	namespace: string
	image?: string
	pullPolicy?: string
	port?: number
	cpu: string
	memory: string
	storage: string
	age: string
	status: string
	editable: string[]
}

export interface DeploymentDetail {
	name: string
	namespace: string
	replicas: number
	age: string
	status: string
}
