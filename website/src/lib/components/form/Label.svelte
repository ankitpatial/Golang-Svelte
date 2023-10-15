<!--
  -  Copyright (c) 2022 - 2023.  Ankit Patial.
  -  Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
  -  Author: Ankit Patial
  -->

<script lang='ts'>
  import IconInformationCircle from '$lib/assets/svg/IconInformationCircle.svelte';

  export let content = '';
  export let htmlContent = '';
  export let uppercase = false;
  export let id = '-none-';
  export let cls = '';
  export let tooltip = '';
</script>

<label class='label p-0 {cls}' for={id}>
	<span class="label p-0 mb-0.5" class:uppercase={uppercase}>
		{#if !!htmlContent}
      {@html htmlContent}
    {:else }
      {content}
    {/if}
    {#if !!tooltip}
      <div class="tooltip pl-1" data-tip={tooltip}>
        <IconInformationCircle cls="w-3 h-3"/>
      </div>
    {/if}
	</span>
  <span class="label-text-alt">
		<slot></slot>
	</span>
</label>
