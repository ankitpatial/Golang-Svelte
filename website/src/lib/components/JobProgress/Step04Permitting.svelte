<script lang='ts'>
  import { progress } from '$lib/components/JobProgress/store';
  import { PATH } from '$lib/enum';
  import StepName from './StepName.svelte';
  import QuestionModal from './QuestionModal.svelte';
  import type { Answer, Question } from './d';

  export let step;

  const questions: Array<Question> = [
    {
      id: 1,
      type: 'buttons',
      text: 'Is a permit required for this job?',
      options: ['Yes', 'No'],
      on: {
        'Yes': {
          enter: () => {
            data.permitRequired = true;
          },
          next: 2
        },
        'No': {
          enter: () => {
            data.permitRequired = false;
          },
          next: 4
        }
      }
    },
    {
      id: 2,
      type: 'buttons',
      text: 'Have you submitted the permit?',
      options: ['Yes', 'No']
    },
    {
      id: 3,
      type: 'buttons',
      text: 'Is a final inspection required for this job?',
      options: ['Yes', 'No'],
      on: {
        Yes: {
          enter: () => {
            data.inspectionRequired = true;
          }
        },
        No: {
          enter: () => {
            data.inspectionRequired = false;
          }
        }
      },
      end: true
    },
    {
      id: 4,
      type: 'radio',
      text: 'Why not?',
      options: [
        'Permits are not required in this market',
        'Other'
      ],
      end: true
    }
  ];

  let open = false,
    question: Question | undefined = questions[0],
    data: { permitRequired?: boolean, inspectionRequired?: boolean } = {};

  async function save({ detail }) {
    const confirmed = (detail.find(p => p.qID == 1) as Answer).answer;
    let saved = false;
    if (confirmed === 'Yes') {
      const submitted = (detail.find(p => p.qID == 2) as Answer).answer;
      const note = 'Permit is required.';
      const note2 = `${submitted === 'Yes' ? 'Permit was submitted.' : 'Permit was not submitted.'}`;
      const note3 = `${submitted === 'Yes' ? 'Final inspection is required.' : 'Final inspection is not required.'}`;
      saved = await progress.saveAndNext(`${note} \n${note2} \n${note3}`, data);
    } else {
      const ans = (detail.find(p => p.qID == 4) as Answer);
      const whyNot = ans.hasOtherTxt ? ans.otherTxt : ans.answer;
      const note = `Permit is not required.\nReason: ${whyNot}`;
      saved = await progress.saveAndNext(note, { permitRequired: false });
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
    title='Acquire the permit.'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}
  >
    Details on permitting are found
    <a href='{PATH.JOB_DOCS}/permitting' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}

