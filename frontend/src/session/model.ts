export interface MyUserData extends Credentials {
  name: string;
}

export interface Credentials {
  id: string;
  secret: string;
}
