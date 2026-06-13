export interface ServiceRow {
	name: string
	namespace: string
	type: string
	intIp: string
	extIp: string
	port: number
	status: string
	age: string
}

export interface IngressRow {
	name: string
	namespace: string
	host: string[]
	path: string[]
	age: string
}
