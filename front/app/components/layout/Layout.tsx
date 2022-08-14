import React from 'react';
import { Footer } from './inc/Footer';
import { Header } from './inc/Header';

export type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout(props: LayoutProps) {
  return (
    <div className="min-h-screen flex flex-col justify-between bg-white">
      <Header />

      <div className="container mx-auto flex-1">
        <div className="mx-4 mt-20 mb-8">{props.children}</div>
      </div>

      <Footer />
    </div>
  );
}
