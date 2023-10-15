<script lang='ts'>
  import { progress } from '$lib/components/JobProgress/store';
  import alerts from '$lib/stores/alerts';
  import { PATH } from '$lib/enum';
  import type { Question } from './d';
  import QuestionModal from './QuestionModal.svelte';
  import StepName from './StepName.svelte';

  export let step;

  const questions: Array<Question> = [
    {
      id: 1,
      type: 'buttons',
      text: 'Is everything complete and you are ready to submit the invoice?',
      options: ['Yes', 'No'],
      on: {
        'Yes': {
          enter: save
        },
        No: {
          enter: hide
        }
      },
      end: true
    },
  ];

  let open = false,
    question: Question | undefined = questions[0],
    data: { permitRequired?: boolean, inspectionRequired?: boolean } = {};

  async function save() {
    const note = 'Submitted the invoice';
    const saved = await progress.saveAndNext(note);
    if (saved) {
      const msg = 'This job will now appear in your payment center dashboard as it processed, verified and then paid out.';
      alerts.success('Your invoice has been submitted.', msg);
      hide();
    }
  }

  function next({ detail }) {
    // no next step
  }

  function show() {
    question = questions[0];
    open = true;
  }

  function hide() {
    question = questions[0];
    open = false;
  }
</script>

<StepName {step} on:click={show} />

{#if (open)}
  <QuestionModal
    title='Time to get paid.'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}
  >
    Details on invoicing are found
    <a href='{PATH.JOB_DOCS}/invoicing' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}

