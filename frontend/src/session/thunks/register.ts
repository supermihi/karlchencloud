import { createAction, createAsyncThunk } from '@reduxjs/toolkit';
import { MyUserData } from '../model';
import { writeLoginDataToLocalStorage, deleteLoginDataInLocalStorage } from '../localstorage';
import { AppThunk } from 'state';
import * as api from 'api/client';

export const localStorageUpdated = createAction<MyUserData | null>('localStorageUpdated');

export const register = createAsyncThunk<MyUserData, string>(
  'register',
  async (name, { dispatch }) => {
    const { id, secret } = await api.register(name);
    const ans = { name, id, secret };
    writeLoginDataToLocalStorage(ans);
    dispatch(localStorageUpdated(ans));
    return ans;
  }
);

export const forgetLogin = (): AppThunk => (dispatch) => {
  deleteLoginDataInLocalStorage();
  dispatch(localStorageUpdated(null));
};
