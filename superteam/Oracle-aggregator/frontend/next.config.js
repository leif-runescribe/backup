// next.config.js
module.exports = {
    async rewrites() {
      return [
        {
          source: '/api/:path*',
          destination: 'https://oracle-api-whvo.onrender.com/:path*', // Proxy to your API
        },
      ];
    },
  };
  