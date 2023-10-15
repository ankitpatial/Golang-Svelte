<script lang='ts'>
  import { afterUpdate, createEventDispatcher } from 'svelte';
  import Modal from '$lib/components/Modal.svelte';
  import { localDate } from '$lib/utils/time';
  import { progress } from '$lib/components/JobProgress/store';

  import type { Answer, Question } from './d';
  import Q from './Q.svelte';

  export let title: string;
  export let question: Question;

  const dispatch = createEventDispatcher();

  let answers: Array<Answer> = [],
    showOther: boolean,
    end = false, otherAns, to;

  afterUpdate(() => {
    if (question && question.type === 'note') {
      const detail = {
        qID: question.id,
        qType: question.type,
        question: question.text,
        answer: question.options?.join('<br/>')
      };
      input({ detail });
    }
  });

  function input({ detail }) {
    detail.order = answers.length;
    if (detail.answer === 'Other' && !detail.otherTxt) {
      showOther = true;
      otherAns = detail;
      otherAns.hasOtherTxt = true;
    } else if (question?.on) {
      answers = answers.concat(detail);
      const on = question.on[detail.answer];
      if (typeof on.enter === 'function') {
        on.enter();
      }
      const nID = on.next != undefined ? on.next : question?.id + 1;
      next(nID);
    } else {
      answers = answers.concat(detail);
      next(question?.id + 1);
    }
  }

  function next(n: number) {
    showOther = false;
    otherAns = undefined;
    end = question?.end || false;
    // console.log('next, Last:', end, 'Q:', question);
    if (!end) {
      dispatch('next', n);
    }
  }

  function save() {
    dispatch('save', answers);
  }

</script>

<Modal
  open
  {title}
  size='sm'
  disableOkBtn={!end}
  okText='SAVE'
  on:ok={save}
  on:cancel
  busy={$progress?.saving}
>
  <slot />

  {#each answers as a}
    <div class='mt-5'>
      <h3>{a.question}</h3>
      <p class='pl-2'>
        {#if (a.qType === 'dtp')}
          {localDate(a.answer, false)}
        {:else }
          {#if (a.hasOtherTxt)}
            {a.otherTxt}
          {:else}
            {@html a.answer}
          {/if}
        {/if}
      </p>
    </div>
  {/each}

  {#if (!end && question)}
    <Q {question} on:input={input} />
    {#if (showOther)}
      <div class='ml-8 mt-4'>
        <input
          type='text'
          class='input input-bordered w-full'
          bind:value={otherAns.otherTxt}
          maxlength='250'
        />
        <button
          class='btn btn-outline mt-2 btn-sm'
          on:click={() => {
            input({detail: otherAns})
          }}>
          NEXT
        </button>
      </div>
    {/if}
  {/if}
</Modal>
