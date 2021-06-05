import { MyUserData } from 'session/model';

const key_id = 'user_id';
const key_token = 'session_token';
const key_name = 'name';
const key_email = 'email';

export function getLoginDataFromLocalStorage(): MyUserData | null {
  const [id, email, name, token] = [key_id, key_email, key_name, key_token].map((key) =>
    window.localStorage.getItem(key)
  );
  if (id !== null && email !== null && name !== null && token !== null) {
    return { id, name, email, token };
  }
  deleteLoginDataInLocalStorage();
  return null;
}

export function writeLoginDataToLocalStorage({ id, name, token, email }: MyUserData): void {
  window.localStorage.setItem(key_id, id);
  window.localStorage.setItem(key_name, name);
  window.localStorage.setItem(key_token, token);
  window.localStorage.setItem(key_email, email);
}

export function deleteLoginDataInLocalStorage(): void {
  window.localStorage.removeItem(key_id);
  window.localStorage.removeItem(key_name);
  window.localStorage.removeItem(key_email);
  window.localStorage.removeItem(key_token);
}
