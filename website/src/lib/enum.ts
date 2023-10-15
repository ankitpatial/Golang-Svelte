import { AccountStatus, Role } from '$lib/graph/graphql';

export enum PATH {
  HOME = '/',
  LOGIN = '/',
  LOGOUT = '/logout',
  PROFILE = '/profile',
  NOTIFICATIONS = '/notifications',

  OVERVIEW = '/overview',

  ESTIMATES = '/estimates',
  ESTIMATES_VIEW = '/estimates/view',
  ESTIMATES_REQ = '/estimates/requests',
  ESTIMATE_RESPONSE_NEARMAP = '/estimates/:id/response-view/:respID/nearmap',

  THEME = '/theme',

  PARTNER_ROOFING = '/partner/roofing',
  PARTNER_SOLAR = '/partner/solar',
  PARTNER_EPC = '/partner/epc',
  PARTNER_INTEGRATION = '/partner/integration',

  PARTNERS = '/partners',
  PARTNER_SAVE = '/partners/:id/save',

  PRICING = '/pricing',

  JOBS_UNASSIGNED = '/jobs/unassigned',
  JOBS_ASSIGNED = '/jobs/assigned',
  JOBS_IN_HOUSE = '/jobs/in-house',

  JOB_DOCS = '/job-docs',

  MPU = '/mpu',
  INSULATION = '/insulation',

  HVAC = '/hvac/packages',
  HVAC_PENDING = '/hvac/pending',
  HVAC_APPROVED = '/hvac/approved',

  TURF = '/turf',
  WINDOWS = '/windows',
  BATTERIES = '/batteries',

  SITE_SURVEYS = '/survey/dashboard',
  SMART_HOME = '/smart-home/packages',
  SMART_HOME_PENDING = '/smart-home/pending',
  SMART_HOME_APPROVED = '/smart-home/approved',


  PAYMENTS_PENDING = '/payments/pending',
  PAYMENTS_APPROVED = '/payments/approved',
  PAYMENTS_COMPLETED = '/payments/completed',

  PARTNER_ADD_USER = '/company-users/add-new-user',
  PARTNER_VIEW_USERS = '/company-users/list',
  PARTNER_EDIT_USERS = '/company-users/:id/edit',

  ROOFING_PARTNER = '/roofing/overview',
  ROOFING_PARTNER_JOBS = '/roofing/jobs',
  ROOFING_PARTNER_TRAINING = '/roofing/training-center',

  SOLAR_PARTNER = '/solar/overview',
  SOLAR_PARTNER_EST = '/solar/estimate',
  SOLAR_PARTNER_EST_REQ = '/solar/estimate/request',
  SOLAR_PARTNER_EST_COMPLETED = '/solar/estimate/completed',
  SOLAR_PARTNER_JOBS = '/solar/jobs',
  SOLAR_PARTNER_SURVEYS = '/solar/survey/dashboard',
  SOLAR_PARTNER_SURVEY_REQUEST = '/solar/survey/request',
  SOLAR_PARTNER_TRAINING = '/solar/training-center',
  SOLAR_PARTNER_SMART_HOME = '/solar/smart-home/packages',
  SOLAR_PARTNER_SMART_HOME_PENDING = '/solar/smart-home/pending',
  SOLAR_PARTNER_SMART_HOME_APPROVED = '/solar/smart-home/approved',
  SOLAR_PARTNER_HVAC = '/solar/hvac/packages',
  SOLAR_PARTNER_HVAC_PENDING = '/solar/hvac/pending',
  SOLAR_PARTNER_HVAC_APPROVED = '/solar/hvac/approved',

  EPC_PARTNER = '/epc/overview',

  INTEGRATION_PARTNER = '/integration/overview',

  UNKNOWN_PARTNER = '/unknown-partner',

  TRAINING_CENTER = '/training',
  TRAINING_VIDEOS = '/training/:kind',
  TRAINING_VIDEOS_SAVE = '/training/:kind/:id',

  USERS = '/users',
  USERS_CREATE = '/users/create',
  USERS_EDIT = '/users/:id/edit',

  API_USERS = '/api-users',
  API_USERS_EDIT = '/api-users/:id/edit',
  API_USERS_AUDIT_LOG = '/api-users/:id/audit-logs',
  API_ACCESS = '/api-access',
  API_ACCESS_CREATE = '/api-access/create',
  API_ACCESS_EDIT = '/api-access/:id/edit',

  SERVICE_AREAS = '/service-areas',
  PRODUCT_PACKAGES = '/product-packages',
  PRODUCT_PACKAGES_ADD = '/product-packages/add',
  PRODUCT_PACKAGES_EDIT = '/product-packages/edit',
  PRODUCTS = '/product-packages/products',
  PRODUCTS_ADD = '/product-packages/products/add',
  PRODUCTS_EDIT = '/product-packages/products/edit',

  AUDIT_LOGS = '/audit-logs',

  DOCS = '/docs',
  DOCS_EXT_API = '/docs/external-api',

  NOT_FOUND = '/404',
}

export const All_ROLE = [Role.Admin, Role.SiteUser];

export const ALL_ACCOUNT_STATUS = [
  AccountStatus.Pending,
  AccountStatus.Active,
  AccountStatus.Disabled
];


export enum API_ACCESS_SOURCE {
  EAGLE_VIEW = 'EagleView',
  RFX = 'RFX',
  NEAR_MAP = 'NearMap',
  REGRID = 'ReGrid'
}
