/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */


export function urlPathName(url: string) {
	if (!document || !url){
		return url
	}

	const el = document.createElement('a');
	el.href = url;
	return el.pathname
}