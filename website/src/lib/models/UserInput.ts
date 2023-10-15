import type {AccountStatus, Role} from "$lib/graph/graphql";

export default interface UserInput {
  id?: string;
  email: string;
  firstName: string;
  lastName: string;
  phone?: string;
  role: Role;
  status?: AccountStatus;
  picture?: string;
}
