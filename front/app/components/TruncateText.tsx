import React, { useEffect, useRef, useState } from 'react';
import { twMerge } from 'tailwind-merge';

type ClampedTextProps = {
  children: string;
  className?: string;
};

export default function TruncateText(props: ClampedTextProps) {
  const [isTruncated, setIsTruncated] = useState(false);
  const ref = useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (ref.current) {
      const { scrollWidth, clientWidth } = ref.current;

      setIsTruncated(scrollWidth > clientWidth);
    }
  }, [ref, ref.current?.scrollWidth, ref.current?.clientWidth, props.children]);

  return (
    <div
      className={twMerge(`truncate`, isTruncated && 'cursor-help', props.className)}
      ref={ref}
      title={isTruncated ? props.children : undefined}
    >
      {props.children}
    </div>
  );
}
