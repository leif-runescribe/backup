'use client'
import React from 'react';
import HeroContent from './HeroContent';
import Stuff from './Stuff';
import Vid from './Vid';

import Base from './Base';

const Home = () => {
  return (
    <div className="relative flex flex-col h-screen w-screen">
      <video
        autoPlay
        loop
        muted
        className="pointer-events-none select-none absolute top-0 left-0 z-[1] h-full w-full object-cover overflow-hidden">
        <source src="/v4.mp4" type="video/mp4" />
      </video>
      

      <div className="absolute top-0 left-0 z-20 h-full w-full bg-black opacity-50"></div>
      <HeroContent/>
      
      
    
     
    </div>
  );
};

export default Home;
