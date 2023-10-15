<!--
  - Copyright (c) Simsaw Software Pvt. Ltd. 2022.
  - Author: Ankit Patial <ankit@simsaw.com>
  -->

<script lang='ts'>
  import { PATH } from '$lib/enum';
  import { authUser, menuCounts } from '$lib/stores/auth';
  import { page } from '$app/stores';
  import NavLink from '$lib/components/NavLink.svelte';
  import IconGraph from '$lib/assets/svg/IconGraph.svelte';
  import IconSolarPower from '$lib/assets/svg/IconSolarPower.svelte';
  import IconJob from '$lib/assets/svg/IconJob.svelte';
  import IconDollar from '$lib/assets/svg/IconDollar.svelte';
  import IconRoof from '$lib/assets/svg/IconRoof.svelte';
  import IconUsers from '$lib/assets/svg/IconUsers.svelte';
  import IconClipboardDocuments from '$lib/assets/svg/IconClipboardDocuments.svelte';
  import IconApiAccess from '$lib/assets/svg/IconApiAccess.svelte';
  import IconApiUser from '$lib/assets/svg/IconApiUser.svelte';
  import IconDocumentMagnifying from '$lib/assets/svg/IconDocumentMagnifying.svelte';
  import IconRequestAnEstimate from '$lib/assets/svg/IconRequestAnEstimate.svelte';
  import IconViewEstimates from '$lib/assets/svg/IconViewEstimates.svelte';
  import IconAssignedJobs from '$lib/assets/svg/IconAssignedJobs.svelte';
  import IconUnassignedJobs from '$lib/assets/svg/IconUnassignedJobs.svelte';
  import IconVideoCamera from '$lib/assets/svg/IconVideoCamera.svelte';
  import IconUser from '$lib/assets/svg/IconUser.svelte';
  import { PartnerType, Role } from '$lib/graph/graphql';
  import IconPlusSmall from '$lib/assets/svg/IconPlusSmall.svelte';
  import IconFolder from '$lib/assets/svg/IconFolder.svelte';

  let adminMenuItems = [];
  let solarPartnerMenu = [];
  let epcPartnerMenu = [];
  let integrationPartnerMenu = [];
  let roofingPartnerMenu = [];

  $: counts = $menuCounts;
  $: if ($authUser?.isAdmin) {
    adminMenuItems = [
      { href: PATH.OVERVIEW, label: 'COMPANY OVERVIEW', icon: IconGraph },
      {
        label: 'SERVICES',
        items: [
          {
            label: 'ROOFING',
            items: [
              [
                {
                  href: PATH.ESTIMATES_REQ,
                  label: 'Request An Estimate',
                  icon: IconRequestAnEstimate
                },
                {
                  href: PATH.ESTIMATES_VIEW,
                  label: counts
                    ? `View Estimates (${counts?.estimateCount})`
                    : 'View Estimates',
                  icon: IconViewEstimates
                }
              ],
              [
                {
                  href: PATH.JOBS_UNASSIGNED,
                  label: counts
                    ? `Unassigned Jobs (${counts?.unassignedJobCount})`
                    : 'Unassigned Jobs',
                  icon: IconUnassignedJobs
                },
                {
                  href: PATH.JOBS_ASSIGNED,
                  label: counts
                    ? `Assigned Jobs (${counts?.assignedJobCount})`
                    : 'Assigned Jobs',
                  icon: IconAssignedJobs
                }
              ],
              [
                {
                  href: PATH.PAYMENTS_PENDING,
                  label: counts
                    ? `Pending Approval (${counts?.paymentsPending})`
                    : 'Pending Approval',
                  icon: IconDollar
                },
                {
                  href: PATH.PAYMENTS_APPROVED,
                  label: counts
                    ? `Approved (${counts?.paymentsApproved})`
                    : 'Approved',
                  icon: IconDollar
                },
                {
                  href: PATH.PAYMENTS_COMPLETED,
                  label: counts
                    ? `Completed (${counts?.paymentsCompleted})`
                    : 'Completed',
                  icon: IconDollar
                }
              ]
            ]
          },
          {
            label: 'SURVEYS',
            items: [
              { href: PATH.SITE_SURVEYS, label: 'VIEW DASHBOARD', icon: IconJob }
            ]
          },
          {
            label: 'SMART HOME',
            items: [
              { href: PATH.SMART_HOME, label: 'Select Package', icon: IconRequestAnEstimate },
              { href: PATH.SMART_HOME_PENDING, label: 'View Pending Jobs', icon: IconViewEstimates },
              { href: PATH.SMART_HOME_APPROVED, label: 'View Approved Jobs', icon: IconUnassignedJobs }
            ]
          },
          { href: PATH.MPU, label: 'MPU\'s', icon: IconJob },
          { href: PATH.INSULATION, label: 'Insulation', icon: IconJob },

          {
            label: 'HVAC',
            items: [
              { href: PATH.HVAC, label: 'Select Package', icon: IconRequestAnEstimate },
              { href: PATH.HVAC_PENDING, label: 'View Pending Jobs', icon: IconViewEstimates },
              { href: PATH.HVAC_APPROVED, label: 'View Approved Jobs', icon: IconUnassignedJobs }
            ]
          },
          { href: PATH.TURF, label: 'Turf', icon: IconJob },
          { href: PATH.WINDOWS, label: 'Windows', icon: IconJob },
          { href: PATH.BATTERIES, label: 'Batteries', icon: IconJob }
        ]
      },
      {
        label: 'PARTNERS',
        items: [
          { href: PATH.PARTNER_ROOFING, label: 'Roofing Partner', icon: IconRoof },
          { href: PATH.PARTNER_SOLAR, label: 'Solar Partner', icon: IconSolarPower },
          { href: PATH.PARTNER_EPC, label: 'EPC', icon: IconSolarPower },
          { href: PATH.PARTNER_INTEGRATION, label: 'Integration Partner', icon: IconSolarPower }
        ]
      },
      {
        label: 'PAYMENTS',
        items: [
          { href: PATH.PAYMENTS_PENDING, label: 'Pending Approval', icon: IconDollar },
          { href: PATH.PAYMENTS_APPROVED, label: 'Approved', icon: IconDollar },
          { href: PATH.PAYMENTS_COMPLETED, label: 'Completed', icon: IconDollar }
        ]
      },
      {
        label: 'TRAINING',
        items: [
          { href: PATH.TRAINING_CENTER, label: 'Training Center', icon: IconVideoCamera }
        ]
      },
      {
        label: 'SETTINGS',
        items: [
          { href: PATH.PRICING, label: 'Pricing', icon: IconDocumentMagnifying },
          { href: PATH.USERS, label: 'Users', icon: IconUsers },
          { href: PATH.API_USERS, label: 'API Users', icon: IconApiUser },
          { href: PATH.API_ACCESS, label: 'API Access', icon: IconApiAccess },
          { href: PATH.SERVICE_AREAS, label: 'Service Areas', icon: IconDocumentMagnifying },
          { href: PATH.PRODUCT_PACKAGES, label: 'Product Packages', icon: IconDocumentMagnifying },
          { href: PATH.AUDIT_LOGS, label: 'Audit Logs', icon: IconDocumentMagnifying },
          { href: PATH.DOCS, label: 'Documentation', icon: IconClipboardDocuments, target: '_blank' }
        ]
      }
    ];
  }

  $: if ($authUser?.partner) {
    const settings = {
      label: 'SETTINGS',
      items: [
        { href: `${PATH.PROFILE}`, label: 'Personal Profile', icon: IconUser }
      ]
    };

    if ($authUser.isCompanyAdmin) {
      settings.items.push({
        href: `/${($authUser?.partner?.type || '').toLowerCase()}/${$authUser?.partner?.id}`,
        label: 'Company Profile',
        icon: IconUsers
      });
      settings.items.push({
        label: 'Users',
        items: [
          { href: PATH.PARTNER_ADD_USER, label: 'Add New User', icon: IconPlusSmall },
          { href: PATH.PARTNER_VIEW_USERS, label: 'View All Users', icon: IconViewEstimates }
        ]
      });
    }

    switch ($authUser?.partner.type) {
      case PartnerType.Roofing:
        roofingPartnerMenu = [
          // { href: PATH.ROOFING_PARTNER, label: 'Overview', icon: IconGraph },
          { href: PATH.ROOFING_PARTNER_JOBS, label: 'Jobs Dashboard', icon: IconAssignedJobs },
          { href: PATH.ROOFING_PARTNER_TRAINING, label: 'Training Center', icon: IconVideoCamera },
          settings
        ];
        break;
      case PartnerType.Solar:
        solarPartnerMenu = [
          // { href: PATH.SOLAR_PARTNER, label: 'Overview', icon: IconGraph },
          {
            label: 'ROOFING ',
            items: [
              {
                href: PATH.SOLAR_PARTNER_EST_REQ,
                label: 'Request An Estimate',
                icon: IconRequestAnEstimate
              },
              {
                href: PATH.SOLAR_PARTNER_EST_COMPLETED,
                label: counts
                  ? `Completed Estimates (${counts?.estimateCount})`
                  : 'Completed Estimates',
                icon: IconViewEstimates
              },
              {
                href: PATH.SOLAR_PARTNER_JOBS,
                label: counts
                  ? `Approved Jobs (${counts?.unassignedJobCount})`
                  : 'Approved Jobs',
                icon: IconAssignedJobs
              }
            ]
          }

        ];

        // HVAC, SITE_SURVEYS, WINDOWS, INSULATION

        if (!($authUser.partner.mobileAppSettings?.hideTabs?.includes('SMART_HOME') || false)) {
          solarPartnerMenu.push({
            label: 'SMART HOME',
            items: [
              { href: PATH.SOLAR_PARTNER_SMART_HOME, label: 'Select Package', icon: IconRequestAnEstimate },
              { href: PATH.SOLAR_PARTNER_SMART_HOME_PENDING, label: 'View Pending Jobs', icon: IconUnassignedJobs },
              { href: PATH.SOLAR_PARTNER_SMART_HOME_APPROVED, label: 'View Approved Jobs', icon: IconAssignedJobs }
            ]
          });
        }

        if (!($authUser.partner.mobileAppSettings?.hideTabs?.includes('HVAC') || false)) {
          solarPartnerMenu.push({
            label: 'HVAC',
            items: [
              { href: PATH.SOLAR_PARTNER_HVAC, label: 'Select Package', icon: IconRequestAnEstimate },
              { href: PATH.SOLAR_PARTNER_HVAC_PENDING, label: 'View Pending Jobs', icon: IconUnassignedJobs },
              { href: PATH.SOLAR_PARTNER_HVAC_APPROVED, label: 'View Approved Jobs', icon: IconAssignedJobs }
            ]
          });
        }

        if (!($authUser.partner.mobileAppSettings?.hideTabs?.includes('TRAINING_CENTER') || false)) {
          solarPartnerMenu.push({
            label: 'TRAINING',
            items: [
              { href: PATH.SOLAR_PARTNER_TRAINING, label: 'Training Center', icon: IconVideoCamera }
            ]
          });
        }

        solarPartnerMenu.push(settings);
        break;
      case PartnerType.Epc:

        epcPartnerMenu = [
          { href: PATH.EPC_PARTNER, label: 'Overview', icon: IconGraph },
          settings
        ];
        break;
      case PartnerType.Integration:
        integrationPartnerMenu = [
          { href: PATH.INTEGRATION_PARTNER, label: 'Overview', icon: IconGraph },
          settings
        ];
        break;
    }
  }


  const unknownUserMenu = [
    { href: PATH.UNKNOWN_PARTNER, label: 'Overview', icon: IconUser }
  ];

  let isAdmin, menuItems = [];
  let aCls = 'link flex text-neutral-focus gap-x-2 p-2 rounded-md border-0 focus:text-primary' +
    ' focus:ease-in focus:duration-100 focus:scale-90 ';
  let ulCss = 'menu text-base-content text-accent p-0';

  $:{
    isAdmin = $authUser?.role === Role.Admin;
    switch ($authUser?.role) {
      case Role.Admin:
        menuItems = adminMenuItems;
        break;
      case Role.SiteUser:
        switch ($authUser?.partner?.type) {
          case PartnerType.Epc:
            menuItems = epcPartnerMenu;
            break;
          case PartnerType.Integration:
            menuItems = integrationPartnerMenu;
            break;
          case PartnerType.Roofing:
            menuItems = roofingPartnerMenu;
            break;
          case PartnerType.Solar:
            menuItems = solarPartnerMenu;
            break;
          default:
            menuItems = unknownUserMenu;
        }
        break;
      default:
        menuItems = unknownUserMenu;
    }
  }
  $:pathname = $page.url.pathname;

