<script>
  import IconEdit from '$lib/assets/svg/IconEdit.svelte';
  import BtnIcon from '$lib/components/form/BtnIcon.svelte';
  import Modal from '$lib/components/Modal.svelte';
  import Input from '$lib/components/form/Input.svelte';
  import { getContextClient } from '@urql/svelte';
  import { UpdateApiAccessSecret } from '$lib/gql/apiAccess';
  import alerts, { modalAlert } from '$lib/stores/alerts';

  export let id;

  const client = getContextClient();
  const title = 'Update API Access Secret';

  let showModal = false;
  let secret = '';
  let busy = false;

  // save function to execute updateApiAccessKey mutation
  async function save() {
    if (busy) return;
    if (secret === '') {
      modalAlert.warning(title, 'Secret is required');
      return;
    }

    busy = true;
    const res = await client.mutation(UpdateApiAccessSecret, { id, secret }).toPromise();
    busy = false;
    if (res.error) {
      modalAlert.error(title, res.error.message);
      return;
    }

    alerts.success(title, 'Successfully updated Access Secret');
    showModal = false;
    secret = '';
  }
</script>

<BtnIcon tooltip={title} on:click={()=> {showModal = true;}}>
  <IconEdit />
</BtnIcon>

<Modal open={showModal} {title} on:ok={save} on:cancel={()=> { showModal = false}} {busy}>
  <Input name='key' bind:value={secret} />
</Modal>
