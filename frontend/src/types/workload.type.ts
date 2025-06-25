export interface IPod {
	name: string
	namespace: string
	replicas: number
	container: IContainer
	age: string
	status: boolean
}

export interface IContainer {
	limit: IResources
	request: IResources
}

export interface IResources {
	cpu: string
	memory: string
}

export interface IDeployment {
	name: string
	namespace: string
	age: string
	status: boolean
}
