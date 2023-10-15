/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

import type {Channel, Topic} from "$lib/graph/graphql";

export interface SocketMsg {
  channel: Channel;
  topic: Topic;
  title: string;
  message: string;
  data: any;
}
