import React from 'react';
import { useValue } from 'react-cosmos/fixture';
import { mockGrpcError } from 'shared/mock';
import LoginView from './LoginView';

const Fixture: React.FC = () => {
  const [error] = useValue('error', { defaultValue: false });
  const [loading] = useValue('loading', { defaultValue: false });
  return (
    <LoginView
      loading={loading}
      error={error && mockGrpcError('mock')}
      login={() => console.log('login')}
    />
  );
};
export default Fixture;
