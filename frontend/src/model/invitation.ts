export function inviteLink(inviteCode: string) {
  return `${window.location.origin}?invitecode=${inviteCode}`;
}
