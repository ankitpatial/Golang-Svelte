<script lang='ts'>
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import UserForm from '../UserForm.svelte';
  import { getContextClient } from '@urql/svelte';
  import { CreateUser } from '$lib/gql/user';
  import type UserInput from '$lib/models/UserInput';
  import alerts from '$lib/stores/alerts';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import { page } from '$app/stores';
  import { Role } from '$lib/graph/graphql';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();

  async function save(input: UserInput): Promise<void> {
    const title = 'Create User';
    const res = await client.mutation(CreateUser, { input }).toPromise();
    if (res.error) {
      alerts.push({ type: 'error', title, body: res.error.message });
      return;
    }

    alerts.push({ type: 'success', title, body: `successfully created new user ${input.firstName} ${input.lastName}` });
    toUsers();
  }

  function toUsers() {
    goto(PATH.USERS + $page.url.search);
  }
</script>

<Page title='Create User'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={toUsers} />
  </svelte:fragment>
  <div class='flex gap-4 flex-col sm:flex-row flex-wrap'>
    <CardWithBorderTitle title='User Detail'>
      <UserForm user={{ role: Role.SiteUser  }} saveHandler={save} />
    </CardWithBorderTitle>
  </div>
</Page>
