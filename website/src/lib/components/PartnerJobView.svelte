<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->
<script lang='ts'>
	import Badge from '$lib/components/Badge.svelte';
	import IconPhone from '$lib/assets/svg/IconPhone.svelte';
	import IconEnvelope from '$lib/assets/svg/IconEnvelope.svelte';
	import IconMapPin from '$lib/assets/svg/IconMapPin.svelte';
	import IconDocumentTextSolid from '$lib/assets/svg/IconDocumentTextSolid.svelte';
	import IconPhotoSolid from '$lib/assets/svg/IconPhotoSolid.svelte';
	import type Job from '$lib/models/Job';
	import { usdFormat } from '$lib/utils/currency.js';

	export let job: Job;
	let partner = '',  stateCls = 'warning';

</script>

<div class='card w-full sm:w-96'>
	<div class='card-body'>
		<h2>
			<span class='tooltip tooltip-top' data-tip='Owner Name'>
				{job?.ownerFirstName || ''  } {job?.ownerLastName || ''}
			</span>
			<Badge type={stateCls}>{job.status}</Badge>
		</h2>
		<ul class='space-y-3 my-4'>
			<li class='flex gap-4'>
				<span><IconPhone /></span>
				<p>
					<span class='tooltip tooltip-top' data-tip='Rep Mobile'>
						{job?.repMobile}
					</span>
				</p>
			</li>
			<li class='flex gap-4'>
				<span><IconEnvelope /></span>
				<p>
					<span class='tooltip tooltip-top' data-tip='Rep Email'>
						{job?.repEmail}
					</span>
				</p>
			</li>
			<li class='flex gap-4'>
				<span><IconMapPin /></span>
				<p>
					<span class='tooltip tooltip-top' data-tip='Job Address'>
						{job?.streetNumber} {job?.streetName}, {job?.city}, {job?.state} {job?.zip}
					</span>
				</p>
			</li>
		</ul>

		<!-- section:ICON -->
		<div class='flex gap-4 text-center'>
			<div class='bg-accent-content rounded-md w-1/2 p-1'>
				<button class='btn btn-square btn-ghost btn-xs text-secondary'>
					<IconDocumentTextSolid />
				</button>
				<button class='btn btn-square btn-ghost btn-xs text-secondary'>
					<IconPhotoSolid />
				</button>
				<h5>Documents & Photos</h5>
			</div>
			<div class='bg-accent-content rounded-md w-1/2 p-1'>
				<button class='btn btn-square btn-ghost btn-xs text-secondary'>
					<IconDocumentTextSolid />
				</button>
				<h5>Completion Docs</h5>
			</div>
		</div>

		<!-- section:LIST/ values -->
		<ul class='my-4 space-y-2'>
			<li class='flex justify-between border-b'>
				<h5>Job Cost</h5>
				<h6>{usdFormat(job?.price)}</h6>
			</li>

			<li class='flex justify-between border-b'>
				<h5>Material Delivery</h5>
				<h6>{job?.materialDate}</h6>
			</li>

			<li class='flex justify-between border-b'>
				<h5>Removal Date</h5>
				<h6>{job?.removeDate}</h6>
			</li>

			<li class='flex justify-between border-b'>
				<h5>Install Date</h5>
				<h6>{job?.installDate}</h6>
			</li>

			<li class='flex justify-between border-b'>
				<h5>Completion Date</h5>
				<h6>{job?.completionDate}</h6>
			</li>
		</ul>

		<h5>Special Note:</h5>
		<p>{job?.note || 'N/A'}</p>
		<div class='card-actions justify-center mt-4'>
			<button class='btn btn-primary btn-block'>Submit For Payment</button>
		</div>
	</div>
</div>

