export interface IPod {
	name: string
	namespace: string
	replicas: number
	cpu: string
	memory: string
	storage: string
	age: string
	status: string
}

export interface IDeployment {
	name: string
	namespace: string
	age: string
	status: string
}
