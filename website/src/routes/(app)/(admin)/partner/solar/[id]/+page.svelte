<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import { page } from '$app/stores';
  import PartnerJobsView from '$lib/components/partner/PartnerJobsView.svelte';

  const title = ($page.url.searchParams.get('name') || '')  +' Jobs';
  const partnerID = $page.params.id;
  const baseURL = $page.url.pathname;

  function back() {
    goto(`${PATH.PARTNER_SOLAR}${$page.url.search}`);
  }

</script>

<PartnerJobsView {title} {partnerID} {baseURL} on:back={back} />
