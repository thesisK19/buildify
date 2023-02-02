import React from 'react';
import YouTube from 'react-youtube';
import { YoutubeDiv } from './styled';

export const Video = (props) => {
  const { videoId, width, height } = props;

  return (
    <YoutubeDiv width={width} height={height}>
      <YouTube
        videoId={videoId}
        opts={{
          width: '100%',
          height: '100%',
        }}
      />
    </YoutubeDiv>
  );
};

