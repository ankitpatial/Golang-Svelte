/*
 * Copyright (c)  2022. All rights reserved
 * Author: Ankit Patial<ankit@simsaw.com>
 */


import type { StatusChangeCB } from './d';

export async function uploadFile(
  id: number, f: File, inpName: string, url: string, meta: object, cbStatusChange: StatusChangeCB
) {
	const isS3Upload = url.startsWith('https://s3.');
	const method = isS3Upload ? 'PUT' : 'POST';
	const xhr = newXHR(method, url, !isS3Upload);
	let body;
	if (isS3Upload) {
		body = f;
		xhr.setRequestHeader('Content-Type', f.type);
		for (const [key, value] of Object.entries(meta || {})) {
			xhr.setRequestHeader(`x-amz-meta-${key}`, value);
		}
	} else {
		body = new FormData();
		body.append('file', f);
	}

	try {
		await upload(xhr, body, id, cbStatusChange);
		cbStatusChange(id, 'completed', undefined);
	} catch (ex: any) {
		cbStatusChange(id, 'error', ex);
	}
}

function upload(
	req: XMLHttpRequest, body: Document | XMLHttpRequestBodyInit | null, fileID: number, cbStatusChange: StatusChangeCB
): Promise<void> {
	return new Promise((resolve, reject) => {
		req.onreadystatechange = function() {
			if (req.readyState === 4 && req.status !== 200) {
				reject(req.responseText || 'Failed with status = ' + req.status);
			}
		};

		req.onload = function() {
			if (req.status === 200) {
				resolve();
			} else {
				reject(`status: ${req.status}, ${req.responseText}`);
			}
		};

		// handle error
		req.upload.onerror = function() {
			reject(`status: ${req.status}, ${req.responseText}`);
		};

		req.upload.ontimeout = function() {
			reject('timeout error');
		};

		req.upload.onprogress = function(e) {
			if (!e.lengthComputable) return;

			const progress = Math.round((e.loaded / e.total) * 100);
			const status = progress === 100 ? 'finalizing' : `uploading ${progress}%`;
			cbStatusChange(fileID, status, undefined);
		};

		try {
			req.send(body);
			cbStatusChange(fileID, 'uploading', undefined);
		} catch (ex) {
			reject(ex);
		}
	});
}

function newXHR(method: string, url: string, withCredentials = false) {
	const xhr = new XMLHttpRequest();
	if (xhr.withCredentials !== null && withCredentials) {
		xhr.withCredentials = true;
	}

	xhr.open(method, url, true);
	return xhr;
}


