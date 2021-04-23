import * as React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Paper, { PaperProps } from '@material-ui/core/Paper';

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(3),
    marginBottom: theme.spacing(3),
    padding: theme.spacing(2),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'stretch',
  },
}));

const MainPaper: React.FC<PaperProps> = (props) => {
  const classes = useStyles();
  return <Paper {...props} className={classes.paper} />;
};
export default MainPaper;
