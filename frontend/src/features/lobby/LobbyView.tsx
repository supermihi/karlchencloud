import React from "react";
import { Table } from "model/table";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Divider from "@material-ui/core/Divider";
import makeStyles from "@material-ui/core/styles/makeStyles";
import CurrentTableView from "./CurrentTableView";

interface Props {
  activeTable?: Table;
  createTable: () => void;
}

const useStyles = makeStyles((theme) => ({
  addTable: {
    marginTop: theme.spacing(2),
    alignSelf: "center",
  },
  main: {
    display: "flex",
    flexDirection: "column",
  },
}));

export default ({ activeTable, createTable }: Props) => {
  const classes = useStyles();
  return (
    <div className={classes.main}>
      {activeTable && <CurrentTableView table={activeTable} />}

      <Divider />
      <Fab
        variant="extended"
        onClick={createTable}
        className={classes.addTable}
      >
        <AddIcon />
        Tisch starten
      </Fab>
    </div>
  );
};
