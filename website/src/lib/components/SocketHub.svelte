<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { socket, wsMessage } from '$lib/stores/socket';
  import alerts from '$lib/stores/alerts';
  import { authUser } from '$lib/stores/auth';
  import { onDestroy } from 'svelte';
  import { Channel, Topic } from '$lib/graph/graphql';

  let host = import.meta.env.VITE_WS_HOST;
  let pingInterval, interval = 1000 * 60 * 2;

  const unsubscribe = wsMessage.subscribe(msg => {
    if (!msg) return;

    if (Channel.Task === msg.channel && Topic.Progress === msg.topic) {
      alerts.info(`${msg.title || msg.data.task}`, msg.message);
    }
  });

  onDestroy(unsubscribe);

  $: if ($authUser) {
    // console.log('** authUser', $authUser);
    socket.connect(host, $authUser?.token || '');
    pingInterval = setInterval(() => {
      socket.ping();
    }, interval);
  } else {
    // console.log('~~ authUser');
    socket.close('user session expired');
  }
</script>
