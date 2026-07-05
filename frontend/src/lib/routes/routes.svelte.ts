export const routes = {
	home: () => '/home',
	coffee: () => '/coffee',
	admin: () => '/admin',
} as const satisfies Record<string, (...args: any[]) => string>;
