import * as React from 'react';
import Button from '@material-ui/core/Button';
import Backdrop from '@material-ui/core/Backdrop';
import CircularProgress from '@material-ui/core/CircularProgress';
import Typography from '@material-ui/core/Typography';
import Alert from '@material-ui/lab/Alert';
import Grid from '@material-ui/core/Grid';
import ForwardIcon from '@material-ui/icons/Forward';
import { formatError } from 'api/client';
import MainPaper from 'components/MainPaper';

import { MyUserData, Credentials } from 'app/auth';
import makeStyles from '@material-ui/core/styles/makeStyles';

interface Props {
  currentLogin: MyUserData;
  loading: boolean;
  error?: any;
  login: (login: Credentials) => void;
  forgetLogin: () => void;
  resetError?: () => void;
}

const useStyles = makeStyles((theme) => ({
  backdrop: {
    zIndex: theme.zIndex.drawer + 1,
    color: '#fff',
  },
}));

export default function LoginView({
  login,
  forgetLogin,
  loading,
  currentLogin,
  error,
  resetError,
}: Props) {
  const classes = useStyles();
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
              onClick={() => login(currentLogin)}
            >
              Als {currentLogin.name} einloggen
            </Button>
          </Grid>
          <Grid item xs={12}>
            <Button onClick={forgetLogin} fullWidth>
              Nicht {currentLogin.name}?
            </Button>
          </Grid>
        </Grid>

        <Backdrop open={loading} className={classes.backdrop}>
          <CircularProgress />
        </Backdrop>
      </MainPaper>
      {error && (
        <Alert onClose={() => resetError && resetError()} severity="error" elevation={6}>
          Error logging in: {formatError(error)}
        </Alert>
      )}
    </>
  );
}
