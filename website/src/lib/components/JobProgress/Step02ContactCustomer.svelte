<script lang='ts'>
  import QuestionModal from './QuestionModal.svelte';
  import { progress } from '$lib/components/JobProgress/store';
  import StepName from './StepName.svelte';
  import { getContextClient } from '@urql/svelte';
  import { PATH } from '$lib/enum';
  import type { Answer, Question } from './d';

  export let step;

  const client = getContextClient();
  const questions: Array<Question> = [
    {
      id: 1,
      type: 'buttons',
      text: 'Have you made contact with the customer?',
      options: ['Yes', 'No'],
      on: {
        Yes: {
          enter: () => {
            data.contacted = true;
          }
        },
        No: {
          enter: () => {
            data.contacted = false;
          }
        }
      }
    },
    {
      id: 2,
      type: 'radio',
      text: 'What methods did you attempt?',
      options: ['Phone Call', 'Text', 'Email'],
      end: true
    }
  ];

  let open = false,
    question: Question | undefined = questions[0],
    data: { contacted?: boolean } = {};

  // save progress step data
  async function save({ detail }) {
    const method = (detail.find(p => p.qID == 2) as Answer).answer;
    const saved = data.contacted
      ? await progress.saveAndNext(`Successfully contacted the customer via ${method}`)
      : await progress.save(`Unsuccessful attempt to contact the customer via ${method}`);
    if (saved) {
      hide();
    }
  }

  function next({ detail }) {
    question = questions.find(o => o.id === detail);
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
    title='Contact the customer to perform the Welcome Call.'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}
  >
    Details on what to cover during the welcome call are found
    <a href='{PATH.JOB_DOCS}/customer-welcome-call' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}
