export const routes = {
	Home: () => '/home',
	Coffee: () => '/coffee',
} as const satisfies Record<string, (...args: any[]) => string>;
