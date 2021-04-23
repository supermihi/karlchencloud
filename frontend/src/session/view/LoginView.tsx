import * as React from "react";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import ForwardIcon from "@material-ui/icons/Forward";
import { formatError } from "api/client";
import MainPaper from "shared/MainPaper";
import SpinBackdrop from "shared/SpinBackdrop";
import ErrorAlert from "shared/ErrorAlert";

interface Props {
  name: string;
  loading: boolean;
  error?: unknown;
  login: () => void;
  forgetLogin: () => void;
  resetError?: () => void;
}

const LoginView: React.FC<Props> = ({
  login,
  forgetLogin,
  loading,
  error,
  name,
  resetError,
}) => (
  <>
    <MainPaper>
      <Typography component="h1" variant="h6" gutterBottom>
        Willkommen zurück, {name}!
      </Typography>
      <Grid container spacing={1}>
        <Grid item xs={12}>
          <Button
            endIcon={<ForwardIcon />}
            fullWidth
            variant="contained"
            color="primary"
            onClick={login}
          >
            Weiter
          </Button>
        </Grid>
        <Grid item xs={12}>
          <Button onClick={forgetLogin} fullWidth>
            Nicht {name}?
          </Button>
        </Grid>
      </Grid>
      <SpinBackdrop open={loading} />
    </MainPaper>
    {error && (
      <ErrorAlert
        message={`Error logging in: ${formatError(error)}`}
        reset={resetError}
      />
    )}
  </>
);
export default LoginView;
