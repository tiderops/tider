export interface Menu {
	name: string
	link: string
	icon: string
}

export interface ClusterLayout {
	name: string
	cluster: string
	status: boolean
	options: Menu[]
}
