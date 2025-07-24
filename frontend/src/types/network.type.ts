export interface IService {
	name: string
	namespace: string
	labels: Map<string, string>
	status: string
	age: string
	spec: string
}

export interface IIngress {
	name: string
	namespace: string
	host: string[]
	path: string[]
	age: string
	labels: Map<string, string>
}

export interface IRules {
	host: string
	path: string
}

export interface IServiceDetail {
	name: string
	namespace: string
	labels: Map<string, string>
	status: string
	age: string
	spec: string
}
