<script lang='ts'>
  import { getContextClient, queryStore } from '@urql/svelte';
  import { QryNearmapResponse } from '$lib/gql/job';
  import { page } from '$app/stores';
  import PlaceholderBasic from '$lib/components/placeholder/PlaceholderBasic.svelte';
  import Page from '$lib/components/Page.svelte';

  const client = getContextClient();
  const store = queryStore({
    client,
    query: QryNearmapResponse,
    variables: { id: $page.params.id, respID: $page.params.respID }
  });

  const rollupDictionary: Array<{ key: string, name: string, break?: boolean }> = [
    { key: '51cbb913-1c5c-5253-b333-3789e34839de', name: 'roof count' },
    { key: '31225e77-2159-5509-a83c-838c18445dbf', name: 'pitch degrees of primary roof' },
    { key: 'b04d6b59-5307-57d9-a763-dfa64100670a', name: 'total unclipped area (sq. ft) of roof', break: true },

    { key: 'a5d04b2f-7f29-5279-80a1-3403788e1911', name: 'flat (roof types) area (sq. ft) of primary roof' },
    { key: '9a6dcdc5-ac39-52a7-9f3c-9b473784a655', name: 'flat (roof types) ratio in primary roof', break: true },

    { key: '9ae46172-8a4b-553e-9f36-e325efedfc61', name: 'shingle (roof material) is dominant in primary roof' },
    { key: 'ea83ca07-b493-5232-a60d-6165d7b8535b', name: 'shingle (roof material) area (sq. ft) of primary roof' },
    { key: '4a5af13b-d1a1-5169-8284-28ffdd1d950c', name: 'shingle (roof material) ratio in primary roof' },
    { key: 'de592b36-efde-54ac-b532-8a4692ca5c09', name: 'total unclipped area (sq. ft) of shingle', break: true },

    { key: '18a30033-ec67-5e98-8c76-918f7bb207c8', name: 'tile (roof material) is dominant in primary roof' },
    { key: '3be5bbdd-fb8a-56c8-9c05-83224e31323b', name: 'tile (roof material) area (sq. ft) of primary roof' },
    { key: 'ebf8f274-19b4-5a1f-9ebf-f3e3d3f1eee9', name: 'tile (roof material) ratio in primary roof' },
    { key: 'e654c744-f3e8-5ec8-a4cd-4157891d9732', name: 'total unclipped area (sq. ft) of tile', break: true },

    { key: '50031e0b-98d1-5524-94b7-bde691eaee76', name: 'metal (roof material) is dominant in primary roof' },
    { key: '97ff0211-98a5-53f9-a07e-0d0ddc5f8fe5', name: 'metal (roof material) area (sq. ft) of primary roof' },
    { key: '0491594d-e331-5161-835b-7b3411072c60', name: 'metal (roof material) ratio in primary roof' },
    { key: 'e59c6f1b-7902-532c-aff2-c9df1461a5bf', name: 'total unclipped area (sq. ft) of metal', break: true },

    { key: '0e107f3a-1942-5172-91ad-512dc1c19978', name: 'gable (roof types) area (sq. ft) of primary roof' },
    { key: '58989579-1f7c-50e7-a5a5-33eab090d0c6', name: 'gable (roof types) ratio in primary roof' },
    { key: '77fbe573-16fb-5217-9595-b6332b444b36', name: 'total unclipped area (sq. ft) of gable', break: true },

    { key: '522cbd32-14ed-5450-b218-19ef92f92edf', name: 'dutch gable (roof types) area (sq. ft) of primary roof' },
    { key: '56a1e713-3675-5bb8-a5f6-d2cd029bd84f', name: 'dutch gable (roof types) ratio in primary roof' },
    { key: '6f55d682-fc06-5ad0-aace-333b9edb1d2e', name: 'total unclipped area (sq. ft) of dutch gable', break: true },

    { key: '2869252c-7972-5228-b620-6c66f24cb964', name: 'hip (roof types) area (sq. ft) of primary roof' },
    { key: '3c600ebb-0f2e-56e0-8b43-131637a019d5', name: 'hip (roof types) ratio in primary roof' },
    { key: '93849242-f355-5886-8180-322824b0eb9e', name: 'total unclipped area (sq. ft) of hip', break: true },

    { key: 'f4971ff5-e254-5aa7-9c33-16f5939044df', name: 'turret (roof types) area (sq. ft) of primary roof' },
    { key: 'a169a6af-7968-56e6-b47a-66a8308b39ad', name: 'turret (roof types) ratio in primary roof' },
    { key: 'ed7effb9-491c-5a38-98a1-695bd7dd768c', name: 'total unclipped area (sq. ft) of turret', break: true },

    { key: '6132706b-622f-5c68-ad85-3b3f7d6b6f5b', name: 'tree overhang present' },
    { key: '6196e9dc-db51-5814-858b-b664a1e160a2', name: 'tree overhang count' },
    {
      key: 'eeaff46c-9a32-515b-a192-57b026f9cec6',
      name: 'tree overhang (roof overhang attributes) area (sq. ft) of primary roof'
    },
    {
      key: '83c0010e-045a-5e35-81bb-6bd0695c1212',
      name: 'tree overhang (roof overhang attributes) ratio in primary roof'
    },
    {
      key: '8b0cc2ee-c529-5439-b7af-d843b7910b0c',
      name: 'tree overhang (roof overhang attributes) confidence in primary roof'
    },
    { key: 'e817a248-b17c-52f2-b6a8-a684f7eab50c', name: 'tree overhang presence confidence' },
    { key: '01b02231-d840-548d-978b-171955619b2b', name: 'total clipped area (sq. ft) of tree overhang' },
    { key: '1d530a3c-efda-54ba-952d-3cf8f1f6f70d', name: 'total unclipped area (sq. ft) of tree overhang' }
  ];

  $:console.log($page.params);
  $:data = $store.data?.nearmapResponse?.detail;
  $:rollups = $store.data?.nearmapResponse?.raw?.rollups || {};
  $:console.log($store.data?.nearmapResponse?.raw?.rollups || {});
