import { MyUserData, SessionPhase } from 'session/model';
import { getLoginDataFromLocalStorage } from './localstorage';

export interface SessionState {
  userData: MyUserData | null;
  phase: SessionPhase;
  error?: unknown;
}

export const initialState = (): SessionState => {
  const rememberedLogin = getLoginDataFromLocalStorage();
  return {
    userData: rememberedLogin,
    phase: rememberedLogin === null ? SessionPhase.NoToken : SessionPhase.TokenObtained,
  };
};
