import { RefObject, useEffect, useRef, useState } from 'react';

export interface Size {
  width: number;
  height: number;
}
export function useResizeAwareRef<TElement extends HTMLElement>(): [RefObject<TElement>, Size] {
  const ref = useRef<TElement>(null);
  const [size, setSize] = useState<Size>({ width: 0, height: 0 });
  useEffect(() => {
    if (!ref.current) {
      return;
    }
    const htmlElement = ref.current;
    const updateSize = () => {
      if (!ref?.current) {
        return;
      }
      setSize({ width: htmlElement.clientWidth, height: htmlElement.clientHeight });
    };
    updateSize();
    window.addEventListener('resize', updateSize);
    return () => window.removeEventListener('resize', updateSize);
  }, [ref]);
  return [ref, size];
}
