<script lang="ts">
  import Input from "$lib/components/form/Input.svelte";

  export let label;
  export let value;
  export let required = false;
  export let handleChange: (event: Event) => any = null;

</script>

<Input type="password" {label} bind:value maxlength="100" {required} {handleChange}>
  <svelte:fragment slot="desc">
    At least 8 characters, including uppercase and lowercase letters, a number, a special character, and no white
    spaces
  </svelte:fragment>
</Input>
