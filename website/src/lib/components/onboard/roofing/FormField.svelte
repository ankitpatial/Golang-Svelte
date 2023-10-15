<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2023.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
	export let label;
	let id = `fm-${Date.now()}`;
</script>

<div class='form-control mb-5'>
	<label for={id} class='label p-0 mb-0.5 uppercase'>{label} </label>
	<div id={id}>
		<slot></slot>
	</div>
</div>
