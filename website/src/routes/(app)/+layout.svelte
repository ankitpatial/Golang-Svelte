<script lang='ts'>
  import {fade} from 'svelte/transition';
  import Header from '$lib/components/Header.svelte';
  import {authUser, menuCounts} from '$lib/stores/auth.js';
  import {goto} from '$app/navigation';
  import {PATH} from '$lib/enum';
  import LeftNav from './LeftNav.svelte';
  import Drawer from '$lib/components/Drawer.svelte';
  import {page} from '$app/stores';
  import SocketHub from '$lib/components/SocketHub.svelte';
  import {onMount} from 'svelte';
  import {wsMessage} from '$lib/stores/socket';
  import {getContextClient} from '@urql/svelte';
  import {QryCounts} from '$lib/gql/stats';
  import {Channel, JobStatus, Topic} from '$lib/graph/graphql';

  const client = getContextClient();

  let status, isOnboarding = false, isOnboardingDone = false;

  if (!$authUser) {
    goto(`${PATH.LOGIN}?returnURL=${$page.url.pathname}`);
  }

  if ($authUser?.partner) {
    status = $authUser.partner.status;
    isOnboarding = status === 'Onboarding';
    isOnboardingDone = status === 'OnboardingDone';
  }

  onMount(() => {
    load();
    return wsMessage.subscribe((msg) => {
      if (!msg) return;
      switch (true) {
        case Channel.Estimate === msg.channel && Topic.StatusChange === msg.topic:
          switch (msg.data.status) {
            case JobStatus.Estimated:
            case JobStatus.Approved:
              load();
          }
          break;

        case Channel.Job === msg.channel && Topic.StatusChange === msg.topic:
          switch (msg.data.status) {
            case JobStatus.Assigned:
            case JobStatus.Accepted:
              load();
          }
      }
    });
  });

  function load() {
    client.query(QryCounts, {}).then((res) => {
      if (res.error) return;
      // set menu counts
      menuCounts.set(res.data.counts);
    });
  }


</script>

<Drawer>
  <Header/>
  {#key ($page.route.id)}
    <div class='sm:px-6 px-4' in:fade>
      {#if ($authUser?.id)}
        <slot></slot>
      {/if}
    </div>
  {/key}
  <svelte:fragment slot='left'>
    {#if (!isOnboarding && !isOnboardingDone)}
      <div class='mt-10'>
        <LeftNav/>
      </div>
    {/if}
  </svelte:fragment>
</Drawer>

<SocketHub/>


