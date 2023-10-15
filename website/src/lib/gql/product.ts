import { gql } from '@urql/svelte';

//
// mutations ==>
//
export const SaveProductPackage = gql`
  mutation SaveProductPackage($input: ProductPackageInput!) {
    saveProductPackage(input: $input)
  }`;

export const SaveProduct = gql`
  mutation SaveProduct($input: ProductInput!) {
    saveProduct(input: $input)
  }`;

//
// queries ==>
//
export const QryProductPackages = gql`
  query ProductPackages($category: ProductType, $search: String $page: PageInput!) {
    productPackages(category: $category, search: $search, page: $page) {
      totalCount
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        node {
          id
          category
          soldAs
          name
          description
          price
          features
          items {
            id
            name
            image {
              publicUrl
            }
          }
        }
      }
    }
  }`;


export const QryProducts = gql`
  query Products($category: ProductType, $search: String $page: PageInput!) {
    products(category: $category, search: $search, page: $page) {
      totalCount
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        node {
          id
          category
          name
          description
          price
          features
          specialNote
          image {
            id
            publicUrl
          }
        }
      }
    }
  }`;

export const QrySmartHomePackages = gql`
  query SmartHomePackages($page: PageInput!) {
    smartHomePackages(page: $page) {
      totalCount
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        node {
          id
          category
          soldAs
          name
          description
          price
          features
          items {
            id
            name
            description
            price
            features
            image {
              publicUrl
            }
          }
        }
      }
    }
  }`;
export const QeyHVACPackages = gql`
  query HVACPackages($page: PageInput!) {
    hvacPackages(page: $page) {
      totalCount
      pageInfo {
        startCursor
        endCursor
        hasNextPage
        hasPreviousPage
      }
      edges {
        node {
          id
          category
          soldAs
          name
          description
          price
          features
          items {
            id
            name
            description
            price
            features
            image {
              publicUrl
            }
          }
        }
      }
    }
  }`;
