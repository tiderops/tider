// TODO-17: Remove mock data
export const delay = (ms = 200): Promise<void> => new Promise((resolve) => setTimeout(resolve, ms))
