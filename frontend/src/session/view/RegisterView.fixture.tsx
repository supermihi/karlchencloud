import React from "react";
import { useValue } from "react-cosmos/fixture";
import { mockGrpcError } from "shared/mock";
import RegisterView from "./RegisterView";

const Fixture: React.FC = () => {
  const [error] = useValue("error", { defaultValue: false });
  const [loading] = useValue("loading", { defaultValue: false });
  return (
    <RegisterView
      loading={loading}
      error={error ? mockGrpcError("this is a mock") : undefined}
      register={() => console.log("login")}
    />
  );
};
export default Fixture;
