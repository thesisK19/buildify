import styled from 'styled-components';

export const YoutubeDiv = styled.div<any>`
width: ${(props) => (props.width)};
height: ${(props) => (props.height)};
> div {
  height: 100%;
}
iframe {
  pointer-events: ${(props) => (props.enabled ? 'none' : 'auto')};
  // width:100%!important;
  // height:100%!important;
}
`;