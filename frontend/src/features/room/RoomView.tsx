import React from "react";
import { Table } from "model/table";
import MainPaper from "core/MainPaper";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemText from "@material-ui/core/ListItemText";
import Fab from "@material-ui/core/Fab";
import AddIcon from "@material-ui/icons/Add";
import Divider from "@material-ui/core/Divider";
import makeStyles from "@material-ui/core/styles/makeStyles";
import Typography from "@material-ui/core/Typography";

interface Props {
  tables: Table[];
  createTable: () => void;
}

const useStyles = makeStyles((theme) => ({
  addTable: {
    marginTop: theme.spacing(2),
    alignSelf: "center",
  },
}));

export default ({ tables, createTable }: Props) => {
  const classes = useStyles();
  return (
    <MainPaper>
      <Typography variant="h6">Aktive Tische</Typography>
      <List>
        {tables.map((t) => (
          <ListItem key={t.id}>
            <ListItemText>{`${t.owner}'s table ${t.id}`}</ListItemText>
          </ListItem>
        ))}
      </List>
      <Divider />
      <Fab color="primary" onClick={createTable} className={classes.addTable}>
        <AddIcon />
      </Fab>
    </MainPaper>
  );
};
