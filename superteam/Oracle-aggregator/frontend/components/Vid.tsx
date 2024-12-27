import React from 'react';

function Vid() {
  return (
    <video
      autoPlay
      loop
      muted
      className="fixed  top-0 left-0 w-full h-full object-cover z-0"
    >
      <source src="v4.mp4" type="video/mp4" />
    </video>
  );
}

export default Vid;
