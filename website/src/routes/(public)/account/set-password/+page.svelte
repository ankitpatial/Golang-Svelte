<script lang='ts'>
  import alerts from '$lib/stores/alerts';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { PATH } from '$lib/enum.js';
  import { getContextClient } from '@urql/svelte';
  import { authUser, authUserDashboard } from '$lib/stores/auth';
  import { validatePwd } from '$lib/utils/validate';
  import Input from '$lib/components/form/Input.svelte';
  import InputPassword from '$lib/components/form/InputPassword.svelte';
  import Btn from '$lib/components/form/Btn.svelte';

  const client = getContextClient();
  const title = 'Set Password';

  let pwd = '', confirm = '', busy = false;

  async function submit() {
    if (busy) return;

    let valErr: string | null = validatePwd(pwd);
    if (valErr) {
      alerts.warning(title, valErr);
      return;
    }

    if (pwd !== confirm) {
      alerts.push({ type: 'error', title: 'Validate', body: 'Password and Confirmation Password do not match.' });
      return;
    }

    try {
      busy = true;
      const resRaw = await fetch('/set-password', {
        method: 'POST',
        body: JSON.stringify({
          token: $page.url.searchParams.get('token'),
          password: pwd,
          confirmPassword: confirm
        })
      });
      const res = await resRaw.json();
      if (res.error) {
        alerts.push({ type: 'error', title, body: res.error });
        return;
      }

      authUser.set(res.data);
      await goto($authUserDashboard);
      alerts.push({ type: 'success', title, body: 'successfully logged in' });
    } catch (ex) {
      alerts.push({ type: 'error', title, body: ex.message || ex });
    } finally {
      busy = false;
    }
  }
</script>


<form on:submit|preventDefault={submit}>
  <h2 class='card-title font-normal justify-center mt-16 mb-6'>
    Set Account Password
  </h2>

  <div class='card-actions justify-center space-y-3'>
    <InputPassword label='Password' bind:value={pwd} required />
    <Input type='password' label='Confirm Password' bind:value={confirm} maxlength='100' required />
    <Btn type='submit' primary shadow {busy}>
      Submit
    </Btn>
  </div>
</form>

<p class='text-xs mt-4 text-center'>
  <a href={PATH.LOGIN} class='link'>Click to Login</a>
</p>
