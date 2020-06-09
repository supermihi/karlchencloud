import React from "react";
import Button from "@material-ui/core/Button";
import Backdrop from "@material-ui/core/Backdrop";
import CircularProgress from "@material-ui/core/CircularProgress";
import Typography from "@material-ui/core/Typography";
import Alert from "@material-ui/lab/Alert";
import Grid from "@material-ui/core/Grid";
import { formatError } from "api/client";
import MainPaper from "core/MainPaper";
import ForwardIcon from "@material-ui/icons/Forward";

interface Props {
  login: () => void;
  name: string;
  loading: boolean;
  error: any;
}

export default ({ name, loading, login, error }: Props) => {
  return (
    <>
      <MainPaper>
        <Typography component="h1" variant="h6" gutterBottom>
          Willkommen zurück!
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
              Als {name} einloggen
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Button fullWidth>Nicht {name}?</Button>
          </Grid>
        </Grid>

        <Backdrop open={loading}>
          <CircularProgress />
        </Backdrop>
      </MainPaper>
      {error && (
        <Alert severity="error" elevation={6}>
          Error logging in: {formatError(error)}
        </Alert>
      )}
    </>
  );
};