</script>


<Page title='Nearmap Estimate Response'>
  {#if $store.fetching}
    <PlaceholderBasic />
  {:else if $store.data?.nearmapResponse}
    <div class='flex flex-row gap-x-4 '>
      <div>
        <table class='table table-compact'>
          <tr>
            <td class='text-xl'>Tile Area</td>
            <td>{data?.tileArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Tile Ratio</td>
            <td>{data?.tileRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Shingle Area</td>
            <td>{data?.shingleArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Shingle Ratio</td>
            <td>{data?.shingleRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Metal Area</td>
            <td>{data?.metalArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Metal Ratio</td>
            <td>{data?.metalRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Flat Area</td>
            <td>{data?.flatArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Flat Ratio</td>
            <td>{data?.flatRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Gable Area</td>
            <td>{data?.gableArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Gable Ratio</td>
            <td>{data?.gableRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Hip Area</td>
            <td>{data?.hipArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Hip Ratio</td>
            <td>{data?.hipRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Dutch Gable Area</td>
            <td>{data?.dutchGableArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Dutch Gable Ratio</td>
            <td>{data?.dutchGableRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Turret Area</td>
            <td>{data?.turretArea}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Turret Ratio</td>
            <td>{data?.turretRatio}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Dominant Roof Type</td>
            <td>{data?.dominantRoofMaterialID}</td>
            <td>none = 0, shingle = 1, metal = 2 & tile = 3</td>
          </tr>
          <tr>
            <td class='text-xl'>Roof Count</td>
            <td>{data?.roofCount}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Primary Pitch (degrees)</td>
            <td>{data?.primaryPitchInDegrees}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Primary Pitch (x/12)</td>
            <td>{data?.primaryPitch}</td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Total Unclipped Area</td>
            <td>{data?.totalUnclippedArea}</td>
            <td></td>
          </tr>
          <tr>
            <td></td>
            <td></td>
            <td></td>
          </tr>
          <tr>
            <td class='text-xl'>Roof Material Ratio Total</td>
            <td>{data?.roofMaterialRatioTotal}</td>
            <td>shingle, metal & tile</td>
          </tr>
          <tr>
            <td class='text-xl'>Roof Type Ratio Total</td>
            <td>{data?.roofTypeRatioTotal}</td>
            <td>flat, turret, gable, dutch gable & hip</td>
          </tr>
          <tr>
            <td class='text-xl'>Roof Material Area Total</td>
            <td>{data?.roofMaterialSurfaceAreaTotal}</td>
            <td>shingle, metal & tile</td>
          </tr>
          <tr>
            <td class='text-xl'>Roof Type Area Total</td>
            <td>{data?.roofTypeSurfaceAreaTotal}</td>
            <td>flat, turret, gable, dutch gable & hip</td>
          </tr>
        </table>
      </div>
      <div>
        <table class='table table-compact'>
          {#each rollupDictionary as d}
            <tr>
              <td class='text-xl'>{d.name}:</td>
              <td>{rollups[d.key]?.value}</td>
            </tr>
            {#if d.break}
              <tr>
                <td colspan='2'></td>
              </tr>
            {/if}
          {/each}
        </table>
      </div>
    </div>
    <br>
    <div>
      <h2>Full Roollup Feilds</h2>
      <table class='table table-compact'>
        {#each Object.values(rollups) as v}
          <tr>
            <td class='text-xl'>{v.description}:</td>
            <td>{v.value}</td>
          </tr>
        {/each}
      </table>
    </div>
  {/if}
</Page>
