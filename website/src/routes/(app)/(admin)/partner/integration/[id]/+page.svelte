<script lang='ts'>
  import { getContextClient } from '@urql/svelte';
  import DataLoader from '$lib/components/data/DataLoader';
  import { QryPartnerJobs } from '$lib/gql/job';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { page } from '$app/stores';
  import IntersectionObserver from '$lib/components/IntersectionObserver.svelte';
  import { humanize } from '$lib/utils/string';
  import PartnerJobView from '$lib/components/PartnerJobView.svelte';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  const loader = new DataLoader(client, QryPartnerJobs, 'partnerJobs', 20);
  const store = loader.store;
  const loading = loader.loading;

  let name, pType;
  $:{
    pType = $page.params.type;
    name = $page.url.searchParams.get('name') || 'Partner';
    loader.setParam('partnerID', $page.params.id);
    loader.setParam('search', '');
    loader.load();
  }


  function back() {
    $page.url.searchParams.delete('name');
    goto(`${PATH.PARTNER_INTEGRATION}${$page.url.search}`);
  }

</script>

<Page title='{humanize(pType)} Partner, {name}' subTitle='({$store.count})'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={back} />
  </svelte:fragment>

  <div class='flex sm:flex-row flex-col flex-wrap gap-4' class:blur={$loading}>
    {#each $store.items || [] as item}
      <PartnerJobView job={item} />
    {/each}
    {#if !$loading && $store?.hasNext}
      <br />
      <div class='flex justify-between'>
        <IntersectionObserver on:intersect={loader.loadMore}>
          <div class='btn btn-outline' class:loading={loading}>loading more users</div>
        </IntersectionObserver>
      </div>
    {/if}
  </div>
</Page>


