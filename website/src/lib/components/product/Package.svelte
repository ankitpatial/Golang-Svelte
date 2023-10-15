<script lang='ts'>
  import { usdFormat } from '$lib/utils/currency.js';
  import Btn from '$lib/components/form/Btn.svelte';
  import IconCheckCircle from '$lib/assets/svg/IconCheckCircle.svelte';
  import type { ProductPackage } from '$lib/graph/graphql';
  import { createEventDispatcher } from 'svelte';

  export let idx: number;
  export let data: ProductPackage;

  const dispatch = createEventDispatcher();

  let len = 0;
  let current = 0;
  $: len = data.items.length;

  function select() {
    dispatch('select', { pkg: data });
  }
</script>

<div class='card'>
  <div class='card-body'>
    <h2 class='card-title'>{data.name}</h2>
    <div class='carousel max-w-md'>
      {#each data.items as prod, j}
        <div id='s{idx}{j}' class='carousel-item relative w-full'>
          <div>
            <figure>
              <img src={prod.image.publicUrl} class='w-96' alt='product' />
            </figure>
            <ul>
              {#each prod.features as f}
                <li class='flex flex-row gap-2'>
                  <div>
                    <IconCheckCircle color='green' />
                  </div>
                  <div>
                    {f}
                  </div>
                </li>
              {/each}
            </ul>
            <div class='absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2'>
              <a href='#s{idx}{j === 0 ? 0 : j-1}' class='btn btn-circle' class:disabled={j===0}>❮</a>
              <a href='#s{idx}{j + 1 < len ? j + 1 : len}' class='btn btn-circle'>❯</a>
            </div>
          </div>
        </div>
      {/each}
    </div>
    <div class='stat'>
      <div class='stat-value'> {usdFormat(data.price)}</div>
      <div class='stat-actions'>
        <Btn primary on:click={select}>
          SELECT
        </Btn>
      </div>
    </div>
  </div>
</div>
