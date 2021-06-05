import React from 'react';
import { useValue } from 'react-cosmos/fixture';
import { mockGrpcError } from 'shared/mock';
import LoginWithTokenView from './LoginWithTokenView';

const Fixture: React.FC = () => {
  const [error] = useValue('error', { defaultValue: false });
  const [loading] = useValue('loading', { defaultValue: false });
  return (
    <LoginWithTokenView
      name="Nils"
      loading={loading}
      error={error && mockGrpcError('mock')}
      login={() => console.log('login')}
      forgetLogin={() => console.log('forget')}
    />
  );
};
export default Fixture;
