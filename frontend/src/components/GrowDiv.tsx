import * as React from 'react';
import makeStyles from '@material-ui/core/styles/makeStyles';
import { HTMLProps } from 'react';

const useStyles = makeStyles((theme) => ({
  grow: {
    flexGrow: 1,
  },
}));

export default function GrowDiv(props: HTMLProps<HTMLDivElement>) {
  const classes = useStyles();
  return <div className={classes.grow} {...props} />;
}
