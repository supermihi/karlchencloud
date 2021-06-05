import React, { useState } from 'react';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { formatError } from 'api/client';
import MainPaper from 'shared/MainPaper';
import SpinBackdrop from 'shared/SpinBackdrop';
import ErrorAlert from 'shared/ErrorAlert';
import { LoginData } from '../model';
import { useStyles } from './formstyle';
import { isValidEmail } from './validation';

interface Props {
  login: (data: LoginData) => void;
  loading: boolean;
  error?: unknown;
}

export default function LoginView({ login, loading, error }: Props): React.ReactElement {
  const classes = useStyles();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const emailValid = isValidEmail(email);
  const passwordValid = password.trim() !== '';

  return (
    <>
      <MainPaper>
        <Typography component='h1' variant='h6'>
          Bei Karlchencloud einloggen
        </Typography>
        <form noValidate className={classes.root} autoComplete='off'>
          <TextField
            required
            autoComplete='email'
            type='email'
            error={!emailValid}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            fullWidth
            placeholder='karlchen@mueller.de'
          />
          <TextField
            required
            autoComplete='current-password'
            error={!passwordValid}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            fullWidth
            type='password'
            placeholder='password'
          />
        </form>
        <div className={classes.buttons}>
          <Button
            disabled={!(emailValid && passwordValid)}
            variant='contained'
            color='primary'
            onClick={() => login({ email, password })}
          >
            Los
          </Button>
        </div>
        <SpinBackdrop open={loading} />
      </MainPaper>
      {error && <ErrorAlert message={`Error logging in: ${formatError(error)}`} />}
    </>
  );
}
