<script lang='ts'>
  import { getContextClient, queryStore } from '@urql/svelte';
  import { EditApiUser, QryApiUser } from '$lib/gql/apiUser';
  import { page } from '$app/stores';
  import Input from '$lib/components/form/Input.svelte';
  import DropDown from '$lib/components/form/DropDown.svelte';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import alerts from '$lib/stores/alerts';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  const store = queryStore<any>({
    client,
    query: QryApiUser,
    variables: { id: $page.params.id }
  });

  let busy = false, data;
  $: data = $store.data?.apiUser || { cbApiEndpoints: {} };


  async function save() {
    if (busy) return;

    const input = {
      id: $page.params.id,
      cbApiAuth: data.cbApiAuth,
      cbApiUrl: data.cbApiUrl,
      cbApiEndpoints: data.cbApiEndpoints
    };

    if (data.cbApiAuth === 'BASIC' || data.cbApiAuth === 'OAUTH') {
      input.cbApiUser = data.cbApiUser;
      input.cbApiPwd = data.cbApiPwd;
    }

    if (data.cbApiAuth === 'TOKEN') {
      input.cbApiToken = data.cbApiToken;
    }

    busy = true;
    const res = await client.mutation(EditApiUser, { input }).toPromise();
    busy = false;
    if (res.error) {
      alerts.error('Save Api User', res.error.message);
      return;
    }

    alerts.success('Save Api User', 'Saved Successfully');
  }
</script>

<Page title='Edit API User'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={() => goto(PATH.API_USERS)} tooltip='API Users' />
  </svelte:fragment>

  <div class='flex flex-col sm:flex-row gap-4 flex-wrap' class:blur-sm={$store.fetching}>

    <CardWithBorderTitle title='Api User'>
      <Input disabled label='Username' value={data.username} />
    </CardWithBorderTitle>

    <CardWithBorderTitle title='Callback API'>
      <Input label='API URL' bind:value={data.cbApiUrl} placeholder='https://api-domain.com/version/extra'>
        <svelte:fragment slot='desc'>
          do not end it with "/"
        </svelte:fragment>
      </Input>
      <DropDown label='API Auth' bind:value={data.cbApiAuth} options={[
        {id: 'NONE', name: 'None'},
        {id: 'BASIC', name: 'Basic Authentication'},
        {id: 'TOKEN', name: 'Api Token/Key'},
        {id: 'OAUTH', name: 'OAuth 2'},
      ]} />
      {#if (data.cbApiAuth === 'BASIC' || data.cbApiAuth === 'OAUTH')}
        <Input label='API User' bind:value={data.cbApiUser} />
        <Input label='API Password' bind:value={data.cbApiPwd} />
      {/if}

      {#if (data.cbApiAuth === 'TOKEN')}
        <Input label='API Token' bind:value={data.cbApiToken} />
      {/if}
    </CardWithBorderTitle>

    <CardWithBorderTitle title='Callback API Endpoints'>
      <Input label='On Estimate Complete' placeholder='/endpoint_name'
             bind:value={data.cbApiEndpoints.OnEstimateComplete}>
        <svelte:fragment slot='desc'>
          must start with "/"
        </svelte:fragment>
      </Input>
    </CardWithBorderTitle>
  </div>
  <button class='btn btn-primary mt-10' on:click={save}>
    SAVE
  </button>
</Page>
