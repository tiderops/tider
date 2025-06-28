export interface INamespace {
	name: string
	version: number
	age: string
	status: string
	labels: Map<string, string>
}

export interface INode {
	name: string
	namespace: string
	memory: string
	cpu: string
	storage: string
	ephemeralStorage: string
	version: string
	age: string
	labels: Map<string, string>
}
