import React, { useState } from "react";
import {
  Typography,
  Grid,
  TextField,
  Button,
  makeStyles,
  Backdrop,
  CircularProgress,
} from "@material-ui/core";
import Alert from "@material-ui/lab/Alert";
import { formatError } from "api/client";
import MainPaper from "core/MainPaper";

interface Props {
  register: (name: string) => void;
  loading: boolean;
  error: any;
}

const useStyles = makeStyles((theme) => ({
  buttons: {
    display: "flex",
    justifyContent: "flex-end",
    marginTop: theme.spacing(3),
  },
}));

export default ({ register, loading, error }: Props) => {
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
        <Backdrop open={loading}>
          <CircularProgress />
        </Backdrop>
      </MainPaper>
      {error && (
        <Alert severity="error" elevation={6}>
          Error signing up: {formatError(error)}
        </Alert>
      )}
    </>
  );
};
