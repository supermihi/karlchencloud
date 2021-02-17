import React, { useEffect, useRef, useState } from 'react';

export default function GameView(): React.ReactElement {
  const dRef = useRef<HTMLDivElement>(null);
  const [width, setWidth] = useState(0);
  useEffect(() => {
    if (dRef.current) {
      setWidth(dRef.current.clientWidth);
    }
  }, [dRef]);
  return (
    <div ref={dRef}>
      <p>width: {width}</p>
    </div>
  );
}
