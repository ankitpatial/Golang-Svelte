/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */


export interface Pager {
	before?: string;
	after?: string;
	first?: number;
	last?: number;
}

export interface PageInfo {
	startCursor?: string;
	endCursor?: string;
	hasNextPage?: boolean;
	hasPreviousPage?: boolean;
}

export type ColumnData = {
	type: 'string' | 'html' | 'component',
	content: any
}