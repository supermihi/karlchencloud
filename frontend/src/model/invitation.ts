export function inviteLink(inviteCode: string): string {
  return `${window.location.origin}?invitecode=${inviteCode}`;
}

export function parseInviteCode(location: string): string | null {
  const re = /\?invitecode=(\w+)$/;
  const match = location.match(re);
  if (match) {
    return match[1];
  }
  return null;
}
