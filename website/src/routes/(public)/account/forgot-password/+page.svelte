<script lang='ts'>
    import {getContextClient} from "@urql/svelte";
    import alerts from "$lib/stores/alerts";
    import {PATH} from '$lib/enum.js';
    import {ForgotPassword} from "$lib/gql/account";
    import Btn from '$lib/components/form/Btn.svelte';

    const client = getContextClient();
    let email = '', busy = false;

    async function submit() {
        const title = 'Forgot Password';

        busy = true;
        const res = await client.mutation(ForgotPassword, {email}).toPromise()
        busy = false;
        if (res.error) {
            alerts.push({type: 'error', title, body: res.error.message});
            return
        }

        alerts.push({type: 'success', title, body: 'please check your email'});
        email = ''
    }
</script>

<form on:submit|preventDefault={submit}>
    <h2 class='card-title font-normal justify-center mt-16 mb-6 uppercase'>
        Forgot Password
    </h2>

	<div class='card-actions justify-center space-y-3'>
		<input
			type='text'
			name='email'
			placeholder='YOUR EMAIL ADDRESS'
			class='input input-bordered border-white w-full bg-white shadow-md text-center focus:border-primary focus:shadow-sm'
			bind:value={email}
			required
		/>
		<Btn type='submit' primary shadow {busy} cls="btn-block">
			SUBMIT
    </Btn>
	</div>
</form>

<p class='text-sm mt-4 text-center'>
	<a href={PATH.LOGIN} class='link'>Go to login</a>
</p>
