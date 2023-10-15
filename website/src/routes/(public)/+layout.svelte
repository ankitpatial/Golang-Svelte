<script lang='ts'>
	import LogoRoofix from '$lib/assets/svg/LogoRoofix.svelte';
</script>

<div class='grid place-content-center min-h-screen bg-base-200'>
	<div class='card ring-0 w-full max-w-sm sm:w-96 h-[430px] shadow-none'>
		<div class='card-body justify-center'>
			<p class='text-center'>
				<LogoRoofix black cls='h-20 inline-flex' />
			</p>
			<div class='card-actions'>
				<slot></slot>
			</div>
		</div>
	</div>
</div>
