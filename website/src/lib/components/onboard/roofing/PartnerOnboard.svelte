<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import {onDestroy, onMount} from 'svelte';
  import {goto} from '$app/navigation';
  import {getContextClient} from '@urql/svelte';
  import {page} from '$app/stores';
  import {QryPartner} from '$lib/gql/partner';
  import {PartnerContactType} from '$lib/models/partner';
  import {activeStep, initialValue, partner, steps} from './index';
  import {PartnerType} from "$lib/graph/graphql";

  export let step: number;
  export let returnURL = '';

  const client = getContextClient();
  const partnerType = PartnerType.Roofing;

  let busy = false;

  $: path = $page.url.pathname;
  $: activeStep.set(step);
  $: if (partnerType) {
    partner.update(p => {
      p.type = partnerType;
      return p;
    });
  }

  onMount(() => {
    busy = true;
    client.query(QryPartner, {id: $page.params.id})
      .toPromise()
      .then((res) => {
        const p = res.data?.partner;
        if (p) {
          if (!p.contacts || p.contacts.length === 0) {
            p.contacts = [
              {type: PartnerContactType.Primary, firstName: '', lastName: '', email: '', phone: ''}
            ];
          } else {
            p.contacts.forEach((c, idx) => {
              if (!c.type) {
                c.type = Object.values(PartnerContactType)[idx];
              }
            });
          }
          partner.set(p);
        } else {
          partner.set({...initialValue, type: partnerType})
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
      if (!returnURL) return;

      $page.url.searchParams.delete('step')
      goto(returnURL + $page.url.search);
    } else {
      $page.url.searchParams.set('step', `${$activeStep - 1}`);
      goto($page.url.pathname + $page.url.search);
    }
  }
</script>

<div class:blur-sm={busy}>
  <svelte:component this={$steps[$activeStep-1].component}/>
</div>


