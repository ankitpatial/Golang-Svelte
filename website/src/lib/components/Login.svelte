<script lang='ts'>
  import {goto} from '$app/navigation';
  import {authBusy, authUser, authUserDashboard} from '$lib/stores/auth';
  import alerts from '$lib/stores/alerts';
  import {page} from '$app/stores';
  import {urlPathName} from '$lib/utils/url';
  import Btn from '$lib/components/form/Btn.svelte';

  let email, pwd, busy = false, err;

  $: if ($authUser?.id) {
    next();
  }

  async function submit() {
    if (busy || $authBusy) return;

    try {
      busy = true;
      const resRaw = await fetch('/login', {method: 'POST', body: JSON.stringify({email, password: pwd})});
      const res = await resRaw.json();
      if (res.error) {
        alerts.push({type: 'error', title: 'Login', body: res.error});
        return;
      }

      authUser.set(res.data);
      alerts.push({type: 'success', title: 'Login', body: 'successfully logged in'});
      next();
    } catch (ex) {
      alerts.push({type: 'error', title: 'Login', body: ex.message || ex});
    } finally {
      busy = false;
    }
  }

  function next() {
    const key = 'returnURL';
    let path = $authUserDashboard.toString();
    if ($page.url.searchParams.has(key)) {
      const p = urlPathName($page.url.searchParams.get(key));
      if (p && p !== '/') {
        path = p;
      }
    }
    goto(path);
  }
</script>

<form method='post' on:submit|preventDefault={submit}>
  <h2 class='card-title font-normal justify-center mt-16 mb-6'>
    SIGN IN
  </h2>

  <div class='card-actions justify-center space-y-3'>
    <input
      type='email'
      name='email'
      placeholder='email'
      class='input input-bordered border-white w-full bg-white shadow-md text-center focus:border-primary focus:shadow-sm'
      bind:value={email}
      required
    />
    <input
      type='password'
      name='password'
      placeholder='PASSWORD'
      class='input input-bordered border-white w-full bg-white shadow-md text-center focus:border-primary focus:shadow-sm'
      bind:value={pwd}
      required
    />
    <Btn type='submit' primary shadow {busy} cls="btn-block">
      LOGIN
    </Btn>
  </div>
</form>

<p class='text-sm mt-4 text-center'>
  <a href='/account/forgot-password' class='link'>Forgot Password?</a>
</p>
