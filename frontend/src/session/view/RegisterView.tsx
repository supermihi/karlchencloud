import React, { useState } from 'react';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { formatError } from 'api/client';
import MainPaper from 'shared/MainPaper';
import SpinBackdrop from 'shared/SpinBackdrop';
import ErrorAlert from 'shared/ErrorAlert';
import { RegisterData } from '../model';
import { isValidEmail } from './validation';
import { useStyles } from './formstyle';

interface Props {
  register: (data: RegisterData) => void;
  loading: boolean;
  error?: unknown;
}

export default function RegisterView({ register, loading, error }: Props): React.ReactElement {
  const classes = useStyles();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const nameValid = name.trim() !== '';
  const emailValid = isValidEmail(email);
  const passwordValid = password.length >= 6;

  return (
    <>
      <MainPaper>
        <Typography component="h1" variant="h6">
          Willkommen!
        </Typography>
        <Typography component="p" variant="body1" gutterBottom>
          Sieht aus als wärst du zum ersten Mal hier. Wie heißt du?
        </Typography>
        <form noValidate className={classes.root} autoComplete="off">
          <TextField
            required
            autoComplete="name"
            error={!nameValid}
            value={name}
            onChange={(e) => setName(e.target.value)}
            fullWidth
            placeholder="Karlchen Müller"
          />
          <TextField
            required
            autoComplete="email"
            type="email"
            error={!emailValid}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            fullWidth
            placeholder="karlchen@mueller.de"
          />
          <TextField
            required
            autoComplete="new-password"
            error={!passwordValid}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            fullWidth
            type="password"
            placeholder="password"
          />
        </form>
        <div className={classes.buttons}>
          <Button
            disabled={!(nameValid && emailValid && passwordValid)}
            variant="contained"
            color="primary"
            onClick={() => register({ name, email, password })}
          >
            Los
          </Button>
        </div>
        <SpinBackdrop open={loading} />
      </MainPaper>
      {error && <ErrorAlert message={`Error signing up: ${formatError(error)}`} />}
    </>
  );
}
