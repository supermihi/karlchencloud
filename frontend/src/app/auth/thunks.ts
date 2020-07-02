import { createAsyncThunk } from '@reduxjs/toolkit';
import { MyUserData } from '.';
import { writeLoginDataToLocalStorage, deleteLoginDataInLocalStorage } from './localstorage';
import { AppThunk } from 'app/store';
import { actions } from './slice';
import * as api from 'api/client';

export const register = createAsyncThunk<MyUserData, string>(
  'model/register',
  async (name, { dispatch }) => {
    const { id, secret } = await api.register(name);
    const ans = { name, id, secret };
    writeLoginDataToLocalStorage(ans);
    dispatch(actions.localStorageUpdated(ans));
    return ans;
  }
);

export const forgetLogin = (): AppThunk => (dispatch) => {
  deleteLoginDataInLocalStorage();
  dispatch(actions.localStorageUpdated(null));
};
