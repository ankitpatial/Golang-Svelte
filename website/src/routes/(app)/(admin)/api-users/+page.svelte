<script lang='ts'>
  import DataLoader from '$lib/components/data/DataLoader';
  import { ChangeApiUserStatus, QryApiUsers, RefreshApiUserPwd } from '$lib/gql/apiUser';
  import { getContextClient } from '@urql/svelte';
  import { TableDataLoader } from '$lib/components/DataTable';
  import alerts from '$lib/stores/alerts';
  import { PATH } from '$lib/enum.js';
  import Form from './Form.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import PwdField from './PwdField.svelte';
  import ConfirmAction from '$lib/components/ConfirmAction.svelte';
  import BtnPlusSmall from '$lib/components/form/BtnPlusSmall.svelte';
  import TableCard from '$lib/components/TableCard.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import IconStack from '$lib/assets/svg/IconStack.svelte';
  import { goto } from '$app/navigation';
  import IconKey from '$lib/assets/svg/IconKey.svelte';
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  const loader = new DataLoader(client, QryApiUsers, 'apiUsers', 10);
  const store = loader.store;

  let search = '',
    showModal = false,
    changeId = '',
    pwd = '',
    changingPwd = false,
    showCopyPwd = false;
  $: {
    if (search) {
      loader.where = {
        usernameContainsFold: search
      };
    } else {
      loader.where = {};
    }

    reload();
  }

  function openModal() {
    showModal = true;
  }

  function closeModal() {
    showModal = false;
  }

  function reload() {
    loader.load();
  }

  async function updateStatus(id: string, isActive: boolean) {
    const res = await client.mutation(ChangeApiUserStatus, { id, isActive }).toPromise();
    if (res.error) {
      alerts.push({ type: 'error', title: 'Update Api User', body: res.error.message });
      await loader.load(false);
      return;
    }

    store.update(s => {
      const u = s.items.find(i => i.id === id);
      if (u) {
        u.active = isActive;
      }
      return s;
    });
  }

  async function newPwd(id: string) {
    if (changingPwd) return;
    changeId = id;
    changingPwd = true;
    const res = await client.mutation(RefreshApiUserPwd, { id }).toPromise();
    changeId = '';
    changingPwd = false;
    if (res.error) {
      alerts.push({ type: 'error', title: 'Change Password', body: res.error.message });
      return;
    }

    // show modal to copy pwd
    pwd = res.data.refreshApiUserPwd;
    showCopyPwd = true;
  }

</script>

<svelte:head>
  <title>Manage API Users</title>
</svelte:head>

<Page title='API Users'>
  <svelte:fragment slot='buttons'>
    <BtnPlusSmall on:click={openModal} />
  </svelte:fragment>

  <TableCard>
    <TableDataLoader searchable header={['NAME', 'ACTIVE', 'ACTION', ]} bind:value={search} {loader}>
      {#each $store.items as item}
        <tr>
          <td class='md:w-96'>{item.username}</td>
          <td>
            <ConfirmAction
              type='toggle'
              title={`${item.active ? 'Deactivate': 'Activate'} API User`}
              message={`Are you sure you want to ${item.active ? 'deactivate': 'activate'} API Access To: <strong>${item.username}</strong>?`}
              bind:value={item.active}
              confirm={()=> updateStatus(item.id, !item.active)}
              tooltip='Change Password' />
          </td>
          <td class='w-52'>
            <div class='flex flex-row flex-wrap space-x-2'>
              <ConfirmAction
                title='Set New Password'
                message='Are you sure you want to set a new password?'
                confirm={() => newPwd(item.id)}
                tooltip='Change Password'
              >
                <IconKey color='red' />
              </ConfirmAction>

              <BtnIcon on:click={()=> goto(PATH.API_USERS_AUDIT_LOG.replace(':id', item.id))} tooltip='Audit Logs'>
                <IconStack />
              </BtnIcon>
              <BtnIcon on:click={()=> goto(PATH.API_USERS_EDIT.replace(':id', item.id))} tooltip='Edit'>
                <IconEdit />
              </BtnIcon>
            </div>
          </td>
        </tr>
      {/each}
    </TableDataLoader>
  </TableCard>

  {#if showModal}
    <Form reloadHandel={reload} closeHandel={closeModal} />
  {/if}

  <Modal open={showCopyPwd} title='New Password' hideOkBtn cancelText='Close' on:cancel={() => { showCopyPwd = false;}}>
    <PwdField value={pwd} />
  </Modal>
</Page>
