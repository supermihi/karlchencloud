import { Credentials, MyUserData } from 'session/model';
import { getLoginDataFromLocalStorage } from './localstorage';

export interface SessionState {
  storedLogin: MyUserData | null;
  session: MyUserData | null;
  currentLoginCredentials: Credentials | null;
  error?: unknown;
  loading: boolean;
}

export const initialState = (): SessionState => {
  const existingLogin = getLoginDataFromLocalStorage();
  return {
    storedLogin: existingLogin,
    session: null,
    currentLoginCredentials: null,
    loading: false,
  };
};
