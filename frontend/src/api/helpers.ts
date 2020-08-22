import { TableId, Timestamp, GameType, BidType } from './karlchen_pb';

export function tableId(id: string): TableId {
  const ans = new TableId();
  ans.setValue(id);
  return ans;
}

export function toDate(ts: Timestamp): Date {
  return new Date(ts.getNanos() / 1e6);
}

export function gameTypeString(t: GameType) {
  switch (t) {
    case GameType.NORMAL_GAME:
      return 'Normalspiel';
    case GameType.MARRIAGE:
      return 'Hochzeit';
    case GameType.DIAMONDS_SOLO:
      return 'Karo Solo';
    case GameType.HEARTS_SOLO:
      return 'Herz Solo';
    case GameType.SPADES_SOLO:
      return 'Pik Solo';
    case GameType.CLUBS_SOLO:
      return 'Kreuz Solo';
    case GameType.MEATLESS_SOLO:
      return 'Fleischlos';
  }
}

export function bidString(b: BidType) {
  switch (b) {
    case BidType.RE_BID:
      return 'Re';
    case BidType.CONTRA_BID:
      return 'Contra';
    case BidType.RE_NO_NINETY:
    case BidType.CONTRA_NO_NINETY:
      return 'Keine 90';
    case BidType.RE_NO_SIXTY:
    case BidType.CONTRA_NO_SIXTY:
      return 'Keine 60';
    case BidType.RE_NO_THIRTY:
    case BidType.CONTRA_NO_THIRTY:
      return 'Keine 30';
    case BidType.RE_SCHWARZ:
    case BidType.CONTRA_SCHWARZ:
      return 'Schwarz';
  }
}
