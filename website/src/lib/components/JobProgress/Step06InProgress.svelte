<script lang='ts'>
  import {progress} from '$lib/components/JobProgress/store';
  import {localDate} from '$lib/utils/time';
  import {PATH} from '$lib/enum';
  import type {Answer, Question} from './d';
  import QuestionModal from './QuestionModal.svelte';
  import StepName from './StepName.svelte';

  export let step;

  let questions: Array<Question>;

  $:if ($progress?.job) {
    const j = $progress.job;
    const permitRequired = j.permitRequired;
    const inspectionRequired = j.inspectionRequired;

    questions = [
      {
        id: 1,
        type: 'buttons',
        text: 'Are crew members and the materials on site and the job is ready to begin?',
        options: ['Yes', 'No'],
        on: {
          Yes: {
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
        text: 'Before you begin this job, you must upload the permit here',
        upload: {
          name: 'PERMIT',
          accept: 'image/*,application/pdf',
          section: 'Docs',
          multiple: false
        },
        inActive: !permitRequired,
        end: !inspectionRequired
      },
      {
        id: 3,
        type: 'note',
        text: 'Great!',
        options: [
          'Remember that quality production photos and a passed final inspection doc (if applicable) will be required to be uploaded in order to be paid.',
          'More details on production photos are found here.'
        ],
        inActive: permitRequired,
        end: !inspectionRequired
      },
      {
        id: 4,
        type: 'dtp',
        text: 'When is the final inspection scheduled for?',
        end: true
      }
    ];
  }

  let
    open = false,
    question: Question,
    data: { permitRequired?: boolean, inspectionRequired?: boolean } = {};

  async function save({detail}) {
    const fileUpload = (detail.find(p => p.qID == 2) as Answer);
    const dtp = (detail.find(p => p.qID == 4) as Answer);
    let note = 'Ready to begin.',
      data = undefined;

    if (fileUpload) {
      note = `\nUploaded File: ${fileUpload.answer}`;
    }

    if (dtp) {
      note = `\nInspection Date: ${localDate(dtp.answer, false)}`;
      data = {inspectionDate: dtp.answer};
    }

    const saved = await progress.saveAndNext(note, data);
    if (saved) {
      hide();
    }
  }

  function next({detail}) {
    let nextID = detail;
    do {
      question = questions?.find(o => o.id === nextID);
      if (!question) {
        break;
      }

      if (question.inActive) {
        nextID++;
      } else {
        break;
      }

    } while (true);
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

<StepName {step} on:click={show}/>

{#if (open)}
  <QuestionModal
    title='Time to install the job.'
    {question}
    on:save={save}
    on:next={next}
    on:cancel={hide}
  >
    Details on installing are found
    <a href='{PATH.JOB_DOCS}/installing' class='link-primary' target='_blank' rel='noopener noreferrer'>
      here
    </a>
  </QuestionModal>
{/if}

