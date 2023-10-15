<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
	export let title = '';
	export let open = false;

	function toggle() {
		open = !open;
	}
</script>

<div
	tabindex='0'
	class='collapse collapse-arrow rounded-box'
	class:collapse-open={open}
	class:collapse-close={!open}
>
	<div class='collapse-title text-lg font-medium min-h-4 bg-blue-50 border-b-2 border-blue-100' on:click|stopPropagation={toggle}>
		{title}
	</div>
	<div class='collapse-content collapse-open mt-2'>
		<slot></slot>
	</div>
</div>