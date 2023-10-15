<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang="ts">
	export let width = 'sm:w-9/12';
</script>

<div class='w-full {width}'>
	<div class='card-body p-0'>
		<slot></slot>
	</div>
</div>