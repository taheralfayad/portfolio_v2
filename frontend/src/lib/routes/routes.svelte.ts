type RouteFn = (...args: any[]) => string;
type RouteNode = RouteFn | { [key: string]: RouteNode };

export const routes = {
  Home: () => "/home",
  Books: () => "/books",
  Blog: () => "/blog",
  Coffee: Object.assign(() => "/coffee", {
    Charts: () => "/coffee/charts",
    SQL: (id: string) => `/coffee/sql`,
  }),
} as const satisfies Record<string, RouteNode>;
