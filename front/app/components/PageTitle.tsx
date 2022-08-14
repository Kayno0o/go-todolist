import React from 'react';
import { twMerge } from 'tailwind-merge';

type PageTitleProps = {
  children: React.ReactNode;
  className?: string;
};

export default function PageTitle(props: PageTitleProps) {
  return <h1 className={twMerge('text-3xl font-bold text-slate-800 ml-4 mb-4', props.className)}>{props.children}</h1>;
}
