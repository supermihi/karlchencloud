import { MyUserData } from 'session/model';
import { getLoginDataFromLocalStorage } from './localstorage';

export interface SessionState {
  storedLogin: MyUserData | null;
  activeSession: MyUserData | null;
  startingSession: MyUserData | null;
  error?: unknown;
  loading: boolean;
}

export const initialState = (): SessionState => {
  const existingLogin = getLoginDataFromLocalStorage();
  return {
    storedLogin: existingLogin,
    activeSession: null,
    startingSession: null,
    loading: false,
  };
};
