<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import UserForm from '../../UserForm.svelte';
  import { getContextClient } from '@urql/svelte';
  import { QryUser, UpdateUser } from '$lib/gql/user';
  import type UserInput from '$lib/models/UserInput';
  import alerts from '$lib/stores/alerts';
  import BtnBackSmall from '$lib/components/form/BtnBackSmall.svelte';
  import CardWithBorderTitle from '$lib/components/CardWithBorderTitle.svelte';
  import UserNotifications from '$lib/components/UserNotifications.svelte';
  import { Role } from '$lib/graph/graphql';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  let user = { role: Role.SiteUser }, busy = true;

  client
    .query(QryUser, { id: $page.params.id })
    .toPromise()
    .then((res) => {
      busy = false;
      if (res.error) {
        alerts.push({ type: 'error', title: 'Get User Info', body: res.error.message });
        return;
      }
      user = res.data.user;
    });

  async function save(input: UserInput): Promise<void> {
    const title = 'Create User';
    const res = await client.mutation(UpdateUser, {
      input: {
        id: input.id,
        firstName: input.firstName,
        lastName: input.lastName,
        phone: input.phone,
        role: input.role,
        status: input.status
      }
    }).toPromise();
    if (res.error) {
      alerts.push({ type: 'error', title, body: res.error.message });
      return;
    }

    alerts.push({ type: 'success', title, body: `successfully saved user ${input.firstName} ${input.lastName}` });
    toUsers();
  }

  function toUsers() {
    goto(PATH.USERS + $page.url.search);
  }
</script>

<Page title='Edit User'>
  <svelte:fragment slot='buttons'>
    <BtnBackSmall on:click={toUsers} />
  </svelte:fragment>
  <div class='flex gap-4 flex-col sm:flex-row flex-wrap' class:blur-sm={busy}>
    <CardWithBorderTitle title='User Detail'>
      <UserForm {user} saveHandler={save} />
    </CardWithBorderTitle>
    <CardWithBorderTitle title='Notification Settings'>
      <UserNotifications userID={$page.params.id} />
    </CardWithBorderTitle>
  </div>
</Page>
