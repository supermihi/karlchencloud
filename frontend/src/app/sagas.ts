import { takeEvery, put } from "redux-saga/effects";
import { tryLogin, register } from "features/auth/slice";
import { setLocation } from "core/routing";

function* toLobby() {
  yield put(setLocation("lobby"));
}

export default function* rootSaga() {
  yield takeEvery(tryLogin.fulfilled.type, toLobby);
  yield takeEvery(register.fulfilled.type, toLobby);
}