</script>

<ul class='menu menu-xs rounded-md max-w-xs w-full space-y-2'>
  {#each menuItems as l1}
    {#if (l1.items || []).length > 0}
      <li>
        <details>
          <summary>
            <IconFolder cls='h-5' />{l1.label}</summary>
          <ul>
            <li>
              {#each l1.items as l2}
                {#if (l2.items || []).length > 0}
                  <details>
                    <summary>
                      <IconFolder cls='h-5' /> {l2.label}</summary>
                    {#if (Array.isArray(l2.items[0]))}
                      {#each l2.items as l3, idx}
                        <ul class:mt-4={idx>0}>
                          {#each l3 as l4}
                            <li>
                              <NavLink item={l4} cls={aCls} />
                            </li>
                          {/each}
                        </ul>
                      {/each}
                    {:else }
                      <ul>
                        {#each l2.items as l3}
                          <li class:border-b={l3.borderB}>
                            <NavLink item={l3} cls={aCls} />
                          </li>
                        {/each}
                      </ul>
                    {/if}

                  </details>
                {:else }
                  <NavLink item={l2} cls={aCls} />
                {/if}
              {/each}
            </li>
          </ul>
        </details>
      </li>
    {:else }
      <li>
        <NavLink item={l1} cls={aCls} />
      </li>
    {/if}
  {/each}
</ul>

<style>
  .collapse-plus .collapse-title:after {
    @apply text-accent -mt-3 text-lg;
  }

  .menu > :where(li) > :where(ul) {
    display: block;
  }

  summary {
    @apply font-medium text-accent capitalize;
    padding: 0.625rem 0.5rem;
    margin-bottom: 5px;
  }

  summary:hover {
    @apply bg-secondary-focus text-white !important ;
  }

  .menu :where(li ul):before {
    top: 0.5rem;
    bottom: 0.5rem;
    background: #64748b;
  }
</style>
