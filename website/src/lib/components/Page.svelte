<script lang='ts'>
  export let title;
  export let subTitle = '';
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<div class='flex flex-col sm:flex-row justify-between mb-3 top-12 z-30 w-full pt-2 sm:pt-0 bg-white sticky'>
  <div class='flex flex-col gap-3'>
    <div class='flex gap-x-2'>
      <h1>
        {@html title}
        {#if !!subTitle}
				<span class='text-sm text-gray-400'>
					{subTitle}
				</span>
        {/if}
      </h1>
      {#if $$slots.buttons}
        <div class='flex flex-row gap-x-2'>
          <slot name='buttons'></slot>
        </div>
      {/if}
    </div>
    {#if $$slots.titleContent}
      <div>
        <slot name='titleContent'></slot>
      </div>
    {/if}
  </div>
  {#if $$slots.filters}
    <div>
      <slot name='filters'></slot>
    </div>
  {/if}
</div>
<div class='mt-8'>
  <slot></slot>
</div>
