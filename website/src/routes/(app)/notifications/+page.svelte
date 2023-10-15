<script lang='ts'>
  import IntersectionObserver from '$lib/components/IntersectionObserver.svelte';
  import { getContextClient } from '@urql/svelte';
  import { QryMyNotifications } from '$lib/gql/notification';
  import DataLoader from '$lib/components/data/DataLoader';
  import Message from './Message.svelte';
  import { onMount } from 'svelte';
  import { wsMessage } from '$lib/stores/socket';
  import Page from '$lib/components/Page.svelte';

  const title = 'My Notifications';
  const client = getContextClient();
  const dataProp = 'myNotifications';

  const loader = new DataLoader(client, QryMyNotifications, 'myNotifications', 5);
  const store = loader.store;
  const loading = loader.loading;

  loader.load();

  onMount(() => {
    return wsMessage.subscribe(msg => {
      if (!msg) return;
      if (msg.data.isNotification) {
        loader.load(true);
      }
    });
  });

</script>


<Page {title}>
  <div class='flex flex-col gap-4 flex-wrap'>
    {#each $store.items as data}
      <Message {data} />
    {/each}
  </div>

  {#if !$loading && $store.hasNext}
    <br />
    <div class='flex justify-between'>
      <IntersectionObserver on:intersect={loader.loadMore}>
        <div class='btn btn-outline' class:loading={$loading}>loading more users</div>
      </IntersectionObserver>
    </div>
  {/if}
</Page>
