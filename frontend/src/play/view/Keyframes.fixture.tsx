import React from 'react';
import Button from '@material-ui/core/Button';

import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';

const useStyles = makeStyles((theme) => ({
  animatedItem: {
    animation: `$myEffect 3000ms ${theme.transitions.easing.easeInOut}`,
  },
  animatedItemExiting: {
    animation: `$myEffectExit 3000ms ${theme.transitions.easing.easeInOut}`,
    opacity: 0,
    transform: 'translateY(-200%)',
  },
  '@keyframes myEffect': {
    '0%': {
      opacity: 0,
      transform: 'translateY(-200%)',
    },
    '100%': {
      opacity: 1,
      transform: 'translateY(0)',
    },
  },
  '@keyframes myEffectExit': {
    '0%': {
      opacity: 1,
      transform: 'translateY(0)',
    },
    '100%': {
      opacity: 0,
      transform: 'translateY(-200%)',
    },
  },
}));
export default function Keyframes(): React.ReactElement {
  const classes = useStyles();
  const [exit, setExit] = React.useState(false);
  return (
    <>
      <div
        className={clsx(classes.animatedItem, {
          [classes.animatedItemExiting]: exit,
        })}
      >
        <h1>Hello CodeSandbox</h1>
        <h2>Start editing to see some magic happen!</h2>
        <Button onClick={() => setExit(true)}>Click to exit</Button>
      </div>
      {exit && <Button onClick={() => setExit(false)}>Click to enter</Button>}
    </>
  );
}
