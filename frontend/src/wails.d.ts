import { database } from '../wailsjs/go/models'

export {}

// TODO: Se agrego para Wails events
declare global {
	import CommonParameterDto = database.CommonParameterDto

	interface Window {
		runtime: {
			EventsOn(event: string, callback: (data: any) => void): void
			EventsEmit(event: string, data?: any): void
			Call(method: string, ...args: any[]): Promise<any>
		}
		go: {
			middleware: {
				ParameterMiddleware: {
					GetCommonParameters(): Promise<CommonParameterDto[]>
				}
			}
		}
	}
}
