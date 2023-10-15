<script lang='ts'>
  import {createEventDispatcher} from 'svelte';
  import {progress, steps} from './store';
  import {getContextClient} from '@urql/svelte';
  import {authUser} from "$lib/stores/auth";
  import type {Job} from "$lib/graph/graphql";

  export let job: Job;
  export let isDirty: boolean = false;

  const client = getContextClient();
  const dispatch = createEventDispatcher();

  let color = '', readonly = false;
  $: isDirty = $progress?.isDirty;
  $: {
    if (!job?.contractor?.id || !$authUser?.partner?.id) {
      readonly = true
    } else {
      readonly = job?.contractor?.id != $authUser?.partner?.id
    }
    progress.init(client, job, readonly);
  }

  function cancel() {
    dispatch('cancel', {hasChanges: $progress?.isDirty});
  }
</script>

{#if ($progress.loading)}
  <p>Please wait...</p>
{:else }
  <div class='p-2 mt-5'>
    <ul class='steps mb-4 steps-vertical w-full'>
      {#each $steps || [] as step}
        <li
          class='step step-error'
          class:step-error={step.active && $progress?.job?.progressFlagged}
          class:step-success={step.completed}
          data-content={step.completed? "✓" : "●"}
        >
          <svelte:component this={step.component} {step}/>
        </li>
      {/each}
    </ul>
  </div>
{/if}

