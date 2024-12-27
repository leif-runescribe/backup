// components/PriceFeedTable.tsx
import React from 'react';

interface Feed {
  asset: string;
  price: number;
  confidence: number;
  lastUpdated: number;
}

interface PriceFeedTableProps {
  feeds: Feed[];
}

const PriceFeedTable: React.FC<PriceFeedTableProps> = ({ feeds }) => {
  return (
    <div className="container mx-auto p-4">
      <div className="overflow-x-auto">
        <table className="min-w-full text-white text-xl divide-y ">
          <thead className="">
            <tr>
              <th
                scope="col"
                className="px-6 py-3 "
              >
                Asset
              </th>
              <th
                scope="col"
                className="px-6 py-3 "
              >
                Aggregated Price
              </th>
              <th
                scope="col"
                className="px-6 py-3 "
              >
                Confidence Interval
              </th>
              <th
                scope="col"
                className="px-6 py-3 "
              >
                Last Updated
              </th>
              <th
                scope="col"
                className="px-6 py-3 "
              >
                Pyth
              </th>
              <th
                scope="col"
                className="px-6 py-3 "
              >
                Band
              </th>
            </tr>
          </thead>
          <tbody className=" divide-y divide-gray-200">
            {feeds.map((feed) => (
              <tr key={feed.asset}>
                <td className="px-6 py-4 whitespace-nowrap  text-gray-500">{feed.asset}</td>
                <td className="px-6 py-4 whitespace-nowrap  text-gray-500">${feed.price.toFixed(2)}</td>
                <td className="px-6 py-4 whitespace-nowrap  text-gray-500">±${feed.confidence}</td>
                <td className="px-6 py-4 whitespace-nowrap  text-gray-500">{new Date(feed.lastUpdated).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default PriceFeedTable;
