<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts' xmlns:svele='http://www.w3.org/1999/html'>
  import { getContextClient } from '@urql/svelte';
  import { afterNavigate } from '$app/navigation';
  import { page } from '$app/stores';
  import { onDestroy } from 'svelte';
  import { activeStep, initialValue, partner, steps } from '$lib/components/onboard/roofing';
  import { authUser } from '$lib/stores/auth';
  import { QryPartner } from '$lib/gql/partner';
  import alerts from '$lib/stores/alerts';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();

  let busy = false;
  $: path = $page.url.pathname;

  afterNavigate((na) => {
    activeStep.set(Number(na.to?.url.searchParams.get('step') || 1));
  });

  onDestroy(() => {
    partner.set(initialValue);
  });

  $: if ($authUser?.partner?.id) {
    busy = true;
    client.query(QryPartner, { id: $authUser.partner.id })
      .toPromise()
      .then((res) => {
        if (res.error) {
          alerts.error('Partner Info', res.error.message);
          return;
        }
        partner.set(res.data.partner);
      })
      .finally(() => {
        busy = false;
      });
  }
</script>

<Page title='Company Profile'>
  <ul class='steps mb-4' slot='titleContent'>
    {#each $steps as step, idx}
      <li
        class='step'
        class:step-primary={idx < $activeStep}
        data-content={idx < $partner.setupStepsCompleted? "✓" : "●"}
      >
        {#if idx <= $partner.setupStepsCompleted}
          <a href={`${path}?step=${idx+1}`} class='link'>{step.name}</a>
        {:else }
          {step.name}
        {/if}
      </li>
    {/each}
  </ul>

  <div class:blur-sm={busy}>
    <svelte:component this={$steps[$activeStep-1].component} />
  </div>
</Page>
