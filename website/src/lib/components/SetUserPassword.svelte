<script lang="ts">
  import {authUser} from "$lib/stores/auth";
  import Input from "$lib/components/form/Input.svelte";
  import Modal from "$lib/components/Modal.svelte";
  import {getContextClient} from "@urql/svelte";
  import {SetUserPwd} from "$lib/gql/account";
  import alerts from "$lib/stores/alerts";
  import {validatePwd} from "$lib/utils/validate";
  import InputPassword from "$lib/components/form/InputPassword.svelte";
  import {Role} from "$lib/graph/graphql";

  export let userID: string;

  const client = getContextClient();
  const title = 'Set New Password'

  let pwd, confirmPwd, open = false, busy = false;
  $:isAdmin = $authUser.role === Role.Admin

  function toggle() {
    open = !open;
    if (open) {
      pwd = '';
      confirmPwd = '';
    }
  }

  async function submit() {
    if (busy) {
      return
    }

    let valErr: string | null = validatePwd(pwd)
    if (valErr) {
      alerts.warning(title, valErr)
      return
    }

    if (pwd !== confirmPwd) {
      alerts.warning(title, 'Confirm Password does not match the New Password')
      return
    }

    busy = true
    const res = await client.mutation(SetUserPwd, {userID, pwd, confirmPwd}).toPromise()
    busy = false
    if (res.error) {
      alerts.error(title, res.error.message)
      return
    }

    alerts.success(title, 'Successfully changed password');
    toggle();
  }
</script>

{#if (isAdmin)}
  <div type="button" class="link link-neutral" on:click={toggle}>
    set new password
  </div>
{/if}

<Modal {open} {busy} {title} okText="SUBMIT" cancelText="CANCEL" size="sm" on:ok={submit} on:cancel={toggle}>
  <InputPassword label="New Password" bind:value={pwd} required/>
  <Input type="password" label="Confirm Password" bind:value={confirmPwd} maxlength="100" required/>
</Modal>
