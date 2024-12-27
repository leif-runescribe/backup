import { useEffect, useState } from 'react';
import axios from 'axios';
import { PriceServiceConnection } from '@pythnetwork/price-service-client';

const PriceFeed = () => {
    const [prices, setPrices] = useState(null);

    useEffect(() => {
        const fetchPrices = async () => {
            try {
                const connection = new PriceServiceConnection("https://hermes.pyth.network");

                const priceIds = [
                    "0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43", // BTC/USD price id
                    "0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace", // ETH/USD price id
                ];

                const currentPrices = await connection.getLatestPriceFeeds(priceIds);
                setPrices(currentPrices);
            } catch (error) {
                console.error('Error fetching prices:', error);
            }
        };

        fetchPrices();
    }, []); // Empty dependency array ensures this runs once

    return (
        <div>
            <h2>Latest Prices</h2>
            {prices ? (
                <div>
                    <p>BTC/USD: ${prices[0].price}</p>
                    <p>ETH/USD: ${prices[1].price}</p>
                </div>
            ) : (
                <p>Loading...</p>
            )}
        </div>
    );
};

export default PriceFeed;
