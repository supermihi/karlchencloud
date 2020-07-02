import { parseInviteCode } from './invitation';

test('can parse invitation code', () => {
  const url = 'https://karlchencloud.provider.net:7090?invitecode=12345';
  const invite = parseInviteCode(url);
  expect(invite).toEqual('12345');
});
