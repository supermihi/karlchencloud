export interface LoginData {
  name: string;
  id: string;
  secret: string;
}

export function getLoginDataFromLocalStorage(): LoginData | null {
  const [id, secret, name] = ["id", "secret", "name"].map((key) =>
    window.localStorage.getItem(key)
  );
  if (id !== null && secret !== null && name !== null) {
    return { id, secret, name };
  }
  return null;
}

export function writeLoginDataToLocalStorage({ id, secret, name }: LoginData) {
  window.localStorage.setItem("id", id);
  window.localStorage.setItem("name", name);
  window.localStorage.setItem("secret", secret);
}

export function deleteLoginDataInLocalStorage() {
  window.localStorage.removeItem("id");
  window.localStorage.removeItem("name");
  window.localStorage.removeItem("secret");
}
