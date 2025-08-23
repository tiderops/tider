export interface INamespace {
	name: string
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
	kubeletVersion: string
	operatingSystem: string
	version: string
	age: string
	labels: Map<string, string>
}
