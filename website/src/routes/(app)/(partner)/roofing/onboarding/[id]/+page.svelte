<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts' xmlns:svele='http://www.w3.org/1999/html'>
  import { getContextClient } from '@urql/svelte';
  import { afterNavigate, goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { page } from '$app/stores';
  import { QryPartner } from '$lib/gql/partner';
  import { onDestroy, onMount } from 'svelte';
  import { PartnerContactType } from '$lib/models/partner';
  import { activeStep, initialValue, partner, steps } from '$lib/components/onboard/roofing';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  let busy = false;

  $: path = $page.url.pathname;

  afterNavigate((na) => {
    activeStep.set(Number(na.to?.url.searchParams.get('step') || 1));
  });

  onMount(() => {
    busy = true;
    client.query(QryPartner, { id: $page.params.id })
      .toPromise()
      .then((res) => {
        const p = res.data?.partner;
        if (p) {
          if (!p.contacts || p.contacts.length === 0) {
            p.contacts = [
              { type: PartnerContactType.Primary, firstName: '', lastName: '', email: '', phone: '' }
            ];
          } else {
            p.contacts.forEach((c, idx) => {
              if (!c.type) {
                c.type = Object.values(PartnerContactType)[idx];
              }
            });
          }
          partner.set(p);
        }
      })
      .finally(() => {
        busy = false;
      });
  });

  onDestroy(() => {
    partner.set(initialValue);
  });

  function back() {
    if ($activeStep === 1) {
      goto(PATH.PARTNERS);
    } else {
      $page.url.searchParams.set('step', `${$activeStep - 1}`);
      goto($page.url.pathname + $page.url.search);
    }
  }
</script>

<Page title='Partner Details'>
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


