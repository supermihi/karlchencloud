import React from "react";
import Alert from "@material-ui/lab/Alert";

interface Props {
  message: React.ReactNode;
  reset?: () => void;
}

const ErrorAlert: React.FC<Props> = ({ message, reset }) => (
  <Alert onClose={() => reset && reset()} severity="error" elevation={6}>
    {message}
  </Alert>
);
export default ErrorAlert;
