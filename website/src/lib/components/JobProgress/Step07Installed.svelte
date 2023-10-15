<script lang='ts'>
  import { progress } from '$lib/components/JobProgress/store';
  import { PATH } from '$lib/enum';
  import StepName from './StepName.svelte';
  import QuestionModal from './QuestionModal.svelte';
  import type { Question } from './d';

  export let step;

  const questions: Array<Question> = [
    {
      id: 1,
      type: 'buttons',
      text: 'Has the job been successfully complete?',
      options: ['Yes', 'No'],
      on: {
        'Yes': {
          next: 2
        },
        No: {
          enter: hide
        }
      }
    },
    {
      id: 2,
      type: 'fileUpload',
      text: 'Great!',
      description: 'Upload the production photos here:',
      upload: {
        accept: 'image/*',
        multiple: true,
        section: 'ProductionPics'
      },
      end: true
    }
  ];

  let open = false,
    question: Question | undefined = questions[0],
    data: { permitRequired?: boolean, inspectionRequired?: boolean } = {};

  async function save({ detail }) {
    const note = 'Completed the job';
    const saved = await progress.saveAndNext(note);
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
    title='Complete the install.'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}
  >
    Details on job completion are found
    <a href='{PATH.JOB_DOCS}/completion' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}

