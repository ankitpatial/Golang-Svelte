/* eslint-disable */
import * as types from './graphql';
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  mutation ForgotPassword($email: String!){\n    accForgotPwd(email: $email)\n  }": types.ForgotPasswordDocument,
    "\n  mutation SetUserPwd($userID: ID!, $pwd: String!, $confirmPwd: String!){\n    setUserPwd(userID: $userID, pwd: $pwd, confirmPwd: $confirmPwd)\n  }": types.SetUserPwdDocument,
    "\n  mutation UpdateProfile($inp: InputUserProfile!) {\n    updateProfile(input: $inp)\n  }": types.UpdateProfileDocument,
    "\n  query Me {\n    me {\n      id\n      email\n      firstName\n      lastName\n      status\n      role\n      phone\n      picture\n      partner {\n        id\n        type\n        partnerName\n        status\n        contactType\n        role\n        mobileAppSettings {\n          hideTabs\n        }\n      }\n      token\n      isAdmin\n      isCompanyAdmin\n    }\n  }": types.MeDocument,
    "\n  query UserEmailAvailable($id: String!, $email: String!){\n    userEmailAvailable(id: $id, email: $email)\n  }": types.UserEmailAvailableDocument,
    "\n  query ApiAccess($search: String, $page: PageInput!) {\n    apiAccess(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          url\n          username\n        }\n      }\n    }\n  }": types.ApiAccessDocument,
    "\n  mutation SaveApiAccess($input: ApiAccessInput!){\n    saveApiAccess(input: $input)\n  }": types.SaveApiAccessDocument,
    "\n  query ApiUsers($page: PageInput! $where: ApiUserWhereInput) {\n    apiUsers(page: $page where: $where) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          username\n          active\n        }\n      }\n    }\n  }": types.ApiUsersDocument,
    "\n  query ApiUser($id: ID!) {\n    apiUser(id: $id) {\n      id\n      username\n      active\n      cbApiAuth\n      cbApiUrl\n      cbApiUser\n      cbApiPwd\n      cbApiToken\n      cbApiEndpoints\n    }\n  }": types.ApiUserDocument,
    "\n  mutation AddApiUser($username: String!) {\n    addApiUser(username: $username)\n  }": types.AddApiUserDocument,
    "\n  mutation EditApiUser($input: ApiUserInput!) {\n    editApiUser(input: $input)\n  }": types.EditApiUserDocument,
    "\n  mutation ChangeApiUserStatus($id: ID! $isActive: Boolean!) {\n    changeApiUserStatus(id: $id isActive: $isActive)\n  }": types.ChangeApiUserStatusDocument,
    "\n  mutation RefreshApiUserPwd($id: ID!) {\n    refreshApiUserPwd(id: $id)\n  }": types.RefreshApiUserPwdDocument,
    "\n  query AuditLogs($search: String, $page: PageInput!, $orderBy: AuditLogOrder) {\n    auditLogs(search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          createdAt\n          action\n          description\n          ip\n          user {\n            firstName\n            lastName\n          }\n          apiUser {\n            username\n          }\n        }\n      }\n    }\n  }": types.AuditLogsDocument,
    "\n  query ApiUserAuditLogs($id: ID!, $search: String, $page: PageInput!, $orderBy: AuditLogOrder) {\n    apiUserAuditLogs(id: $id, search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          createdAt\n          action\n          description\n          ip\n        }\n      }\n    }\n  }": types.ApiUserAuditLogsDocument,
    "\n  query Overview($filter: Filter!){\n    overview(f: $filter) {\n      id\n      total\n      items {\n        id\n        name\n        count\n      }\n    }\n  }\n": types.OverviewDocument,
    "\n  fragment document on Document {\n    id\n    key\n    name\n    contentType\n    contentSize\n    ready\n  }\n": types.DocumentFragmentDoc,
    "\n  mutation documentUploadUrl(\n    $folder: DocumentFolder!\n    $dir: String!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n    $overwrite: Boolean!\n  ) {\n    documentUploadUrl(\n      doc: {\n        folder: $folder\n        dir: $dir\n        section: $section\n        name: $name,\n        fileName: $fileName\n        contentType: $fileType,\n        contentSize: $fileSize\n        overwrite: $overwrite\n      }\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } ": types.DocumentUploadUrlDocument,
    "\n  mutation PublicDataUploadUrl(\n    $dir: ID!\n    $name: String!\n    $section: PublicDataSection!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    publicDataUploadUrl(\n      entityID: $dir,\n      section: $section\n      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      id\n      key\n      meta\n      publicUrl\n      uploadUrl\n    }\n  }": types.PublicDataUploadUrlDocument,
    "\n  mutation PartnerDocUploadUrl(\n    $partnerID: ID!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    partnerDocUploadUrl(\n      partnerID: $partnerID\n      section: $section\n      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } ": types.PartnerDocUploadUrlDocument,
    "\n  mutation JobDocUploadUrl(\n    $jobID: ID!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    jobDocUploadUrl(\n      jobID: $jobID\n      section: $section\n      doc: {name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } ": types.JobDocUploadUrlDocument,
    "\n  mutation  DeleteDoc($id: ID!){\n    deleteDoc(id: $id)\n  }": types.DeleteDocDocument,
    "\n  query JobDocs($jobID: ID!) {\n    jobDocs(jobID: $jobID) {\n      id\n      name\n      filename\n      section\n      meta\n    }\n  }": types.JobDocsDocument,
    "\n  mutation createEstimate($input: CreateEstimateInput!) {\n    createEstimate(input: $input)\n  }": types.CreateEstimateDocument,
    "\n  mutation ApproveEstimate($input: ApproveEstimateInput!) {\n    approveEstimate(input: $input)\n  }": types.ApproveEstimateDocument,
    "\n  mutation DenyEstimate($input: DenyEstimateInput!) {\n    denyEstimate(input: $input)\n  }": types.DenyEstimateDocument,
    "\n  mutation RemoveDenied($id: ID!) {\n    removeDenied(id: $id)\n  }": types.RemoveDeniedDocument,
    "\n  mutation TestPricing($job: CreateEstimateInput!, $measure: [Measurement!]!) {\n    testPricing(job: $job, measure: $measure) {\n      total\n      summary\n    }\n  }": types.TestPricingDocument,
    "\n  query estimate($id: ID!) {\n    estimate(id: $id) {\n      id\n      createdAt\n      status\n      companyName\n      ...material\n      partial\n      measurementType\n      price\n      priceSummary\n      totalSquares\n      primaryPitch\n      failureReason\n      bounds {\n        lat\n        lng\n      }\n      homeOwner {\n        ...homeOwner\n        latitude\n        longitude\n      }\n      salesRep {\n        ...salesRep\n      }\n      pdf {\n        id\n        key\n        name\n        contentType\n        contentSize\n        ready\n      }\n      creatorName\n    }\n  }   ": types.EstimateDocument,
    "\n  query Estimates(\n    $status: EstimateStatus, $search: String, $dates: [String!], $page: PageInput!,\n  ) {\n    estimates(status: $status, search: $search, dtRange: $dates, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        node {\n          id\n          createdAt\n          status\n          companyName\n          measurementType\n          price\n          homeOwner {\n            ...homeOwner\n          }\n          salesRep {\n            ...salesRep\n          }\n        }\n      }\n    }\n  }\n    ": types.EstimatesDocument,
    "\n  fragment page on PageInfo {\n    startCursor\n    hasNextPage\n    endCursor\n    hasPreviousPage\n  }": types.PageFragmentDoc,
    "\n  fragment userBasic on User {\n    id\n    email\n    name\n    phone\n    role\n    status\n    picture\n  }": types.UserBasicFragmentDoc,
    "\n  fragment homeOwner on HomeOwner {\n    id\n    firstName\n    lastName\n    email\n    phone\n    streetNumber\n    streetName\n    city\n    state\n    zip\n  }": types.HomeOwnerFragmentDoc,
    "\n  fragment salesRep on UserInfo {\n    id\n    firstName\n    lastName\n    email\n    phone\n  }": types.SalesRepFragmentDoc,
    "\n  fragment createdBy on UserInfo {\n    id\n    firstName\n    lastName\n  }": types.CreatedByFragmentDoc,
    "\n  fragment material on Estimate {\n    currentMaterial\n    newRoofingMaterial\n    currentMaterialLowSlope\n    newRoofingMaterialLowSlope\n    redeck\n    layers\n    layer2Material\n    layer3Material\n  }": types.MaterialFragmentDoc,
    "\n  query NewULID{\n    newULID\n  }\n": types.NewUlidDocument,
    "\n  mutation BookInstallation($type: InstallationType!, $pkgID: ID!, $productID: ID,$owner: InstallationOwnerInput!) {\n    bookInstallation(type: $type, pkgID: $pkgID, productID: $productID, owner: $owner)\n  }": types.BookInstallationDocument,
    "\n  mutation ApproveInstallation($input: InstallationApproveInput!) {\n    approveInstallation(input: $input)\n  }": types.ApproveInstallationDocument,
    "\n  mutation DenyInstallation($id: ID!, $reason: String!) {\n    denyInstallation(id: $id, reason: $reason)\n  }": types.DenyInstallationDocument,
    "\n  mutation UndoDenyInstallation($id: ID!) {\n    undoDenyInstallation(id: $id)\n  }": types.UndoDenyInstallationDocument,
    "\n  query PendingInstallations($type: InstallationType!, $approval: Approval, $search: String, $betweenDates: [String!], $page: PageInput!){\n    pendingInstallations(type: $type, approval: $approval, search:  $search, betweenDates: $betweenDates, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      },\n      edges {\n        node {\n          id\n          ownerName\n          ownerAddress\n          ownerEmail\n          ownerPhone\n          pkg\n          price\n          approval\n          status\n          createdAt\n        }\n      }\n    }\n  }": types.PendingInstallationsDocument,
    "\n  query ApprovedInstallations($type: InstallationType!, $status: InstallationStatus, $search: String, $betweenDates: [String!], $page: PageInput!){\n    approvedInstallations(type: $type, status: $status, search: $search, betweenDates: $betweenDates,  page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      },\n      edges {\n        node {\n          id\n          ownerName\n          ownerAddress\n          pkg\n          price\n          approval\n          status\n          approvalAt\n        }\n      }\n    }\n  }": types.ApprovedInstallationsDocument,
    "\n  mutation creatJobNote($jobID: ID!, $note: String!)  {\n    creatJobNote(jobID: $jobID, note:  $note)\n  }": types.CreatJobNoteDocument,
    "\n  mutation editJobNote($jobID: ID!, $noteID: ID!, $note: String!)  {\n    editJobNote(jobID: $jobID, noteID: $noteID, note:  $note)\n  }": types.EditJobNoteDocument,
    "\n  query JobNotes($jobID: ID!) {\n    jobNotes(jobID: $jobID) {\n      id\n      note\n      createdAt\n      updatedAt\n      creator {\n        id\n        firstName\n        lastName\n        picture\n      }\n    }\n  }": types.JobNotesDocument,
    "\n  mutation JobProgressUpdate(\n    $id: ID!, $step: JobProgress!, $stepComplete: Boolean!, $note: String!,$data: ProgressInput\n  ) {\n    jobProgressUpdate(id: $id, step: $step, stepComplete: $stepComplete, note: $note, data: $data)\n  }": types.JobProgressUpdateDocument,
    "\n  mutation JobDelay($id: ID!, $flag: Boolean!, $reason: String!) {\n    jobDelay(id: $id, flag: $flag, reason: $reason)\n  }": types.JobDelayDocument,
    "\n  query JobProgress($id: ID!, $page: PageInput!) {\n    jobProgress(id: $id, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          status\n          statusAt\n          note\n        }\n      }\n    }\n  }": types.JobProgressDocument,
    "\n  query JobCompletedProgress($id: ID!) {\n    jobCompletedProgress(id: $id) {\n      id\n      status\n      statusAt\n      complete\n      note\n    }\n  }": types.JobCompletedProgressDocument,
    "\n\tmutation AssignPartnerToJob($jobID: ID!, $partnerID: ID!) {\n\t\tassignPartnerToJob(jobID: $jobID, partnerID: $partnerID)\n\t}\n": types.AssignPartnerToJobDocument,
    "\n\tquery Job($id: ID!) {\n\t\tjob(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tupdatedAt\n\t\t\thomeOwner {\n\t\t\t\t...homeOwner\n        latitude\n        longitude\n\t\t\t}\n\t\t\tcompanyName\n\t\t\tsalesRep {\n\t\t\t\t...salesRep\n\t\t\t}\n\t\t\testimate {\n\t\t\t\t...material\n\t\t\t\tpartial\n\t\t\t\tmeasurementType\n\t\t\t\tpriceSummary\n\t\t\t}\n\t\t\tprice\n\t\t\tworkOrderPrice\n\t\t\tcreatorName\n\t\t}\n\t}\n\t\n\t\n\t\n": types.JobDocument,
    "\n\tquery UnassignedJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {\n\t\tunassignedJobs(search: $search, betweenDates: $betweenDates, page: $page, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tcreatedAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tepcName\n\t\t\t\t\tprice\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n": types.UnassignedJobsDocument,
    "\n\tquery AssignedJobs(\n\t\t$progress: JobProgress\n\t\t$search: String\n\t\t$betweenDates: [String!]\n\t\t$page: PageInput!\n\t\t$orderBy: JobOrder\n\t) {\n\t\tassignedJobs(progress: $progress, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tcreatedAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tepcName\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t\tprice\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressFlagged\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n": types.AssignedJobsDocument,
    "\n\tquery PendingInvoiceJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {\n\t\tjobsByProgress(status: Invoiced, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tsalesRep {\n\t\t\t\t\t\t...salesRep\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tprice\n\t\t\t\t\tinstallDate\n\t\t\t\t\tinspectionDate\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n": types.PendingInvoiceJobsDocument,
    "\n\tquery PartnerJobStats($partnerType: PartnerType!, $search: String, $skip: Int!, $take: Int!) {\n\t\tpartnerJobStats(partnerType: $partnerType, search: $search, skip: $skip, take: $take) {\n\t\t\tid\n\t\t\tname\n\t\t\tstatus\n\t\t\tnewCount\n\t\t\tnewCountFlagged\n\t\t\tcontactedCount\n\t\t\tcontactedCountFlagged\n\t\t\tconfirmedCount\n\t\t\tconfirmedCountFlagged\n\t\t\tpermittingCount\n\t\t\tpermittingCountFlagged\n\t\t\tscheduledCount\n\t\t\tscheduledCountFlagged\n\t\t\tinProgressCount\n\t\t\tinProgressCountFlagged\n\t\t\tinstalledCount\n\t\t\tinstalledCountFlagged\n\t\t\tinvoicedCount\n\t\t\tinvoicedCountFlagged\n\t\t\tdelayedCount\n\t\t\ttotal\n\t\t\ttotalFlagged\n\t\t}\n\t}\n": types.PartnerJobStatsDocument,
    "\n\tquery PartnerJobs(\n\t\t$partnerID: ID!\n\t\t$search: String\n\t\t$flagged: Boolean\n\t\t$progress: JobProgress\n\t\t$dates: [String!]\n\t\t$page: PageInput!\n\t) {\n\t\tpartnerJobs(\n\t\t\tpartnerID: $partnerID\n\t\t\tsearch: $search\n\t\t\tprogress: $progress\n\t\t\tflagged: $flagged\n\t\t\tdates: $dates\n\t\t\tpage: $page\n\t\t) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tsalesRep {\n\t\t\t\t\t\t...salesRep\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tprice\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressAt\n\t\t\t\t\tprogressFlagged\n\t\t\t\t\tpermitRequired\n\t\t\t\t\tinstallDate\n\t\t\t\t\tinspectionDate\n\t\t\t\t\tinspectionRequired\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t\tepcName\n\t\t\t\t\testimate {\n\t\t\t\t\t\t...material\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n\t\n": types.PartnerJobsDocument,
    "\n\tquery JobEstimates($jobID: ID!) {\n\t\tjobEstimates(jobID: $jobID) {\n\t\t\tid\n\t\t\tstatus\n\t\t\testimates {\n\t\t\t\tid\n\t\t\t\ttotalSquares\n\t\t\t\tprimaryPitch\n\t\t\t\tprice\n\t\t\t\tpriceSummary\n\t\t\t\tbounds {\n\t\t\t\t\tlat\n\t\t\t\t\tlng\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n": types.JobEstimatesDocument,
    "\n\tquery EstimateResponses($id: ID!) {\n\t\testimateResponses(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tdescription\n\t\t\tneed\n\t\t\traw\n\t\t}\n\t}\n": types.EstimateResponsesDocument,
    "\n\tquery NearmapResponse($id: ID!, $respID: ID!) {\n\t\tnearmapResponse(id: $id, respID: $respID) {\n\t\t\traw\n\t\t\tdetail {\n\t\t\t\tprice\n\t\t\t\tprimaryPitchInDegrees\n\t\t\t\tprimaryPitch\n\t\t\t\ttileArea\n\t\t\t\ttileRatio\n\t\t\t\tshingleArea\n\t\t\t\tshingleRatio\n\t\t\t\tmetalArea\n\t\t\t\tmetalRatio\n\t\t\t\tflatArea\n\t\t\t\tflatRatio\n\t\t\t\tgableArea\n\t\t\t\tgableRatio\n\t\t\t\thipArea\n\t\t\t\thipRatio\n\t\t\t\tdutchGableArea\n\t\t\t\tdutchGableRatio\n\t\t\t\tturretArea\n\t\t\t\tturretRatio\n\t\t\t\tdominantRoofMaterial\n\t\t\t\tdominantRoofMaterialID\n\t\t\t\troofCount\n\t\t\t\ttotalUnclippedArea\n\t\t\t\troofMaterialRatioTotal\n\t\t\t\troofTypeRatioTotal\n\t\t\t\troofMaterialSurfaceAreaTotal\n\t\t\t\troofTypeSurfaceAreaTotal\n\t\t\t\ttreeOverhangRatioPrimaryRoof\n\t\t\t\ttreeOverhangConfidencePrimaryRoof\n\t\t\t\ttreeOverhangPresenceConfidence\n\t\t\t\ttreeOverhangAreaPrimaryRoof\n\t\t\t\ttreeOverhangTotalClippedArea\n\t\t\t\ttreeOverhangTotalUnClippedArea\n\t\t\t\ttreeOverhangPresent\n\t\t\t\ttreeOverhangCount\n\t\t\t}\n\t\t}\n\t}\n": types.NearmapResponseDocument,
    "\n  mutation NotificationRead($messageID: ID!) {\n    notificationRead(messageID: $messageID)\n  }": types.NotificationReadDocument,
    "\n  query MyUnreadNotificationsCount {\n    myUnreadNotificationsCount\n  }": types.MyUnreadNotificationsCountDocument,
    "\n  query MyNotifications($page: PageInput!) {\n    myNotifications(page: $page) {\n      totalCount\n      pageInfo {\n        hasPreviousPage\n        startCursor\n        hasNextPage\n        endCursor\n      }\n      edges {\n        node {\n          id\n          channel\n          topic\n          refID\n          title\n          message\n          from\n          unread\n          createdAt\n        }\n      }\n    }\n  }": types.MyNotificationsDocument,
    "\n  query OptionList($types: [OptionListType!])  {\n    optionList(types: $types) {\n      type\n      options {\n        id\n        name\n      }\n    }\n  }\n": types.OptionListDocument,
    "\n  mutation SavePartnerContact($partnerID: ID!, $contact: PartnerContactInput!){\n    savePartnerContact(\n      partnerID: $partnerID\n      contact: $contact\n    )}": types.SavePartnerContactDocument,
    "\n    mutation PartnerContactSendPwdResetEmail($partnerID: ID! $userID: ID!){\n        partnerContactSendPwdResetEmail(\n            partnerID: $partnerID\n            userID: $userID\n        )}": types.PartnerContactSendPwdResetEmailDocument,
    "\n  query partnerContacts(\n    $partnerID: ID!\n    $search: String\n    $page: PageInput!\n  ) {\n    partnerContacts(\n      partnerID: $partnerID\n      search: $search\n      page: $page\n    ) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        node {\n          id\n          userID\n          firstName\n          lastName\n          email\n          phone\n          role\n          accountStatus\n          picture\n          type\n          title\n          description\n        }\n      }\n    }\n  } ": types.PartnerContactsDocument,
    "\n  mutation SavePartnerMobileSettings($id: ID!, $inp: InputMobileAppSettings!) {\n    savePartnerMobileSettings( id: $id, inp: $inp)\n  }": types.SavePartnerMobileSettingsDocument,
    "\n  query PartnerMobileSettings($id: ID!) {\n    partnerMobileSettings( id: $id) {\n      logoURL\n      primaryColor\n      hideTabs\n    }\n  }": types.PartnerMobileSettingsDocument,
    "\n\tfragment partner on Partner {\n\t\tid\n\t\tcreatedAt\n\t\ttype\n\t\tname\n\t\taddress\n\t\twebsite\n\t\tisNationWide\n\t\tstatus\n\t}\n": types.PartnerFragmentDoc,
    "\n\tfragment contact on PartnerContact {\n\t\tid\n\t\tuserID\n\t\ttype\n\t\tfirstName\n\t\tlastName\n\t\temail\n\t\tphone\n\t\totherEmail\n\t\ttitle\n\t\tdescription\n\t}\n": types.ContactFragmentDoc,
    "\n\tmutation SavePartner($input: PartnerInput!) {\n\t\tsavePartner(input: $input)\n\t}\n": types.SavePartnerDocument,
    "\n\tmutation invitePartner($input: InvitePartnerInput!) {\n\t\tinvitePartner(input: $input) {\n\t\t\tid\n\t\t\ttype\n\t\t\tcompanyName\n\t\t\tcontactID\n\t\t\tuserID\n\t\t\tfirstName\n\t\t\tlastName\n\t\t\temail\n\t\t\tphone\n\t\t\tcreatedAt\n\t\t}\n\t}\n": types.InvitePartnerDocument,
    "\n\tmutation SavePartnerCompletedSteps($partnerID: ID!, $step: Int!, $done: Boolean) {\n\t\tsavePartnerCompletedSteps(partnerID: $partnerID, step: $step, done: $done)\n\t}\n": types.SavePartnerCompletedStepsDocument,
    "\n\tmutation SetPartnerActive($partnerID: ID!, $value: Boolean!) {\n\t\tsetPartnerActive(partnerID: $partnerID, value: $value)\n\t}\n": types.SetPartnerActiveDocument,
    "\n\tmutation savePartnerContacts($partnerID: ID!, $contacts: [PartnerContactInput!]!) {\n\t\tsavePartnerContacts(partnerID: $partnerID, contacts: $contacts) {\n\t\t\t...contact\n\t\t}\n\t}\n\t\n": types.SavePartnerContactsDocument,
    "\n\tmutation savePartnerOperations($partnerID: ID!, $input: PartnerOperationInput!) {\n\t\tsavePartnerOperations(partnerID: $partnerID, inp: $input)\n\t}\n": types.SavePartnerOperationsDocument,
    "\n\tmutation SaveService($id: ID!, $partnerID: ID!, $service: Service!, $active: Boolean!) {\n\t\tsaveService(id: $id, partnerID: $partnerID, service: $service, active: $active)\n\t}\n": types.SaveServiceDocument,
    "\n\tmutation SaveLeadTime($partnerID: ID!, $asphalt: String, $metal: String, $tile: String) {\n\t\tsaveLeadTime(partnerID: $partnerID, asphalt: $asphalt, metal: $metal, tile: $tile)\n\t}\n": types.SaveLeadTimeDocument,
    "\n\tmutation SaveServiceState($partnerID: ID!, $state: String!, $licNo: String, $expDate: Time, $proofDocID: String) {\n\t\tsaveServiceState(partnerID: $partnerID, state: $state, licNo: $licNo, expDate: $expDate, proofDocID: $proofDocID)\n\t}\n": types.SaveServiceStateDocument,
    "\n\tmutation SaveServiceCity($partnerID: ID!, $postalID: ID!, $active: Boolean, $licNo: String, $proofDocID: String) {\n\t\tsaveServiceCity(partnerID: $partnerID, postalID: $postalID, active: $active, licNo: $licNo, proofDocID: $proofDocID)\n\t}\n": types.SaveServiceCityDocument,
    "\n\tquery PartnerNameAvailable($id: ID!, $type: PartnerType!, $name: String!) {\n\t\tpartnerNameAvailable(id: $id, type: $type, name: $name)\n\t}\n": types.PartnerNameAvailableDocument,
    "\n\tquery Partner($id: ID!) {\n\t\tpartner(id: $id) {\n\t\t\t...partner\n\t\t\tcrewCount\n\t\t\tjobCapacity\n\t\t\tsetupStepsCompleted\n\t\t\tasphaltLeadT\n\t\t\tmetalLeadT\n\t\t\ttileLeadT\n\t\t\tcontacts {\n\t\t\t\t...contact\n\t\t\t\totherEmail\n\t\t\t\ttitle\n\t\t\t\tdescription\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n": types.PartnerDocument,
    "\n\tquery SolarPartner($id: ID!) {\n\t\tpartner(id: $id, type: SOLAR) {\n\t\t\tid\n\t\t\ttype\n\t\t\tstatus\n\t\t\tname\n\t\t\taddress\n\t\t\twebsite\n\t\t\tyearsInBusiness\n\t\t\tsalesVolume\n\t\t\tfinanceOptions\n\t\t\tdownPayment\n\t\t\tpifDate\n\t\t\tinstallInHouse\n\t\t\tepcOptions\n\t\t\tsetupStepsCompleted\n\t\t\tcontacts {\n\t\t\t\t...contact\n\t\t\t\totherEmail\n\t\t\t\ttitle\n\t\t\t\tdescription\n\t\t\t}\n\t\t}\n\t}\n\t\n": types.SolarPartnerDocument,
    "\n\tquery PartnerOptions($partnerID: ID!) {\n\t\tpartnerOptions(partnerID: $partnerID) {\n\t\t\ttype\n\t\t\toptions {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n": types.PartnerOptionsDocument,
    "\n\tquery Partners($search: String, $partnerType: PartnerType, $status: String, $page: PageInput!) {\n\t\tpartners(search: $search, partnerType: $partnerType, status: $status, page: $page) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\thasPreviousPage\n\t\t\t\thasNextPage\n\t\t\t\tstartCursor\n\t\t\t\tendCursor\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\t...partner\n\t\t\t\t\tcrewCount\n\t\t\t\t\tjobCapacity\n\t\t\t\t\tisActive\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n": types.PartnersDocument,
    "\n\tquery MyJob($id: ID!) {\n\t\tmyJob(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tupdatedAt\n\t\t\tcompanyName\n\t\t\thomeOwner {\n\t\t\t\t...homeOwner\n\t\t\t}\n\t\t\tsalesRep {\n\t\t\t\t...salesRep\n\t\t\t}\n\t\t\testimate {\n\t\t\t\t...material\n\t\t\t\tmeasurementType\n\t\t\t}\n\t\t\tprice\n\t\t\tworkOrderPrice\n\t\t\tcontractor {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n": types.MyJobDocument,
    "\n\tquery PartnerDocs($partnerID: ID!, $section: DocumentSection!) {\n\t\tpartnerDocs(partnerID: $partnerID, section: $section) {\n\t\t\t...document\n\t\t}\n\t}\n\t\n": types.PartnerDocsDocument,
    "\n\tquery PartnerServiceStates($partnerID: ID!) {\n\t\tpartnerServiceStates(partnerID: $partnerID) {\n\t\t\tid\n\t\t\tname\n\t\t\tlicenseNo\n\t\t\tlicenseExpDate\n\t\t\texpand\n\t\t\tcities {\n\t\t\t\tid\n\t\t\t\tactive\n\t\t\t\tlicenseNo\n\t\t\t\tlicenseProof\n\t\t\t\tcityZip\n\t\t\t\tcityName\n\t\t\t}\n\t\t}\n\t\tpartnerDocs(partnerID: $partnerID, section: Proof) {\n\t\t\t...document\n\t\t}\n\t}\n\t\n": types.PartnerServiceStatesDocument,
    "\n\tquery PartnerServices($partnerID: ID!) {\n\t\tpartnerServices(partnerID: $partnerID) {\n\t\t\tid\n\t\t\tservice\n\t\t\tdescription\n\t\t\tactive\n\t\t}\n\t}\n": types.PartnerServicesDocument,
    "\n  mutation MarkServiceArea($id: ID!, $value: Boolean!) {\n    markServiceArea(id: $id, value: $value)\n  }": types.MarkServiceAreaDocument,
    "\n  query AllServiceAreas {\n    allServiceAreas {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }": types.AllServiceAreasDocument,
    "\n  query ServiceStates($q: String!) {\n    serviceStates(q: $q) {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }": types.ServiceStatesDocument,
    "\n  query States($q: String!) {\n    states(q: $q) {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }": types.StatesDocument,
    "\n  query Cities($state: String!, $q: String!, $skip: Int!, $take: Int!) {\n    cities(state: $state, q: $q, skip: $skip, take: $take) {\n      id\n      zip\n      name\n    }\n  }": types.CitiesDocument,
    "\n  query GetPricing{\n    getPricing {\n      id\n      items {\n        id\n        country\n        state\n        stateAbr\n        city\n        zip\n        productId\n        price\n        pricePer\n      }\n      products {\n        id\n        name\n      }\n\n    }\n  }\n": types.GetPricingDocument,
    "\n  mutation SaveProductPackage($input: ProductPackageInput!) {\n    saveProductPackage(input: $input)\n  }": types.SaveProductPackageDocument,
    "\n  mutation SaveProduct($input: ProductInput!) {\n    saveProduct(input: $input)\n  }": types.SaveProductDocument,
    "\n  query ProductPackages($category: ProductType, $search: String $page: PageInput!) {\n    productPackages(category: $category, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }": types.ProductPackagesDocument,
    "\n  query Products($category: ProductType, $search: String $page: PageInput!) {\n    products(category: $category, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          name\n          description\n          price\n          features\n          specialNote\n          image {\n            id\n            publicUrl\n          }\n        }\n      }\n    }\n  }": types.ProductsDocument,
    "\n  query SmartHomePackages($page: PageInput!) {\n    smartHomePackages(page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            description\n            price\n            features\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }": types.SmartHomePackagesDocument,
    "\n  query HVACPackages($page: PageInput!) {\n    hvacPackages(page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            description\n            price\n            features\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }": types.HvacPackagesDocument,
    "\n  query Surveys($search: String, $page: PageInput!, $orderBy: SurveyOrder) {\n    surveys(search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      edges {\n        node {\n          id\n          date\n          slot\n          from\n          to\n          name\n          phone\n          address\n          notes\n          status\n        }\n      }\n    }\n  }\n": types.SurveysDocument,
    "\n  query SurveySlotAvailability($date: String!) {\n    surveySlotAvailability(date: $date) {\n      id\n      name\n      available\n    }\n  }\n": types.SurveySlotAvailabilityDocument,
    "\n  mutation SurveyRequest($date: String!, $slotID: ID!) {\n    surveyRequest(date: $date, slotID: $slotID)\n  }\n": types.SurveyRequestDocument,
    "\n  mutation SurveyReserve($input: SurveyInput!) {\n    surveyReserve(input: $input) {\n      id\n      date\n      slot\n      from\n      to\n      name\n      phone\n      address\n      notes\n      status\n    }\n  }\n": types.SurveyReserveDocument,
    "\n  query SurveyDetails($id: ID!) {\n    surveyDetails(id: $id) {\n      id\n      date\n      slot\n      from\n      to\n      name\n      phone\n      address\n      notes\n      status\n    }\n  }\n": types.SurveyDetailsDocument,
    "\n  fragment videoInfo on TrainingVideo {\n    id\n    title\n    description\n    posterURL\n    videoURL\n  }": types.VideoInfoFragmentDoc,
    "\n  mutation CreateTrainingCourse($name: String!){\n    createTrainingCourse(name: $name) {\n      id\n      name\n    }\n  }": types.CreateTrainingCourseDocument,
    "\n  mutation PartnerTrainingVideoAccess($partnerID: ID!, $videoID: ID!, $enabled: Boolean!) {\n    partnerTrainingVideoAccess(partnerID: $partnerID, videoID: $videoID, enabled: $enabled)\n  }": types.PartnerTrainingVideoAccessDocument,
    "\n  mutation SaveTrainingVideo($inp: InputTrainingVideo!) {\n    saveTrainingVideo(inp: $inp)\n  }": types.SaveTrainingVideoDocument,
    "\n  query TrainingCourses($search: String, $page: PageInput!) {\n    trainingCourses(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          id\n          name\n        }\n      }\n    }\n  }": types.TrainingCoursesDocument,
    "\n  query TrainingVideoKinds {\n    trainingVideoKinds {\n      id\n      name\n    }\n  }": types.TrainingVideoKindsDocument,
    "\n  query TrainingVideoCourses($kind: TrainingType!, $partnerID: ID,  $pageSize: Int) {\n    trainingVideoCourses(kind: $kind, partnerID: $partnerID, pageSize: $pageSize) {\n      id\n      name\n      videos {\n        totalCount\n        pageInfo {\n          ...page\n        }\n        edges {\n          cursor\n          node {\n            ...videoInfo\n            assigned\n          }\n        }\n      }\n    }\n  }  ": types.TrainingVideoCoursesDocument,
    "\n  query TrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $partnerID: ID, $page: PageInput!) {\n    trainingVideos(kind: $kind, courseID: $courseID, search: $search, partnerID: $partnerID, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          ...videoInfo\n          kind\n          assigned\n        }\n      }\n    }\n  } ": types.TrainingVideosDocument,
    "\n  query MyTrainingVideoKinds {\n    myTrainingVideoKinds {\n      id\n      name\n    }\n  }": types.MyTrainingVideoKindsDocument,
    "\n  query MyTrainingVideoCourses($kind: TrainingType!,  $pageSize: Int) {\n    myTrainingVideoCourses(kind: $kind,  pageSize: $pageSize) {\n      id\n      name\n      videos {\n        totalCount\n        pageInfo {\n          ...page\n        }\n        edges {\n          cursor\n          node {\n            ...videoInfo\n          }\n        }\n      }\n    }\n  }  ": types.MyTrainingVideoCoursesDocument,
    "\n  query MyTrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $page: PageInput!) {\n    myTrainingVideos(kind: $kind, courseID: $courseID, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          ...videoInfo\n          kind\n        }\n      }\n    }\n  } ": types.MyTrainingVideosDocument,
    "\n  query User($id: ID!) {\n    user(id: $id) {\n      id\n      email\n      firstName\n      lastName\n      phone\n      role\n      status\n      note\n    }\n  }": types.UserDocument,
    "\n  query Users($page: PageInput! $where: UserWhereInput) {\n    users(page: $page where: $where) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          ...userBasic\n        }\n      }\n    }\n  }": types.UsersDocument,
    "\n  query usersSearch($search: String) {\n    usersSearch(search: $search) {\n      id\n      firstName\n      lastName\n      email\n      phone\n      partnerID\n      partnerName\n      partnerContactTypeID\n      partnerContactTitle\n    }\n  }": types.UsersSearchDocument,
    "\n  query myCompanyUsers($search: String, $page: PageInput!) {\n    myCompanyUsers(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          firstName\n          lastName\n          email\n          phone\n          picture\n        }\n      }\n    }\n  }": types.MyCompanyUsersDocument,
    "\n  mutation CreateUser($input: CreateUserInput!) {\n    createUser(input: $input)\n  }": types.CreateUserDocument,
    "\n  mutation UpdateUser($input: UpdateUserInput!) {\n    updateUser(input: $input)\n  }": types.UpdateUserDocument,
    "\n  mutation SaveNotifySettings($id: ID! $topicID: String! $email: Boolean!) {\n    saveNotifySettings(userID: $id topicID: $topicID email: $email)\n  }": types.SaveNotifySettingsDocument,
    "\n  query UserNotifySettings($id: ID!) {\n    userNotifySettings(id: $id) {\n      id\n      topic\n      receiveEmail\n      receiveSMS\n    }\n  }": types.UserNotifySettingsDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation ForgotPassword($email: String!){\n    accForgotPwd(email: $email)\n  }"): (typeof documents)["\n  mutation ForgotPassword($email: String!){\n    accForgotPwd(email: $email)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SetUserPwd($userID: ID!, $pwd: String!, $confirmPwd: String!){\n    setUserPwd(userID: $userID, pwd: $pwd, confirmPwd: $confirmPwd)\n  }"): (typeof documents)["\n  mutation SetUserPwd($userID: ID!, $pwd: String!, $confirmPwd: String!){\n    setUserPwd(userID: $userID, pwd: $pwd, confirmPwd: $confirmPwd)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation UpdateProfile($inp: InputUserProfile!) {\n    updateProfile(input: $inp)\n  }"): (typeof documents)["\n  mutation UpdateProfile($inp: InputUserProfile!) {\n    updateProfile(input: $inp)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Me {\n    me {\n      id\n      email\n      firstName\n      lastName\n      status\n      role\n      phone\n      picture\n      partner {\n        id\n        type\n        partnerName\n        status\n        contactType\n        role\n        mobileAppSettings {\n          hideTabs\n        }\n      }\n      token\n      isAdmin\n      isCompanyAdmin\n    }\n  }"): (typeof documents)["\n  query Me {\n    me {\n      id\n      email\n      firstName\n      lastName\n      status\n      role\n      phone\n      picture\n      partner {\n        id\n        type\n        partnerName\n        status\n        contactType\n        role\n        mobileAppSettings {\n          hideTabs\n        }\n      }\n      token\n      isAdmin\n      isCompanyAdmin\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query UserEmailAvailable($id: String!, $email: String!){\n    userEmailAvailable(id: $id, email: $email)\n  }"): (typeof documents)["\n  query UserEmailAvailable($id: String!, $email: String!){\n    userEmailAvailable(id: $id, email: $email)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ApiAccess($search: String, $page: PageInput!) {\n    apiAccess(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          url\n          username\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query ApiAccess($search: String, $page: PageInput!) {\n    apiAccess(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          url\n          username\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SaveApiAccess($input: ApiAccessInput!){\n    saveApiAccess(input: $input)\n  }"): (typeof documents)["\n  mutation SaveApiAccess($input: ApiAccessInput!){\n    saveApiAccess(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ApiUsers($page: PageInput! $where: ApiUserWhereInput) {\n    apiUsers(page: $page where: $where) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          username\n          active\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query ApiUsers($page: PageInput! $where: ApiUserWhereInput) {\n    apiUsers(page: $page where: $where) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          username\n          active\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ApiUser($id: ID!) {\n    apiUser(id: $id) {\n      id\n      username\n      active\n      cbApiAuth\n      cbApiUrl\n      cbApiUser\n      cbApiPwd\n      cbApiToken\n      cbApiEndpoints\n    }\n  }"): (typeof documents)["\n  query ApiUser($id: ID!) {\n    apiUser(id: $id) {\n      id\n      username\n      active\n      cbApiAuth\n      cbApiUrl\n      cbApiUser\n      cbApiPwd\n      cbApiToken\n      cbApiEndpoints\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation AddApiUser($username: String!) {\n    addApiUser(username: $username)\n  }"): (typeof documents)["\n  mutation AddApiUser($username: String!) {\n    addApiUser(username: $username)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation EditApiUser($input: ApiUserInput!) {\n    editApiUser(input: $input)\n  }"): (typeof documents)["\n  mutation EditApiUser($input: ApiUserInput!) {\n    editApiUser(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation ChangeApiUserStatus($id: ID! $isActive: Boolean!) {\n    changeApiUserStatus(id: $id isActive: $isActive)\n  }"): (typeof documents)["\n  mutation ChangeApiUserStatus($id: ID! $isActive: Boolean!) {\n    changeApiUserStatus(id: $id isActive: $isActive)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation RefreshApiUserPwd($id: ID!) {\n    refreshApiUserPwd(id: $id)\n  }"): (typeof documents)["\n  mutation RefreshApiUserPwd($id: ID!) {\n    refreshApiUserPwd(id: $id)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query AuditLogs($search: String, $page: PageInput!, $orderBy: AuditLogOrder) {\n    auditLogs(search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          createdAt\n          action\n          description\n          ip\n          user {\n            firstName\n            lastName\n          }\n          apiUser {\n            username\n          }\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query AuditLogs($search: String, $page: PageInput!, $orderBy: AuditLogOrder) {\n    auditLogs(search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          createdAt\n          action\n          description\n          ip\n          user {\n            firstName\n            lastName\n          }\n          apiUser {\n            username\n          }\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ApiUserAuditLogs($id: ID!, $search: String, $page: PageInput!, $orderBy: AuditLogOrder) {\n    apiUserAuditLogs(id: $id, search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          createdAt\n          action\n          description\n          ip\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query ApiUserAuditLogs($id: ID!, $search: String, $page: PageInput!, $orderBy: AuditLogOrder) {\n    apiUserAuditLogs(id: $id, search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          createdAt\n          action\n          description\n          ip\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Overview($filter: Filter!){\n    overview(f: $filter) {\n      id\n      total\n      items {\n        id\n        name\n        count\n      }\n    }\n  }\n"): (typeof documents)["\n  query Overview($filter: Filter!){\n    overview(f: $filter) {\n      id\n      total\n      items {\n        id\n        name\n        count\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment document on Document {\n    id\n    key\n    name\n    contentType\n    contentSize\n    ready\n  }\n"): (typeof documents)["\n  fragment document on Document {\n    id\n    key\n    name\n    contentType\n    contentSize\n    ready\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation documentUploadUrl(\n    $folder: DocumentFolder!\n    $dir: String!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n    $overwrite: Boolean!\n  ) {\n    documentUploadUrl(\n      doc: {\n        folder: $folder\n        dir: $dir\n        section: $section\n        name: $name,\n        fileName: $fileName\n        contentType: $fileType,\n        contentSize: $fileSize\n        overwrite: $overwrite\n      }\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } "): (typeof documents)["\n  mutation documentUploadUrl(\n    $folder: DocumentFolder!\n    $dir: String!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n    $overwrite: Boolean!\n  ) {\n    documentUploadUrl(\n      doc: {\n        folder: $folder\n        dir: $dir\n        section: $section\n        name: $name,\n        fileName: $fileName\n        contentType: $fileType,\n        contentSize: $fileSize\n        overwrite: $overwrite\n      }\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation PublicDataUploadUrl(\n    $dir: ID!\n    $name: String!\n    $section: PublicDataSection!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    publicDataUploadUrl(\n      entityID: $dir,\n      section: $section\n      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      id\n      key\n      meta\n      publicUrl\n      uploadUrl\n    }\n  }"): (typeof documents)["\n  mutation PublicDataUploadUrl(\n    $dir: ID!\n    $name: String!\n    $section: PublicDataSection!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    publicDataUploadUrl(\n      entityID: $dir,\n      section: $section\n      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      id\n      key\n      meta\n      publicUrl\n      uploadUrl\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation PartnerDocUploadUrl(\n    $partnerID: ID!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    partnerDocUploadUrl(\n      partnerID: $partnerID\n      section: $section\n      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } "): (typeof documents)["\n  mutation PartnerDocUploadUrl(\n    $partnerID: ID!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    partnerDocUploadUrl(\n      partnerID: $partnerID\n      section: $section\n      doc: { name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation JobDocUploadUrl(\n    $jobID: ID!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    jobDocUploadUrl(\n      jobID: $jobID\n      section: $section\n      doc: {name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } "): (typeof documents)["\n  mutation JobDocUploadUrl(\n    $jobID: ID!\n    $section: DocumentSection!\n    $name: String!\n    $fileName: String!\n    $fileType: String!\n    $fileSize: Int64!\n  ) {\n    jobDocUploadUrl(\n      jobID: $jobID\n      section: $section\n      doc: {name: $name, fileName: $fileName, contentType: $fileType, contentSize: $fileSize}\n    ) {\n      ...document\n      uploadUrl\n      meta\n    }\n  } "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation  DeleteDoc($id: ID!){\n    deleteDoc(id: $id)\n  }"): (typeof documents)["\n  mutation  DeleteDoc($id: ID!){\n    deleteDoc(id: $id)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query JobDocs($jobID: ID!) {\n    jobDocs(jobID: $jobID) {\n      id\n      name\n      filename\n      section\n      meta\n    }\n  }"): (typeof documents)["\n  query JobDocs($jobID: ID!) {\n    jobDocs(jobID: $jobID) {\n      id\n      name\n      filename\n      section\n      meta\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation createEstimate($input: CreateEstimateInput!) {\n    createEstimate(input: $input)\n  }"): (typeof documents)["\n  mutation createEstimate($input: CreateEstimateInput!) {\n    createEstimate(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation ApproveEstimate($input: ApproveEstimateInput!) {\n    approveEstimate(input: $input)\n  }"): (typeof documents)["\n  mutation ApproveEstimate($input: ApproveEstimateInput!) {\n    approveEstimate(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation DenyEstimate($input: DenyEstimateInput!) {\n    denyEstimate(input: $input)\n  }"): (typeof documents)["\n  mutation DenyEstimate($input: DenyEstimateInput!) {\n    denyEstimate(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation RemoveDenied($id: ID!) {\n    removeDenied(id: $id)\n  }"): (typeof documents)["\n  mutation RemoveDenied($id: ID!) {\n    removeDenied(id: $id)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation TestPricing($job: CreateEstimateInput!, $measure: [Measurement!]!) {\n    testPricing(job: $job, measure: $measure) {\n      total\n      summary\n    }\n  }"): (typeof documents)["\n  mutation TestPricing($job: CreateEstimateInput!, $measure: [Measurement!]!) {\n    testPricing(job: $job, measure: $measure) {\n      total\n      summary\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query estimate($id: ID!) {\n    estimate(id: $id) {\n      id\n      createdAt\n      status\n      companyName\n      ...material\n      partial\n      measurementType\n      price\n      priceSummary\n      totalSquares\n      primaryPitch\n      failureReason\n      bounds {\n        lat\n        lng\n      }\n      homeOwner {\n        ...homeOwner\n        latitude\n        longitude\n      }\n      salesRep {\n        ...salesRep\n      }\n      pdf {\n        id\n        key\n        name\n        contentType\n        contentSize\n        ready\n      }\n      creatorName\n    }\n  }   "): (typeof documents)["\n  query estimate($id: ID!) {\n    estimate(id: $id) {\n      id\n      createdAt\n      status\n      companyName\n      ...material\n      partial\n      measurementType\n      price\n      priceSummary\n      totalSquares\n      primaryPitch\n      failureReason\n      bounds {\n        lat\n        lng\n      }\n      homeOwner {\n        ...homeOwner\n        latitude\n        longitude\n      }\n      salesRep {\n        ...salesRep\n      }\n      pdf {\n        id\n        key\n        name\n        contentType\n        contentSize\n        ready\n      }\n      creatorName\n    }\n  }   "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Estimates(\n    $status: EstimateStatus, $search: String, $dates: [String!], $page: PageInput!,\n  ) {\n    estimates(status: $status, search: $search, dtRange: $dates, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        node {\n          id\n          createdAt\n          status\n          companyName\n          measurementType\n          price\n          homeOwner {\n            ...homeOwner\n          }\n          salesRep {\n            ...salesRep\n          }\n        }\n      }\n    }\n  }\n    "): (typeof documents)["\n  query Estimates(\n    $status: EstimateStatus, $search: String, $dates: [String!], $page: PageInput!,\n  ) {\n    estimates(status: $status, search: $search, dtRange: $dates, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        node {\n          id\n          createdAt\n          status\n          companyName\n          measurementType\n          price\n          homeOwner {\n            ...homeOwner\n          }\n          salesRep {\n            ...salesRep\n          }\n        }\n      }\n    }\n  }\n    "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment page on PageInfo {\n    startCursor\n    hasNextPage\n    endCursor\n    hasPreviousPage\n  }"): (typeof documents)["\n  fragment page on PageInfo {\n    startCursor\n    hasNextPage\n    endCursor\n    hasPreviousPage\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment userBasic on User {\n    id\n    email\n    name\n    phone\n    role\n    status\n    picture\n  }"): (typeof documents)["\n  fragment userBasic on User {\n    id\n    email\n    name\n    phone\n    role\n    status\n    picture\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment homeOwner on HomeOwner {\n    id\n    firstName\n    lastName\n    email\n    phone\n    streetNumber\n    streetName\n    city\n    state\n    zip\n  }"): (typeof documents)["\n  fragment homeOwner on HomeOwner {\n    id\n    firstName\n    lastName\n    email\n    phone\n    streetNumber\n    streetName\n    city\n    state\n    zip\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment salesRep on UserInfo {\n    id\n    firstName\n    lastName\n    email\n    phone\n  }"): (typeof documents)["\n  fragment salesRep on UserInfo {\n    id\n    firstName\n    lastName\n    email\n    phone\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment createdBy on UserInfo {\n    id\n    firstName\n    lastName\n  }"): (typeof documents)["\n  fragment createdBy on UserInfo {\n    id\n    firstName\n    lastName\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment material on Estimate {\n    currentMaterial\n    newRoofingMaterial\n    currentMaterialLowSlope\n    newRoofingMaterialLowSlope\n    redeck\n    layers\n    layer2Material\n    layer3Material\n  }"): (typeof documents)["\n  fragment material on Estimate {\n    currentMaterial\n    newRoofingMaterial\n    currentMaterialLowSlope\n    newRoofingMaterialLowSlope\n    redeck\n    layers\n    layer2Material\n    layer3Material\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query NewULID{\n    newULID\n  }\n"): (typeof documents)["\n  query NewULID{\n    newULID\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation BookInstallation($type: InstallationType!, $pkgID: ID!, $productID: ID,$owner: InstallationOwnerInput!) {\n    bookInstallation(type: $type, pkgID: $pkgID, productID: $productID, owner: $owner)\n  }"): (typeof documents)["\n  mutation BookInstallation($type: InstallationType!, $pkgID: ID!, $productID: ID,$owner: InstallationOwnerInput!) {\n    bookInstallation(type: $type, pkgID: $pkgID, productID: $productID, owner: $owner)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation ApproveInstallation($input: InstallationApproveInput!) {\n    approveInstallation(input: $input)\n  }"): (typeof documents)["\n  mutation ApproveInstallation($input: InstallationApproveInput!) {\n    approveInstallation(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation DenyInstallation($id: ID!, $reason: String!) {\n    denyInstallation(id: $id, reason: $reason)\n  }"): (typeof documents)["\n  mutation DenyInstallation($id: ID!, $reason: String!) {\n    denyInstallation(id: $id, reason: $reason)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation UndoDenyInstallation($id: ID!) {\n    undoDenyInstallation(id: $id)\n  }"): (typeof documents)["\n  mutation UndoDenyInstallation($id: ID!) {\n    undoDenyInstallation(id: $id)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query PendingInstallations($type: InstallationType!, $approval: Approval, $search: String, $betweenDates: [String!], $page: PageInput!){\n    pendingInstallations(type: $type, approval: $approval, search:  $search, betweenDates: $betweenDates, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      },\n      edges {\n        node {\n          id\n          ownerName\n          ownerAddress\n          ownerEmail\n          ownerPhone\n          pkg\n          price\n          approval\n          status\n          createdAt\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query PendingInstallations($type: InstallationType!, $approval: Approval, $search: String, $betweenDates: [String!], $page: PageInput!){\n    pendingInstallations(type: $type, approval: $approval, search:  $search, betweenDates: $betweenDates, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      },\n      edges {\n        node {\n          id\n          ownerName\n          ownerAddress\n          ownerEmail\n          ownerPhone\n          pkg\n          price\n          approval\n          status\n          createdAt\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ApprovedInstallations($type: InstallationType!, $status: InstallationStatus, $search: String, $betweenDates: [String!], $page: PageInput!){\n    approvedInstallations(type: $type, status: $status, search: $search, betweenDates: $betweenDates,  page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      },\n      edges {\n        node {\n          id\n          ownerName\n          ownerAddress\n          pkg\n          price\n          approval\n          status\n          approvalAt\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query ApprovedInstallations($type: InstallationType!, $status: InstallationStatus, $search: String, $betweenDates: [String!], $page: PageInput!){\n    approvedInstallations(type: $type, status: $status, search: $search, betweenDates: $betweenDates,  page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      },\n      edges {\n        node {\n          id\n          ownerName\n          ownerAddress\n          pkg\n          price\n          approval\n          status\n          approvalAt\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation creatJobNote($jobID: ID!, $note: String!)  {\n    creatJobNote(jobID: $jobID, note:  $note)\n  }"): (typeof documents)["\n  mutation creatJobNote($jobID: ID!, $note: String!)  {\n    creatJobNote(jobID: $jobID, note:  $note)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation editJobNote($jobID: ID!, $noteID: ID!, $note: String!)  {\n    editJobNote(jobID: $jobID, noteID: $noteID, note:  $note)\n  }"): (typeof documents)["\n  mutation editJobNote($jobID: ID!, $noteID: ID!, $note: String!)  {\n    editJobNote(jobID: $jobID, noteID: $noteID, note:  $note)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query JobNotes($jobID: ID!) {\n    jobNotes(jobID: $jobID) {\n      id\n      note\n      createdAt\n      updatedAt\n      creator {\n        id\n        firstName\n        lastName\n        picture\n      }\n    }\n  }"): (typeof documents)["\n  query JobNotes($jobID: ID!) {\n    jobNotes(jobID: $jobID) {\n      id\n      note\n      createdAt\n      updatedAt\n      creator {\n        id\n        firstName\n        lastName\n        picture\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation JobProgressUpdate(\n    $id: ID!, $step: JobProgress!, $stepComplete: Boolean!, $note: String!,$data: ProgressInput\n  ) {\n    jobProgressUpdate(id: $id, step: $step, stepComplete: $stepComplete, note: $note, data: $data)\n  }"): (typeof documents)["\n  mutation JobProgressUpdate(\n    $id: ID!, $step: JobProgress!, $stepComplete: Boolean!, $note: String!,$data: ProgressInput\n  ) {\n    jobProgressUpdate(id: $id, step: $step, stepComplete: $stepComplete, note: $note, data: $data)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation JobDelay($id: ID!, $flag: Boolean!, $reason: String!) {\n    jobDelay(id: $id, flag: $flag, reason: $reason)\n  }"): (typeof documents)["\n  mutation JobDelay($id: ID!, $flag: Boolean!, $reason: String!) {\n    jobDelay(id: $id, flag: $flag, reason: $reason)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query JobProgress($id: ID!, $page: PageInput!) {\n    jobProgress(id: $id, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          status\n          statusAt\n          note\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query JobProgress($id: ID!, $page: PageInput!) {\n    jobProgress(id: $id, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          status\n          statusAt\n          note\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query JobCompletedProgress($id: ID!) {\n    jobCompletedProgress(id: $id) {\n      id\n      status\n      statusAt\n      complete\n      note\n    }\n  }"): (typeof documents)["\n  query JobCompletedProgress($id: ID!) {\n    jobCompletedProgress(id: $id) {\n      id\n      status\n      statusAt\n      complete\n      note\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation AssignPartnerToJob($jobID: ID!, $partnerID: ID!) {\n\t\tassignPartnerToJob(jobID: $jobID, partnerID: $partnerID)\n\t}\n"): (typeof documents)["\n\tmutation AssignPartnerToJob($jobID: ID!, $partnerID: ID!) {\n\t\tassignPartnerToJob(jobID: $jobID, partnerID: $partnerID)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery Job($id: ID!) {\n\t\tjob(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tupdatedAt\n\t\t\thomeOwner {\n\t\t\t\t...homeOwner\n        latitude\n        longitude\n\t\t\t}\n\t\t\tcompanyName\n\t\t\tsalesRep {\n\t\t\t\t...salesRep\n\t\t\t}\n\t\t\testimate {\n\t\t\t\t...material\n\t\t\t\tpartial\n\t\t\t\tmeasurementType\n\t\t\t\tpriceSummary\n\t\t\t}\n\t\t\tprice\n\t\t\tworkOrderPrice\n\t\t\tcreatorName\n\t\t}\n\t}\n\t\n\t\n\t\n"): (typeof documents)["\n\tquery Job($id: ID!) {\n\t\tjob(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tupdatedAt\n\t\t\thomeOwner {\n\t\t\t\t...homeOwner\n        latitude\n        longitude\n\t\t\t}\n\t\t\tcompanyName\n\t\t\tsalesRep {\n\t\t\t\t...salesRep\n\t\t\t}\n\t\t\testimate {\n\t\t\t\t...material\n\t\t\t\tpartial\n\t\t\t\tmeasurementType\n\t\t\t\tpriceSummary\n\t\t\t}\n\t\t\tprice\n\t\t\tworkOrderPrice\n\t\t\tcreatorName\n\t\t}\n\t}\n\t\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery UnassignedJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {\n\t\tunassignedJobs(search: $search, betweenDates: $betweenDates, page: $page, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tcreatedAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tepcName\n\t\t\t\t\tprice\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n"): (typeof documents)["\n\tquery UnassignedJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {\n\t\tunassignedJobs(search: $search, betweenDates: $betweenDates, page: $page, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tcreatedAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tepcName\n\t\t\t\t\tprice\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery AssignedJobs(\n\t\t$progress: JobProgress\n\t\t$search: String\n\t\t$betweenDates: [String!]\n\t\t$page: PageInput!\n\t\t$orderBy: JobOrder\n\t) {\n\t\tassignedJobs(progress: $progress, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tcreatedAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tepcName\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t\tprice\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressFlagged\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n"): (typeof documents)["\n\tquery AssignedJobs(\n\t\t$progress: JobProgress\n\t\t$search: String\n\t\t$betweenDates: [String!]\n\t\t$page: PageInput!\n\t\t$orderBy: JobOrder\n\t) {\n\t\tassignedJobs(progress: $progress, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tcreatedAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tepcName\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t\tprice\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressFlagged\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PendingInvoiceJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {\n\t\tjobsByProgress(status: Invoiced, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tsalesRep {\n\t\t\t\t\t\t...salesRep\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tprice\n\t\t\t\t\tinstallDate\n\t\t\t\t\tinspectionDate\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n"): (typeof documents)["\n\tquery PendingInvoiceJobs($page: PageInput!, $search: String, $betweenDates: [String!], $orderBy: JobOrder) {\n\t\tjobsByProgress(status: Invoiced, page: $page, search: $search, betweenDates: $betweenDates, orderBy: $orderBy) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressAt\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tsalesRep {\n\t\t\t\t\t\t...salesRep\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tprice\n\t\t\t\t\tinstallDate\n\t\t\t\t\tinspectionDate\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerJobStats($partnerType: PartnerType!, $search: String, $skip: Int!, $take: Int!) {\n\t\tpartnerJobStats(partnerType: $partnerType, search: $search, skip: $skip, take: $take) {\n\t\t\tid\n\t\t\tname\n\t\t\tstatus\n\t\t\tnewCount\n\t\t\tnewCountFlagged\n\t\t\tcontactedCount\n\t\t\tcontactedCountFlagged\n\t\t\tconfirmedCount\n\t\t\tconfirmedCountFlagged\n\t\t\tpermittingCount\n\t\t\tpermittingCountFlagged\n\t\t\tscheduledCount\n\t\t\tscheduledCountFlagged\n\t\t\tinProgressCount\n\t\t\tinProgressCountFlagged\n\t\t\tinstalledCount\n\t\t\tinstalledCountFlagged\n\t\t\tinvoicedCount\n\t\t\tinvoicedCountFlagged\n\t\t\tdelayedCount\n\t\t\ttotal\n\t\t\ttotalFlagged\n\t\t}\n\t}\n"): (typeof documents)["\n\tquery PartnerJobStats($partnerType: PartnerType!, $search: String, $skip: Int!, $take: Int!) {\n\t\tpartnerJobStats(partnerType: $partnerType, search: $search, skip: $skip, take: $take) {\n\t\t\tid\n\t\t\tname\n\t\t\tstatus\n\t\t\tnewCount\n\t\t\tnewCountFlagged\n\t\t\tcontactedCount\n\t\t\tcontactedCountFlagged\n\t\t\tconfirmedCount\n\t\t\tconfirmedCountFlagged\n\t\t\tpermittingCount\n\t\t\tpermittingCountFlagged\n\t\t\tscheduledCount\n\t\t\tscheduledCountFlagged\n\t\t\tinProgressCount\n\t\t\tinProgressCountFlagged\n\t\t\tinstalledCount\n\t\t\tinstalledCountFlagged\n\t\t\tinvoicedCount\n\t\t\tinvoicedCountFlagged\n\t\t\tdelayedCount\n\t\t\ttotal\n\t\t\ttotalFlagged\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerJobs(\n\t\t$partnerID: ID!\n\t\t$search: String\n\t\t$flagged: Boolean\n\t\t$progress: JobProgress\n\t\t$dates: [String!]\n\t\t$page: PageInput!\n\t) {\n\t\tpartnerJobs(\n\t\t\tpartnerID: $partnerID\n\t\t\tsearch: $search\n\t\t\tprogress: $progress\n\t\t\tflagged: $flagged\n\t\t\tdates: $dates\n\t\t\tpage: $page\n\t\t) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tsalesRep {\n\t\t\t\t\t\t...salesRep\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tprice\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressAt\n\t\t\t\t\tprogressFlagged\n\t\t\t\t\tpermitRequired\n\t\t\t\t\tinstallDate\n\t\t\t\t\tinspectionDate\n\t\t\t\t\tinspectionRequired\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t\tepcName\n\t\t\t\t\testimate {\n\t\t\t\t\t\t...material\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n\t\n"): (typeof documents)["\n\tquery PartnerJobs(\n\t\t$partnerID: ID!\n\t\t$search: String\n\t\t$flagged: Boolean\n\t\t$progress: JobProgress\n\t\t$dates: [String!]\n\t\t$page: PageInput!\n\t) {\n\t\tpartnerJobs(\n\t\t\tpartnerID: $partnerID\n\t\t\tsearch: $search\n\t\t\tprogress: $progress\n\t\t\tflagged: $flagged\n\t\t\tdates: $dates\n\t\t\tpage: $page\n\t\t) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\t...page\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\tid\n\t\t\t\t\thomeOwner {\n\t\t\t\t\t\t...homeOwner\n\t\t\t\t\t}\n\t\t\t\t\tsalesRep {\n\t\t\t\t\t\t...salesRep\n\t\t\t\t\t}\n\t\t\t\t\tcompanyName\n\t\t\t\t\tprice\n\t\t\t\t\tprogress\n\t\t\t\t\tprogressAt\n\t\t\t\t\tprogressFlagged\n\t\t\t\t\tpermitRequired\n\t\t\t\t\tinstallDate\n\t\t\t\t\tinspectionDate\n\t\t\t\t\tinspectionRequired\n\t\t\t\t\tworkOrderPrice\n\t\t\t\t\tcontractor {\n\t\t\t\t\t\tid\n\t\t\t\t\t\tname\n\t\t\t\t\t}\n\t\t\t\t\tepcName\n\t\t\t\t\testimate {\n\t\t\t\t\t\t...material\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery JobEstimates($jobID: ID!) {\n\t\tjobEstimates(jobID: $jobID) {\n\t\t\tid\n\t\t\tstatus\n\t\t\testimates {\n\t\t\t\tid\n\t\t\t\ttotalSquares\n\t\t\t\tprimaryPitch\n\t\t\t\tprice\n\t\t\t\tpriceSummary\n\t\t\t\tbounds {\n\t\t\t\t\tlat\n\t\t\t\t\tlng\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n"): (typeof documents)["\n\tquery JobEstimates($jobID: ID!) {\n\t\tjobEstimates(jobID: $jobID) {\n\t\t\tid\n\t\t\tstatus\n\t\t\testimates {\n\t\t\t\tid\n\t\t\t\ttotalSquares\n\t\t\t\tprimaryPitch\n\t\t\t\tprice\n\t\t\t\tpriceSummary\n\t\t\t\tbounds {\n\t\t\t\t\tlat\n\t\t\t\t\tlng\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery EstimateResponses($id: ID!) {\n\t\testimateResponses(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tdescription\n\t\t\tneed\n\t\t\traw\n\t\t}\n\t}\n"): (typeof documents)["\n\tquery EstimateResponses($id: ID!) {\n\t\testimateResponses(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tdescription\n\t\t\tneed\n\t\t\traw\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery NearmapResponse($id: ID!, $respID: ID!) {\n\t\tnearmapResponse(id: $id, respID: $respID) {\n\t\t\traw\n\t\t\tdetail {\n\t\t\t\tprice\n\t\t\t\tprimaryPitchInDegrees\n\t\t\t\tprimaryPitch\n\t\t\t\ttileArea\n\t\t\t\ttileRatio\n\t\t\t\tshingleArea\n\t\t\t\tshingleRatio\n\t\t\t\tmetalArea\n\t\t\t\tmetalRatio\n\t\t\t\tflatArea\n\t\t\t\tflatRatio\n\t\t\t\tgableArea\n\t\t\t\tgableRatio\n\t\t\t\thipArea\n\t\t\t\thipRatio\n\t\t\t\tdutchGableArea\n\t\t\t\tdutchGableRatio\n\t\t\t\tturretArea\n\t\t\t\tturretRatio\n\t\t\t\tdominantRoofMaterial\n\t\t\t\tdominantRoofMaterialID\n\t\t\t\troofCount\n\t\t\t\ttotalUnclippedArea\n\t\t\t\troofMaterialRatioTotal\n\t\t\t\troofTypeRatioTotal\n\t\t\t\troofMaterialSurfaceAreaTotal\n\t\t\t\troofTypeSurfaceAreaTotal\n\t\t\t\ttreeOverhangRatioPrimaryRoof\n\t\t\t\ttreeOverhangConfidencePrimaryRoof\n\t\t\t\ttreeOverhangPresenceConfidence\n\t\t\t\ttreeOverhangAreaPrimaryRoof\n\t\t\t\ttreeOverhangTotalClippedArea\n\t\t\t\ttreeOverhangTotalUnClippedArea\n\t\t\t\ttreeOverhangPresent\n\t\t\t\ttreeOverhangCount\n\t\t\t}\n\t\t}\n\t}\n"): (typeof documents)["\n\tquery NearmapResponse($id: ID!, $respID: ID!) {\n\t\tnearmapResponse(id: $id, respID: $respID) {\n\t\t\traw\n\t\t\tdetail {\n\t\t\t\tprice\n\t\t\t\tprimaryPitchInDegrees\n\t\t\t\tprimaryPitch\n\t\t\t\ttileArea\n\t\t\t\ttileRatio\n\t\t\t\tshingleArea\n\t\t\t\tshingleRatio\n\t\t\t\tmetalArea\n\t\t\t\tmetalRatio\n\t\t\t\tflatArea\n\t\t\t\tflatRatio\n\t\t\t\tgableArea\n\t\t\t\tgableRatio\n\t\t\t\thipArea\n\t\t\t\thipRatio\n\t\t\t\tdutchGableArea\n\t\t\t\tdutchGableRatio\n\t\t\t\tturretArea\n\t\t\t\tturretRatio\n\t\t\t\tdominantRoofMaterial\n\t\t\t\tdominantRoofMaterialID\n\t\t\t\troofCount\n\t\t\t\ttotalUnclippedArea\n\t\t\t\troofMaterialRatioTotal\n\t\t\t\troofTypeRatioTotal\n\t\t\t\troofMaterialSurfaceAreaTotal\n\t\t\t\troofTypeSurfaceAreaTotal\n\t\t\t\ttreeOverhangRatioPrimaryRoof\n\t\t\t\ttreeOverhangConfidencePrimaryRoof\n\t\t\t\ttreeOverhangPresenceConfidence\n\t\t\t\ttreeOverhangAreaPrimaryRoof\n\t\t\t\ttreeOverhangTotalClippedArea\n\t\t\t\ttreeOverhangTotalUnClippedArea\n\t\t\t\ttreeOverhangPresent\n\t\t\t\ttreeOverhangCount\n\t\t\t}\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation NotificationRead($messageID: ID!) {\n    notificationRead(messageID: $messageID)\n  }"): (typeof documents)["\n  mutation NotificationRead($messageID: ID!) {\n    notificationRead(messageID: $messageID)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyUnreadNotificationsCount {\n    myUnreadNotificationsCount\n  }"): (typeof documents)["\n  query MyUnreadNotificationsCount {\n    myUnreadNotificationsCount\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyNotifications($page: PageInput!) {\n    myNotifications(page: $page) {\n      totalCount\n      pageInfo {\n        hasPreviousPage\n        startCursor\n        hasNextPage\n        endCursor\n      }\n      edges {\n        node {\n          id\n          channel\n          topic\n          refID\n          title\n          message\n          from\n          unread\n          createdAt\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query MyNotifications($page: PageInput!) {\n    myNotifications(page: $page) {\n      totalCount\n      pageInfo {\n        hasPreviousPage\n        startCursor\n        hasNextPage\n        endCursor\n      }\n      edges {\n        node {\n          id\n          channel\n          topic\n          refID\n          title\n          message\n          from\n          unread\n          createdAt\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query OptionList($types: [OptionListType!])  {\n    optionList(types: $types) {\n      type\n      options {\n        id\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  query OptionList($types: [OptionListType!])  {\n    optionList(types: $types) {\n      type\n      options {\n        id\n        name\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SavePartnerContact($partnerID: ID!, $contact: PartnerContactInput!){\n    savePartnerContact(\n      partnerID: $partnerID\n      contact: $contact\n    )}"): (typeof documents)["\n  mutation SavePartnerContact($partnerID: ID!, $contact: PartnerContactInput!){\n    savePartnerContact(\n      partnerID: $partnerID\n      contact: $contact\n    )}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation PartnerContactSendPwdResetEmail($partnerID: ID! $userID: ID!){\n        partnerContactSendPwdResetEmail(\n            partnerID: $partnerID\n            userID: $userID\n        )}"): (typeof documents)["\n    mutation PartnerContactSendPwdResetEmail($partnerID: ID! $userID: ID!){\n        partnerContactSendPwdResetEmail(\n            partnerID: $partnerID\n            userID: $userID\n        )}"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query partnerContacts(\n    $partnerID: ID!\n    $search: String\n    $page: PageInput!\n  ) {\n    partnerContacts(\n      partnerID: $partnerID\n      search: $search\n      page: $page\n    ) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        node {\n          id\n          userID\n          firstName\n          lastName\n          email\n          phone\n          role\n          accountStatus\n          picture\n          type\n          title\n          description\n        }\n      }\n    }\n  } "): (typeof documents)["\n  query partnerContacts(\n    $partnerID: ID!\n    $search: String\n    $page: PageInput!\n  ) {\n    partnerContacts(\n      partnerID: $partnerID\n      search: $search\n      page: $page\n    ) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        node {\n          id\n          userID\n          firstName\n          lastName\n          email\n          phone\n          role\n          accountStatus\n          picture\n          type\n          title\n          description\n        }\n      }\n    }\n  } "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SavePartnerMobileSettings($id: ID!, $inp: InputMobileAppSettings!) {\n    savePartnerMobileSettings( id: $id, inp: $inp)\n  }"): (typeof documents)["\n  mutation SavePartnerMobileSettings($id: ID!, $inp: InputMobileAppSettings!) {\n    savePartnerMobileSettings( id: $id, inp: $inp)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query PartnerMobileSettings($id: ID!) {\n    partnerMobileSettings( id: $id) {\n      logoURL\n      primaryColor\n      hideTabs\n    }\n  }"): (typeof documents)["\n  query PartnerMobileSettings($id: ID!) {\n    partnerMobileSettings( id: $id) {\n      logoURL\n      primaryColor\n      hideTabs\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tfragment partner on Partner {\n\t\tid\n\t\tcreatedAt\n\t\ttype\n\t\tname\n\t\taddress\n\t\twebsite\n\t\tisNationWide\n\t\tstatus\n\t}\n"): (typeof documents)["\n\tfragment partner on Partner {\n\t\tid\n\t\tcreatedAt\n\t\ttype\n\t\tname\n\t\taddress\n\t\twebsite\n\t\tisNationWide\n\t\tstatus\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tfragment contact on PartnerContact {\n\t\tid\n\t\tuserID\n\t\ttype\n\t\tfirstName\n\t\tlastName\n\t\temail\n\t\tphone\n\t\totherEmail\n\t\ttitle\n\t\tdescription\n\t}\n"): (typeof documents)["\n\tfragment contact on PartnerContact {\n\t\tid\n\t\tuserID\n\t\ttype\n\t\tfirstName\n\t\tlastName\n\t\temail\n\t\tphone\n\t\totherEmail\n\t\ttitle\n\t\tdescription\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SavePartner($input: PartnerInput!) {\n\t\tsavePartner(input: $input)\n\t}\n"): (typeof documents)["\n\tmutation SavePartner($input: PartnerInput!) {\n\t\tsavePartner(input: $input)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation invitePartner($input: InvitePartnerInput!) {\n\t\tinvitePartner(input: $input) {\n\t\t\tid\n\t\t\ttype\n\t\t\tcompanyName\n\t\t\tcontactID\n\t\t\tuserID\n\t\t\tfirstName\n\t\t\tlastName\n\t\t\temail\n\t\t\tphone\n\t\t\tcreatedAt\n\t\t}\n\t}\n"): (typeof documents)["\n\tmutation invitePartner($input: InvitePartnerInput!) {\n\t\tinvitePartner(input: $input) {\n\t\t\tid\n\t\t\ttype\n\t\t\tcompanyName\n\t\t\tcontactID\n\t\t\tuserID\n\t\t\tfirstName\n\t\t\tlastName\n\t\t\temail\n\t\t\tphone\n\t\t\tcreatedAt\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SavePartnerCompletedSteps($partnerID: ID!, $step: Int!, $done: Boolean) {\n\t\tsavePartnerCompletedSteps(partnerID: $partnerID, step: $step, done: $done)\n\t}\n"): (typeof documents)["\n\tmutation SavePartnerCompletedSteps($partnerID: ID!, $step: Int!, $done: Boolean) {\n\t\tsavePartnerCompletedSteps(partnerID: $partnerID, step: $step, done: $done)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SetPartnerActive($partnerID: ID!, $value: Boolean!) {\n\t\tsetPartnerActive(partnerID: $partnerID, value: $value)\n\t}\n"): (typeof documents)["\n\tmutation SetPartnerActive($partnerID: ID!, $value: Boolean!) {\n\t\tsetPartnerActive(partnerID: $partnerID, value: $value)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation savePartnerContacts($partnerID: ID!, $contacts: [PartnerContactInput!]!) {\n\t\tsavePartnerContacts(partnerID: $partnerID, contacts: $contacts) {\n\t\t\t...contact\n\t\t}\n\t}\n\t\n"): (typeof documents)["\n\tmutation savePartnerContacts($partnerID: ID!, $contacts: [PartnerContactInput!]!) {\n\t\tsavePartnerContacts(partnerID: $partnerID, contacts: $contacts) {\n\t\t\t...contact\n\t\t}\n\t}\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation savePartnerOperations($partnerID: ID!, $input: PartnerOperationInput!) {\n\t\tsavePartnerOperations(partnerID: $partnerID, inp: $input)\n\t}\n"): (typeof documents)["\n\tmutation savePartnerOperations($partnerID: ID!, $input: PartnerOperationInput!) {\n\t\tsavePartnerOperations(partnerID: $partnerID, inp: $input)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SaveService($id: ID!, $partnerID: ID!, $service: Service!, $active: Boolean!) {\n\t\tsaveService(id: $id, partnerID: $partnerID, service: $service, active: $active)\n\t}\n"): (typeof documents)["\n\tmutation SaveService($id: ID!, $partnerID: ID!, $service: Service!, $active: Boolean!) {\n\t\tsaveService(id: $id, partnerID: $partnerID, service: $service, active: $active)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SaveLeadTime($partnerID: ID!, $asphalt: String, $metal: String, $tile: String) {\n\t\tsaveLeadTime(partnerID: $partnerID, asphalt: $asphalt, metal: $metal, tile: $tile)\n\t}\n"): (typeof documents)["\n\tmutation SaveLeadTime($partnerID: ID!, $asphalt: String, $metal: String, $tile: String) {\n\t\tsaveLeadTime(partnerID: $partnerID, asphalt: $asphalt, metal: $metal, tile: $tile)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SaveServiceState($partnerID: ID!, $state: String!, $licNo: String, $expDate: Time, $proofDocID: String) {\n\t\tsaveServiceState(partnerID: $partnerID, state: $state, licNo: $licNo, expDate: $expDate, proofDocID: $proofDocID)\n\t}\n"): (typeof documents)["\n\tmutation SaveServiceState($partnerID: ID!, $state: String!, $licNo: String, $expDate: Time, $proofDocID: String) {\n\t\tsaveServiceState(partnerID: $partnerID, state: $state, licNo: $licNo, expDate: $expDate, proofDocID: $proofDocID)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tmutation SaveServiceCity($partnerID: ID!, $postalID: ID!, $active: Boolean, $licNo: String, $proofDocID: String) {\n\t\tsaveServiceCity(partnerID: $partnerID, postalID: $postalID, active: $active, licNo: $licNo, proofDocID: $proofDocID)\n\t}\n"): (typeof documents)["\n\tmutation SaveServiceCity($partnerID: ID!, $postalID: ID!, $active: Boolean, $licNo: String, $proofDocID: String) {\n\t\tsaveServiceCity(partnerID: $partnerID, postalID: $postalID, active: $active, licNo: $licNo, proofDocID: $proofDocID)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerNameAvailable($id: ID!, $type: PartnerType!, $name: String!) {\n\t\tpartnerNameAvailable(id: $id, type: $type, name: $name)\n\t}\n"): (typeof documents)["\n\tquery PartnerNameAvailable($id: ID!, $type: PartnerType!, $name: String!) {\n\t\tpartnerNameAvailable(id: $id, type: $type, name: $name)\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery Partner($id: ID!) {\n\t\tpartner(id: $id) {\n\t\t\t...partner\n\t\t\tcrewCount\n\t\t\tjobCapacity\n\t\t\tsetupStepsCompleted\n\t\t\tasphaltLeadT\n\t\t\tmetalLeadT\n\t\t\ttileLeadT\n\t\t\tcontacts {\n\t\t\t\t...contact\n\t\t\t\totherEmail\n\t\t\t\ttitle\n\t\t\t\tdescription\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n"): (typeof documents)["\n\tquery Partner($id: ID!) {\n\t\tpartner(id: $id) {\n\t\t\t...partner\n\t\t\tcrewCount\n\t\t\tjobCapacity\n\t\t\tsetupStepsCompleted\n\t\t\tasphaltLeadT\n\t\t\tmetalLeadT\n\t\t\ttileLeadT\n\t\t\tcontacts {\n\t\t\t\t...contact\n\t\t\t\totherEmail\n\t\t\t\ttitle\n\t\t\t\tdescription\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery SolarPartner($id: ID!) {\n\t\tpartner(id: $id, type: SOLAR) {\n\t\t\tid\n\t\t\ttype\n\t\t\tstatus\n\t\t\tname\n\t\t\taddress\n\t\t\twebsite\n\t\t\tyearsInBusiness\n\t\t\tsalesVolume\n\t\t\tfinanceOptions\n\t\t\tdownPayment\n\t\t\tpifDate\n\t\t\tinstallInHouse\n\t\t\tepcOptions\n\t\t\tsetupStepsCompleted\n\t\t\tcontacts {\n\t\t\t\t...contact\n\t\t\t\totherEmail\n\t\t\t\ttitle\n\t\t\t\tdescription\n\t\t\t}\n\t\t}\n\t}\n\t\n"): (typeof documents)["\n\tquery SolarPartner($id: ID!) {\n\t\tpartner(id: $id, type: SOLAR) {\n\t\t\tid\n\t\t\ttype\n\t\t\tstatus\n\t\t\tname\n\t\t\taddress\n\t\t\twebsite\n\t\t\tyearsInBusiness\n\t\t\tsalesVolume\n\t\t\tfinanceOptions\n\t\t\tdownPayment\n\t\t\tpifDate\n\t\t\tinstallInHouse\n\t\t\tepcOptions\n\t\t\tsetupStepsCompleted\n\t\t\tcontacts {\n\t\t\t\t...contact\n\t\t\t\totherEmail\n\t\t\t\ttitle\n\t\t\t\tdescription\n\t\t\t}\n\t\t}\n\t}\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerOptions($partnerID: ID!) {\n\t\tpartnerOptions(partnerID: $partnerID) {\n\t\t\ttype\n\t\t\toptions {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n"): (typeof documents)["\n\tquery PartnerOptions($partnerID: ID!) {\n\t\tpartnerOptions(partnerID: $partnerID) {\n\t\t\ttype\n\t\t\toptions {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery Partners($search: String, $partnerType: PartnerType, $status: String, $page: PageInput!) {\n\t\tpartners(search: $search, partnerType: $partnerType, status: $status, page: $page) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\thasPreviousPage\n\t\t\t\thasNextPage\n\t\t\t\tstartCursor\n\t\t\t\tendCursor\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\t...partner\n\t\t\t\t\tcrewCount\n\t\t\t\t\tjobCapacity\n\t\t\t\t\tisActive\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n"): (typeof documents)["\n\tquery Partners($search: String, $partnerType: PartnerType, $status: String, $page: PageInput!) {\n\t\tpartners(search: $search, partnerType: $partnerType, status: $status, page: $page) {\n\t\t\ttotalCount\n\t\t\tpageInfo {\n\t\t\t\thasPreviousPage\n\t\t\t\thasNextPage\n\t\t\t\tstartCursor\n\t\t\t\tendCursor\n\t\t\t}\n\t\t\tedges {\n\t\t\t\tnode {\n\t\t\t\t\t...partner\n\t\t\t\t\tcrewCount\n\t\t\t\t\tjobCapacity\n\t\t\t\t\tisActive\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery MyJob($id: ID!) {\n\t\tmyJob(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tupdatedAt\n\t\t\tcompanyName\n\t\t\thomeOwner {\n\t\t\t\t...homeOwner\n\t\t\t}\n\t\t\tsalesRep {\n\t\t\t\t...salesRep\n\t\t\t}\n\t\t\testimate {\n\t\t\t\t...material\n\t\t\t\tmeasurementType\n\t\t\t}\n\t\t\tprice\n\t\t\tworkOrderPrice\n\t\t\tcontractor {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n"): (typeof documents)["\n\tquery MyJob($id: ID!) {\n\t\tmyJob(id: $id) {\n\t\t\tid\n\t\t\tcreatedAt\n\t\t\tupdatedAt\n\t\t\tcompanyName\n\t\t\thomeOwner {\n\t\t\t\t...homeOwner\n\t\t\t}\n\t\t\tsalesRep {\n\t\t\t\t...salesRep\n\t\t\t}\n\t\t\testimate {\n\t\t\t\t...material\n\t\t\t\tmeasurementType\n\t\t\t}\n\t\t\tprice\n\t\t\tworkOrderPrice\n\t\t\tcontractor {\n\t\t\t\tid\n\t\t\t\tname\n\t\t\t}\n\t\t}\n\t}\n\t\n\t\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerDocs($partnerID: ID!, $section: DocumentSection!) {\n\t\tpartnerDocs(partnerID: $partnerID, section: $section) {\n\t\t\t...document\n\t\t}\n\t}\n\t\n"): (typeof documents)["\n\tquery PartnerDocs($partnerID: ID!, $section: DocumentSection!) {\n\t\tpartnerDocs(partnerID: $partnerID, section: $section) {\n\t\t\t...document\n\t\t}\n\t}\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerServiceStates($partnerID: ID!) {\n\t\tpartnerServiceStates(partnerID: $partnerID) {\n\t\t\tid\n\t\t\tname\n\t\t\tlicenseNo\n\t\t\tlicenseExpDate\n\t\t\texpand\n\t\t\tcities {\n\t\t\t\tid\n\t\t\t\tactive\n\t\t\t\tlicenseNo\n\t\t\t\tlicenseProof\n\t\t\t\tcityZip\n\t\t\t\tcityName\n\t\t\t}\n\t\t}\n\t\tpartnerDocs(partnerID: $partnerID, section: Proof) {\n\t\t\t...document\n\t\t}\n\t}\n\t\n"): (typeof documents)["\n\tquery PartnerServiceStates($partnerID: ID!) {\n\t\tpartnerServiceStates(partnerID: $partnerID) {\n\t\t\tid\n\t\t\tname\n\t\t\tlicenseNo\n\t\t\tlicenseExpDate\n\t\t\texpand\n\t\t\tcities {\n\t\t\t\tid\n\t\t\t\tactive\n\t\t\t\tlicenseNo\n\t\t\t\tlicenseProof\n\t\t\t\tcityZip\n\t\t\t\tcityName\n\t\t\t}\n\t\t}\n\t\tpartnerDocs(partnerID: $partnerID, section: Proof) {\n\t\t\t...document\n\t\t}\n\t}\n\t\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n\tquery PartnerServices($partnerID: ID!) {\n\t\tpartnerServices(partnerID: $partnerID) {\n\t\t\tid\n\t\t\tservice\n\t\t\tdescription\n\t\t\tactive\n\t\t}\n\t}\n"): (typeof documents)["\n\tquery PartnerServices($partnerID: ID!) {\n\t\tpartnerServices(partnerID: $partnerID) {\n\t\t\tid\n\t\t\tservice\n\t\t\tdescription\n\t\t\tactive\n\t\t}\n\t}\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation MarkServiceArea($id: ID!, $value: Boolean!) {\n    markServiceArea(id: $id, value: $value)\n  }"): (typeof documents)["\n  mutation MarkServiceArea($id: ID!, $value: Boolean!) {\n    markServiceArea(id: $id, value: $value)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query AllServiceAreas {\n    allServiceAreas {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }"): (typeof documents)["\n  query AllServiceAreas {\n    allServiceAreas {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ServiceStates($q: String!) {\n    serviceStates(q: $q) {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }"): (typeof documents)["\n  query ServiceStates($q: String!) {\n    serviceStates(q: $q) {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query States($q: String!) {\n    states(q: $q) {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }"): (typeof documents)["\n  query States($q: String!) {\n    states(q: $q) {\n      id\n      name\n      cities {\n        id\n        zip\n        name\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Cities($state: String!, $q: String!, $skip: Int!, $take: Int!) {\n    cities(state: $state, q: $q, skip: $skip, take: $take) {\n      id\n      zip\n      name\n    }\n  }"): (typeof documents)["\n  query Cities($state: String!, $q: String!, $skip: Int!, $take: Int!) {\n    cities(state: $state, q: $q, skip: $skip, take: $take) {\n      id\n      zip\n      name\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query GetPricing{\n    getPricing {\n      id\n      items {\n        id\n        country\n        state\n        stateAbr\n        city\n        zip\n        productId\n        price\n        pricePer\n      }\n      products {\n        id\n        name\n      }\n\n    }\n  }\n"): (typeof documents)["\n  query GetPricing{\n    getPricing {\n      id\n      items {\n        id\n        country\n        state\n        stateAbr\n        city\n        zip\n        productId\n        price\n        pricePer\n      }\n      products {\n        id\n        name\n      }\n\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SaveProductPackage($input: ProductPackageInput!) {\n    saveProductPackage(input: $input)\n  }"): (typeof documents)["\n  mutation SaveProductPackage($input: ProductPackageInput!) {\n    saveProductPackage(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SaveProduct($input: ProductInput!) {\n    saveProduct(input: $input)\n  }"): (typeof documents)["\n  mutation SaveProduct($input: ProductInput!) {\n    saveProduct(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ProductPackages($category: ProductType, $search: String $page: PageInput!) {\n    productPackages(category: $category, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query ProductPackages($category: ProductType, $search: String $page: PageInput!) {\n    productPackages(category: $category, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Products($category: ProductType, $search: String $page: PageInput!) {\n    products(category: $category, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          name\n          description\n          price\n          features\n          specialNote\n          image {\n            id\n            publicUrl\n          }\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query Products($category: ProductType, $search: String $page: PageInput!) {\n    products(category: $category, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          name\n          description\n          price\n          features\n          specialNote\n          image {\n            id\n            publicUrl\n          }\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query SmartHomePackages($page: PageInput!) {\n    smartHomePackages(page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            description\n            price\n            features\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query SmartHomePackages($page: PageInput!) {\n    smartHomePackages(page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            description\n            price\n            features\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query HVACPackages($page: PageInput!) {\n    hvacPackages(page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            description\n            price\n            features\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query HVACPackages($page: PageInput!) {\n    hvacPackages(page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        endCursor\n        hasNextPage\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          category\n          soldAs\n          name\n          description\n          price\n          features\n          items {\n            id\n            name\n            description\n            price\n            features\n            image {\n              publicUrl\n            }\n          }\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Surveys($search: String, $page: PageInput!, $orderBy: SurveyOrder) {\n    surveys(search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      edges {\n        node {\n          id\n          date\n          slot\n          from\n          to\n          name\n          phone\n          address\n          notes\n          status\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query Surveys($search: String, $page: PageInput!, $orderBy: SurveyOrder) {\n    surveys(search: $search, page: $page, orderBy: $orderBy) {\n      totalCount\n      edges {\n        node {\n          id\n          date\n          slot\n          from\n          to\n          name\n          phone\n          address\n          notes\n          status\n        }\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query SurveySlotAvailability($date: String!) {\n    surveySlotAvailability(date: $date) {\n      id\n      name\n      available\n    }\n  }\n"): (typeof documents)["\n  query SurveySlotAvailability($date: String!) {\n    surveySlotAvailability(date: $date) {\n      id\n      name\n      available\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SurveyRequest($date: String!, $slotID: ID!) {\n    surveyRequest(date: $date, slotID: $slotID)\n  }\n"): (typeof documents)["\n  mutation SurveyRequest($date: String!, $slotID: ID!) {\n    surveyRequest(date: $date, slotID: $slotID)\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SurveyReserve($input: SurveyInput!) {\n    surveyReserve(input: $input) {\n      id\n      date\n      slot\n      from\n      to\n      name\n      phone\n      address\n      notes\n      status\n    }\n  }\n"): (typeof documents)["\n  mutation SurveyReserve($input: SurveyInput!) {\n    surveyReserve(input: $input) {\n      id\n      date\n      slot\n      from\n      to\n      name\n      phone\n      address\n      notes\n      status\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query SurveyDetails($id: ID!) {\n    surveyDetails(id: $id) {\n      id\n      date\n      slot\n      from\n      to\n      name\n      phone\n      address\n      notes\n      status\n    }\n  }\n"): (typeof documents)["\n  query SurveyDetails($id: ID!) {\n    surveyDetails(id: $id) {\n      id\n      date\n      slot\n      from\n      to\n      name\n      phone\n      address\n      notes\n      status\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  fragment videoInfo on TrainingVideo {\n    id\n    title\n    description\n    posterURL\n    videoURL\n  }"): (typeof documents)["\n  fragment videoInfo on TrainingVideo {\n    id\n    title\n    description\n    posterURL\n    videoURL\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateTrainingCourse($name: String!){\n    createTrainingCourse(name: $name) {\n      id\n      name\n    }\n  }"): (typeof documents)["\n  mutation CreateTrainingCourse($name: String!){\n    createTrainingCourse(name: $name) {\n      id\n      name\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation PartnerTrainingVideoAccess($partnerID: ID!, $videoID: ID!, $enabled: Boolean!) {\n    partnerTrainingVideoAccess(partnerID: $partnerID, videoID: $videoID, enabled: $enabled)\n  }"): (typeof documents)["\n  mutation PartnerTrainingVideoAccess($partnerID: ID!, $videoID: ID!, $enabled: Boolean!) {\n    partnerTrainingVideoAccess(partnerID: $partnerID, videoID: $videoID, enabled: $enabled)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SaveTrainingVideo($inp: InputTrainingVideo!) {\n    saveTrainingVideo(inp: $inp)\n  }"): (typeof documents)["\n  mutation SaveTrainingVideo($inp: InputTrainingVideo!) {\n    saveTrainingVideo(inp: $inp)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query TrainingCourses($search: String, $page: PageInput!) {\n    trainingCourses(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          id\n          name\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query TrainingCourses($search: String, $page: PageInput!) {\n    trainingCourses(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          id\n          name\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query TrainingVideoKinds {\n    trainingVideoKinds {\n      id\n      name\n    }\n  }"): (typeof documents)["\n  query TrainingVideoKinds {\n    trainingVideoKinds {\n      id\n      name\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query TrainingVideoCourses($kind: TrainingType!, $partnerID: ID,  $pageSize: Int) {\n    trainingVideoCourses(kind: $kind, partnerID: $partnerID, pageSize: $pageSize) {\n      id\n      name\n      videos {\n        totalCount\n        pageInfo {\n          ...page\n        }\n        edges {\n          cursor\n          node {\n            ...videoInfo\n            assigned\n          }\n        }\n      }\n    }\n  }  "): (typeof documents)["\n  query TrainingVideoCourses($kind: TrainingType!, $partnerID: ID,  $pageSize: Int) {\n    trainingVideoCourses(kind: $kind, partnerID: $partnerID, pageSize: $pageSize) {\n      id\n      name\n      videos {\n        totalCount\n        pageInfo {\n          ...page\n        }\n        edges {\n          cursor\n          node {\n            ...videoInfo\n            assigned\n          }\n        }\n      }\n    }\n  }  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query TrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $partnerID: ID, $page: PageInput!) {\n    trainingVideos(kind: $kind, courseID: $courseID, search: $search, partnerID: $partnerID, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          ...videoInfo\n          kind\n          assigned\n        }\n      }\n    }\n  } "): (typeof documents)["\n  query TrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $partnerID: ID, $page: PageInput!) {\n    trainingVideos(kind: $kind, courseID: $courseID, search: $search, partnerID: $partnerID, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          ...videoInfo\n          kind\n          assigned\n        }\n      }\n    }\n  } "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyTrainingVideoKinds {\n    myTrainingVideoKinds {\n      id\n      name\n    }\n  }"): (typeof documents)["\n  query MyTrainingVideoKinds {\n    myTrainingVideoKinds {\n      id\n      name\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyTrainingVideoCourses($kind: TrainingType!,  $pageSize: Int) {\n    myTrainingVideoCourses(kind: $kind,  pageSize: $pageSize) {\n      id\n      name\n      videos {\n        totalCount\n        pageInfo {\n          ...page\n        }\n        edges {\n          cursor\n          node {\n            ...videoInfo\n          }\n        }\n      }\n    }\n  }  "): (typeof documents)["\n  query MyTrainingVideoCourses($kind: TrainingType!,  $pageSize: Int) {\n    myTrainingVideoCourses(kind: $kind,  pageSize: $pageSize) {\n      id\n      name\n      videos {\n        totalCount\n        pageInfo {\n          ...page\n        }\n        edges {\n          cursor\n          node {\n            ...videoInfo\n          }\n        }\n      }\n    }\n  }  "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyTrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $page: PageInput!) {\n    myTrainingVideos(kind: $kind, courseID: $courseID, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          ...videoInfo\n          kind\n        }\n      }\n    }\n  } "): (typeof documents)["\n  query MyTrainingVideos($kind: TrainingType!, $courseID: ID!, $search: String, $page: PageInput!) {\n    myTrainingVideos(kind: $kind, courseID: $courseID, search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        ...page\n      }\n      edges {\n        cursor\n        node {\n          ...videoInfo\n          kind\n        }\n      }\n    }\n  } "];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query User($id: ID!) {\n    user(id: $id) {\n      id\n      email\n      firstName\n      lastName\n      phone\n      role\n      status\n      note\n    }\n  }"): (typeof documents)["\n  query User($id: ID!) {\n    user(id: $id) {\n      id\n      email\n      firstName\n      lastName\n      phone\n      role\n      status\n      note\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query Users($page: PageInput! $where: UserWhereInput) {\n    users(page: $page where: $where) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          ...userBasic\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query Users($page: PageInput! $where: UserWhereInput) {\n    users(page: $page where: $where) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          ...userBasic\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query usersSearch($search: String) {\n    usersSearch(search: $search) {\n      id\n      firstName\n      lastName\n      email\n      phone\n      partnerID\n      partnerName\n      partnerContactTypeID\n      partnerContactTitle\n    }\n  }"): (typeof documents)["\n  query usersSearch($search: String) {\n    usersSearch(search: $search) {\n      id\n      firstName\n      lastName\n      email\n      phone\n      partnerID\n      partnerName\n      partnerContactTypeID\n      partnerContactTitle\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query myCompanyUsers($search: String, $page: PageInput!) {\n    myCompanyUsers(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          firstName\n          lastName\n          email\n          phone\n          picture\n        }\n      }\n    }\n  }"): (typeof documents)["\n  query myCompanyUsers($search: String, $page: PageInput!) {\n    myCompanyUsers(search: $search, page: $page) {\n      totalCount\n      pageInfo {\n        startCursor\n        hasNextPage\n        endCursor\n        hasPreviousPage\n      }\n      edges {\n        node {\n          id\n          firstName\n          lastName\n          email\n          phone\n          picture\n        }\n      }\n    }\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateUser($input: CreateUserInput!) {\n    createUser(input: $input)\n  }"): (typeof documents)["\n  mutation CreateUser($input: CreateUserInput!) {\n    createUser(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation UpdateUser($input: UpdateUserInput!) {\n    updateUser(input: $input)\n  }"): (typeof documents)["\n  mutation UpdateUser($input: UpdateUserInput!) {\n    updateUser(input: $input)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation SaveNotifySettings($id: ID! $topicID: String! $email: Boolean!) {\n    saveNotifySettings(userID: $id topicID: $topicID email: $email)\n  }"): (typeof documents)["\n  mutation SaveNotifySettings($id: ID! $topicID: String! $email: Boolean!) {\n    saveNotifySettings(userID: $id topicID: $topicID email: $email)\n  }"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query UserNotifySettings($id: ID!) {\n    userNotifySettings(id: $id) {\n      id\n      topic\n      receiveEmail\n      receiveSMS\n    }\n  }"): (typeof documents)["\n  query UserNotifySettings($id: ID!) {\n    userNotifySettings(id: $id) {\n      id\n      topic\n      receiveEmail\n      receiveSMS\n    }\n  }"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;
