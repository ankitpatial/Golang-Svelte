<!--
 Copyright (c) 2022. Ankit Patial.
 Unauthorized copying of this file, via any medium is strictly prohibited. Proprietary and confidential.
 Author: Ankit Patial
  -->
<script lang='ts'>
	export let cls = '';
</script>

<div class='card p-0 sm:w-96 {cls} bg-white ease-in-out duration-500 transition hover:-translate-y-1 hover:shadow-lg'>
	<div class='flex'>
		<div class='flex-none w-28 bg-slate-200 rounded-md overflow-hidden'>
			{#if $$slots.title}
				<div class='card-title'>
					<slot name='title'></slot>
				</div>
			{/if}
		</div>
		<div class='flex-1 text-left p-3'>
			<div class='flex justify-between items-start mb-1.5'>
				<slot name='header'></slot>
				{#if $$slots.actions}
					<slot name='actions'></slot>
				{/if}
			</div>
			<slot></slot>
		</div>
	</div>
</div>
