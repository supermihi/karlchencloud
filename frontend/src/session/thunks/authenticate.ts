import { createAction, createAsyncThunk } from '@reduxjs/toolkit';
import { LoginData, MyUserData, RegisterData } from '../model';
import { writeLoginDataToLocalStorage, deleteLoginDataInLocalStorage } from '../localstorage';
import { AppThunk } from 'state';
import * as api from 'api/client';

export const localStorageUpdated = createAction<MyUserData | null>('localStorageUpdated');

export const register = createAsyncThunk<MyUserData, RegisterData>(
  'register',
  async ({ name, email, password }, { dispatch }) => {
    const { id, token } = await api.register(email, name, password);
    const ans = { name, id, token, email };
    writeLoginDataToLocalStorage(ans);
    dispatch(localStorageUpdated(ans));
    return ans;
  }
);

export const login = createAsyncThunk<MyUserData, LoginData>(
  'login',
  async ({email, password}, {dispatch}) => {
    const { id, name, token} = await api.login(email, password);
    const ans = {name, id, token, email};
    writeLoginDataToLocalStorage(ans);
    dispatch(localStorageUpdated(ans));
    return ans;
  }

)
export const forgetLogin = (): AppThunk => (dispatch) => {
  deleteLoginDataInLocalStorage();
  dispatch(localStorageUpdated(null));
};
