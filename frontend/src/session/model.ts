export interface MyUserData {
  name: string;
  email: string;
  id: string;
  token: string;
}

export interface RegisterData {
  name: string;
  email: string;
  password: string;
}
export interface LoginData {
  email: string;
  password: string;
}

export enum LoginPage {
  Register,
  Login,
}
export enum SessionPhase {
  NoToken,
  ObtainingToken,
  TokenObtained,
  Starting,
  Established,
}
