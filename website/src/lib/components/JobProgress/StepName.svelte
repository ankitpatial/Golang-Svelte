<script lang='ts'>
  import {localDate} from '$lib/utils/time';
  import {createEventDispatcher} from 'svelte';
  import ConfirmAction from '$lib/components/ConfirmAction.svelte';
  import {progress} from '$lib/components/JobProgress/store';
  import {getContextClient} from '@urql/svelte';
  import {JobDelay} from '$lib/gql/job.progress';
  import alerts from '$lib/stores/alerts';
  import {JobProgress} from "$lib/graph/graphql";

  export let step;

  const client = getContextClient();
  const dispatch = createEventDispatcher();

  let delayed: boolean, delayReason: string, readonly: boolean;

  $:delayed = $progress?.delayed || false;
  $:delayReason = $progress?.history?.findLast(h => h.status === JobProgress.Delayed)?.note || '';
  $:readonly = $progress?.readonly;

  function click() {
    dispatch('click');
  }

  async function saveDelay(reason) {
    const flag = !delayed;
    const res = await client.mutation(JobDelay, {id: $progress?.job?.id, flag, reason: reason}).toPromise();
    if (res.error) {
      alerts.error('Error', res.error.message);
      return;
    }

    await progress.pullHistory();
    progress.setJobDelay(flag);
  }
</script>
<div class='flex flex-col gap-y-1 justify-start text-left'>
  {#if (readonly || step.completed)}
    <div class='font-bold'>{step.name}</div>
    {#if (step.completedAt)}
      <div>{localDate(step.completedAt, false)}</div>
    {/if}
  {:else}
    <button
      class='btn btn-outline w-56 '
      disabled={!step.active || delayed}
      on:click={click}
    >
      {step.name}
    </button>

    {#if (step.active)}
      <div>
        <div class='form-control'>
          <label class='label cursor-pointer'>
            <span class='label-text'>{(delayed ? 'Remove delay' : 'Set as delayed')}</span>
            <ConfirmAction
              type='toggle'
              size='sm'
              title={( delayed ? 'Remove Delay' :'Delayed'  )}
              message={(delayed ? 'How was delayed resolved?' :'Please specify the delay reason:'  )}
              value={delayed}
              noteTitle=''
              noteRequired
              confirm={saveDelay}
            />
          </label>
        </div>
        {#if delayed && delayReason}
          <p>{delayReason}</p>
        {/if}
      </div>

    {/if}
  {/if}

</div>
