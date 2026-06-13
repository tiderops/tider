import { config } from '../wailsjs/go/models'

export {}

// TODO: Se agrego para Wails events
declare global {
	import CommonParameterDto = config.CommonParameterDto

	interface Window {
		runtime: {
			EventsOn(event: string, callback: (data: unknown) => void): void
			EventsEmit(event: string, data?: unknown): void
			Call(method: string, ...args: unknown[]): Promise<unknown>
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
