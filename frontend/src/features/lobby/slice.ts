import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { parseInviteCode } from 'model/invitation';
import { AppThunk, RootState } from 'app/store';
import { joinTable } from 'app/game/thunks';

export interface LobbyState {
  suppliedInviteCode: string | null;
}

const initialState: LobbyState = {
  suppliedInviteCode: parseInviteCode(window.location.href),
};

const lobbySlice = createSlice({
  name: 'lobby',
  initialState,
  reducers: {
    setInviteCode: (state, action: PayloadAction<string>) => {
      state.suppliedInviteCode = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(joinTable.fulfilled, (state) => {
      state.suppliedInviteCode = null;
    });
  },
});
const { setInviteCode } = lobbySlice.actions;

export const clearInviteCode = (): AppThunk => (dispatch) => {
  dispatch(setInviteCode(''));
  window.history.pushState('', '', '/');
};
export const selectLobby = (state: RootState) => state.lobby;
export default lobbySlice.reducer;
