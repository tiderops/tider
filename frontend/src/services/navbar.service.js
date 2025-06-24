import { GetK8sObjects } from '../../wailsjs/go/middleware/ParameterMiddleware'
export const fetchObjectsParameter = async () => GetK8sObjects()
