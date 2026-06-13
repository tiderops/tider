import { test, expect } from '@playwright/test'

// The app runs outside Wails here, so window.go is stubbed with the
// minimal binding surface the visited pages call.
const goStub = `
window.go = {
	binding: {
		Environment: {
			GetClusters: async () => [
				{ Name: 'minikube', Cluster: 'minikube-cluster', Server: '', User: '', Namespace: '', Status: true, Source: '' },
			],
		},
		Parameter: {
			GetKubernetesParameters: async () => [{ Name: 'Workload', Link: 'workload', Icon: '' }],
			GetCommonParameters: async () => [{ Name: 'Settings', Link: 'settings', Icon: '' }],
			GetK8sObjects: async () => [
				{
					Name: 'Workload',
					IsVisible: true,
					IsEditable: true,
					K8sObject: [{ Name: 'Pod', Link: 'pod', IsVisible: true, IsEditable: true }],
				},
			],
			GetHeaderParams: async () => [
				{ Title: 'Name', Key: 'name', Align: 'start', Sortable: true },
				{ Title: 'Namespace', Key: 'namespace', Align: 'start', Sortable: false },
				{ Title: 'Status', Key: 'status', Align: 'start', Sortable: false },
				{ Title: 'Actions', Key: 'actions', Align: 'start', Sortable: false },
			],
		},
		Workload: {
			GetPods: async () => [
				{
					Name: 'web-1',
					Namespace: 'default',
					Containers: [],
					Node: 'node-1',
					Age: '3h',
					CreatedAt: 0,
					Status: 'Running',
					Editable: [],
					Labels: {},
				},
			],
		},
		General: {
			GetNamespaces: async () => [{ Name: 'default', Version: '', Age: '1d', CreatedAt: 0, Labels: {}, Status: 'Active' }],
		},
	},
}
`

test.beforeEach(async ({ page }) => {
	await page.addInitScript(goStub)
})

test('root redirects to the home page', async ({ page }) => {
	await page.goto('/')
	await expect(page).toHaveURL(/\/home$/)
	await expect(page.locator('h2')).toHaveText('Welcome to Kubexplorer')
})

test('pod grid lists pods from the backend', async ({ page }) => {
	await page.goto('/minikube/workload/pod')

	await expect(page.getByText('web-1')).toBeVisible()
	await expect(page.getByText('Running')).toBeVisible()
})

test('unknown routes render the 404 page', async ({ page }) => {
	await page.goto('/totally/unknown/route/here')
	await expect(page.getByText('Page not found')).toBeVisible()
})

test('sidebar lists clusters by context name', async ({ page }) => {
	await page.goto('/home')
	await expect(page.getByText('minikube', { exact: false }).first()).toBeVisible()
})
