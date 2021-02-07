import { MyUserData } from 'model/session';

const key_id = 'auth_id';
const key_secret = 'auth_secret';
const key_name = 'auth_name';

export function getLoginDataFromLocalStorage(): MyUserData | null {
  const [id, secret, name] = [key_id, key_secret, key_name].map((key) =>
    window.localStorage.getItem(key)
  );
  if (id !== null && secret !== null && name !== null) {
    return { id, secret, name };
  }
  return null;
}

export function writeLoginDataToLocalStorage({ id, name, secret }: MyUserData) {
  window.localStorage.setItem(key_id, id);
  window.localStorage.setItem(key_name, name);
  window.localStorage.setItem(key_secret, secret);
}

export function deleteLoginDataInLocalStorage() {
  window.localStorage.removeItem(key_id);
  window.localStorage.removeItem(key_name);
  window.localStorage.removeItem(key_secret);
}
