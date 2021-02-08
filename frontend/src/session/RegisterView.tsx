import React, { useState } from "react";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import makeStyles from "@material-ui/core/styles/makeStyles";
import { formatError } from "api/client";
import MainPaper from "shared/MainPaper";
import SpinBackdrop from "shared/SpinBackdrop";
import ErrorAlert from "shared/ErrorAlert";

interface Props {
  register: (name: string) => void;
  loading: boolean;
  error?: unknown;
}

const useStyles = makeStyles((theme) => ({
  buttons: {
    display: "flex",
    justifyContent: "flex-end",
    marginTop: theme.spacing(3),
  },
}));

export default function RegisterView({ register, loading, error }: Props) {
  const classes = useStyles();
  const [name, setName] = useState("");
  const valid = name.trim() !== "";
  return (
    <>
      <MainPaper>
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
              onSubmit={() => valid && register(name)}
              placeholder="Gorm"
            />
          </Grid>
        </Grid>
        <div className={classes.buttons}>
          <Button
            disabled={!valid}
            variant="contained"
            color="primary"
            onClick={(_) => register(name)}
          >
            Los
          </Button>
        </div>
        <SpinBackdrop open={loading} />
      </MainPaper>
      {error && (
        <ErrorAlert message={`Error signing up: ${formatError(error)}`} />
      )}
    </>
  );
}
