/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

export interface Document {
	id: string;
	key: string;
	name: string;
	contentType?: string;
	contentSize: number;
	ready: boolean;
	uploadUrl: string;
	meta?: object;
}

export type  SignedUrlCB = (inpName: string, fileInfo: FileInfo) => Promise<Document>;

export  type  StatusChangeCB = (id: number, status: string, err: undefined | string | Error) => void;

export interface FileInfo {
	name: string;
	type: string;
	size: number;
}