import Image from 'next/image';
import React from 'react';
import HeroContent from './HeroContent';
import Background from './Background';

const Stuff = () => {
  return (
    <div className='max-h-full mt-72 '>
      
      <div className=' w-full flex md:flex-row flex-col '>
        <div className='  bg-black py-40 px-20  '>
        <h1 className=' text-6xl text-white'>Latest Price Feeds from the Leading Oracles in the Industry
        
        </h1>
        </div>
        <div className='bg-black pt-32 p-20 '>
          <Background/>
          <Image
          src="/8.png"
          alt='cool img'
          height={1000}
          width={300}
        />
        <h1 className='text-4xl  text-white'>AGRIGATO serves as an Oracle Aggregator which aims to fetch latest and most accurate price feeds for the Defi ecosystem
        
        </h1>
        </div>
      </div>
    </div>
  );
};

export default Stuff;
