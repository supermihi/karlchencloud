import { getAuthMeta } from "./client";

test("auth string correct", () => {
  const { authorization: x } = getAuthMeta("michael", "geheim");
  expect(x).toEqual("basic bWljaGFlbDpnZWhlaW0=");
});

test("auth string with special char", () => {
  const { authorization: x } = getAuthMeta("Yårkl→nd", "`564ΣΛ");
  expect(x).toEqual("basic WcOlcmts4oaSbmQ6YDU2NM6jzps=");
});
