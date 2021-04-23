import React from 'react';
import { Match } from 'model/match';
import tabletop from './resources/Pine_wood_Table_Top.jpg';
import { useResizeAwareRef } from 'shared/resizeHook';
import { Table } from 'model/table';
import MatchView from 'play/view/MatchView';
import InGameView from 'play/view/InGameView';

interface Props {
  match: Match | null;
  table: Table;
}
export default function PlayView({ match, table }: Props): React.ReactElement {
  const [ref, size] = useResizeAwareRef<HTMLDivElement>();
  return (
    <div
      style={{
        position: 'relative',
        width: '100%',
        height: '100%',
        backgroundImage: `url(${tabletop})`,
        backgroundSize: 'cover',
      }}
      ref={ref}
    >
      {/*<DeclarationDialogContainer />*/}
      <div
        style={{
          position: 'absolute',
          bottom: 0,
          width: '100%',
          height: '100%',
          overflow: 'hidden',
        }}
      >
        {match?.game && <InGameView game={match.game} tableSize={size} />}
        {match && <MatchView match={match} table={table} tableSize={size} />}
      </div>
    </div>
  );
}
