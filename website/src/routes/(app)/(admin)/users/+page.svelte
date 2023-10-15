<script lang='ts'>
  import { page } from '$app/stores';
  import DataLoader from '$lib/components/data/DataLoader';
  import { QryUsers } from '$lib/gql/user';
  import { getContextClient } from '@urql/svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import UserCard from '$lib/components/UserCard.svelte';
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import InputSearch from '$lib/components/form/InputSearch.svelte';
  import { Role } from '$lib/graph/graphql';
  import Page from '$lib/components/Page.svelte';
  import DataIntersection from '$lib/components/data/DataIntersection.svelte';

  const client = getContextClient();
  const loader = new DataLoader(client, QryUsers, 'users', 50);
  const store = loader.store;
  const loading = loader.loading;

  let username = '', pwd = '', inp, showModal = false, busy = false,
    role = $page.url.searchParams.get('role') || '',
    q = $page.url.searchParams.get('q') || '';

  $: {
    $page.url.searchParams.set('q', q);
    $page.url.searchParams.set('role', role);
    goto($page.url.search);

    loader.where = {
      role: role || undefined
    };
    if (q) {
      loader.where.and = {
        role: role || undefined,
        or: [
          { emailContainsFold: q },
          { firstNameContainsFold: q },
          { lastNameContainsFold: q }
        ]
      };
    }
    loader.load();
  }
</script>

<Page title='Users' subTitle='({$store.count})'>
  <svelte:fragment slot='buttons'>
    <BtnPlusSmall on:click={() => goto(PATH.USERS_CREATE+ $page.url.search)} />
    <InputSearch bind:value={q} />
  </svelte:fragment>
  <svelte:fragment slot='filters'>
    <select id='ddlStatus' class='select select-ghost select-sm' bind:value={role}>
      <option value=''>Role: Any</option>
      <option value={Role.Admin}>Role: Admin</option>
      <option value={Role.SiteUser}>Role: Site User</option>
    </select>
  </svelte:fragment>

  <div class='flex flex-col sm:flex-row gap-4 flex-wrap'>
    {#each $store.items as user}
      <UserCard {user} />
    {/each}
  </div>

  <DataIntersection {loader} />
</Page>
