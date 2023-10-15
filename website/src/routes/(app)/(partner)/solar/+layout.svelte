<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
  import {authBusy, authUser} from '$lib/stores/auth.js';
  import {goto} from '$app/navigation';
  import {page} from '$app/stores';
  import PlaceholderBasic from "$lib/components/placeholder/PlaceholderBasic.svelte";
  import {onMount} from "svelte";
  import {PartnerStatus, PartnerType} from "$lib/graph/graphql";

  let isOnboarding = false, isOnboardingDone = false;

  if ($authUser?.partner) {
    let status = $authUser.partner.status;
    isOnboarding = status === PartnerStatus.Onboarding;
    isOnboardingDone = status === PartnerStatus.OnboardingDone
  }

  onMount(() => {
    if (!$authUser?.partner) {
      return;
    }

    const pt = $authUser.partner.type?.toLowerCase()
    if (isOnboarding && !$page.route.id?.endsWith('/onboarding/[id]')) {
      goto(`/${pt}/onboarding/${$authUser?.partner?.id}`);
      return
    }

    if (isOnboardingDone && !$page.route.id?.endsWith('/onboarding/done')) {
      goto(`/${pt}/onboarding/done`);
      return
    }
  })

</script>

{#if ($authBusy)}
  <PlaceholderBasic/>
{:else if $authUser?.partner?.type === PartnerType.Solar}
  <slot></slot>
{:else}
  <p>
    Only a solar partner account has access to this page.
  </p>
{/if}


