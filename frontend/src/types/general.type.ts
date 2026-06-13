export interface NamespaceRow {
	name: string
	age: string
	status: string
}

export interface NodeRow {
	name: string
	memory: number
	cpu: number
	storage: number
	ephemeralStorage: number
	kubeletVersion: string
	operatingSystem: string
	version: string
	age: string
}

export interface ResourceSummary {
	name: string
	namespace?: string
	age?: string
	status?: string
}
