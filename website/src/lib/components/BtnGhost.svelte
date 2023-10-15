<script lang='ts'>
	export let tooltip = '';
	export let cls = '';
	export let btnCls = '';
</script>
<div class='text-sm z-30 tooltip-left {cls}' class:tooltip={!!tooltip} data-tip={tooltip}>
	<button class='btn btn-ghost btn-square btn-sm text-primary hover:bg-primary hover:text-white {btnCls}' on:click>
		<slot />
	</button>
</div>