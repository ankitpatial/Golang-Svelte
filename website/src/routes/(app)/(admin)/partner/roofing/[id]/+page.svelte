<script lang='ts'>
  import {goto} from '$app/navigation';
  import BtnSmall from "$lib/components/form/BtnSmall.svelte";
  import {page} from '$app/stores';
  import PartnerJobsView from "$lib/components/partner/PartnerJobsView.svelte";
  import {PATH} from "$lib/enum";

  let
    partnerID = $page.params.id,
    baseURL = $page.url.pathname,
    name = $page.url.searchParams.get('name') || 'Partner',
    title = '';

  $:title = `${name} Jobs`;

  function back() {
    goto(`${PATH.PARTNER_ROOFING}${$page.url.search}`);
  }

  function profile() {
    goto(`${$page.url.pathname}/profile${$page.url.search}`);
  }

</script>

<PartnerJobsView {title} {partnerID} {baseURL} on:back={back}/>
