<script>
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import { getContextClient } from '@urql/svelte';
  import { UpdateApiAccessKey } from '$lib/gql/apiAccess';
  import alerts, { modalAlert } from '$lib/stores/alerts';

  export let id;

  const client = getContextClient();
  const title = 'Update API Access Key';

  let showModal = false;
  let key = '';
  let busy = false;

  // save function to execute updateApiAccessKey mutation
  async function save() {
    if (busy) return;
    if (key === '') {
      modalAlert.warning(title, 'Key is required');
      return;
    }

    busy = true;
    const res = await client.mutation(UpdateApiAccessKey, { id, key }).toPromise();
    busy = false;
    if (res.error) {
      modalAlert.error(title, res.error.message);
      return;
    }

    alerts.success(title, 'Successfully updated Access Key');
    showModal = false;
    key = '';
  }
</script>

<BtnIcon tooltip={title} on:click={()=> {showModal = true;}}>
  <IconEdit />
</BtnIcon>

<Modal open={showModal} {title} on:ok={save} on:cancel={()=> { showModal = false}} {busy}>
  <Input name='key' bind:value={key} />
</Modal>
