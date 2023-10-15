<!--
  - Copyright (c) 2022. SimSaw Software Private Limited.
  - Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  - Author: Ankit Patial
  - Dated:  20/05/22, 3:06 PM
  -->
<script lang='ts'>
  import IconSortAsc from '$icons/IconSortAsc.svelte';
  import IconSortDesc from '$icons/IconSortDesc.svelte';

  export let asc;
</script>

<label class='btn btn-square btn-sm swap swap-rotate btn-outline bg-white border-slate-200 hover:bg-neutral hover:border-slate-400'>
  <input type='checkbox' bind:checked={asc} />
  <IconSortAsc className='swap-on' />
  <IconSortDesc className='swap-off' />
</label>
