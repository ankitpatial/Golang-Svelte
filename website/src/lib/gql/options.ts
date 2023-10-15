import {gql} from "@urql/svelte";

export const QryOptionList = gql`
  query OptionList($types: [OptionListType!])  {
    optionList(types: $types) {
      type
      options {
        id
        name
      }
    }
  }
`
