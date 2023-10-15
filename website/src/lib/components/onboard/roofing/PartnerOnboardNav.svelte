<script lang='ts'>
  import {page} from '$app/stores';
  import {goto} from '$app/navigation';
  import {activeStep, partner, steps} from './index';

  let path
  $: path = $page.url.pathname

</script>

<ul class='steps mb-4'>
  {#each $steps as step, idx}
    <li
      class='step'
      class:step-primary={idx < $activeStep}
      data-content={idx < $partner.setupStepsCompleted? "✓" : "●"}
    >
      {#if idx <= $partner.setupStepsCompleted}
        <div on:click={()=> {
          $page.url.searchParams.set('step', idx + 1);
          goto(path + $page.url.search)
        }} class='link link-primary'>
          {step.name}
        </div>
      {:else }
        {step.name}
      {/if}
    </li>
  {/each}
</ul>

