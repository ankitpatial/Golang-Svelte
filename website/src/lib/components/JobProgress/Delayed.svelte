<script lang='ts'>
  import Modal from '$lib/components/Modal.svelte';

  export let step;
  let open = false;

  function yes() {

  }

  function no() {

  }
</script>

<button class='btn btn-outline' disabled={!step.active}>
  {step.name}
</button>

<Modal
  {open}
  title='Contact the customer to perform the welcome call.'
  okText='Yes'
  cancelText='No'
  on:ok={yes}
  on:cancel={no}
>
  <p>
    Details on what to cover during the welcome call are found <a href='#'>here</a>.
    Have you made contact with the customer?
  </p>
</Modal>
