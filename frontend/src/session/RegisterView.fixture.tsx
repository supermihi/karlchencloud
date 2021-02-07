import React from "react";
import { useValue } from "react-cosmos/fixture";
import { mockGrpcError } from "shared/mock";
import RegisterView from "./RegisterView";

const Fixture: React.FC = () => {
  const [error] = useValue("boolean", { defaultValue: false });
  return (
    <RegisterView
      loading={false}
      error={error ? mockGrpcError("error logging in") : undefined}
      register={() => console.log("login")}
    />
  );
};
export default Fixture;
