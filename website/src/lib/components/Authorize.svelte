<script lang='ts'>
  import {PATH} from '$lib/enum';
  import {authBusy, authUser} from '$lib/stores/auth';
  import {goto} from '$app/navigation';
  import {Role} from "$lib/graph/graphql";

  export let role: Role;
  let busy = $authBusy;

  function toLogin() {
    goto(PATH.LOGIN);
  }
</script>

{#if (!$authUser) }
  <div>
    <p>Session expired.</p>
    <div class='btn btn-outline' on:click={toLogin}>
      Click to login again
    </div>
  </div>
{:else if ($authUser?.role !== role)}
  <div>
    <h1>Not Authorized </h1>
    <p class="mt-5">This page is for user's with {role} role.</p>
    <p>You are not authorize to view this page.</p>
  </div>
{:else }
  <slot></slot>
{/if}
