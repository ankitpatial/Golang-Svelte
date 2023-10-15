<script lang='ts'>
  import { PATH } from '$lib/enum';
  import { progress } from '$lib/components/JobProgress/store';
  import QuestionModal from './QuestionModal.svelte';
  import StepName from './StepName.svelte';
  import type { Answer, Question } from './d';

  export let step;

  const questions: Array<Question> = [
    {
      id: 1,
      type: 'buttons',
      text: 'Have you confirmed the shingle color with the customer?',
      options: ['Yes', 'No'],
      on: {
        'Yes': {
          next: 2
        },
        'No': {
          next: 3,
        }
      }
    },
    {
      id: 2,
      type: 'radio',
      text: 'What is the confirmed shingle color?',
      options: ['Brown', 'Sandstorm', 'Blue', 'Other'],
      end: true
    },
    {
      id: 3,
      type: 'radio',
      text: 'Why not?',
      options: ['Still waiting on customer to decide', 'Other'],
      end: true
    }
  ];

  let open = false,
    question: Question | undefined = questions[0],
    data: { shingleColor?: string } = {};

  async function save({ detail }) {
    const confirmed = (detail.find(p => p.qID == 1) as Answer).answer;
    let saved:boolean;
    if (confirmed === 'Yes') {
      let color = '';
      const currentQuestion = (detail.find(p => p.qID == 2) as Answer);
      if(currentQuestion.answer === 'Other')
        color = currentQuestion.otherTxt || '';
      else 
        color = currentQuestion.answer;
      const note = `Shingle color confirmed. \nColor: ${color}`;
      saved = await progress.saveAndNext(note, { shingleColor: color });
    } else {
      const ans = (detail.find(p => p.qID == 3) as Answer);
      const whyNot = ans.hasOtherTxt ? ans.otherTxt : ans.answer;
      const note = `Shingle color not confirmed. \nReason: ${whyNot}`;
      saved = await progress.save(note);
    }

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
    title='Confirm the details of the job with the customer.'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}

  >
    Requirements on what job details to confirm are found
    <a href='{PATH.JOB_DOCS}/details-to-confirm' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}

