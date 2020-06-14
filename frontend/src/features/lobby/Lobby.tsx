import React from "react";
import { Table } from "model/table";
import MainPaper from "core/MainPaper";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Divider from "@material-ui/core/Divider";
import makeStyles from "@material-ui/core/styles/makeStyles";
import Typography from "@material-ui/core/Typography";

interface Props {
  activeTable: Table | null;
  createTable: () => void;
}

const useStyles = makeStyles((theme) => ({
  addTable: {
    marginTop: theme.spacing(2),
    alignSelf: "center",
  },
}));

export default ({ activeTable, createTable }: Props) => {
  const classes = useStyles();
  return (
    <MainPaper>
      {activeTable && <Typography variant="h6">Aktiver Tisch:</Typography>}

      <Divider />
      <Fab color="primary" onClick={createTable} className={classes.addTable}>
        <AddIcon />
      </Fab>
    </MainPaper>
  );
};
