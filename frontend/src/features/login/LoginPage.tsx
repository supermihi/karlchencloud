import React, { useState } from "react";
import {
  Typography,
  Paper,
  Grid,
  TextField,
  Button,
  makeStyles,
  Backdrop,
  CircularProgress,
  Snackbar,
} from "@material-ui/core";
import Alert from "@material-ui/lab/Alert";
import { formatError } from "api/client";

interface Props {
  signup: (name: string) => void;
  loading: boolean;
  error: any;
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(3),
    padding: theme.spacing(2),
  },
  buttons: {
    display: "flex",
    justifyContent: "flex-end",
    marginTop: theme.spacing(3),
  },
}));

export default ({ signup, loading, error }: Props) => {
  const classes = useStyles();
  const [name, setName] = useState("");
  const valid = name.trim() !== "";
  return (
    <>
      <Paper className={classes.paper}>
        <Typography component="h1" variant="h6">
          Willkommen!
        </Typography>
        <Typography component="p" variant="body1" gutterBottom>
          Sieht aus als wärst du zum ersten Mal hier. Wie heißt du?
        </Typography>
        <Grid container>
          <Grid item xs={12}>
            <TextField
              required
              error={!valid}
              value={name}
              onChange={(e) => setName(e.target.value)}
              fullWidth
              placeholder="Gorm"
            />
          </Grid>
        </Grid>
        <div className={classes.buttons}>
          <Button
            disabled={!valid}
            variant="contained"
            color="primary"
            onClick={(_) => signup(name)}
          >
            Los
          </Button>
        </div>
        <Backdrop open={loading}>
          <CircularProgress />
        </Backdrop>
      </Paper>
      {error && (
        <Alert severity="error" elevation={6}>
          Error signing up: {formatError(error)}
        </Alert>
      )}
    </>
  );
};
