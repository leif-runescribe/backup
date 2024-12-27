'use client'
import Background from "@/components/Background";

import PriceFeedTable from "@/components/PriceFeedTable";
import axios from "axios";
import React, { useState } from "react";

const sampleData = [
  {
    asset: "BTC/USD",
    price: 34250.75,
    confidence: 25.34,
    lastUpdated: Date.now() - 1000 * 60 * 1, // 1 minute ago
  },
  {
    asset: "ETH/USD",
    price: 2100.5,
    confidence: 10.21,
    lastUpdated: Date.now() - 1000 * 60 * 2, // 2 minutes ago
  },
  {
    asset: "SOL/USD",
    price: 35.67,
    confidence: 0.75,
    lastUpdated: Date.now() - 1000 * 60 * 3, // 3 minutes ago
  },
  {
    asset: "BTC/USD",
    price: 34250.75,
    confidence: 25.34,
    lastUpdated: Date.now() - 1000 * 60 * 1, // 1 minute ago
  },
  {
    asset: "ETH/USD",
    price: 2100.5,
    confidence: 10.21,
    lastUpdated: Date.now() - 1000 * 60 * 2, // 2 minutes ago
  },
  {
    asset: "SOL/USD",
    price: 35.67,
    confidence: 0.75,
    lastUpdated: Date.now() - 1000 * 60 * 3, // 3 minutes ago
  },
  {
    asset: "BTC/USD",
    price: 34250.75,
    confidence: 25.34,
    lastUpdated: Date.now() - 1000 * 60 * 1, // 1 minute ago
  },
  {
    asset: "ETH/USD",
    price: 2100.5,
    confidence: 10.21,
    lastUpdated: Date.now() - 1000 * 60 * 2, // 2 minutes ago
  },
  {
    asset: "SOL/USD",
    price: 35.67,
    confidence: 0.75,
    lastUpdated: Date.now() - 1000 * 60 * 3, // 3 minutes ago
  },
  {
    asset: "BTC/USD",
    price: 34250.75,
    confidence: 25.34,
    lastUpdated: Date.now() - 1000 * 60 * 1, // 1 minute ago
  },
  {
    asset: "ETH/USD",
    price: 2100.5,
    confidence: 10.21,
    lastUpdated: Date.now() - 1000 * 60 * 2, // 2 minutes ago
  },
  {
    asset: "SOL/USD",
    price: 35.67,
    confidence: 0.75,
    lastUpdated: Date.now() - 1000 * 60 * 3, // 3 minutes ago
  },
  
];

const page = () => {

  const [id, setId] = useState('');
  const [pair, setPair] = useState('');
  const [response1, setResponse1] = useState(null);
  const [response2, setResponse2] = useState(null);

  const handleKeyPress = (event) => {
    if (event.key === 'Enter') {
      performSearch();
    }
  };

  const performSearch = () => {
    if (id.trim() === '' || pair.trim() === '') {
      alert('Please fill in both ID and pair fields.');
      return;
    }

    const pairRegex = /^[a-zA-Z0-9]+\/[a-zA-Z0-9]+$/;
    if (!pairRegex.test(pair)) {
      alert('Pair must be in the format abc/def.');
      return;
    }

    axios.get(`https://oracle-api-whvo.onrender.com/pyth/${id}`)
      .then(response => {
        setResponse1(response.data);
      })
      .catch(error => {
        console.error('Error fetching response 1:', error);
        // Handle error
      });

    axios.get(`https://oracle-api-whvo.onrender.com/band/${pair}`)
      .then(response => {
        setResponse2(response.data);
      })
      .catch(error => {
        console.error('Error fetching response 2:', error);
        // Handle error
      });
  };

  return (
    <div>
      <Background />
      <div className="container mx-auto py-8">
      <div className="flex items-center text-black text-xl space-x-4">
        <input
          type="text"
          value={id}
          onChange={(e) => setId(e.target.value)}
          onKeyPress={handleKeyPress}
          placeholder="Enter ID"
          className="outline-none px-2 py-1 border border-gray-300 text-black rounded-md"
        />
        <input
          type="text"
          value={pair}
          onChange={(e) => setPair(e.target.value)}
          onKeyPress={handleKeyPress}
          placeholder="Enter Pair"
          className="outline-none px-2 py-1 border border-gray-300 rounded-md"
        />
        <button
          onClick={performSearch}
          className="px-4 py-1 bg-blue-500 text-white rounded-md hover:bg-blue-600"
        >
          Search
        </button>
      </div>
      {/* Your other content here */}
    </div>
      <div className="py-6">
        <h1 className="text-center text-5xl text-white  mb-4">
          Agrigato
        </h1>
        <PriceFeedTable feeds={sampleData} />
      </div>
      page
    </div>
  );
};

export default page;
