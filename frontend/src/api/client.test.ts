import { getAuthHeader } from "./client";

test("auth string correct", () => {
  const { authorization: x } = getAuthHeader("michael", "geheim");
  expect(x).toEqual("basic bWljaGFlbDpnZWhlaW0=");
});

test("auth string with special char", () => {
  const { authorization: x } = getAuthHeader("Yårkl→nd", "`564ΣΛ");
  expect(x).toEqual("basic WcOlcmts4oaSbmQ6YDU2NM6jzps=");
});
