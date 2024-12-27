"use client";
import React from "react";
import { motion } from "framer-motion";
import { slideInFromLeft, slideInFromRight, slideInFromTop } from "@/utils/motion";
import Image from "next/image";
import Link from "next/link";

const HeroContent = () => {
  return (
    <motion.div
      initial="hidden"
      animate="visible"
      className="flex flex-row items-center justify-center w-full px-40 mt-20 z-[20]"
    >
      <div className="h-full text-start w-full flex flex-col gap-5 justify-center">
        <motion.div variants={slideInFromTop} className=" py-4 mr-40 opacity-[0.9]">
          <div className=" justify-center ">
            <div className="w-40 border-2 py-2 px-2 flex items-center justify-center shadow-md">
              
              <p className="text-center  text-3xl font-bold bg-gradient-to-r from-white to-teal-700 text-transparent bg-clip-text">Agrigato</p>
            </div>
          </div>
        </motion.div>

        <motion.div 
        variants={slideInFromLeft(0.5)}
        className="flex flex-col gap-2 text-7xl text-bold text-white w-auto h-auto">
          Providing<span className="text-7xl font-extrabold bg-gradient-to-r from-cyan-800 to-cyan-100 text-transparent bg-clip-text">the best</span>
          price feeds        </motion.div>

        <motion.p
        variants={slideInFromLeft(0.8)}
        className="text-3xl text-gray-300 my-5 ">
          A blockchain aggregator you can trust to get the <br/>latest price updates from various oracles
        </motion.p>

        <motion.a
          variants={slideInFromLeft(1)}
          className="bg-black py-2 px-2 text-2xl button-primary text-center text-white cursor-pointer rounded-lg max-w-[180px]"
        >
          <Link href='/feed'>
          Get Started</Link>
        </motion.a>
      </div>
      
      {/* <motion.div
        variants={slideInFromRight(0.8)}
        className="justify-center items-center"
      >
        <Image
          src="/9.png"
          alt='cool img'
          height={650}
          width={650}
        />
      </motion.div> */}
      
    </motion.div>
  );
};

export default HeroContent;
