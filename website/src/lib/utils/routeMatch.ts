/**
 * regular expression to wrap words stating with ':' []
 * e.g. /some/:id/other => /some/[:id]/other
 */
const wrapRegEx = /(:\w+)/gi;

export function routeMatch(href: string, routeId: string, startsWith?: string): boolean {
	if (!href || !routeId) return false;

	routeId = routeId
		.replace('/(app)', '')
		.replace('/(admin)', '')
		.replace('/(user)', '')
		.replace('/(public)', '');

	const route = href.replace(wrapRegEx, '[$1]').replace(':', '');
	const s = startsWith || route;

	if (s === routeId) return true;

	return routeId.startsWith(`${s}/`);
}
