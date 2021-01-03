import React from 'react';
import * as d from '@storybook/react';
import { Provider } from 'react-redux';
import { store } from 'app/store';

export const withProvider: d.DecoratorFn = (story) => <Provider store={store}>{story()}</Provider>;
