<script>
  import {DataColumn, DataTable} from "$lib/components/DataTable";
  import {PartnerContactSendPwdResetEmail, QryPartnerContacts} from "$lib/gql/partner.contact";
  import {onMount} from "svelte";
  import {wsMessage} from "$lib/stores/socket";
  import {Channel, PartnerContactSendPwdResetEmailDocument, Topic} from "$lib/graph/graphql";
  import BtnView from "$lib/components/form/BtnView.svelte";
  import {goto} from "$app/navigation";
  import {humanize} from "$lib/utils/string";
  import Avatar from "$lib/components/Avatar.svelte";
  import {page} from "$app/stores";
  import {editUser} from "./store";
  import ConfirmAction from "$lib/components/ConfirmAction.svelte";
  import IconEmail from "$lib/assets/svg/IconEmail.svelte";
  import {getContextClient} from "@urql/svelte";
  import alerts from "$lib/stores/alerts";

  export let partnerID;
  export let editPath;

  const client = getContextClient();
  const titleResetPwd = "Password Reset Email";

  let count;
  let variables;
  $: variables = {
    partnerID,
    search: $page.url.searchParams.get('qu') || ''
  };

  onMount(() => {
    editUser.set(null);
    return wsMessage.subscribe(msg => {
      if (!msg) return;
      if (Channel.PartnerUser === msg.channel && Topic.Created === msg.topic) {
        variables._v = Date.now();
      }
    })
  })

  async function resetPwd(userID) {
    const res = await client.mutation(PartnerContactSendPwdResetEmail, {partnerID, userID}).toPromise();
    if (res.error) {
      return Promise.reject(res.error.message);
    }
    alerts.success(titleResetPwd, "Successfully sent email to reset password")
  }

</script>

<div class='w-full sm:w-full'>
  <div class='card-body p-0'>
    <DataTable
      headers={[
        'NAME',
        'EMAIL',
        'PHONE',
        'TYPE',
        'ROLE',
        'ACCOUNT STATUS',
        'PASSWORD RESET',
        'VIEW'
      ]}
      query={QryPartnerContacts}
      dataProp='partnerContacts'
      {variables}
      searchable
      searchProp='qu'
      bind:totalCount={count}
    >
      <svelte:fragment slot='row' let:item let:headers>
        <DataColumn header={headers[0]}>
          <div class="flex flex-row space-x-2">
            <div>
              <Avatar name="{item.firstName} {item.lastName}" src={item.picture} className=""/>
            </div>
            <div class="mt-2">
              {item.firstName} {item.lastName}
            </div>
          </div>
        </DataColumn>
        <DataColumn header={headers[1]}>
          {item.email}
        </DataColumn>
        <DataColumn header={headers[2]}>
          {item.phone}
        </DataColumn>
        <DataColumn header={headers[3]}>
          {humanize(item.type)}
        </DataColumn>
        <DataColumn header={headers[3]}>
          {humanize(item.role)}
        </DataColumn>
        <DataColumn header={headers[4]}>
          {item.accountStatus}
        </DataColumn>
        <DataColumn header={headers[5]}>
          <ConfirmAction
            title={titleResetPwd}
            message="Are you sure you want to send password-reset email to this user?"
            okText="SEND"
            confirm={() => resetPwd(item.userID)}
          >
            <IconEmail/>
            Password Reset
          </ConfirmAction>
        </DataColumn>
        <DataColumn header={headers[6]}>
          <BtnView
            tooltip="View User Details"
            on:click={() => {
              editUser.set(item);
              goto(editPath.replace(':id', item.id)+$page.url.search)
          }}
          />
        </DataColumn>
      </svelte:fragment>
    </DataTable>
  </div>
</div>
