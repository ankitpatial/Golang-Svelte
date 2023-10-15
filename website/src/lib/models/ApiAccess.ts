/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import type { API_ACCESS_SOURCE } from '$lib/enum';

export default interface ApiAccess {
	id: string;
	name: API_ACCESS_SOURCE;
	url: string;
	username: string;
	password: string;
	key: string;
	secret: string;
}
