<script lang="ts">

  import InputNumber from "$lib/components/form/InputNumber.svelte";
  import DropDown from "$lib/components/form/DropDown.svelte";
  import H3WithBorder from "$lib/components/H3WithBorder.svelte";
  import {getContextClient} from "@urql/svelte";
  import {QryOptionList} from "$lib/gql/options";
  import alerts from "$lib/stores/alerts";
  import {SavePartnerCompletedSteps, SavePartnerOperations} from "$lib/gql/partner";
  import {isProfilePage, partner, stepName} from "$lib/components/onboard/solar/index";
  import {page} from "$app/stores";
  import {goto} from "$app/navigation";
  import {authUser} from "$lib/stores/auth";
  import BtnSubmit from "$lib/components/form/BtnSubmit.svelte";
  import DropDownMultiple from "$lib/components/form/DropDownMultiple.svelte";
  import {PartnerStatus} from "$lib/graph/graphql";
  import InputChk from "$lib/components/form/InputChk.svelte";
  import InputChkMulti from "$lib/components/form/InputChkMulti.svelte";

  const client = getContextClient();
  const stepNo = 4;
  const insOptions = [
    {id: undefined, name: ''},
    {id: true, name: 'Yes'},
    {id: false, name: 'No'},
  ]

  let busy = false;
  let loading = true;
  let finances = [];
  let epc = [];
  let data;

  $:partnerID = $page.params.id;
  $:asProfile = $isProfilePage;
  $:data = $partner || {};

  client.query(QryOptionList, {types: ["FINANCE", "EPC"]}).toPromise().then((res) => {
    loading = false;
    if (res.error) {
      alerts.error('Options', res.error.message);
      return
    }

    const d = res.data.optionList
    finances = d.find(ol => ol.type === 'FINANCE').options
    epc = d.find(ol => ol.type === 'EPC').options
  })


  async function save() {
    if (busy) return;

    busy = true;
    const res = await client.mutation(SavePartnerOperations, {
      partnerID, input: {
        salesVolume: data.salesVolume,
        financeOptions: data.financeOptions,
        downPayment: data.downPayment,
        pifDate: data.pifDate,
        installInHouse: data.installInHouse,
        epcOptions: data.epcOptions,
      }
    }).toPromise();

    if (res.error) {
      busy = false;
      alerts.error(`Error:${$stepName}`, res.error.message);
      return;
    }


    // mark step
    if ($authUser?.partner?.status === PartnerStatus.Onboarding) {
      await client.mutation(SavePartnerCompletedSteps, {partnerID, step: stepNo, done: true}).toPromise();
    }

    busy = false;
    alerts.success('Save', 'Saved Successfully');
    partner.update(p => {
      if (p.setupStepsCompleted < stepNo) {
        p.setupStepsCompleted = stepNo;
      }
      return p;
    });

    if ($authUser?.partner && !asProfile) {
      await goto(`/${$authUser?.partner.type.toLowerCase()}/onboarding/done`);
    }
  }

</script>


<div class="md:w-1/3">
  <form on:submit|preventDefault={save}>
    <H3WithBorder>Operations</H3WithBorder>
    <InputNumber label="Average Monthly Sales Volume" max="99999" bind:value={data.salesVolume} required/>
    <InputChkMulti label="What finance options do you use?" options={finances} bind:value={data.financeOptions}/>
    <br>
    <H3WithBorder>Payment Terms</H3WithBorder>
    <InputNumber label="Down Payment %" min="0" bind:value={data.downPayment}/>
    <InputNumber label='PIF Date' min="0" max="30" bind:value={data.pifDate}/>
    <br>
    <H3WithBorder>Fulfillment</H3WithBorder>
    <DropDown label="Do you install your solar projects in house?" options={insOptions}
              bind:value={data.installInHouse}/>
    {#if (data.installInHouse === false)}
      <InputChkMulti label="What EPC’s do you use?" options={epc} bind:value={data.epcOptions}/>
    {/if}

    <BtnSubmit cls='mt-5' loading={busy}>
      {asProfile ? 'SAVE' : 'DONE'}
    </BtnSubmit>
  </form>
</div>
