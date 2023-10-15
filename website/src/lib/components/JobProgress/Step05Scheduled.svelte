<script lang='ts'>
  import { progress } from '$lib/components/JobProgress/store';
  import { localDate } from '$lib/utils/time';
  import { PATH } from '$lib/enum';
  import type { Answer, Question } from './d';
  import StepName from './StepName.svelte';
  import QuestionModal from './QuestionModal.svelte';

  export let step;

  const questions: Array<Question> = [
    {
      id: 1,
      type: 'dtp',
      text: 'When is the job scheduled for?',
      end: true
    }
  ];

  let open = false,
    question: Question | undefined = questions[0],
    data: { permitRequired?: boolean, inspectionRequired?: boolean } = {};

  async function save({ detail }) {
    const dt = (detail.find(p => p.qID == 1) as Answer).answer;
    const saved = await progress.saveAndNext(`Scheduled date ${localDate(dt, false)}`, { installDate: dt });
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
    title='Schedule the job'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}
  >
    Details on scheduling are found
    <a href='{PATH.JOB_DOCS}/scheduling' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}

