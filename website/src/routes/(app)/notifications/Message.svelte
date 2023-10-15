<script>
  import { fade } from 'svelte/transition';
  import { localDate } from '$lib/utils/time';
  import { onMount } from 'svelte';
  import { getContextClient } from '@urql/svelte';
  import { NotificationRead } from '$lib/gql/notification';
  import { unreadNotificationsCount } from '$lib/stores/app';

  export let data;

  const client = getContextClient();

  onMount(() => {
    if (data.unread) {
      markRead(data.id).then(success => {
        if (!success) return;
        data.unread = false;
        unreadNotificationsCount.update(p => (p > 1 ? p - 1 : 0));
      });
    }
  });

  async function markRead(messageID) {
    const res = await client.mutation(NotificationRead, { messageID });
    return res.data?.notificationRead || false;
  }

</script>

<div class='card card-compact w-full md:w-8/12' transition:fade>
  <div class='card-body'>
    <div class='text-xl' class:font-bold={data.unread}>{data.title}</div>
    <p class='ml-2 whitespace-pre'>
      {data.message}
    </p>
    <small class='ml-2'>{data.from}, {localDate(data.createdAt)}</small>
  </div>
</div>
