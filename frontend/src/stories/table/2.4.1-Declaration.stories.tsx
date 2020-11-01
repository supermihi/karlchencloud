import React from 'react';
import { action } from '@storybook/addon-actions';
import DeclarationDialog from 'features/table/DeclarationDialog';

export default {
  title: 'Match/Declaration',
  component: DeclarationDialog,
};

export const Dialog = () => {
  return <DeclarationDialog open declare={action('declare')} />;
};
