import { createAction, createAsyncThunk } from '@reduxjs/toolkit';
import { LoginData, MyUserData, RegisterData } from '../model';
import { writeLoginDataToLocalStorage, deleteLoginDataInLocalStorage } from '../localstorage';
import { AppDispatch, AppThunk } from 'state';
import * as api from 'api/client';
import { startSession } from './session';

export const register = createAsyncThunk<MyUserData, RegisterData, { dispatch: AppDispatch }>(
  'register',
  async ({ name, email, password }, { dispatch }) => {
    const { id, token } = await api.register(email, name, password);
    const ans = { name, id, token, email };
    writeLoginDataToLocalStorage(ans);
    dispatch(startSession(ans));
    return ans;
  }
);

export const login = createAsyncThunk<MyUserData, LoginData, { dispatch: AppDispatch }>(
  'login',
  async ({ email, password }, { dispatch }) => {
    const { id, name, token } = await api.login(email, password);
    const ans = { name, id, token, email };
    writeLoginDataToLocalStorage(ans);
    dispatch(startSession(ans));
    return ans;
  }
);
export const forgetLogin = createAction('authenticate/forgetLogin');

export const forgetLoginThunk = (): AppThunk => (dispatch) => {
  deleteLoginDataInLocalStorage();
  dispatch(forgetLogin());
};
