import React from 'react';
import { action } from '@storybook/addon-actions';
import DeclarationDialog from 'features/table/DeclarationDialog';

/* eslint import/no-anonymous-default-export: [2, {"allowObject": true}] */
export default {
  title: 'Match/Declaration',
  component: DeclarationDialog,
};

export const Dialog = () => {
  return <DeclarationDialog open declare={action('declare')} />;
};
